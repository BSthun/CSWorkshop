package model

type TrackLyric struct {
	TrackId *uint64 `gorm:"unique"`
	Track   *Track  `gorm:"foreignKey:TrackId"`
	Lyric   *string `gorm:"not null"`
	Order   *uint64 `gorm:"not null"`
}
