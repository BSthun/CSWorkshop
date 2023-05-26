package music

import (
	functions "backend/functions/music"
	"backend/types/payload"
	"backend/types/response"
	"backend/utils/text"
	"backend/utils/value"
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
	spotifySearchTrack, err := functions.SpotifySearchTrack(*query.Query)
	if err != nil {
		return err
	}
	//spew.Dump(spotifySearchTrack)

	// * Map Spotify search API to response
	searchList, _ := value.Iterate(spotifySearchTrack.Tracks.Items, func(track *payload.SpotifyTrack) (*payload.MusicSearchItem, *response.ErrorInstance) {
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

	return c.JSON(response.Success(c, map[string]any{
		"track": searchList,

		//"list": []*payload.MusicSearchItem{
		//	{
		//		MusicItem: payload.MusicItem{
		//			ArtworkURL: value.Ptr("https://i.scdn.co/image/ab67616d0000b273c492874e96f19148018e759e"),
		//			Title:      value.Ptr("Baby Steps"),
		//			Album:      value.Ptr("'Twinkle' Mini Album"),
		//			Artist:     value.Ptr("Girls' Generation-TTS"),
		//		},
		//		SpotifyId: value.Ptr("02wkrfy00sbUO5DxirKMeV"),
		//	},
		//},
	}))
}
