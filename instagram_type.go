package main

type IGMedia struct {
	Graphql struct {
		ShortcodeMedia struct {
			Typename             string `json:"__typename"`
			AccessibilityCaption string `json:"accessibility_caption"`
			CaptionIsEdited      bool   `json:"caption_is_edited"`
			CommentsDisabled     bool   `json:"comments_disabled"`
			Dimensions           struct {
				Height int `json:"height"`
				Width  int `json:"width"`
			} `json:"dimensions"`
			DisplayResources []struct {
				ConfigHeight int    `json:"config_height"`
				ConfigWidth  int    `json:"config_width"`
				Src          string `json:"src"`
			} `json:"display_resources"`
			DisplayURL           string `json:"display_url"`
			EdgeMediaPreviewLike struct {
				Count int           `json:"count"`
				Edges []interface{} `json:"edges"`
			} `json:"edge_media_preview_like"`
			EdgeMediaToCaption struct {
				Edges []struct {
					Node struct {
						Text string `json:"text"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"edge_media_to_caption"`
			EdgeMediaToComment struct {
				Count int `json:"count"`
				Edges []struct {
					Node struct {
						CreatedAt       int  `json:"created_at"`
						DidReportAsSpam bool `json:"did_report_as_spam"`
						EdgeLikedBy     struct {
							Count int `json:"count"`
						} `json:"edge_liked_by"`
						ID    string `json:"id"`
						Owner struct {
							ID            string `json:"id"`
							IsVerified    bool   `json:"is_verified"`
							ProfilePicURL string `json:"profile_pic_url"`
							Username      string `json:"username"`
						} `json:"owner"`
						Text           string `json:"text"`
						ViewerHasLiked bool   `json:"viewer_has_liked"`
					} `json:"node"`
				} `json:"edges"`
				PageInfo struct {
					EndCursor   string `json:"end_cursor"`
					HasNextPage bool   `json:"has_next_page"`
				} `json:"page_info"`
			} `json:"edge_media_to_comment"`
			EdgeMediaToSponsorUser struct {
				Edges []interface{} `json:"edges"`
			} `json:"edge_media_to_sponsor_user"`
			EdgeMediaToTaggedUser struct {
				Edges []struct {
					Node struct {
						User struct {
							FullName      string `json:"full_name"`
							ID            string `json:"id"`
							IsVerified    bool   `json:"is_verified"`
							ProfilePicURL string `json:"profile_pic_url"`
							Username      string `json:"username"`
						} `json:"user"`
						X float64 `json:"x"`
						Y float64 `json:"y"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"edge_media_to_tagged_user"`
			EdgeWebMediaToRelatedMedia struct {
				Edges []interface{} `json:"edges"`
			} `json:"edge_web_media_to_related_media"`
			GatingInfo        interface{} `json:"gating_info"`
			HasRankedComments bool        `json:"has_ranked_comments"`
			ID                string      `json:"id"`
			IsAd              bool        `json:"is_ad"`
			IsVideo           bool        `json:"is_video"`
			Location          interface{} `json:"location"`
			MediaPreview      string      `json:"media_preview"`
			Owner             struct {
				BlockedByViewer   bool   `json:"blocked_by_viewer"`
				FollowedByViewer  bool   `json:"followed_by_viewer"`
				FullName          string `json:"full_name"`
				HasBlockedViewer  bool   `json:"has_blocked_viewer"`
				ID                string `json:"id"`
				IsPrivate         bool   `json:"is_private"`
				IsUnpublished     bool   `json:"is_unpublished"`
				IsVerified        bool   `json:"is_verified"`
				ProfilePicURL     string `json:"profile_pic_url"`
				RequestedByViewer bool   `json:"requested_by_viewer"`
				Username          string `json:"username"`
			} `json:"owner"`
			Shortcode                  string `json:"shortcode"`
			ShouldLogClientEvent       bool   `json:"should_log_client_event"`
			TakenAtTimestamp           int    `json:"taken_at_timestamp"`
			TrackingToken              string `json:"tracking_token"`
			ViewerCanReshare           bool   `json:"viewer_can_reshare"`
			ViewerHasLiked             bool   `json:"viewer_has_liked"`
			ViewerHasSaved             bool   `json:"viewer_has_saved"`
			ViewerHasSavedToCollection bool   `json:"viewer_has_saved_to_collection"`
			ViewerInPhotoOfYou         bool   `json:"viewer_in_photo_of_you"`
		} `json:"shortcode_media"`
	} `json:"graphql"`
}
