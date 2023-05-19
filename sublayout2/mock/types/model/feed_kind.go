package model

type FeedKind struct {
	Id     *uint64 `gorm:"primaryKey"`
	UserId *uint64 `gorm:"not null"`
	User   *User   `gorm:"foreignKey:UserId"`
	Name   *string `gorm:"type:VARCHAR(256); not null"`
}
