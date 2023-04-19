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
