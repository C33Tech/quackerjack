package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	//"github.com/kr/pretty"
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
	Thumbnail     string
	PublishedAt   string
}

var igMedia IGMedia

func (ig *InstagramPic) FetchMedia(maxID string) IGMedia {
	var err error
	igMedia := IGMedia{}

	igClient := http.Client{
		Timeout: time.Second * 5, // Maximum of 2 secs
	}

	url := fmt.Sprintf("https://www.instagram.com/p/%v?__a=1&maxid=%v", ig.ShortCode, maxID)

	//fmt.Println(url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		LogMsg(err.Error())
		return igMedia
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 5.1; rv:7.0.1) Gecko/20100101 Firefox/7.0.1")
	res, err := igClient.Do(req)
	if err != nil {
		LogMsg(err.Error())
		return igMedia
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		LogMsg(err.Error())
		return igMedia
	}

	err = json.Unmarshal(body, &igMedia)
	if err != nil {
		LogMsg(err.Error())
		return igMedia
	}

	return igMedia
}

func (ig *InstagramPic) GetMetadata() bool {
	igMedia := ig.FetchMedia("")

	ig.ID = igMedia.Graphql.ShortcodeMedia.ID
	ig.ShortCode = igMedia.Graphql.ShortcodeMedia.Shortcode
	if igMedia.Graphql.ShortcodeMedia.IsVideo {
		ig.Type = "video"
	} else {
		ig.Type = "photo"
	}
	ig.Caption = igMedia.Graphql.ShortcodeMedia.EdgeMediaToCaption.Edges[0].Node.Text
	ig.TotalLikes = int64(igMedia.Graphql.ShortcodeMedia.EdgeMediaPreviewLike.Count)
	ig.UserID = igMedia.Graphql.ShortcodeMedia.Owner.ID
	ig.UserName = igMedia.Graphql.ShortcodeMedia.Owner.Username
	ig.TotalComments = uint64(igMedia.Graphql.ShortcodeMedia.EdgeMediaToComment.Count)
	ig.PublishedAt = string(igMedia.Graphql.ShortcodeMedia.TakenAtTimestamp)
	ig.Thumbnail = igMedia.Graphql.ShortcodeMedia.DisplayURL

	return true
}

func (ig InstagramPic) CommentLoop(comments []*Comment, maxID string) CommentList {
	igMedia := ig.FetchMedia(maxID)

	for _, c := range igMedia.Graphql.ShortcodeMedia.EdgeMediaToComment.Edges {
		thisComment := &Comment{
			ID:         c.Node.ID,
			Published:  string(c.Node.CreatedAt),
			Content:    c.Node.Text,
			AuthorName: c.Node.Owner.Username,
		}

		comments = append(comments, thisComment)
	}

	if igMedia.Graphql.ShortcodeMedia.EdgeMediaToComment.PageInfo.HasNextPage {
		//return ig.CommentLoop(comments, igMedia.Graphql.ShortcodeMedia.EdgeMediaToComment.PageInfo.EndCursor)
	}

	return CommentList{Comments: comments}
}

func (ig InstagramPic) GetComments() CommentList {
	return ig.CommentLoop([]*Comment{}, "")
}
