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
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func QueuePostHandler(c *fiber.Ctx) error {
	// * Parse user
	u := c.Locals("u").(*jwt.Token).Claims.(*common.UserClaims)

	// * Parse body
	body := new(payload.MusicQueueBody)
	if err := c.BodyParser(body); err != nil {
		return response.Error(c, false, "Unable to parse body", err)
	}

	// * Validate body
	if err := text.Validator.Struct(body); err != nil {
		return response.Error(c, false, "Unable to validate body", err)
	}

	// * Get spotify track
	spotifyTrack, err := functions.SpotifyTrack(*body.TrackId)
	if err != nil {
		return err
	}

	// * Check track exists
	// TODO: Change spotifyTrack.Artists[0].Name to name of all artists concatenated with comma
	track := &model.Track{
		Id:        nil,
		SpotifyId: spotifyTrack.Id,
		Name:      spotifyTrack.Name,
		Artist:    spotifyTrack.Artists[0].Name,
		Album:     spotifyTrack.Album.Name,
		CoverUrl:  spotifyTrack.Album.Images[0].Url,
		CreatedAt: nil,
		UpdatedAt: nil,
	}
	if result := modules.DB.Where("spotify_id = ?", *spotifyTrack.Id).Attrs(track).FirstOrCreate(track); result.Error != nil {
		return response.Error(c, false, "Unable to check track exists", result.Error)
	}

	// * Create queue
	queue := &model.Queue{
		Id:        nil,
		UserId:    u.UserId,
		User:      nil,
		TrackId:   track.Id,
		Track:     nil,
		Played:    value.Ptr(false),
		CreatedAt: nil,
		UpdatedAt: nil,
	}
	if result := modules.DB.Create(queue); result.Error != nil {
		return response.Error(c, false, "Unable to create queue", result.Error)
	}

	return c.JSON(response.Success(c, "Successfully created queue"))
}
