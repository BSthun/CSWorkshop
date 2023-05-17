package model

type PodcastShow struct {
	Id          *uint64 `gorm:"primaryKey"`
	Title       *string `gorm:"type:VARCHAR(256); not null"`
	Author      *string `gorm:"type:VARCHAR(256); not null"`
	Description *string `gorm:"type:TEXT; not null"`
}
