package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math"
	"net/http"
	"os"
	"strings"

	"github.com/skratchdot/open-golang/open"
)

// Flags
var VideoId = flag.String("video", "", "The YouTube video id.")
var YouTubeKey = flag.String("ytkey", "", "Google API key.")
var StopWordFiles = flag.String("stopwords", "", "A list of file paths, comma delimited, of stop word files.")
var Verbose = flag.Bool("verbose", false, "Extra logging to std out")
var RedisServer = flag.String("redis", "127.0.0.1:6379", "Redis server and port.")
var WebServer = flag.Bool("server", false, "Run as a web server.")
var Port = flag.String("port", "8000", "Port for web server to run.")
var TrainingFiles = flag.String("training", "", "Training text files.")

func LogMsg(msg string) {
	if *Verbose {
		fmt.Printf(msg)
	}
}

type WebError struct {
	Error string
}

type Report struct {
	VideoId                string
	TotalComments          uint64
	CollectedComments      int
	CommentCoveragePercent float64
	ChannelTitle           string
	VideoTitle             string
	Keywords               []string
	Sentiment              []SentimentTag
}

func webHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path[1:] == "api" {
		vid := r.URL.Query().Get("vid")

		var jsonBytes []byte

		if vid != "" {
			jsonBytes = runReport(vid)
		} else {
			jsonBytes, _ = json.Marshal(WebError{Error: "Missing video id."})
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBytes)
	} else {
		http.ServeFile(w, r, "static/gui/index.html")
	}
}

func runReport(vid string) []byte {

	// Create Report
	report := Report{}

	done := make(chan bool)

	// Poll the data sources...
	go func() {
		metadata := GetVideoInfo(vid)
		// Set video metadata
		report.VideoId = vid
		report.TotalComments = metadata.TotalComments
		report.ChannelTitle = metadata.ChannelTitle
		report.VideoTitle = metadata.Title

		done <- true
	}()

	// Fetch the comments
	comments := GetComments_v2(vid)

	// Set comments returned
	report.CollectedComments = len(comments)
	report.CommentCoveragePercent = math.Ceil((float64(report.CollectedComments) / float64(report.TotalComments)) * float64(100))

	// Set Keywords
	go func() {
		report.Keywords = GetKeywords(comments)

		done <- true
	}()

	// Sentiment Tagging
	go func() {
		if *RedisServer != "" {
			report.Sentiment = GetSentimentSummary(comments)
		}

		done <- true
	}()

	// Wait for everything to finish up
	for i := 0; i < 3; i++ {
		<-done
	}

	reportJson, _ := json.Marshal(report)

	// Output Report
	return reportJson
}

func main() {
	flag.Parse()

	// Check if they want to upload training data to the semantic engine
	if *RedisServer != "" && *TrainingFiles != "" {
		trainingFiles := strings.Split(*TrainingFiles, ",")
		for _, path := range trainingFiles {
			LoadTrainingData(path)
		}
		LogMsg("Training data uploaded.")
		os.Exit(1)
	}

	// Check for required params and run
	if *YouTubeKey == "" {
		fmt.Println("A Google API key with YouTube API access is required.")
		os.Exit(1)
	}

	if !*WebServer && *VideoId == "" {
		fmt.Println("Video ID is required.")
		os.Exit(1)
	}

	if *StopWordFiles != "" {
		swFiles := strings.Split(*StopWordFiles, ",")
		for _, path := range swFiles {
			LoadStopWords(path)
		}
	}

	if *WebServer {
		fmt.Println("Web server running on " + *Port)

		open.Start("http://localhost:" + *Port)

		http.HandleFunc("/", webHandler)
		http.ListenAndServe(":"+*Port, nil)
	} else {
		fmt.Println(string(runReport(*VideoId)))
	}
}
