package login

import (
	"net/url"

	"github.com/gofiber/fiber/v2"

	"backend/functions"
	"backend/modules/db"
	"backend/modules/db/model"
	"backend/types/payload"
	"backend/types/response"
	"backend/utils/text"
)

func RedirectHandler(c *fiber.Ctx) error {
	// * Parse body
	var body *payload.LoginSpotifyRedirect
	if err := c.BodyParser(&body); err != nil {
		return response.Error(c, true, "Unable to parse body", err)
	}

	// * Check state
	var state *model.ClientState
	if result := db.DB.Preload("Client").First(&state, "email = ?", body.Email); result.Error != nil {
		return response.Error(c, true, "Email is not exist in whitelist", result.Error)
	}

	// * Refresh state
	st := text.Random(text.StringSet.MixedAlphaNum, 12)

	// * Update state
	if result := db.DB.Model(&state).Update("state", st); result.Error != nil {
		return response.Error(c, true, "Unable to update state", result.Error)
	}

	values := url.Values{}
	values.Set("response_type", "code")
	values.Set("client_id", *state.Client.SpotifyClientId)
	values.Set("scope", "user-read-playback-state user-modify-playback-state")
	values.Set("redirect_uri", functions.SpotifyRedirectUri())
	values.Set("state", *st)

	target := "https://accounts.spotify.com/authorize"
	target += "?" + values.Encode()

	return c.JSON(response.Info(c, map[string]any{
		"redirect": target,
	}))
}
