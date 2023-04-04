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

func RecordLv1(db *gorm.DB, state *payload.SpotifyPlaybackState) *response.ErrorInstance {
	if result := db.Clauses(clause.OnConflict{
		DoUpdates: clause.Assignments(
			map[string]any{
				"count": gorm.Expr("count + 1"),
			},
		),
	}).Create(&branchModel.Lv1Track{
		Id:         nil,
		SpotifyId:  state.Item.Id,
		Name:       state.Item.Name,
		Album:      &state.Item.Album.Name,
		Artist:     &state.Item.Album.Artists[0].Name,
		ArtworkUrl: state.Item.Album.Images[0].Url,
		Duration:   state.Item.DurationMs,
		Popularity: state.Item.Popularity,
		Explicit:   state.Item.Explicit,
		Year:       functions.ExtractYear(state.Item.Album.ReleaseDate),
		Count:      value.Ptr[int64](1),
	}); result.Error != nil {
		logrus.Warnf("Unable to update spotify analytics (lv1): %v", result.Error)
		return response.Error(nil, true, "Unable to update spotify analytics (lv1)")
	}
	return nil
}
