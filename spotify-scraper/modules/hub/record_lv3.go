package hub

import (
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"backend/functions"
	"backend/modules/db/branchModel"
	"backend/types/payload"
	"backend/types/response"
	"backend/utils/value"
)

func RecordLv3(db *gorm.DB, state *payload.SpotifyPlaybackState) *response.ErrorInstance {
	// * Check artist
	var artists []*branchModel.Lv3Artist
	for _, a := range state.Item.Artists {
		var artist *branchModel.Lv3Artist
		if result := db.Where("spotify_id = ?", a.Id).First(&artist); result.Error == gorm.ErrRecordNotFound {
			artist = &branchModel.Lv3Artist{
				Id:        nil,
				SpotifyId: &a.Id,
				Name:      &a.Name,
				Href:      &a.Href,
				CreatedAt: nil,
				UpdatedAt: nil,
			}
			if result := db.Create(artist); result.Error != nil {
				logrus.Warnf("Unable to create lv3 artist: %v", result.Error)
				return response.Error(nil, true, "Unable to create lv3 artist: %v", result.Error)
			}
		} else if result.Error != nil {
			logrus.Warnf("Unable to get lv3 artist: %v", result.Error)
			return response.Error(nil, true, "Unable to get lv3 artist: %v", result.Error)
		}
		artists = append(artists, artist)
	}

	// * Check album
	var album *branchModel.Lv3Album
	if result := db.Where("spotify_id = ?", state.Item.Album.Id).First(&album); result.Error == gorm.ErrRecordNotFound {
		album = &branchModel.Lv3Album{
			Id:         nil,
			SpotifyId:  &state.Item.Album.Id,
			Name:       &state.Item.Album.Name,
			ArtworkUrl: state.Item.Album.Images[0].Url,
			Year:       functions.ExtractYear(state.Item.Album.ReleaseDate),
			CreatedAt:  nil,
			UpdatedAt:  nil,
		}
		if result := db.Create(album); result.Error != nil {
			logrus.Warnf("Unable to create lv3 album: %v", result.Error)
			return response.Error(nil, true, "Unable to create lv3 album: %v", result.Error)
		}
	} else if result.Error != nil {
		logrus.Warnf("Unable to get lv2 album: %v", result.Error)
		return response.Error(nil, true, "Unable to get lv2 album: %v", result.Error)
	}

	// * Check album-artist
	var albumArtistCount int64
	if result := db.Model(new(branchModel.Lv3AlbumArtist)).Where("album_id = ?", album.Id).Count(&albumArtistCount); result.Error != nil {
		logrus.Warnf("Unable to get lv3 album-artist count: %v", result.Error)
		return response.Error(nil, true, "Unable to get lv3 album-artist count: %v", result.Error)
	}

	if albumArtistCount == 0 {
		for _, artist := range artists {
			if result := db.Create(&branchModel.Lv3AlbumArtist{
				AlbumId:  album.Id,
				Album:    nil,
				ArtistId: artist.Id,
				Artist:   nil,
			}); result.Error != nil {
				logrus.Warnf("Unable to create lv3 album-artist: %v", result.Error)
				return response.Error(nil, true, "Unable to create lv3 album-artist: %v", result.Error)
			}
		}
	}

	// * Check track
	var track *branchModel.Lv3Track
	if result := db.Where("spotify_id = ?", state.Item.Id).First(&track); result.Error == gorm.ErrRecordNotFound {
		track = &branchModel.Lv3Track{
			Id:         nil,
			SpotifyId:  state.Item.Id,
			Name:       state.Item.Name,
			AlbumId:    album.Id,
			Album:      nil,
			Duration:   state.Item.DurationMs,
			Popularity: state.Item.Popularity,
			Explicit:   state.Item.Explicit,
			PreviewUrl: state.Item.PreviewUrl,
			CreatedAt:  nil,
			UpdatedAt:  nil,
		}
		if result := db.Create(track); result.Error != nil {
			logrus.Warnf("Unable to create lv3 track: %v", result.Error)
			return response.Error(nil, true, "Unable to create lv3 track: %v", result.Error)
		}
	}

	// * Create activity
	var context *string
	if state.Context != nil {
		context = &state.Context.Uri
	}
	if result := db.Create(&branchModel.Lv3Activity{
		Timestamp:    value.Ptr(time.Now()),
		TrackId:      track.Id,
		Track:        nil,
		ShuffleState: &state.ShuffleState,
		RepeatState:  &state.RepeatState,
		DeviceName:   &state.Device.Name,
		Context:      context,
	}); result.Error != nil {
		logrus.Warnf("Unable to create lv3 activity: %v", result.Error)
		return response.Error(nil, true, "Unable to create lv3 activity: %v", result.Error)
	}

	return nil
}
