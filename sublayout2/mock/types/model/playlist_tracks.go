package model

type PlaylistTrack struct {
	PlaylistId *uint64 `gorm:"primaryKey;"`
	TrackId    *uint64 `gorm:"primaryKey;"`
	Tracks     *Track  `gorm:"foreignKey:TrackId"`
}
