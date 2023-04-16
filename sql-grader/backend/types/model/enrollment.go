package model

import "time"

type Enrollment struct {
	Id        *uint64    `gorm:"primaryKey"`
	UserId    *uint64    `gorm:"not null"`
	User      *User      `gorm:"foreignKey:UserId"`
	LabId     *uint64    `gorm:"not null"`
	Lab       *Lab       `gorm:"foreignKey:LabId"`
	DbName    *string    `gorm:"type:VARCHAR(255); not null"`
	DbValid   *bool      `gorm:"not null"`
	CreatedAt *time.Time `gorm:"not null"` // Embedded field
	UpdatedAt *time.Time `gorm:"not null"` // Embedded field
}

const (
	EnrollmentFieldId     = "id"
	EnrollmentFieldUserId = "user_id"
	EnrollmentFieldLabId  = "lab_id"
	EnrollmentFieldDb     = "db"
)
