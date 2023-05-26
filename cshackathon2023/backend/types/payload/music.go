package payload

import "time"

type MusicQueueBody struct {
	TrackId *string `json:"trackId"`
}

type MusicSearchQuery struct {
	Query *string `query:"query"`
}

type MusicItem struct {
	ArtworkURL *string `json:"artwork_url,omitempty"`
	Title      *string `json:"title,omitempty"`
	Album      *string `json:"album,omitempty"`
	Artist     *string `json:"artist,omitempty"`
}

type MusicSearchItem struct {
	MusicItem
	SpotifyId *string `json:"spotify_id,omitempty"`
}

type MusicQueueItem struct {
	MusicItem
	ID        *uint64    `json:"id,omitempty"`
	QueueBy   *string    `json:"queue_by,omitempty"`
	QueueAt   *time.Time `json:"queue_at,omitempty"`
	IsPlaying *bool      `json:"is_playing,omitempty"`
	IsOwned   *bool      `json:"is_owned,omitempty"`
}
