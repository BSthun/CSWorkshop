package branchModel

import "time"

type Lv3Artist struct {
	Id        *uint64    `gorm:"primaryKey"`
	SpotifyId *string    `gorm:"type:VARCHAR(255); index:spotify_id,unique; not null"`
	Name      *string    `gorm:"type:VARCHAR(255); not null"`
	Href      *string    `gorm:"type:TEXT; not null"`
	CreatedAt *time.Time `gorm:"not null"`
	UpdatedAt *time.Time `gorm:"not null"`
}

func (r *Lv3Artist) TableName() string {
	return "lv3_artists"
}
