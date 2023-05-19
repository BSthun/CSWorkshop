package model

type FeedItem struct {
	Id         *uint64   `gorm:"primaryKey"`
	KindId     *uint64   `gorm:"not null"`
	Kind       *FeedKind `gorm:"foreignKey:KindId"`
	PlaylistId *uint64   `gorm:"not null"`
	Order      *int64    `gorm:"not null"`
}
