package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// LogMsg takes a message and pipes it to stdout if the verbose flag is set.
func LogMsg(msg string) {
	if GetConfigBool("verbose") {
		log.Print(msg)
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
	Keywords               map[string]uint64
	Sentiment              map[string]uint64
	Metadata               Post
	SampleComments         []*Comment
	TopComments            []*Comment
	DailySentiment         map[string]map[string]uint64
	EmojiCount             map[string]uint64
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

func parseURL(url string) (string, []string, string) {
	sites := map[string]map[string]string{
		"instagram": {
			"default": "instag\\.?ram(\\.com)?/p/([\\w\\-]*)/?",
		},
		"youtube": {
			"default": "youtu\\.?be(\\.?com)?/(watch\\?v=)?([\\w\\-_]*)",
		},
		"facebook": {
			"default": "facebook\\.com/([\\w]*)/posts/([\\d]*)/?",
			"videos":  "facebook\\.com/([\\w]+)/videos/\\w{2}\\.([\\d]+)/([\\d]*)/?",
		},
	}

	var domain string
	var matches []string
	var format string

	for d, rgxs := range sites {
		for f, rstr := range rgxs {
			r, _ := regexp.Compile(rstr)
			matches = r.FindStringSubmatch(url)
			if len(matches) > 0 {
				domain = d
				format = f
				break
			}
		}

		if domain != "" {
			break
		}
	}

	return domain, matches, format
}

func verifyToken(token string) bool {
	// Custom token verification logic.

	return true
}

func webHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path[1:] == "api" {
		postURL := r.URL.Query().Get("vid")
		userToken := r.URL.Query().Get("token")

		var jsonBytes []byte

		if verifyToken(userToken) {
			if postURL != "" {
				jsonBytes = runReport(postURL)
			} else {
				jsonBytes, _ = json.Marshal(webError{Error: "Missing post URL."})
			}
		} else {
			jsonBytes, _ = json.Marshal(webError{Error: "Invalid token."})
		}

		// w.Header().Set("Access-Control-Allow-Origin", "*") /// USEFUL FOR DEV ONLY
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBytes)
	} else if r.URL.Path[1:] == "oauth" {
		// Custom oauth link generation logic.

	} else if r.URL.Path[1:] == "login" {
		// Custom login logic

	} else {
		data, err := Asset("static/gui/index.html")
		if err != nil {
			LogMsg("Web GUI asset not found!")
			os.Exit(1)
		}

		htmlPath := GetConfigString("html")
		if htmlPath != "" {
			data, err = ioutil.ReadFile(htmlPath)
			if err != nil {
				LogMsg("Unable to read html file path.")
				os.Exit(1)
			}
		}

		w.Header().Set("Content-Type", "text/html")
		w.Write(data)
	}
}

func runReport(postURL string) []byte {
	// Parse URL
	domain, urlParts, urlFormat := parseURL(postURL)
	if domain == "" || len(urlParts) == 0 {
		return jsonError("Unable to parse post url.")
	}

	// Create Report
	theReport := report{URL: postURL}
	var thePost Post

	switch domain {
	case "youtube":
		if GetConfigString("ytkey") == "" {
			return jsonError("API key for YouTube not configured.")
		}

		thePost = &YouTubeVideo{ID: urlParts[len(urlParts)-1]}
	case "instagram":
		thePost = &InstagramPic{ShortCode: urlParts[len(urlParts)-1]}
	case "facebook":
		if GetConfigString("fbkey") == "" || GetConfigString("fbsecret") == "" {
			return jsonError("Missing Facebook API credentials.")
		}

		switch urlFormat {
		case "default":
			thePost = &FacebookPost{PageName: urlParts[len(urlParts)-2], ID: urlParts[len(urlParts)-1]}
		case "videos":
			thePost = &FacebookPost{PageName: urlParts[len(urlParts)-3], PageID: urlParts[len(urlParts)-2], ID: urlParts[len(urlParts)-1]}
		}
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
	case *FacebookPost:
		theReport.Type = "FacebookPost"
		theReport.ID = p.ID
		//theReport.Title = p.Title
		theReport.PublishedAt = p.PublishedAt
		theReport.TotalComments = p.TotalComments
		theReport.Metadata = p
	}

	// Fetch the comments
	LogMsg("Fetching the comments...")
	comments := thePost.GetComments()

	// If we don't get an comments back, wait for the metadata call to return and send an error.
	if comments.IsEmpty() {
		return jsonError("No comments found for this post.")
	} else {
		LogMsg(fmt.Sprintf("Collected %d comments", comments.GetTotal()))
	}

	// Set comments returned
	theReport.CollectedComments = comments.GetTotal()
	theReport.CommentCoveragePercent = math.Ceil((float64(theReport.CollectedComments) / float64(theReport.TotalComments)) * float64(100))

	done := make(chan bool)

	// Count Keywords
	go func() {
		theReport.Keywords = comments.GetKeywords()

		done <- true
	}()

	// Count Emoji
	go func() {
		theReport.EmojiCount = comments.GetEmojiCount()

		done <- true
	}()

	// Sentiment Tagging
	go func() {
		LogMsg("Starting sentiment analysis...")
		theReport.Sentiment = comments.GetSentimentSummary()
		theReport.DailySentiment = comments.GetDailySentiment()
		done <- true
	}()

	// Wait for everything to finish up
	for i := 0; i < 3; i++ {
		<-done
	}

	// Pull a few sample comments
	theReport.SampleComments = comments.GetRandom(3)

	// Pull the top liked comments
	if theReport.Type == "YouTubeVideo" {
		theReport.TopComments = comments.GetMostLiked(3)
	}

	// Calculate Average Daily Comments
	timestamp, _ := strconv.ParseInt(theReport.PublishedAt, 10, 64)
	t := time.Unix(timestamp, 0)
	delta := time.Now().Sub(t)
	theReport.CommentAvgPerDay = float64(theReport.TotalComments) / (float64(delta.Hours()) / float64(24))

	reportJSON, err := json.Marshal(theReport)
	if err != nil {
		LogMsg(err.Error())
	}

	// Output Report
	return reportJSON
}

func main() {
	LoadConfig()

	// Train the classifier
	InitClassifier()

	if !GetConfigBool("server") && GetConfigString("post") == "" {
		LogMsg("Post URL is required.")
		os.Exit(1)
	}

	if GetConfigString("stopwords") != "" {
		swFiles := strings.Split(GetConfigString("stopwords"), ",")
		for _, path := range swFiles {
			LoadStopWords(path)
		}
	}

	if GetConfigBool("server") {
		LogMsg("Web server running on " + GetConfigString("port"))

		http.HandleFunc("/", webHandler)
		http.ListenAndServe(":"+GetConfigString("port"), nil)
	} else {
		LogMsg(string(runReport(GetConfigString("post"))))
	}
}
