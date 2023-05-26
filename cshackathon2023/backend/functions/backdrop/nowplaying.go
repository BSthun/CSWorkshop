package backdrop

import (
	functions "backend/functions/music"
	"backend/types/payload"
	"backend/utils/text"
	"github.com/davecgh/go-spew/spew"
)

func GetNowPlaying() *payload.BackdropNowPlaying {
	np, _ := functions.SpotifyNowPlaying()
	spew.Dump(np)
	var artists []string
	for _, s := range np.Item.Artists {
		artists = append(artists, *s.Name)
	}

	artistName := text.FormattedText(artists)
	return &payload.BackdropNowPlaying{
		CoverURL: np.Item.Album.Images[0].Url,
		Title:    np.Item.Name,
		Artist:   &artistName,
		Album:    np.Item.Album.Name,
		//QueueBy:  value.Ptr("APISIT MANEERAT"),
	}
}
