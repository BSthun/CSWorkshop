package model

type UserTopArtist struct {
	UserId      *uint64 `gorm:"primaryKey"`
	User        *User   `gorm:"foreignKey:UserId"`
	ArtistId    *uint64 `gorm:"primaryKey"`
	Artist      *Artist `gorm:"foreignKey:ArtistId"`
	MinuteCount *int64  `gorm:"not null"`
}
