package model

type SectionItem struct {
	Id         *uint64      `gorm:"primaryKey"`
	KindId     *uint64      `gorm:"not null"`
	Kind       *SectionKind `gorm:"foreignKey:KindId"`
	PlaylistId *uint64      `gorm:"not null"`
	Playlist   *Playlist    `gorm:"foreignKey:PlaylistId"`
	Order      *int64       `gorm:"not null"`
}
