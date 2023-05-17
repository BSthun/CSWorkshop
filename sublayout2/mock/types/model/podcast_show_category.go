package model

type PodcastShowCategory struct {
	ShowId      *uint64          `gorm:"primaryKey"`
	CategoryId  *uint64          `gorm:"primaryKey"`
	PodcastShow *PodcastShow     `gorm:"foreignKey:ShowId"`
	Category    *PodcastCategory `gorm:"foreignKey:CategoryId"`
}
