package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

// Flags
var video_id = flag.String("video", "", "The YouTube video id.")
var output = flag.String("out", "", "The output file (default: std out).")
var stopWordFiles = flag.String("stopwords", "", "A list of file paths, comma delimited, of stop word files.")
var verbose = flag.Bool("verbose", false, "")

func log(msg string) {
	if *verbose {
		fmt.Printf(msg)
	}
}

var YouTubeKey = os.Getenv("YOUTUBE_KEY")

type Tag struct {
	Name    string
	Percent float32
}

type Report struct {
	TotalComments     uint64
	CollectedComments int
	ChannelTitle      string
	VideoTitle        string
	Keywords          []string
	Sentiment         []Tag
}

func main() {
	flag.Parse()

	if YouTubeKey == "" {
		fmt.Println("Missing YOUTUBE_KEY environment variable.")
		return
	}

	if *video_id == "" {
		fmt.Println("Video ID is required.")
		return
	}

	if *stopWordFiles != "" {
		swFiles := strings.Split(*stopWordFiles, ",")
		for _, path := range swFiles {
			LoadStopWords(path)
		}
	}

	// Poll the data sources...
	metadata := GetVideoInfo(*video_id)
	comments := GetComments_v2(*video_id)

	// Create Report
	report := Report{}

	// Set video metadata
	report.TotalComments = metadata.TotalComments
	report.ChannelTitle = metadata.ChannelTitle
	report.VideoTitle = metadata.Title

	// Set comments returned
	report.CollectedComments = len(comments)

	// Set Keywords
	report.Keywords = GetKeywords(comments)

	// Sentiment Tagging
	// ...

	reportJson, _ := json.Marshal(report)
	// Output Report
	fmt.Println(string(reportJson))
}
