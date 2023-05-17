package model

import "time"

type UserRecentArtist struct {
	UserId       *uint64    `gorm:"primaryKey"`
	User         *User      `gorm:"foreignKey:UserId"`
	ArtistId     *uint64    `gorm:"primaryKey"`
	Artist       *Artist    `gorm:"foreignKey:ArtistId"`
	LastListened *time.Time `gorm:"not null"`
}
