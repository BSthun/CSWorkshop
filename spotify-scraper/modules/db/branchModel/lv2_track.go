package branchModel

import "time"

type Lv2Track struct {
	Id         *uint64    `gorm:"primaryKey"`
	SpotifyId  *string    `gorm:"type:VARCHAR(255); index:spotify_id,unique; not null"`
	Name       *string    `gorm:"type:VARCHAR(255); not null"`
	AlbumId    *uint64    `gorm:"not null"`
	Album      *Lv2Album  `gorm:"foreignKey:AlbumId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Duration   *int64     `gorm:"not null"`
	Popularity *int64     `gorm:"not null"`
	Explicit   *bool      `gorm:"not null"`
	Count      *int64     `gorm:"not null"`
	CreatedAt  *time.Time `gorm:"not null"`
	UpdatedAt  *time.Time `gorm:"not null"`
}

func (r *Lv2Track) TableName() string {
	return "lv2_tracks"
}
