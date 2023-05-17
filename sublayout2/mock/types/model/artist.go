package model

type Artist struct {
	Id         *uint64 `gorm:"primaryKey;autoIncrement"`
	Name       string  `gorm:"type:varchar(256)"`
	Href       string  `gorm:"not null"`
	profileUrl string  `gorm:"not null"`
	bannerUrl  string  `gorm:"not null"`
}
