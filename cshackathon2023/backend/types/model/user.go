package model

import "time"

type User struct {
	Id          *uint64    `gorm:"primaryKey"`
	FirebaseUid *string    `gorm:"type:VARCHAR(255); index:idx_firebase_uid,unique; not null"`
	Email       *string    `gorm:"type:VARCHAR(255); not null"`
	Name        *string    `gorm:"type:VARCHAR(255); not null"`
	Avatar      *string    `gorm:"type:VARCHAR(255); not null"`
	CreatedAt   *time.Time `gorm:"not null"` // Embedded field
	UpdatedAt   *time.Time `gorm:"not null"` // Embedded field
}

const (
	UserFieldId          = "id"
	UserFieldFirebaseUid = "firebase_uid"
	UserFieldEmail       = "email"
	UserFieldName        = "name"
	UserFieldAvatar      = "avatar"
	UserFieldCredential  = "credential"
)
