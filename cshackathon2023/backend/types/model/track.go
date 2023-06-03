package model

import "time"

type Track struct {
	Id        *uint64    `gorm:"primaryKey"`
	SpotifyId *string    `gorm:"type:VARCHAR(255); index:idx_spotify_id,unique; not null"`
	Name      *string    `gorm:"type:VARCHAR(255); not null"`
	Artist    *string    `gorm:"type:VARCHAR(255); not null"`
	Album     *string    `gorm:"type:VARCHAR(255); not null"`
	CoverUrl  *string    `gorm:"type:TEXT; not null"`
	CreatedAt *time.Time `gorm:"not null"` // Embedded field
	UpdatedAt *time.Time `gorm:"not null"` // Embedded field
}
