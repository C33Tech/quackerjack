package main

import (
	"context"
	"strconv"
	"time"

	option "google.golang.org/api/option"
	youtube "google.golang.org/api/youtube/v3"
)

// YouTubeVideo is a distilled record of YouTube video metadata
type YouTubeVideo struct {
	ID            string
	Title         string
	VideoViews    uint64
	ChannelID     string
	ChannelTitle  string
	TotalComments uint64
	TotalLikes    uint64
	TotalDislikes uint64
	Thumbnail     string
	PublishedAt   string
}

// YouTubeGetCommentsV2 pulls the comments for a given YouTube video
func (ytv YouTubeVideo) GetComments() CommentList {
	videoID := ytv.ID
	var comments = []*Comment{}

	ctx := context.Background()
	youtubeService, err := youtube.NewService(ctx, option.WithAPIKey(GetConfigString("ytkey")))
	if err != nil {
		panic(err)
	}

	pageToken := ""
	errCount := 0
	for pageToken != "EOL" {
		results, err := youtubeService.CommentThreads.List([]string{"id", "snippet", "replies"}).TextFormat("plainText").MaxResults(100).VideoId(videoID).PageToken(pageToken).Do()

		if err != nil {
			LogMsg(err.Error())

			if errCount > 3 {
				LogMsg("YouTube API error limit hit.")
				break
			}

			errCount = errCount + 1
			continue
		}

		if len(results.Items) > 0 {
			for _, item := range results.Items {

				tempComments := []*youtube.Comment{}
				tempComments = append(tempComments, item.Snippet.TopLevelComment)

				if item.Replies != nil {
					for _, reply := range item.Replies.Comments {
						tempComments = append(tempComments, reply)
					}
				}

				for _, c := range tempComments {
					thisComment := &Comment{
						ID:         c.Id,
						Published:  c.Snippet.PublishedAt,
						Title:      "",
						Content:    c.Snippet.TextDisplay,
						AuthorName: c.Snippet.AuthorDisplayName,
						Likes:      c.Snippet.LikeCount,
					}

					comments = append(comments, thisComment)
				}
			}
		}

		pageToken = results.NextPageToken
		if pageToken == "" {
			pageToken = "EOL"
		}
	}

	return CommentList{Comments: comments}
}

// GetMetadata returns a subset of video information from the YouTube API
func (ytv *YouTubeVideo) GetMetadata() bool {
	videoID := ytv.ID

	ctx := context.Background()
	youtubeService, err := youtube.NewService(ctx, option.WithAPIKey(GetConfigString("ytkey")))
	if err != nil {
		panic(err)
	}

	call := youtubeService.Videos.List([]string{"id", "snippet", "statistics"}).Id(videoID)
	resp, err := call.Do()
	if err != nil {
		panic(err)
	}

	if len(resp.Items) > 0 {
		video := resp.Items[0]

		t, _ := time.Parse(time.RFC3339Nano, video.Snippet.PublishedAt)

		ytv.Title = video.Snippet.Title
		ytv.ChannelID = video.Snippet.ChannelId
		ytv.ChannelTitle = video.Snippet.ChannelTitle
		ytv.TotalComments = video.Statistics.CommentCount
		ytv.PublishedAt = strconv.FormatInt(t.Unix(), 10)
		ytv.VideoViews = video.Statistics.ViewCount
		ytv.Thumbnail = video.Snippet.Thumbnails.High.Url
		ytv.TotalLikes = video.Statistics.LikeCount
		ytv.TotalDislikes = video.Statistics.DislikeCount

		return true
	}

	return false
}
