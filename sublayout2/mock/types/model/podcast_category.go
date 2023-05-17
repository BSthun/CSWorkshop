package model

import "time"

type PodcastCategory struct {
	Id        *uint64         `gorm:"primaryKey"`
	SectionId *uint64         `gorm:"not null"`
	Section   *PodcastSection `gorm:"foreignKey:SectionId"`
	Name      *string         `gorm:"type:VARCHAR(256); not null"`
	Featured  *bool           `gorm:"not null"`
	CoverUrl  *string         `gorm:"not null"`
	CreatedAt *time.Time      `gorm:"not null"`
	UpdatedAt *time.Time      `gorm:"not null"`
}
