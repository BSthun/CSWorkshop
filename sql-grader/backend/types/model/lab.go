package model

import "time"

type Lab struct {
	Id          *uint64    `json:"id"`
	Name        *string    `json:"name"`
	Description *string    `json:"description"`
	CreatedAt   *time.Time `gorm:"not null"` // Embedded field
	UpdatedAt   *time.Time `gorm:"not null"` // Embedded field
}
