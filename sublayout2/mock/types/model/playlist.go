package model

import "time"

type Playlist struct {
	Id          *uint64    `gorm:"primaryKey"`
	OwnerId     *uint64    `gorm:"null"`
	Owner       *User      `gorm:"foreignKey:OwnerId"`
	Name        *string    `gorm:"unique"`
	Description *string    `gorm:"not null"`
	CoverUrl    *string    `gorm:"not null"`
	CreatedAt   *time.Time `gorm:"not null"`
	UpdatedAt   *time.Time `gorm:"not null"`
}
