package model

type Country struct {
	Id   *uint64 `gorm:"primaryKey"`
	Name *string `gorm:"unique"`
}
