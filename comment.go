package main

import (
	"math"
	"math/rand"
	"sort"
)

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

func (this *Comment) GetSentiment() string {
	if this.Sentiment == "" {
		this.Sentiment = GetSentiment(this.Content)
	}

	return this.Sentiment
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

func (this *CommentList) GetKeywords() map[string]int {
	return GetKeywords(this.Comments)
}

func (this *CommentList) GetSentimentSummary() []SentimentTag {
	tags := map[string]int{}

	for _, comment := range this.Comments {
		tag := comment.GetSentiment()
		tags[tag]++
	}

	summary := []SentimentTag{}

	for tag, count := range tags {
		st := SentimentTag{
			Name:    tag,
			Percent: math.Ceil((float64(count) / float64(len(this.Comments))) * float64(100)),
		}

		summary = append(summary, st)
	}

	return summary
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
