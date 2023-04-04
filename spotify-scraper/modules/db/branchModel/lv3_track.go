package branchModel

import "time"

type Lv3Track struct {
	Id         *uint64    `gorm:"primaryKey"`
	SpotifyId  *string    `gorm:"type:VARCHAR(255); index:spotify_id,unique; not null"`
	Name       *string    `gorm:"type:VARCHAR(255); not null"`
	AlbumId    *uint64    `gorm:"not null"`
	Album      *Lv3Album  `gorm:"foreignKey:AlbumId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Duration   *int64     `gorm:"not null"`
	Popularity *int64     `gorm:"not null"`
	Explicit   *bool      `gorm:"not null"`
	PreviewUrl *string    `gorm:"type:TEXT; null"`
	CreatedAt  *time.Time `gorm:"not null"`
	UpdatedAt  *time.Time `gorm:"not null"`
}

func (r *Lv3Track) TableName() string {
	return "lv3_tracks"
}
