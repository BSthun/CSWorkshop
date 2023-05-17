package model

type Country struct {
	Id       *uint64    `gorm:"primaryKey"`
	Name     *string    `gorm:"unique"`
	Users    []*User    `gorm:"foreignKey:Id;references:Id"`
	Concerts []*Concert `gorm:"foreignKey:Id;references:Id"`
}
