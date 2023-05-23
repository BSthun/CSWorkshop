package model

type SearchHistory struct {
	UserId    *string `gorm:"not null"`
	User      *User   `gorm:"foreignKey:UserId"`
	Query     *string `gorm:"type:VARCHAR(255); not null"`
	CreatedAt *string `gorm:"not null"` // Embedded field
}
