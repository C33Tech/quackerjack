package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"strings"
	"time"

	//"github.com/kr/pretty"
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
	Type                   string
	Title                  string
	PublishedAt            string
	TotalComments          uint64
	CollectedComments      int
	CommentCoveragePercent float64
	CommentAvgPerDay       float64
	Keywords               []string
	Sentiment              []SentimentTag
	Metadata               Post
}

// Post is the interface for all the various post types (YouTubeVideo, etc...)
type Post interface {
	GetComments() []Comment
	GetMetadata() bool
}

// Comment is the distilled comment dataset
type Comment struct {
	ID         string
	Published  string
	Title      string
	Content    string
	AuthorName string
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
		postUrl := r.URL.Query().Get("vid")
		domain, pid := parseURL(postUrl)

		var jsonBytes []byte

		if pid != "" {
			jsonBytes = runReport(domain, pid)
		} else {
			jsonBytes, _ = json.Marshal(webError{Error: "Missing video id."})
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBytes)
	} else {
		http.ServeFile(w, r, "static/gui/index.html")
	}
}

func runReport(domain string, pid string) []byte {
	// Create Report
	theReport := report{}
	var thePost Post

	switch domain {
	case "youtube":
		thePost = &YouTubeVideo{ID: pid}
	case "instagram":
		thePost = &InstagramPic{ShortCode: pid}
	}

	done := make(chan bool)

	// Poll the data sources...
	go func() {
		_ = thePost.GetMetadata()
		theReport.Type = reflect.TypeOf(thePost).String()

		switch p := thePost.(type) {
		case *YouTubeVideo:
			theReport.ID = p.ID
			theReport.Title = p.Title
			theReport.PublishedAt = p.PublishedAt
			theReport.TotalComments = p.TotalComments
			theReport.Metadata = p
		case *InstagramPic:
			theReport.ID = p.ID
			theReport.Title = p.Caption
			theReport.PublishedAt = p.PublishedAt
			theReport.TotalComments = p.TotalComments
			theReport.Metadata = p
		}
		done <- true
	}()

	// Fetch the comments
	comments := thePost.GetComments()

	// If we don't get an comments back, wait for the metadata call to return and send an error.
	if len(comments) == 0 {
		<-done

		noCommentsError := "No comments found for this video."
		if theReport.Title == "" {
			noCommentsError = "Invalid YouTube video ID."
		}

		errorJSON, _ := json.Marshal(webError{Error: noCommentsError})
		return errorJSON
	}

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

	// Parse the PostURL to discover the domain and ID.
	postDomain, postID := parseURL(*PostURL)
	if postDomain == "" || postID == "" {
		fmt.Println("Unable to parse post url.")
		os.Exit(1)
	}

	// Check for required params and run
	if postDomain == "youtube" {
		if *YouTubeKey == "" {
			fmt.Println("A Google API key with YouTube API access is required.")
			os.Exit(1)
		}
	} else if postDomain == "instagram" {
		if *InstagramKey == "" {
			fmt.Println("An Instagram API key is required.")
			os.Exit(1)
		}
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
		fmt.Println(string(runReport(postDomain, postID)))
	}
}
