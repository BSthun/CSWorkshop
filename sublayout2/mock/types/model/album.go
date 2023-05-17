package model

import "gorm.io/gorm"

type Album struct {
	gorm.Model
	Id           *uint64       `gorm:"primaryKey;autoIncrement"`
	Name         string        `gorm:"type:varchar(256)"`
	ArtworkUrl   string        `gorm:"not null"`
	year         uint8         `gorm:"not null"`
	Tracks       []Track       `gorm:"foreignKey:Id;references:Id"`
	AlbumArtists []AlbumArtist `gorm:"foreignKey:AlbumId;references:Id"`
}
