package model

import "time"

type Log struct {
	UserId    *uint64    `gorm:"not null"`
	User      *User      `gorm:"foreignKey:UserId"`
	Type      *string    `gorm:"type:VARCHAR(255); not null"`
	Payload   *string    `gorm:"type:JSON; not null"`
	CreatedAt *time.Time `gorm:"not null"` // Embedded field
	UpdatedAt *time.Time `gorm:"not null"` // Embedded field
}
