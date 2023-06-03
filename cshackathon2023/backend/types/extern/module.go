package extern

import (
	"database/sql"

	"firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Base struct {
	SqlDB        *sql.DB
	DB           *gorm.DB
	FirebaseApp  *firebase.App
	FirebaseAuth *auth.Client
	Fiber        *fiber.App
}
