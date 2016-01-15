package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	//"io/ioutil"
)

type VineSession struct {
	SessionID string
}

var vineSession *VineSession = new(VineSession)

func InitVine() bool {
	if vineSession.SessionID == "" {
		vineSession, _ = authVine(*VineUsername, *VinePassword)
	}

	if vineSession.SessionID != "" {
		return true
	}

	return false
}

type VineVideo struct {
	ID            string // data.records[].postID
	ShortCode     string
	Caption       string // data.records[].description
	VideoViews    uint64 // data.records[].loops.count
	UserID        int64  // data.records[].userID
	UserName      string // data.records[].username
	TotalComments uint64 // data.records[].comments.count
	TotalLikes    uint64 // data.records[].likes.count
	Thumbnail     string // data.records[].thumbnailUrl
	PublishedAt   string // data.records[].created
}

func (this *VineVideo) GetMetadata() bool {
	if InitVine() == false {
		return false
	}

	resp, err := vineSession.vineRequest("/timelines/posts/s/" + this.ShortCode)
	if err == nil {
		defer resp.Body.Close()

		var respJson VinePostResp

		//contents, _ := ioutil.ReadAll(resp.Body)
		//fmt.Println(string(contents))

		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&respJson)

		if err == nil {
			this.ID = strconv.FormatInt(respJson.Data.Records[0].PostID, 10)
			this.Caption = respJson.Data.Records[0].Description
			this.VideoViews = uint64(respJson.Data.Records[0].Loops.Count)
			this.UserID = respJson.Data.Records[0].UserID
			this.UserName = respJson.Data.Records[0].Username
			this.TotalComments = respJson.Data.Records[0].Comments.Count
			this.TotalLikes = respJson.Data.Records[0].Likes.Count
			this.Thumbnail = respJson.Data.Records[0].ThumbnailUrl
			this.PublishedAt = respJson.Data.Records[0].Created

			return true
		}
	}

	return false
}

func (this VineVideo) GetComments() CommentList {
	var comments = []*Comment{}

	if InitVine() {
		anchor := ""
		page := 1

		for {
			var respJson VineCommentResp
			resp, err := vineSession.vineRequest("/posts/" + this.ID + "/comments?size=100&page=" + strconv.Itoa(page) + "&anchor=" + anchor)

			if err != nil {
				break
			}

			defer resp.Body.Close()

			decoder := json.NewDecoder(resp.Body)
			err = decoder.Decode(&respJson)

			if err == nil {
				for _, entry := range respJson.Data.Records {
					thisComment := &Comment{
						ID:         strconv.FormatUint(entry.CommentID, 10),
						Published:  entry.Created,
						Content:    entry.Comment,
						AuthorName: entry.Username,
					}

					comments = append(comments, thisComment)
				}

				if respJson.Data.Count <= len(comments) {
					break
				} else {
					if respJson.Data.AnchorStr != "" {
						anchor = respJson.Data.AnchorStr
					} else {
						break
					}

					page = page + 1
				}
			}
		}
	}

	return CommentList{Comments: comments}
}

func authVine(username string, password string) (*VineSession, error) {
	data := url.Values{}
	data.Set("username", username)
	data.Set("password", password)

	r, _ := http.NewRequest("POST", "https://vine.co/api/users/authenticate", bytes.NewBufferString(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	client := &http.Client{}
	resp, _ := client.Do(r)
	if resp.Status == "200 OK" {
		var respJson VineAuthResp
		defer resp.Body.Close()

		decoder := json.NewDecoder(resp.Body)
		decoder.Decode(&respJson)

		return &VineSession{SessionID: respJson.Data.Key}, nil
	} else {
		return &VineSession{}, errors.New("Auth failed.")
	}
}

func (this *VineSession) vineRequest(path string) (*http.Response, error) {
	u, err := url.Parse(path)
	if err != nil {
		return nil, errors.New("Vine request path invalid.")
	}

	u.Scheme = "https"
	u.Host = "vine.co/api"

	fmt.Println(u.String())

	r, _ := http.NewRequest("GET", u.String(), nil)
	r.Header.Add("vine-session-id", this.SessionID)
	client := &http.Client{}
	resp, err := client.Do(r)

	if resp.Status != "200 OK" {
		return nil, err
	} else {
		return resp, nil
	}
}

type VinePostResp struct {
	Code string `json:"code"`
	Data struct {
		Count     int    `json:"count"`
		Anchorstr string `json:"anchorStr"`
		Records   []struct {
			Liked             int         `json:"liked"`
			VideodashUrl      string      `json:"videoDashUrl"`
			Foursquarevenueid string      `json:"foursquareVenueId"`
			UserID            int64       `json:"userId"`
			Private           int         `json:"private"`
			VideowebmUrl      interface{} `json:"videoWebmUrl"`
			Loops             struct {
				Count    float64 `json:"count"`
				Velocity float64 `json:"velocity"`
				Onfire   int     `json:"onFire"`
			} `json:"loops"`
			ThumbnailUrl    string `json:"thumbnailUrl"`
			Explicitcontent int    `json:"explicitContent"`
			Blocked         int    `json:"blocked"`
			Verified        int    `json:"verified"`
			AvatarUrl       string `json:"avatarUrl"`
			Comments        struct {
				Count     uint64 `json:"count"`
				Anchorstr string `json:"anchorStr"`
				Records   []struct {
					Username          string        `json:"username"`
					Comment           string        `json:"comment"`
					Verified          int           `json:"verified"`
					VanityUrls        []interface{} `json:"vanityUrls"`
					Created           string        `json:"created"`
					Userid            int64         `json:"userId"`
					Profilebackground string        `json:"profileBackground"`
					Entities          []interface{} `json:"entities"`
					User              struct {
						Username          string        `json:"username"`
						Verified          int           `json:"verified"`
						Description       string        `json:"description"`
						Twitterverified   int           `json:"twitterVerified"`
						Avatarurl         string        `json:"avatarUrl"`
						Notporn           int           `json:"notPorn"`
						Userid            int64         `json:"userId"`
						Profilebackground string        `json:"profileBackground"`
						Hidefrompopular   int           `json:"hideFromPopular"`
						Private           int           `json:"private"`
						Location          interface{}   `json:"location"`
						Unflaggable       int           `json:"unflaggable"`
						Vanityurls        []interface{} `json:"vanityUrls"`
					} `json:"user"`
					Commentid int64 `json:"commentId"`
					Postid    int64 `json:"postId"`
				} `json:"records"`
				Previouspage interface{} `json:"previousPage"`
				Backanchor   string      `json:"backAnchor"`
				Anchor       int64       `json:"anchor"`
				Nextpage     int         `json:"nextPage"`
				Size         int         `json:"size"`
			} `json:"comments"`
			Entities []struct {
				Link  string `json:"link"`
				Range []int  `json:"range"`
				Type  string `json:"type"`
				ID    int64  `json:"id"`
				Title string `json:"title"`
			} `json:"entities"`
			Videolowurl       string        `json:"videoLowURL"`
			Vanityurls        []string      `json:"vanityUrls"`
			Username          string        `json:"username"`
			Description       string        `json:"description"`
			Tags              []interface{} `json:"tags"`
			Permalinkurl      string        `json:"permalinkUrl"`
			Promoted          int           `json:"promoted"`
			PostID            int64         `json:"postId"`
			Profilebackground string        `json:"profileBackground"`
			Videourl          string        `json:"videoUrl"`
			Followrequested   int           `json:"followRequested"`
			Created           string        `json:"created"`
			Hassimilarposts   int           `json:"hasSimilarPosts"`
			Shareurl          string        `json:"shareUrl"`
			Myrepostid        int           `json:"myRepostId"`
			Following         int           `json:"following"`
			Reposts           struct {
				Count        int           `json:"count"`
				Anchorstr    string        `json:"anchorStr"`
				Records      []interface{} `json:"records"`
				Previouspage interface{}   `json:"previousPage"`
				Backanchor   string        `json:"backAnchor"`
				Anchor       interface{}   `json:"anchor"`
				Nextpage     interface{}   `json:"nextPage"`
				Size         int           `json:"size"`
			} `json:"reposts"`
			Likes struct {
				Count     uint64 `json:"count"`
				Anchorstr string `json:"anchorStr"`
				Records   []struct {
					Username   string        `json:"username"`
					Verified   int           `json:"verified"`
					Vanityurls []interface{} `json:"vanityUrls"`
					Created    string        `json:"created"`
					Userid     int64         `json:"userId"`
					User       struct {
						Private int `json:"private"`
					} `json:"user"`
					Likeid int64 `json:"likeId"`
				} `json:"records"`
				Previouspage interface{} `json:"previousPage"`
				Backanchor   string      `json:"backAnchor"`
				Anchor       int64       `json:"anchor"`
				Nextpage     int         `json:"nextPage"`
				Size         int         `json:"size"`
			} `json:"likes"`
		} `json:"records"`
		Previouspage interface{} `json:"previousPage"`
		Backanchor   string      `json:"backAnchor"`
		Anchor       interface{} `json:"anchor"`
		Nextpage     interface{} `json:"nextPage"`
		Size         int         `json:"size"`
	} `json:"data"`
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type VineAuthResp struct {
	Code string `json:"code"`
	Data struct {
		Username    string      `json:"username"`
		AvatarURL   string      `json:"avatarUrl"`
		clientFlags interface{} `json:"clientFlags"`
		UserID      int64       `json:"userId"`
		Edition     string      `json:"edition"`
		Key         string      `json:"key"`
	} `json:"data"`
}

type VineCommentResp struct {
	Code string `json:"code"`
	Data struct {
		Anchor       int    `json:"anchor"`
		AnchorStr    string `json:"anchorStr"`
		BackAnchor   string `json:"backAnchor"`
		Count        int    `json:"count"`
		NextPage     int    `json:"nextPage"`
		PreviousPage int    `json:"previousPage"`
		Records      []struct {
			AvatarURL    string `json:"avatarUrl"`
			Comment      string `json:"comment"`
			CommentID    uint64 `json:"commentId"`
			CommentIDStr string `json:"commentIdStr"`
			Created      string `json:"created"`
			Entities     []struct {
				ID         int           `json:"id"`
				IDStr      string        `json:"idStr"`
				Link       string        `json:"link"`
				Range      []int         `json:"range"`
				Title      string        `json:"title"`
				Type       string        `json:"type"`
				VanityUrls []interface{} `json:"vanityUrls"`
			} `json:"entities"`
			Flags_platformHi int         `json:"flags|platform_hi"`
			Flags_platformLo int         `json:"flags|platform_lo"`
			Location         string      `json:"location"`
			PostID           int         `json:"postId"`
			PostIDStr        string      `json:"postIdStr"`
			SourceID         int         `json:"sourceId"`
			SourceIDStr      string      `json:"sourceIdStr"`
			SourceType       interface{} `json:"sourceType"`
			TwitterVerified  int         `json:"twitterVerified"`
			User             struct {
				AvatarURL         string        `json:"avatarUrl"`
				Description       string        `json:"description"`
				Location          string        `json:"location"`
				Private           int           `json:"private"`
				ProfileBackground string        `json:"profileBackground"`
				TwitterVerified   int           `json:"twitterVerified"`
				UserID            int           `json:"userId"`
				UserIDStr         string        `json:"userIdStr"`
				Username          string        `json:"username"`
				VanityUrls        []interface{} `json:"vanityUrls"`
				Verified          int           `json:"verified"`
			} `json:"user"`
			UserID     int           `json:"userId"`
			UserIDStr  string        `json:"userIdStr"`
			Username   string        `json:"username"`
			VanityUrls []interface{} `json:"vanityUrls"`
			Verified   int           `json:"verified"`
		} `json:"records"`
		Size int `json:"size"`
	} `json:"data"`
	Error   string `json:"error"`
	Success bool   `json:"success"`
}
