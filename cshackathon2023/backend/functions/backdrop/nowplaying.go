package backdrop

import (
	functions "backend/functions/music"
	"backend/modules"
	"backend/types/model"
	"backend/types/payload"
	"backend/utils/text"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

var lastQueued time.Time

var nowPlayingData *payload.BackdropNowPlaying
var nowPlayingLastUpdated time.Time
var lastQueueId string

func GetNowPlaying() *payload.BackdropNowPlaying {
	if time.Now().Sub(nowPlayingLastUpdated) > 5*time.Second {
		nowPlayingData = FetchNowPlaying()
		nowPlayingLastUpdated = time.Now()
	}
	return nowPlayingData
}

func FetchNowPlaying() *payload.BackdropNowPlaying {
	// * Get now playing state
	nowPlaying, _ := functions.SpotifyNowPlaying()

	// * Format artist name
	var artists []string
	for _, s := range nowPlaying.Item.Artists {
		artists = append(artists, *s.Name)
	}
	artistName := text.FormattedText(artists)

	// * Empty now playing fallback
	if nowPlaying == nil || nowPlaying.Item == nil {
		logrus.Warn("UNABLE TO FETCH NOW PLAYING")

		// * Get last played track
		var queue *model.Queue
		if result := modules.DB.Preload("User").Preload("Track").Order("created_at DESC").First(&queue, "played = 1"); result.Error != nil {
			logrus.Fatal("UNABLE TO FETCH LAST PLAYED TRACK: " + result.Error.Error())
		}

		return &payload.BackdropNowPlaying{
			QueueId:  queue.Id,
			CoverURL: queue.Track.CoverUrl,
			Title:    queue.Track.Name,
			Artist:   queue.Track.Artist,
			Album:    queue.Track.Album,
			QueueBy:  queue.User.Name,
		}
	}

	// * Check next track
	var nextQueue *model.Queue
	if result := modules.DB.Preload("Track").Order("created_at ASC").First(&nextQueue, "played = 0"); result.Error == gorm.ErrRecordNotFound {
		// * Get last played track
		var queue *model.Queue
		if result := modules.DB.Preload("User").Preload("Track").Order("created_at DESC").First(&queue, "played = 1"); result.Error != nil {
			logrus.Warn("UNABLE TO FETCH LAST PLAYED TRACK 2: " + result.Error.Error())
		}

		return &payload.BackdropNowPlaying{
			QueueId:  nil,
			CoverURL: nowPlaying.Item.Album.Images[0].Url,
			Title:    nowPlaying.Item.Name,
			Artist:   &artistName,
			Album:    nowPlaying.Item.Album.Name,
			QueueBy:  nil,
		}
	} else if result.Error != nil {
		logrus.Fatal("UNABLE TO FETCH NEXT QUEUE TRACK: " + result.Error.Error())
	}

	if *nowPlaying.Item.Id == *nextQueue.Track.SpotifyId {
		// * Update queue status
		if result := modules.DB.Model(&nextQueue).Update("played", true); result.Error != nil {
			logrus.Fatal("UNABLE TO UPDATE QUEUE STATUS: " + result.Error.Error())
		}
	}

	// * Check if track is going to be played next
	if *nowPlaying.ProgressMs > *nowPlaying.Item.DurationMs-15000 {
		if time.Now().Sub(lastQueued) > 15*time.Second {
			// * Add next track to queue
			if lastQueueId != *nextQueue.Track.SpotifyId {
				_, err := functions.SpotifyAddQueue(*nextQueue.Track.SpotifyId)
				if err != nil {
					logrus.Fatal("UNABLE TO ADD QUEUE: " + err.Error())
				}
				lastQueued = time.Now()
				lastQueueId = *nextQueue.Track.SpotifyId
			}
		}
	}

	// * Get current track
	var currentQueue *model.Queue
	if result := modules.DB.Preload("User").Preload("Track").Order("created_at DESC").First(&currentQueue, "played = 1"); result.Error != nil {
		logrus.Warn("UNABLE TO FETCH CURRENT QUEUE TRACK: " + result.Error.Error())
		return &payload.BackdropNowPlaying{
			QueueId:  nil,
			CoverURL: nowPlaying.Item.Album.Images[0].Url,
			Title:    nowPlaying.Item.Name,
			Artist:   &artistName,
			Album:    nowPlaying.Item.Album.Name,
			QueueBy:  nil,
		}
	}

	// * Check if current track is same as now playing
	if *nowPlaying.Item.Id == *currentQueue.Track.SpotifyId {
		return &payload.BackdropNowPlaying{
			QueueId:  currentQueue.Id,
			CoverURL: currentQueue.Track.CoverUrl,
			Title:    currentQueue.Track.Name,
			Artist:   currentQueue.Track.Artist,
			Album:    currentQueue.Track.Album,
			QueueBy:  currentQueue.User.Name,
		}
	}

	return &payload.BackdropNowPlaying{
		QueueId:  nil,
		CoverURL: nowPlaying.Item.Album.Images[0].Url,
		Title:    nowPlaying.Item.Name,
		Artist:   &artistName,
		Album:    nowPlaying.Item.Album.Name,
		QueueBy:  nil,
	}
}
