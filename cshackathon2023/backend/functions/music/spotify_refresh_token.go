package functions

import (
	"backend/modules"
	"backend/types/payload"
	"backend/types/response"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"net/url"
	"strings"
)

func SpotifyRefreshToken(c *fiber.Ctx, id string) (*payload.SpotifyCredentials, *response.ErrorInstance) {
	// * Construct form data values
	values := url.Values{}
	values.Set("grant_type", "refresh_token")
	values.Set("refresh_token", modules.Conf.SpotifyRefreshToken)

	var data *payload.SpotifyCredentials
	if err := DoRequest(
		c,
		"POST",
		"https://accounts.spotify.com/api/token",
		strings.NewReader(values.Encode()),
		func(r *http.Request) {
			r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			r.Header.Set("Authorization", "Basic "+modules.Conf.SpotifyAuthorization)
		},
		&data,
	); err != nil {
		return nil, err
	}

	return data, nil
}
