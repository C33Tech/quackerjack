package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/mikeflynn/sentiment"
	emoji "github.com/tmdvs/Go-Emoji-Utils"
)

// Comment is the distilled comment dataset
type Comment struct {
	ID         string
	Published  string
	Title      string
	Content    string
	AuthorName string
	Sentiment  *sentiment.Analysis
	Likes      int64
}

type CommentList struct {
	Comments []*Comment
}

var sentimentModel sentiment.Models

// Comment methods

func (this *Comment) GetSentiment() (*sentiment.Analysis, error) {
	if this.Sentiment == nil {
		if sentimentModel == nil {
			LogMsg("Init sentiment model...")
			var err error
			sentimentModel, err = sentiment.Restore()
			if err != nil {
				LogMsg(fmt.Sprintf("%v", err))
				return &sentiment.Analysis{}, err
			}
		}

		this.Sentiment = sentimentModel.SentimentAnalysis(this.CleanContent(), sentiment.English)
		LogMsg(fmt.Sprintf("%s ===> %d", this.Content, this.Sentiment.Score))
	}

	return this.Sentiment, nil
}

func (this *Comment) CleanContent() string {
	cleanedContent := this.Content

	// Replace emoji with descriptor.

	res := emoji.FindAll(cleanedContent)
	for _, x := range res {
		cleanedContent = strings.ReplaceAll(cleanedContent, x.Match.(emoji.Emoji).Value, x.Match.(emoji.Emoji).Descriptor)
	}

	// Lowercase the string
	cleanedContent = strings.ToLower(cleanedContent)

	// Strip out stopwords
	words := GetWords(cleanedContent)

	for _, word := range words {
		if IsStopWord(word) {
			re := regexp.MustCompile(`\b` + word + `\b`)
			cleanedContent = re.ReplaceAllString(cleanedContent, "")
		}
	}

	// Clean multiple spaces
	cleanedContent = regexp.MustCompile(`\s+|,|\.+`).ReplaceAllString(cleanedContent, " ")

	return cleanedContent
}

func (this *Comment) GetEmoji() map[string]uint64 {
	ret := map[string]uint64{}

	res := emoji.FindAll(this.Content)
	for _, x := range res {
		if _, ok := ret[x.Match.(emoji.Emoji).Value]; !ok {
			ret[x.Match.(emoji.Emoji).Value] = 0
		}

		ret[x.Match.(emoji.Emoji).Value]++
	}

	return ret
}

func (this *Comment) GetPublishedDay() (string, error) {
	t, err := time.Parse("2006-01-02T15:04:05.000Z", this.Published)
	if err != nil {
		return "", err
	}

	return t.Format("2006-01-02"), nil
}

// CommentList methods

func (this *CommentList) IsEmpty() bool {
	if len(this.Comments) == 0 {
		return true
	}

	return false
}

func (this *CommentList) GetTotal() uint64 {
	return uint64(len(this.Comments))
}

func (this *CommentList) GetKeywords() map[string]uint64 {
	idx := make(map[string]uint64)

	for _, comment := range this.Comments {
		words := GetWords(comment.Content)

		for _, word := range words {
			if !IsStopWord(word) {
				idx[strings.ToLower(word)]++
			}
		}
	}

	sorted := SortedKeys(idx)

	max := 50
	if len(sorted) < max {
		max = len(sorted)
	}

	limited := sorted[:max]

	ret := make(map[string]uint64)
	for _, w := range limited {
		ret[w] = idx[w]
	}

	return ret
}

func (this *CommentList) GetEmojiCount() map[string]uint64 {
	emoji := map[string]uint64{}

	for _, comment := range this.Comments {
		res := comment.GetEmoji()

		for k, v := range res {
			if _, ok := emoji[k]; !ok {
				emoji[k] = 0
			}

			emoji[k] = emoji[k] + v
		}
	}

	sorted := SortedKeys(emoji)

	max := 50
	if len(sorted) < max {
		max = len(sorted)
	}

	limited := sorted[:max]

	ret := make(map[string]uint64)
	for _, w := range limited {
		ret[w] = emoji[w]
	}

	return ret
}

func (this *CommentList) GetSentimentSummary() map[string]uint64 {
	tags := map[string]uint64{}

	for _, comment := range this.Comments {
		res, err := comment.GetSentiment()

		if err == nil {
			if res.Score == 1 {
				tags["positive"]++
			} else {
				tags["negative"]++
			}
		}
	}

	return tags
}

func (this *CommentList) GetDailySentiment() map[string]map[string]uint64 {
	days := map[string]map[string]uint64{}

	for _, comment := range this.Comments {
		res, err := comment.GetSentiment()

		if err == nil {
			sentiment := "negative"
			if res.Score == 1 {
				sentiment = "positive"
			}

			date, err := comment.GetPublishedDay()
			if err == nil {
				if _, ok := days[date]; !ok {
					days[date] = map[string]uint64{
						"positive": 0,
						"negative": 0,
					}
				}

				days[date][sentiment]++
			}
		}
	}

	return days
}

func (this *CommentList) GetRandom(count int) []*Comment {
	seed := rand.NewSource(42)
	rnum := rand.New(seed)

	resp := []*Comment{}

	for i := 0; i < count; i++ {
		resp = append(resp, this.Comments[rnum.Intn(len(this.Comments))])
	}

	return resp
}

type ByLikes []*Comment

func (a ByLikes) Len() int           { return len(a) }
func (a ByLikes) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLikes) Less(i, j int) bool { return a[i].Likes > a[j].Likes }

func (this *CommentList) GetMostLiked(count int) []*Comment {
	bl := ByLikes(this.Comments)
	sort.Sort(bl)

	resp := []*Comment{}

	for i := 0; i < count; i++ {
		resp = append(resp, bl[i])
	}

	return resp
}
