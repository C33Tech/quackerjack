package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/skratchdot/open-golang/open"
)

// Command Line Flags:

// The target YouTube video ID
var VideoID = flag.String("video", "", "The YouTube video id.")

// The Google API key
var YouTubeKey = flag.String("ytkey", "", "Google API key.")

// A list of stopword files to assist in keyword extraction
var StopWordFiles = flag.String("stopwords", "", "A list of file paths, comma delimited, of stop word files.")

// Standard verbose flag.
var Verbose = flag.Bool("verbose", false, "Extra logging to std out")

// The server and port for the required redis server
var RedisServer = flag.String("redis", "127.0.0.1:6379", "Redis server and port.")

// Changes the output from stdout and starts a web server
var WebServer = flag.Bool("server", false, "Run as a web server.")

// The preferred web server port
var Port = flag.String("port", "8000", "Port for web server to run.")

// Sentiment training file paths to be ingested in to redis.
var TrainingFiles = flag.String("training", "", "Training text files.")

// stdout messaging if the verbose flag is set.
func LogMsg(msg string) {
	if *Verbose {
		fmt.Printf(msg)
	}
}

type webError struct {
	Error string
}

type report struct {
	VideoID                string
	VideoViews             uint64
	TotalComments          uint64
	PublishedAt            string
	CollectedComments      int
	CommentCoveragePercent float64
	CommentAvgPerDay       float64
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
			jsonBytes, _ = json.Marshal(webError{Error: "Missing video id."})
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBytes)
	} else {
		http.ServeFile(w, r, "static/gui/index.html")
	}
}

func runReport(vid string) []byte {

	// Create Report
	theReport := report{}

	done := make(chan bool)

	// Poll the data sources...
	go func() {
		metadata := GetVideoInfo(vid)
		// Set video metadata
		theReport.VideoID = vid
		theReport.VideoViews = metadata.VideoViews
		theReport.TotalComments = metadata.TotalComments
		theReport.ChannelTitle = metadata.ChannelTitle
		theReport.VideoTitle = metadata.Title
		theReport.PublishedAt = metadata.PublishedAt

		done <- true
	}()

	// Fetch the comments
	comments := GetCommentsV2(vid)

	// Set comments returned
	theReport.CollectedComments = len(comments)
	theReport.CommentCoveragePercent = math.Ceil((float64(theReport.CollectedComments) / float64(theReport.TotalComments)) * float64(100))

	// Set Keywords
	go func() {
		theReport.Keywords = GetKeywords(comments)

		done <- true
	}()

	// Sentiment Tagging
	go func() {
		if *RedisServer != "" {
			theReport.Sentiment = GetSentimentSummary(comments)
		}

		done <- true
	}()

	// Wait for everything to finish up
	for i := 0; i < 3; i++ {
		<-done
	}

	// Calculate Average Daily Comments
	t, _ := time.Parse(time.RFC3339Nano, theReport.PublishedAt)
	delta := time.Now().Sub(t)
	theReport.CommentAvgPerDay = float64(theReport.TotalComments) / (float64(delta.Hours()) / float64(24))

	reportJSON, _ := json.Marshal(theReport)

	// Output Report
	return reportJSON
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

	if !*WebServer && *VideoID == "" {
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
		fmt.Println(string(runReport(*VideoID)))
	}
}
