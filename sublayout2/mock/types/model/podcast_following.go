package model

type PodcastFollowing struct {
	UserId *uint64      `gorm:"primaryKey"`
	ShowId *uint64      `gorm:"primaryKey"`
	User   *User        `gorm:"foreignKey:UserId"`
	Show   *PodcastShow `gorm:"foreignKey:ShowId"`
}
