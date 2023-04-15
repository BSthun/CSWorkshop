package modules

import (
	"database/sql"

	"firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"backend/modules/config"
	"backend/modules/hub"
)

var Conf *iconfig.Config
var Hub *ihub.Hub
var SqlDB *sql.DB
var DB *gorm.DB
var FirebaseApp *firebase.App
var FirebaseAuth *auth.Client
var Fiber *fiber.App

type Base struct {
	Conf         *iconfig.Config
	Hub          *ihub.Hub
	DB           *gorm.DB
	FirebaseApp  *firebase.App
	FirebaseAuth *auth.Client
}

func NewBase() *Base {
	base := &Base{
		Conf:         Conf,
		DB:           DB,
		FirebaseApp:  FirebaseApp,
		FirebaseAuth: FirebaseAuth,
	}

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
