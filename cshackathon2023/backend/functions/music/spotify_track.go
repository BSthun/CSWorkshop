package functions

import (
	"backend/types/payload"
	"backend/types/response"
	"fmt"
	"net/http"
)

func SpotifyTrack(id string) (*payload.SpotifyTrack, *response.ErrorInstance) {
	var data *payload.SpotifyTrack
	if err := DoRequest(
		nil,
		"GET",
		fmt.Sprintf("https://api.spotify.com/v1/tracks/%s", id),
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
