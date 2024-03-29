package main

import (
	"encoding/json"
	"errors"

	//"fmt"
	"net/http"
	"net/url"
	//"github.com/kr/pretty"
)

type FacebookPost struct {
	ID            string // id
	PageName      string
	PageID        string
	Type          string // type
	Caption       string // message / description
	TotalLikes    uint64 // likes.summary.total_count
	TotalComments uint64 // comments.summary.total_count
	Thumbnail     string // picture
	PublishedAt   string // created_time
}

func (fp *FacebookPost) GetMetadata() bool {
	fp.GetPageID()

	var respTyped postMetaResp
	resp, _ := fbRequest("/" + fp.PageID + "_" + fp.ID + "?fields=id,name,caption,description,picture,created_time,type,message,properties,insights,likes.limit(1).summary(true),comments.limit(1).summary(true)")

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err := decoder.Decode(&respTyped)
	if err != nil {
		return false
	}

	fp.Type = respTyped.Type
	fp.Caption = respTyped.Message
	fp.TotalLikes = respTyped.Likes.Summery.TotalCount
	fp.TotalComments = respTyped.Comments.Summary.TotalCount
	fp.Thumbnail = respTyped.Picture
	fp.PublishedAt = respTyped.CreatedTime

	return true
}

func (fp FacebookPost) GetComments() CommentList {
	fp.GetPageID()

	var comments = []*Comment{}
	after := ""
	max := 10000

	for {
		if len(comments) >= max {
			break
		}

		var respTyped postCommentListResp
		resp, _ := fbRequest("/" + fp.PageID + "_" + fp.ID + "/comments?limit=100&order=reverse_chronological&after=" + after)

		//fmt.Println("/" + this.PageID + "_" + this.ID + "/comments?limit=100&order=reverse_chronological&after=" + after)

		defer resp.Body.Close()

		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&respTyped)

		//LogMsg(fmt.Sprintf("Total comments: %d", len(comments)))

		if err == nil {
			for _, entry := range respTyped.Data {
				thisComment := &Comment{
					ID:         entry.ID,
					Published:  entry.CreatedOn,
					Content:    entry.Message,
					AuthorName: entry.From.Name,
				}

				comments = append(comments, thisComment)
			}

			if respTyped.Pagination.Cursors.After != "" {
				after = respTyped.Pagination.Cursors.After
			} else {
				break
			}
		}
	}

	return CommentList{Comments: comments}
}

func (fp *FacebookPost) GetPageID() *FacebookPost {
	if fp.PageID != "" {
		return fp
	}

	var respTyped pageNameResp
	resp, _ := fbRequest("/" + fp.PageName)

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err := decoder.Decode(&respTyped)
	if err != nil {
		return fp
	}

	fp.PageID = respTyped.ID

	return fp
}

func fbRequest(path string) (*http.Response, error) {
	u, err := url.Parse(path)
	if err != nil {
		return nil, errors.New("FB request path invalid")
	}

	u.Scheme = "https"
	u.Host = "graph.facebook.com"

	query := u.Query()
	query.Add("access_token", GetConfigString("fbkey")+"|"+GetConfigString("fbsecret"))
	u.RawQuery = query.Encode()

	//LogMsg(u.String())

	response, err := http.Get(u.String())
	if err != nil {
		return nil, err
	} else {
		return response, nil
	}
}

type pageNameResp struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type postCommentResp struct {
	From struct {
		Name string `json:"name"`
		ID   string `json:"id"`
	} `json:"from"`
	Message   string `json:"message"`
	CreatedOn string `json:"created_time"`
	ID        string `json:"id"`
}

type postCommentListResp struct {
	Data       []postCommentResp `json:"data"`
	Pagination struct {
		Cursors struct {
			After  string `json:"after,omitempty"`
			Before string `json:"before,omitempty"`
		} `json:"cursors"`
		Next string `json:"next,omitempty"`
	} `json:"paging"`
}

type postMetaProps struct {
	Name string `json:"name"`
	Text string `json:"text"`
}

type likesSummary struct {
	TotalCount uint64 `json:"total_count"`
	CanLike    bool   `json:"can_like"`
	HasLiked   bool   `json:"has_liked"`
}

type commentSummary struct {
	Order      string `json:"order"`
	TotalCount uint64 `json:"total_count"`
	CanComment bool   `json:"can_comment"`
}

type postMetaResp struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Picture     string          `json:"picture"`
	CreatedTime string          `json:"created_time"`
	Type        string          `json:"type"`
	Message     string          `json:"message"`
	Properties  []postMetaProps `json:"properties"`
	Likes       struct {
		Data []struct {
			ID string `json:"id"`
		} `json:"data"`
		Paging struct {
			Cursors struct {
				After  string `json:"after"`
				Before string `json:"before"`
			}
			Next string `json:"next"`
		} `json:"paging"`
		Summery likesSummary `json:"summary"`
	} `json:"likes"`
	Comments struct {
		Data   []postCommentResp
		Paging struct {
			Cursors struct {
				After  string `json:"after"`
				Before string `json:"before"`
			}
			Next string `json:"next"`
		} `json:"paging"`
		Summary commentSummary `json:"summary"`
	}
}
