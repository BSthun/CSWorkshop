package functions

import (
	"backend/types/response"
	"fmt"
	"net/http"
	"net/url"
)

func SpotifyAddQueue(spotifyId string) (map[string]any, *response.ErrorInstance) {
	data := map[string]any{}
	if err := DoRequest(
		nil,
		"POST",
		fmt.Sprintf("https://api.spotify.com/v1/me/player/queue?uri=%s", url.QueryEscape("spotify:track:"+spotifyId)),
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
