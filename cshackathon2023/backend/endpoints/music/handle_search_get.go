package music

import (
	functions "backend/functions/music"
	"backend/modules"
	"backend/types/common"
	"backend/types/model"
	"backend/types/payload"
	"backend/types/response"
	"backend/utils/text"
	"backend/utils/value"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"net/url"
)

func SearchGetHandler(c *fiber.Ctx) error {
	// * Parse user
	u := c.Locals("u").(*jwt.Token).Claims.(*common.UserClaims)

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
	spotifySearchTrack, err := functions.SpotifySearchTrack(url.QueryEscape(*query.Query))
	if err != nil {
		return err
	}

	// * Map Spotify search API to response
	searchItems, _ := value.Iterate(spotifySearchTrack.Tracks.Items, func(track *payload.SpotifyTrack) (*payload.MusicSearchItem, *response.ErrorInstance) {
		var artists []string
		for _, s := range track.Artists {
			artists = append(artists, *s.Name)
		}

		artistName := text.FormattedText(artists)
		return &payload.MusicSearchItem{
			MusicItem: payload.MusicItem{
				ArtworkURL: track.Album.Images[0].Url,
				Title:      track.Name,
				Album:      track.Album.Name,
				Artist:     &artistName,
			},
			SpotifyId: track.Id,
		}, nil
	})

	// * Create logs
	logPayload, _ := json.Marshal(map[string]any{
		"query": *query.Query,
	})
	log := &model.Log{
		UserId:    u.UserId,
		User:      nil,
		Type:      value.Ptr("search"),
		Payload:   value.Ptr(string(logPayload)),
		CreatedAt: nil,
		UpdatedAt: nil,
	}
	if result := modules.DB.Create(log); result.Error != nil {
		return response.Error(c, false, "Unable to create log", result.Error)
	}

	return c.JSON(response.Success(c, map[string]any{
		"items": searchItems,
	}))
}
