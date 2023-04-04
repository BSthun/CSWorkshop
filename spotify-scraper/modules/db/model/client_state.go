package model

import "time"

type ClientState struct {
	Id        *uint64    `gorm:"primaryKey"`
	ClientId  *uint64    `gorm:"not null"`
	Client    *Client    `gorm:"foreignKey:ClientId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Email     *string    `gorm:"type:varchar(255); index:email,unique; not null"`
	State     *string    `gorm:"type:varchar(255); index:state,unique; null"`
	CreatedAt *time.Time `gorm:"not null"`
	UpdatedAt *time.Time `gorm:"not null"`
}
