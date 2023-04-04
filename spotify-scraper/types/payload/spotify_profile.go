package payload

import (
	"database/sql/driver"
	"encoding/json"
)

type SpotifyProfile struct {
	DisplayName  *string              `json:"display_name"`
	ExternalUrls *SpotifyExternalUrls `json:"external_urls"`
	Followers    *SpotifyFollowers    `json:"followers"`
	Href         *string              `json:"href"`
	Id           *string              `json:"id"`
	Images       []*SpotifyImage      `json:"images"`
	Type         *string              `json:"type"`
	Uri          *string              `json:"uri"`
}

func (r *SpotifyProfile) Value() (driver.Value, error) {
	return json.Marshal(map[string]any{
		"display_name": r.DisplayName,
		"followers":    r.Followers,
		"href":         r.Href,
		"images":       r.Images,
		"uri":          r.Uri,
	})
}

func (r *SpotifyProfile) Scan(src interface{}) error {
	return json.Unmarshal(src.([]byte), r)
}

type SpotifyExternalUrls struct {
	Spotify string `json:"spotify"`
}

type SpotifyFollowers struct {
	Href  *string `json:"href"`
	Total *int    `json:"total"`
}

type SpotifyImage struct {
	Height *int    `json:"height"`
	Url    *string `json:"url"`
	Width  *int    `json:"width"`
}

type SpotifyError struct {
	Error *string `json:"error"`
}

type SpotifyApiError struct {
	Error struct {
		Status  *int    `json:"status"`
		Message *string `json:"message"`
	} `json:"error"`
}
