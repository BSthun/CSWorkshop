package model

import "time"

type Submission struct {
	EnrollmentId *uint64     `gorm:"not null"`
	Enrollment   *Enrollment `gorm:"foreignKey:EnrollmentId"`
	TaskId       *uint64     `gorm:"null"`
	Task         *Task       `gorm:"foreignKey:TaskId"`
	Query        *string     `gorm:"type:TEXT; not null"`
	Passed       *bool       `gorm:"not null"`
	EventTime    *time.Time  `gorm:"not null"`
	CreatedAt    *time.Time  `gorm:"not null"` // Embedded field
	UpdatedAt    *time.Time  `gorm:"not null"` // Embedded field
}
