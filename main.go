package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math"
	"os"
	"strings"
)

// Flags
var VideoId = flag.String("video", "", "The YouTube video id.")
var OutputDir = flag.String("out", "", "The output file (default: std out).")
var YouTubeKey = flag.String("ytkey", "", "Google API key.")
var StopWordFiles = flag.String("stopwords", "", "A list of file paths, comma delimited, of stop word files.")
var Verbose = flag.Bool("verbose", false, "Extra logging to std out")
var TrainingFiles = flag.String("training", "", "Training text files.")
var RedisServer = flag.String("redis", "", "Redis server and port.")
var ConfigFile = flag.String("config", "", "Config file.")

func LogMsg(msg string) {
	if *Verbose {
		fmt.Printf(msg)
	}
}

type Report struct {
	TotalComments          uint64
	CollectedComments      int
	CommentCoveragePercent float64
	ChannelTitle           string
	VideoTitle             string
	Keywords               []string
	Sentiment              []SentimentTag
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

	if *VideoId == "" {
		fmt.Println("Video ID is required.")
		os.Exit(1)
	}

	if *StopWordFiles != "" {
		swFiles := strings.Split(*StopWordFiles, ",")
		for _, path := range swFiles {
			LoadStopWords(path)
		}
	}

	// Poll the data sources...
	metadata := GetVideoInfo(*VideoId)
	comments := GetComments_v2(*VideoId)

	// Create Report
	report := Report{}

	// Set video metadata
	report.TotalComments = metadata.TotalComments
	report.ChannelTitle = metadata.ChannelTitle
	report.VideoTitle = metadata.Title

	// Set comments returned
	report.CollectedComments = len(comments)
	report.CommentCoveragePercent = math.Ceil((float64(report.CollectedComments) / float64(report.TotalComments)) * float64(100))

	// Set Keywords
	report.Keywords = GetKeywords(comments)

	// Sentiment Tagging
	if *RedisServer != "" {
		report.Sentiment = GetSentimentSummary(comments)
	}

	reportJson, _ := json.Marshal(report)
	// Output Report
	fmt.Println(string(reportJson))
}
