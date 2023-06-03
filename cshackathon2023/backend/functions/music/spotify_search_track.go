package functions

import (
	"backend/types/payload"
	"backend/types/response"
	"fmt"
	"net/http"
)

func SpotifySearchTrack(trackName string) (*payload.SpotifyTrackSearch, *response.ErrorInstance) {
	var data *payload.SpotifyTrackSearch
	if err := DoRequest(
		nil,
		"GET",
		fmt.Sprintf("https://api.spotify.com/v1/search?q=%s&type=track", trackName),
		nil,
		func(r *http.Request) {
			r.Header.Set("Authorization", "Bearer "+GetSpotifyToken())
		},
		&data,
	); err != nil {
		return nil, err
	}

	return data, nil
}
