package model

import (
	"time"

	"backend/types/payload"
)

type User struct {
	Id           *uint64                 `gorm:"primaryKey"`
	ClientId     *uint64                 `gorm:"not null"`
	Client       *Client                 `gorm:"foreignKey:ClientId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	SpotifyId    *string                 `gorm:"type:VARCHAR(255); index:spotify_id,unique; not null"`
	Profile      *payload.SpotifyProfile `gorm:"type:json;not null"`
	RefreshToken *string                 `gorm:"not null"`
	CreatedAt    *time.Time              `gorm:"not null"`
	UpdatedAt    *time.Time              `gorm:"not null"`
}
