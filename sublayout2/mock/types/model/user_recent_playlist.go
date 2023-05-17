package model

import "time"

type UserRecentPlaylist struct {
	UserId       *uint64    `gorm:"primaryKey"`
	User         *User      `gorm:"foreignKey:UserId"`
	PlaylistId   *uint64    `gorm:"primaryKey"`
	Playlist     *Playlist  `gorm:"foreignKey:PlaylistId"`
	LastListened *time.Time `gorm:"not null"`
}
