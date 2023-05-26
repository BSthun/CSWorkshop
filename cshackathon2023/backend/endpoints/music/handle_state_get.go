package music

import (
	"backend/functions/backdrop"
	"backend/modules"
	"backend/types/common"
	"backend/types/model"
	"backend/types/payload"
	"backend/types/response"
	"backend/utils/value"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func StateGetHandler(c *fiber.Ctx) error {
	// * Parse user
	u := c.Locals("u").(*jwt.Token).Claims.(*common.UserClaims)

	// Fetch now playing
	nowPlaying := backdrop.GetNowPlaying()

	// Fetch all queues
	var queues []*model.Queue
	if result := modules.DB.Preload("Track").Preload("User").Find(&queues); result.Error != nil {
		return response.Error(c, false, "Unable to fetch queues", result.Error)
	}

	// Map queues
	mappedQueues, _ := value.Iterate(queues, func(queue *model.Queue) (*payload.MusicQueueItem, *response.ErrorInstance) {
		isPlaying := false
		if nowPlaying != nil && nowPlaying.QueueId != nil && *nowPlaying.QueueId == *queue.Id {
			isPlaying = true
		}

		isOwned := false
		if *queue.UserId == *u.UserId {
			isOwned = true
		}

		return &payload.MusicQueueItem{
			MusicItem: payload.MusicItem{
				ArtworkURL: queue.Track.CoverUrl,
				Title:      queue.Track.Name,
				Album:      queue.Track.Album,
				Artist:     queue.Track.Artist,
			},
			ID:        queue.Id,
			QueueBy:   queue.User.Name,
			QueueAt:   queue.CreatedAt,
			IsPlaying: &isPlaying,
			IsOwned:   &isOwned,
		}, nil
	})

	logPayload, _ := json.Marshal(map[string]any{
		"device": c.Get("User-Agent"),
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
		"nowPlaying": nowPlaying,
		"queues":     mappedQueues,
	}))
}
