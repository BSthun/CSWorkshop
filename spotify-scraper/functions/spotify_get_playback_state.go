package functions

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"backend/types/payload"
	"backend/types/response"
)

func SpotifyGetPlaybackState(c *fiber.Ctx, accessToken string) (*payload.SpotifyPlaybackState, *response.ErrorInstance) {
	var data *payload.SpotifyPlaybackState
	if err := DoRequest(
		c,
		"GET",
		"https://api.spotify.com/v1/me/player",
		nil,
		func(r *http.Request) {
			r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			r.Header.Set("Authorization", "Bearer "+accessToken)
		},
		&data,
	); err != nil {
		return nil, err
	}

	return data, nil
}
