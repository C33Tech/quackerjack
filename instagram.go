package main

import (
	"net/url"
	"strings"

	//"github.com/kr/pretty"
	"github.com/mikeflynn/golang-instagram/instagram"
)

type InstagramPic struct {
	ID            string
	ShortCode     string
	Type          string
	Caption       string
	TotalLikes    int64
	UserID        string
	UserName      string
	TotalComments uint64
	PublishedAt   string
}

var instagramApi *instagram.Api
var instagramPostResponse *instagram.MediaResponse

func (ig *InstagramPic) GetMetadata() bool {
	if instagramApi == nil {
		instagramApi = instagram.New(*InstagramKey, "")
	}

	var resp *instagram.MediaResponse

	if ig.ShortCode != "" {
		resp, _ = instagramApi.GetMediaByShortcode(ig.ShortCode, url.Values{})
	} else if ig.ID != "" {
		resp, _ = instagramApi.GetMedia(ig.ID, url.Values{})
	} else {
		return false
	}

	if resp != (*instagram.MediaResponse)(nil) {
		ig.ID = resp.Media.Id
		parts := strings.Split(resp.Media.Link, "/")
		ig.ShortCode = parts[len(parts)-2]
		ig.Type = resp.Media.Type
		ig.Caption = resp.Media.Caption.Text
		ig.TotalLikes = resp.Media.Likes.Count
		ig.UserID = resp.Media.User.Id
		ig.UserName = resp.Media.User.Username
		ig.TotalComments = uint64(resp.Media.Comments.Count)
		ig.PublishedAt = string(resp.Media.CreatedTime)

		return true
	}

	return false
}

func (ig InstagramPic) GetComments() CommentList {
	if instagramApi == nil {
		instagramApi = instagram.New(*InstagramKey, "")
	}

	var resp *instagram.CommentsResponse
	var comments = []*Comment{}

	resp, _ = instagramApi.GetMediaComments(ig.ID, url.Values{})

	if resp != new(instagram.CommentsResponse) {
		for _, entry := range resp.Comments {
			thisComment := &Comment{
				ID:         entry.Id,
				Published:  string(entry.CreatedTime),
				Content:    entry.Text,
				AuthorName: entry.From.Username,
			}

			comments = append(comments, thisComment)
		}
	}

	return CommentList{Comments: comments}
}
