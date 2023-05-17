package model

type Artist struct {
	Id         *uint64 `gorm:"primaryKey"`
	Name       *string `gorm:"type:VARCHAR(256)"`
	Biography  *string `gorm:"type:TEXT"`
	Href       *string `gorm:"not null"`
	ProfileUrl *string `gorm:"not null"`
	BannerUrl  *string `gorm:"not null"`
}
