package model

import "time"

type Queue struct {
	Id        *uint64    `gorm:"primaryKey"`
	UserId    *uint64    `gorm:"not null"`
	User      *User      `gorm:"foreignKey:UserId"`
	TrackId   *uint64    `gorm:"not null"`
	Track     *Track     `gorm:"foreignKey:TrackId"`
	Played    *bool      `gorm:"not null; default:false"`
	CreatedAt *time.Time `gorm:"not null"` // Embedded field
	UpdatedAt *time.Time `gorm:"not null"` // Embedded field
	DeletedAt *time.Time `gorm:"index"`    // Embedded field
}
