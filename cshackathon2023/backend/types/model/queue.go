package model

type Queue struct {
	Id        *uint64 `gorm:"primaryKey"`
	UserId    *uint64 `gorm:"not null"`
	User      *User   `gorm:"foreignKey:UserId"`
	Track     *Track  `gorm:"foreignKey:TrackId"`
	TrackId   *uint64 `gorm:"not null"`
	Played    *bool   `gorm:"not null; default:false"`
	CreatedAt *string `gorm:"not null"` // Embedded field
	UpdatedAt *string `gorm:"not null"` // Embedded field
}
