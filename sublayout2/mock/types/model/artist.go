package model

type Artist struct {
	Id         *uint64 `gorm:"primaryKey"`
	Name       *string `gorm:"type:VARCHAR(256)"`
	Href       *string `gorm:"not null"`
	profileUrl *string `gorm:"not null"`
	bannerUrl  *string `gorm:"not null"`
}
