package model

type AlbumArtist struct {
	AlbumId  *uint64 `gorm:"primaryKey"`
	ArtistId *uint64 `gorm:"primaryKey"`
	Artists  Artist  `gorm:"foreignKey:Id;references:ArtistId"`
	Album    Album   `gorm:"foreignKey:Id;references:AlbumId"`
}
