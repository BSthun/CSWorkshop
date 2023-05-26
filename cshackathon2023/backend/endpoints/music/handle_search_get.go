package music

import (
	"backend/types/payload"
	"backend/types/response"
	"backend/utils/text"
	"backend/utils/value"
	"github.com/davecgh/go-spew/spew"
	"github.com/gofiber/fiber/v2"
)

func SearchGetHandler(c *fiber.Ctx) error {
	// * Parse user
	// u := c.Locals("u").(*jwt.Token).Claims.(*common.UserClaims)

	// * Parse query
	query := new(payload.MusicSearchQuery)
	if err := c.QueryParser(query); err != nil {
		return response.Error(c, false, "Unable to parse query", err)
	}

	// * Validate query
	if err := text.Validator.Struct(query); err != nil {
		return response.Error(c, false, "Unable to validate query", err)
	}

	// TODO: Search spotify track
	// * Get Spotify search API by query
	spew.Dump(query.Query)

	// * Map Spotify search API to response

	return c.JSON(response.Success(c, map[string]any{
		"list": []*payload.MusicSearchItem{
			{
				MusicItem: payload.MusicItem{
					ArtworkURL: value.Ptr("https://i.scdn.co/image/ab67616d0000b273c492874e96f19148018e759e"),
					Title:      value.Ptr("Baby Steps"),
					Album:      value.Ptr("'Twinkle' Mini Album"),
					Artist:     value.Ptr("Girls' Generation-TTS"),
				},
				SpotifyId: value.Ptr("02wkrfy00sbUO5DxirKMeV"),
			},
		},
	}))
}
