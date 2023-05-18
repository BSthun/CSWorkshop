package model

type Album struct {
	Id           *uint64        `gorm:"primaryKey"`
	Name         *string        `gorm:"type:varchar(256)"`
	ArtworkUrl   *string        `gorm:"not null"`
	Year         *uint16        `gorm:"not null"`
	Tracks       []*Track       `gorm:"foreignKey:Id; references:Id"`
	AlbumArtists []*AlbumArtist `gorm:"foreignKey:AlbumId; references:Id"`
}
