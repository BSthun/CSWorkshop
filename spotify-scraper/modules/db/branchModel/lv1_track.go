package branchModel

type Lv1Track struct {
	Id         *uint64 `gorm:"primaryKey"`
	SpotifyId  *string `gorm:"type:VARCHAR(255); index:spotify_id,unique; not null"`
	Name       *string `gorm:"type:VARCHAR(255); not null"`
	Album      *string `gorm:"type:VARCHAR(255); not null"`
	Artist     *string `gorm:"type:VARCHAR(255); not null"`
	ArtworkUrl *string `gorm:"type:TEXT; not null"`
	Duration   *int64  `gorm:"not null"`
	Popularity *int64  `gorm:"not null"`
	Explicit   *bool   `gorm:"not null"`
	Year       *int64  `gorm:"not null"`
	Count      *int64  `gorm:"not null"`
}

func (r *Lv1Track) TableName() string {
	return "lv1_tracks"
}
