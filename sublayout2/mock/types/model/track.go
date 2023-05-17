package model

type Track struct {
	Id             *uint64         `gorm:"primaryKey;autoIncrement"`
	Name           string          `gorm:"type:varchar(256)"`
	AlbumId        *uint64         `gorm:"not null"`
	Duration       string          `gorm:"type:varchar(256)"`
	Popularity     *uint64         `gorm:"default:0"`
	Explicit       bool            `gorm:"default:false"`
	PreviewUrl     string          `gorm:"not null"`
	PlaylistTracks []PlaylistTrack `gorm:"foreignKey:TrackId;references:Id"`
}
