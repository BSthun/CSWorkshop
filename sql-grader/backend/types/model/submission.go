package model

import "time"

type Submission struct {
	UserId    *uint64    `gorm:"not null"`
	User      *User      `gorm:"foreignKey:UserId"`
	TaskId    *uint64    `gorm:"not null"`
	Task      *Task      `gorm:"foreignKey:TaskId"`
	Sql       *string    `gorm:"type:TEXT; not null"`
	Passed    *bool      `gorm:"not null"`
	CreatedAt *time.Time `gorm:"not null"` // Embedded field
	UpdatedAt *time.Time `gorm:"not null"` // Embedded field
}
