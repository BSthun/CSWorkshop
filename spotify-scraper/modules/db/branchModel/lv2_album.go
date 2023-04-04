package branchModel

import "time"

type Lv2Album struct {
	Id         *uint64    `gorm:"primaryKey"`
	SpotifyId  *string    `gorm:"type:VARCHAR(255); index:spotify_id,unique; not null"`
	Name       *string    `gorm:"type:VARCHAR(255); not null"`
	ArtistId   *uint64    `gorm:"not null"`
	Artist     *Lv2Artist `gorm:"foreignKey:ArtistId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ArtworkUrl *string    `gorm:"type:TEXT; not null"`
	Year       *int64     `gorm:"not null"`
	CreatedAt  *time.Time `gorm:"not null"`
	UpdatedAt  *time.Time `gorm:"not null"`
}

func (r *Lv2Album) TableName() string {
	return "lv2_albums"
}
