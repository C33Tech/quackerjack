package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Command Line Flags:

// The PostURL from the target YouTube video.
var PostURL = flag.String("post", "", "The target post url (YouTube or Instagram).")

// The YouTubeKey is a Google API key with access to YouTube's Data API
var YouTubeKey = flag.String("ytkey", "", "Google API key.")

// The YouTubeKey is a Google API key with access to YouTube's Data API
var InstagramKey = flag.String("igkey", "", "Instagram API key.")

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

// LogMsg takes a message and pipes it to stdout if the verbose flag is set.
func LogMsg(msg string) {
	if *Verbose {
		fmt.Printf(msg)
	}
}

type webError struct {
	Error string
}

type report struct {
	ID                     string
	URL                    string
	Type                   string
	Title                  string
	PublishedAt            string
	TotalComments          uint64
	CollectedComments      uint64
	CommentCoveragePercent float64
	CommentAvgPerDay       float64
	Keywords               []string
	Sentiment              []SentimentTag
	Metadata               Post
	SampleComments         []*Comment
}

// Post is the interface for all the various post types (YouTubeVideo, etc...)
type Post interface {
	GetComments() CommentList
	GetMetadata() bool
}

func jsonError(msg string) []byte {
	errorJSON, _ := json.Marshal(webError{Error: msg})
	return errorJSON
}

func parseURL(url string) (string, string) {
	sites := map[string]string{
		"instagram": "instag\\.?ram(\\.com)?/p/([\\w]*)/?",
		"youtube":   "youtu\\.?be(\\.?com)?/(watch\\?v=)?([\\w\\-_]*)",
	}

	var domain, id string

	for d, rstr := range sites {
		r, _ := regexp.Compile(rstr)
		matches := r.FindStringSubmatch(url)
		if len(matches) > 0 {
			domain = d
			id = matches[len(matches)-1]
			break
		}
	}

	return domain, id
}

func webHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path[1:] == "api" {
		postURL := r.URL.Query().Get("vid")

		var jsonBytes []byte

		if postURL != "" {
			jsonBytes = runReport(postURL)
		} else {
			jsonBytes, _ = json.Marshal(webError{Error: "Missing video id."})
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBytes)
	} else {
		data, err := Asset("static/gui/index.html")
		if err != nil {
			fmt.Println("Web GUI asset not found!")
			os.Exit(1)
		}

		w.Header().Set("Content-Type", "text/html")
		w.Write(data)
	}
}

func runReport(postURL string) []byte {
	// Parse URL
	domain, pid := parseURL(postURL)
	if domain == "" || pid == "" {
		return jsonError("Unable to parse post url.")
	}

	// Create Report
	theReport := report{URL: postURL}
	var thePost Post

	switch domain {
	case "youtube":
		if *YouTubeKey == "" {
			return jsonError("API key for YouTube not configured.")
		}

		thePost = &YouTubeVideo{ID: pid}
	case "instagram":
		if *InstagramKey == "" {
			return jsonError("API key for Instagram not configured.")
		}

		thePost = &InstagramPic{ShortCode: pid}
	}

	// Fetch the metadata
	flag := thePost.GetMetadata()

	if !flag {
		return jsonError("Could not fetch metadata.")
	}

	switch p := thePost.(type) {
	case *YouTubeVideo:
		theReport.Type = "YouTubeVideo"
		theReport.ID = p.ID
		theReport.Title = p.Title
		theReport.PublishedAt = p.PublishedAt
		theReport.TotalComments = p.TotalComments
		theReport.Metadata = p
	case *InstagramPic:
		theReport.Type = "InstagramPic"
		theReport.ID = p.ID
		theReport.Title = p.Caption
		theReport.PublishedAt = p.PublishedAt
		theReport.TotalComments = p.TotalComments
		theReport.Metadata = p
	}

	// Fetch the comments
	comments := thePost.GetComments()

	// If we don't get an comments back, wait for the metadata call to return and send an error.
	if comments.IsEmpty() {
		return jsonError("No comments found for this post.")
	}

	// Set comments returned
	theReport.CollectedComments = comments.GetTotal()
	theReport.CommentCoveragePercent = math.Ceil((float64(theReport.CollectedComments) / float64(theReport.TotalComments)) * float64(100))

	done := make(chan bool)

	// Set Keywords
	go func() {
		theReport.Keywords = comments.GetKeywords()

		done <- true
	}()

	// Sentiment Tagging
	go func() {
		if *RedisServer != "" {
			theReport.Sentiment = comments.GetSentimentSummary()
		}

		done <- true
	}()

	// Wait for everything to finish up
	for i := 0; i < 2; i++ {
		<-done
	}

	// Pull a few sample comments
	theReport.SampleComments = comments.GetRandom(3)

	// Calculate Average Daily Comments
	timestamp, _ := strconv.ParseInt(theReport.PublishedAt, 10, 64)
	t := time.Unix(timestamp, 0)
	delta := time.Now().Sub(t)
	theReport.CommentAvgPerDay = float64(theReport.TotalComments) / (float64(delta.Hours()) / float64(24))

	reportJSON, err := json.Marshal(theReport)
	if err != nil {
		fmt.Println(err)
	}

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

	if !*WebServer && *PostURL == "" {
		fmt.Println("Post URL is required.")
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

		http.HandleFunc("/", webHandler)
		http.ListenAndServe(":"+*Port, nil)
	} else {
		fmt.Println(string(runReport(*PostURL)))
	}
}
