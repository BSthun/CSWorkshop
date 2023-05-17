package model

import "time"

type Concert struct {
	Id             *uint64         `gorm:"primaryKey"`
	ArtistId       *uint64         `gorm:"not null"`
	Name           string          `gorm:"type:VARCHAR(256)"`
	CountryId      *uint64         `gorm:"not null"`
	Location       *uint64         `gorm:"not null"`
	Time           *time.Time      `gorm:"not null"`
	ConcertTickets []ConcertTicket `gorm:"foreignKey:ConcertId; references:Id"`
}
