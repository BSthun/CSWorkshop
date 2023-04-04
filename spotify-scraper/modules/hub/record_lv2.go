package hub

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"backend/functions"
	"backend/modules/db/branchModel"
	"backend/types/payload"
	"backend/types/response"
	"backend/utils/value"
)

func RecordLv2(db *gorm.DB, state *payload.SpotifyPlaybackState) *response.ErrorInstance {
	// * Check artist
	var artist *branchModel.Lv2Artist
	if result := db.Where("spotify_id = ?", state.Item.Artists[0].Id).First(&artist); result.Error == gorm.ErrRecordNotFound {
		artist = &branchModel.Lv2Artist{
			Id:        nil,
			SpotifyId: &state.Item.Artists[0].Id,
			Name:      &state.Item.Artists[0].Name,
			Href:      &state.Item.Artists[0].Href,
			CreatedAt: nil,
			UpdatedAt: nil,
		}
		if result := db.Create(artist); result.Error != nil {
			logrus.Warnf("Unable to create lv2 artist: %v", result.Error)
			return response.Error(nil, true, "Unable to create lv2 artist: %v", result.Error)
		}
	} else if result.Error != nil {
		logrus.Warnf("Unable to get lv2 artist: %v", result.Error)
		return response.Error(nil, true, "Unable to get lv2 artist: %v", result.Error)
	}

	// * Check album
	var album *branchModel.Lv2Album
	if result := db.Where("spotify_id = ?", state.Item.Album.Id).First(&album); result.Error == gorm.ErrRecordNotFound {
		album = &branchModel.Lv2Album{
			Id:         nil,
			SpotifyId:  &state.Item.Album.Id,
			Name:       &state.Item.Album.Name,
			ArtistId:   artist.Id,
			Artist:     nil,
			ArtworkUrl: state.Item.Album.Images[0].Url,
			Year:       functions.ExtractYear(state.Item.Album.ReleaseDate),
			CreatedAt:  nil,
			UpdatedAt:  nil,
		}
		if result := db.Create(album); result.Error != nil {
			logrus.Warnf("Unable to create lv2 album: %v", result.Error)
			return response.Error(nil, true, "Unable to create lv2 album: %v", result.Error)
		}
	} else if result.Error != nil {
		logrus.Warnf("Unable to get lv2 album: %v", result.Error)
		return response.Error(nil, true, "Unable to get lv2 album: %v", result.Error)
	}

	// * Check track
	if result := db.Clauses(clause.OnConflict{
		DoUpdates: clause.Assignments(
			map[string]any{
				"count": gorm.Expr("count + 1"),
			},
		),
	}).Create(&branchModel.Lv2Track{
		Id:         nil,
		SpotifyId:  state.Item.Id,
		Name:       state.Item.Name,
		AlbumId:    album.Id,
		Album:      nil,
		Duration:   state.Item.DurationMs,
		Popularity: state.Item.Popularity,
		Explicit:   state.Item.Explicit,
		Count:      value.Ptr[int64](1),
		CreatedAt:  nil,
		UpdatedAt:  nil,
	}); result.Error != nil {
		logrus.Warnf("Unable to create lv2 track: %v", result.Error)
		return response.Error(nil, true, "Unable to create lv2 track: %v", result.Error)
	}
	return nil
}
