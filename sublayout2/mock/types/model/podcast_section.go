package model

type PodcastSection struct {
	Id   *uint64 `gorm:"primaryKey"`
	Name *string `gorm:"type:VARCHAR(256); not null"`
}
