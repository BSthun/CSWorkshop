package model

import "time"

type Concert struct {
	Id             *uint64         `gorm:"primaryKey"`
	ArtistId       *uint64         `gorm:"not null"`
	Artist         *Artist         `gorm:"foreignKey:ArtistId"`
	CountryId      *uint64         `gorm:"not null"`
	Country        *Country        `gorm:"foreignKey:CountryId"`
	Location       *string         `gorm:"not null"`
	Date           *string         `gorm:"not null"`
	ConcertTickets []ConcertTicket `gorm:"foreignKey:ConcertId; references:Id"`
	CreatedAt      *time.Time      `gorm:"not null"`
	UpdatedAt      *time.Time      `gorm:"not null"`
}
