package main

import "testing"

func TestParseURL(t *testing.T) {
	cases := []struct {
		name   string
		url    string
		domain string
		id     string
	}{
		{
			name:   "youtube watch",
			url:    "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
			domain: "youtube",
			id:     "dQw4w9WgXcQ",
		},
		{
			name:   "youtube short",
			url:    "https://youtu.be/dQw4w9WgXcQ",
			domain: "youtube",
			id:     "dQw4w9WgXcQ",
		},
		{
			name:   "instagram",
			url:    "https://www.instagram.com/p/CI6lO1FJMfi/",
			domain: "instagram",
			id:     "CI6lO1FJMfi",
		},
		{
			name:   "facebook post",
			url:    "https://www.facebook.com/testpage/posts/1234567890",
			domain: "facebook",
			id:     "1234567890",
		},
		{
			name:   "facebook video",
			url:    "https://www.facebook.com/testpage/videos/vb.987654321/1234567890/",
			domain: "facebook",
			id:     "1234567890",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			domain, parts, _ := parseURL(tc.url)
			if domain != tc.domain {
				t.Fatalf("domain = %s, want %s", domain, tc.domain)
			}
			if len(parts) == 0 {
				t.Fatalf("no parts returned for %s", tc.url)
			}
			gotID := parts[len(parts)-1]
			if gotID != tc.id {
				t.Errorf("id = %s, want %s", gotID, tc.id)
			}
		})
	}
}
