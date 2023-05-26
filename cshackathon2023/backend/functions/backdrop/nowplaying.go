package backdrop

import (
	functions "backend/functions/music"
	"backend/types/payload"
	"backend/utils/text"
)

func GetNowPlaying() *payload.BackdropNowPlaying {
	nowPlaying, _ := functions.SpotifyNowPlaying()
	var artists []string
	for _, s := range nowPlaying.Item.Artists {
		artists = append(artists, *s.Name)
	}

	artistName := text.FormattedText(artists)
	return &payload.BackdropNowPlaying{
		CoverURL: nowPlaying.Item.Album.Images[0].Url,
		Title:    nowPlaying.Item.Name,
		Artist:   &artistName,
		Album:    nowPlaying.Item.Album.Name,
		//QueueBy:  value.Ptr("APISIT MANEERAT"),
	}
}
