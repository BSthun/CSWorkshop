package functions

import (
	"backend/types/payload"
	"backend/types/response"
	"fmt"
	"net/http"
)

func SpotifyNowPlaying() (*payload.SpotifyPlaybackState, *response.ErrorInstance) {
	var data *payload.SpotifyPlaybackState
	if err := DoRequest(
		nil,
		"GET",
		fmt.Sprintf("https://api.spotify.com/v1/me/player/currently-playing"),
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
