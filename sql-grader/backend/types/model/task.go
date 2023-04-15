package model

import (
	"time"

	"backend/types/embed"
)

type Task struct {
	Code        *string         `gorm:"TYPE:VARCHAR(255); primaryKey"`
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
