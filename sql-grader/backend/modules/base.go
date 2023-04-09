package modules

import (
	"firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"gorm.io/gorm"

	"backend/modules/config"
)

var B *Base

type Base struct {
	Conf         *iconfig.Config
	DB           *gorm.DB
	FirebaseApp  *firebase.App
	FirebaseAuth *auth.Client
}

func NewBase() *Base {
	base := &Base{
		Conf:         nil,
		DB:           nil,
		FirebaseApp:  nil,
		FirebaseAuth: nil,
	}

	B = base

	return base
}

func (b *Base) Clone() *Base {
	return &Base{
		Conf:         b.Conf,
		DB:           b.DB,
		FirebaseApp:  b.FirebaseApp,
		FirebaseAuth: b.FirebaseAuth,
	}
}
