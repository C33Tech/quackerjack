package main

import (
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"sort"
	"strings"
	"time"

	emoji "github.com/tmdvs/Go-Emoji-Utils"
)

const MaxKeywords = 100

// Comment is the distilled comment dataset
type Comment struct {
	ID         string
	Published  string
	Title      string
	Content    string
	AuthorName string
	Sentiment  string
	Likes      int64
}

type CommentList struct {
	Comments []*Comment
}

// Comment methods

func (c *Comment) GetSentiment() (string, error) {
	if c.Sentiment == "" {
		if SentimentClassifier.totalWords > 0 {
			c.Sentiment = SentimentClassifier.Classify(c.CleanContent())
			LogMsg(fmt.Sprintf("%s ===> %s", c.Content, c.Sentiment))
		} else {
			LogMsg(fmt.Sprintf("Classifier not trained. Word count: %d", SentimentClassifier.totalWords))
		}
	}

	return c.Sentiment, nil
}

func (c *Comment) CleanContent() string {
	cleanedContent := c.Content

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

func (cl *Comment) GetEmoji() map[string]uint64 {
	ret := map[string]uint64{}

	res := emoji.FindAll(cl.Content)
	for _, x := range res {
		if _, ok := ret[x.Match.(emoji.Emoji).Value]; !ok {
			ret[x.Match.(emoji.Emoji).Value] = 0
		}

		ret[x.Match.(emoji.Emoji).Value]++
	}

	return ret
}

func (cl *Comment) GetPublishedDay() (string, error) {
	t, err := time.Parse("2006-01-02T15:04:05Z", cl.Published)
	if err != nil {
		return "", err
	}

	return t.Format("2006-01-02"), nil
}

// CommentList methods

func (cl *CommentList) IsEmpty() bool {
	return len(cl.Comments) == 0
}

func (cl *CommentList) GetTotal() uint64 {
	return uint64(len(cl.Comments))
}

func (cl *CommentList) GetKeywords() map[string]uint64 {
	idx := make(map[string]uint64)

	for _, comment := range cl.Comments {
		words := GetWords(comment.Content)

		for _, word := range words {
			if !IsStopWord(word) {
				idx[strings.ToLower(word)]++
			}
		}
	}

	sorted := SortedKeys(idx)

	max := MaxKeywords
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

func (cl *CommentList) GetEmojiCount() map[string]uint64 {
	emoji := map[string]uint64{}

	for _, comment := range cl.Comments {
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

func (cl *CommentList) GetSentimentSummary() map[string]uint64 {
	tags := map[string]uint64{}

	for _, comment := range cl.Comments {
		res, err := comment.GetSentiment()

		if err == nil {
			if res == "1" {
				tags["positive"]++
			} else if res == "0" {
				tags["negative"]++
			} else {
				tags["unknown"]++
			}
		}
	}

	return tags
}

func (cl *CommentList) GetDailySentiment() map[string]map[string]uint64 {
	days := map[string]map[string]uint64{}

	for _, comment := range cl.Comments {
		res, err := comment.GetSentiment()

		if err == nil {
			sentiment := "unknown"
			if res == "1" {
				sentiment = "positive"
			} else if res == "0" {
				sentiment = "negative"
			}

			date, err := comment.GetPublishedDay()
			if err == nil {
				if _, ok := days[date]; !ok {
					days[date] = map[string]uint64{
						"positive": 0,
						"negative": 0,
						"unknown":  0,
					}
				}

				days[date][sentiment]++
			} else {
				log.Println(err)
			}
		}
	}

	return days
}

func (cl *CommentList) GetRandom(count int) []*Comment {
	seed := rand.NewSource(42)
	rnum := rand.New(seed)

	resp := []*Comment{}

	for i := 0; i < count; i++ {
		resp = append(resp, cl.Comments[rnum.Intn(len(cl.Comments))])
	}

	return resp
}

type ByLikes []*Comment

func (a ByLikes) Len() int           { return len(a) }
func (a ByLikes) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLikes) Less(i, j int) bool { return a[i].Likes > a[j].Likes }

func (cl *CommentList) GetMostLiked(count int) []*Comment {
	bl := ByLikes(cl.Comments)
	sort.Sort(bl)

	resp := []*Comment{}

	for i := 0; i < count; i++ {
		resp = append(resp, bl[i])
	}

	return resp
}
