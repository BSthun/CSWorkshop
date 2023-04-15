package model

import "time"

type Lab struct {
	Id          *uint64    `gorm:"primaryKey"`
	Code        *string    `gorm:"type:VARCHAR(255); index:idx_code,unique; not null"`
	Name        *string    `gorm:"type:VARCHAR(255); not null"`
	Description *string    `gorm:"type:TEXT; not null"`
	TemplateDb  *string    `gorm:"type:VARCHAR(255); not null"`
	CreatedAt   *time.Time `gorm:"not null"` // Embedded field
	UpdatedAt   *time.Time `gorm:"not null"` // Embedded field
}
