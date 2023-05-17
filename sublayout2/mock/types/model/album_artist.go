package model

type AlbumArtist struct {
	AlbumId  *uint64 `gorm:"primaryKey"`
	ArtistId *uint64 `gorm:"primaryKey"`
	Album    *Album  `gorm:"foreignKey:AlbumId"`
	Artist   *Artist `gorm:"foreignKey:ArtistId"`
}
