package track

import (
	"gopkg.in/loremipsum.v1"
	"gorm.io/gorm"
	"mock/modules"
	"mock/types/model"
	"mock/types/scraper_model"
	"mock/utils/value"
	"time"
)

func Fetch(db *gorm.DB) {
	artistIdRef := make(map[uint64]*uint64)

	var artists []*branchModel.Lv3Artist
	if result := db.Find(&artists); result.Error != nil {
		panic(result.Error)
	}

	for _, artist := range artists {
		if _, ok := artistRef[*artist.SpotifyId]; !ok {
			loremIpsumGeneratoe := loremipsum.NewWithSeed(time.Now().UnixNano())
			biography := loremIpsumGeneratoe.Sentence()
			a := &model.Artist{
				Id:         nil,
				Name:       artist.Name,
				Biography:  &biography,
				Href:       nil,
				ProfileUrl: value.Ptr("https://i.scdn.co/image/ab67616d00001e02ff9ca10b55ce82ae553c8228"),
				BannerUrl:  value.Ptr("https://i.scdn.co/image/ab67618600001016e78af2b2865d4ce269e84838"),
			}
			if result := modules.DB.Create(a); result.Error != nil {
				panic(result.Error)
			}
			artistRef[*artist.SpotifyId] = a.Id
		}
		artistIdRef[*artist.Id] = artistRef[*artist.SpotifyId]

		var albums []*branchModel.Lv3Album
		if result := db.Find(&albums); result.Error != nil {
			panic(result.Error)
		}
		for _, album := range albums {
			if _, ok := albumRef[*album.SpotifyId]; ok {

			}
		}
	}

	var tracks []*branchModel.Lv3Track
	if result := db.Find(&tracks); result.Error != nil {
		panic(result.Error)
	}

	for _, track := range tracks {
		if _, ok := trackRef[*track.SpotifyId]; ok {
			continue
		}
		trackRef[*track.SpotifyId] = true

		t := &model.Track{
			Id:             nil,
			Name:           track.Name,
			AlbumId:        nil,
			Duration:       nil,
			Popularity:     nil,
			Explicit:       nil,
			PreviewUrl:     nil,
			PlaylistTracks: nil,
		}
		if result := db.Create(t); result.Error != nil {
			panic(result.Error)
		}
	}
}
