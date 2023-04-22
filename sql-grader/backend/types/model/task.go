package model

import (
	"time"

	"backend/types/embed"
)

type Task struct {
	Id          *uint64         `gorm:"primaryKey"`
	Code        *string         `gorm:"TYPE:VARCHAR(255); index:idx_code,unique; not null"`
	LabId       *uint64         `gorm:"not null"`
	Lab         *Lab            `gorm:"foreignKey:LabId"`
	Title       *string         `gorm:"type:VARCHAR(255); not null"`
	Description *string         `gorm:"type:TEXT; not null"`
	Tags        *embed.TaskTags `gorm:"type:JSON; not null"`
	Query       *string         `gorm:"type:TEXT; not null"`
	Hint        *string         `gorm:"type:TEXT; not null"`
	CreatedAt   *time.Time      `gorm:"not null"` // Embedded field
	UpdatedAt   *time.Time      `gorm:"not null"` // Embedded field
}

type TaskPassed struct {
	Task
	Passed *bool
}
