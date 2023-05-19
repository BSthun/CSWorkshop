package model

type TrackArtist struct {
	TrackId  *uint64 `gorm:"primaryKey"`
	ArtistId *uint64 `gorm:"primaryKey"`
	Track    *Track  `gorm:"foreignKey:TrackId"`
	Artist   *Artist `gorm:"foreignKey:ArtistId"`
}
