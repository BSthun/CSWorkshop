package functions

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"backend/types/payload"
	"backend/types/response"
)

func SpotifyGetProfile(c *fiber.Ctx, accessToken string) (*payload.SpotifyProfile, *response.ErrorInstance) {
	var data *payload.SpotifyProfile
	if err := DoRequest(
		c,
		"GET",
		"https://api.spotify.com/v1/me",
		nil,
		func(r *http.Request) {
			r.Header.Set("Authorization", "Bearer "+accessToken)
		},
		&data,
	); err != nil {
		return nil, err
	}

	return data, nil
}
