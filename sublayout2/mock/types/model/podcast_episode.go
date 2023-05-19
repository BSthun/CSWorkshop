package model

type PodcastEpisode struct {
	Id          *uint64      `gorm:"primaryKey"`
	ShowId      *uint64      `gorm:"not null"`
	Title       *string      `gorm:"type:VARCHAR(256); not null"`
	Description *string      `gorm:"type:TEXT; not null"`
	Show        *PodcastShow `gorm:"foreignKey:ShowId; references:Id"`
}
