package model

type ArtistListener struct {
	ArtistId  *uint64  `gorm:"primaryKey"`
	Artist    *Artist  `gorm:"foreignKey:ArtistId"`
	CountryId *uint64  `gorm:"primaryKey"`
	Country   *Country `gorm:"foreignKey:CountryId"`
	Count     *int64   `gorm:"not null"`
}
