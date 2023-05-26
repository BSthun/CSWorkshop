package model

type Log struct {
	UserId    *uint64         `gorm:"not null"`
	User      *User           `gorm:"foreignKey:UserId"`
	Type      *string         `gorm:"type:VARCHAR(255); not null"`
	Payload   *map[string]any `gorm:"type:TEXT; not null"`
	CreatedAt *string         `gorm:"not null"` // Embedded field
	UpdatedAt *string         `gorm:"not null"` // Embedded field
}
