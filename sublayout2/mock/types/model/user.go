package model

import "time"

type User struct {
	Id        *uint64    `gorm:"primaryKey"`
	Name      *string    `gorm:"unique"`
	CountryId *uint64    `gorm:"not null"`
	Country   *Country   `gorm:"foreignKey:CountryId"`
	CreatedAt *time.Time `gorm:"not null"`
	UpdatedAt *time.Time `gorm:"not null"`
}
