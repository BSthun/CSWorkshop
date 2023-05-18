package branchModel

import "time"

type Lv3Activity struct {
	Timestamp    *time.Time `gorm:"primaryKey"`
	TrackId      *uint64    `gorm:"not null"`
	Track        *Lv3Track  `gorm:"foreignKey:TrackId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ShuffleState *bool      `gorm:"not null"`
	RepeatState  *string    `gorm:"type:enum('off','track','context'); not null"`
	DeviceName   *string    `gorm:"type:VARCHAR(255); not null"`
	Context      *string    `gorm:"type:VARCHAR(255); null"`
}
