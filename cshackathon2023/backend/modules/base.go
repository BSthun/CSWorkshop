package modules

import (
	hubModel "backend/modules/hub/model"
	"database/sql"

	"firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"backend/modules/config"
)

var Conf *iconfig.Config
var Hub *hubModel.Hub
var SqlDB *sql.DB
var DB *gorm.DB
var FirebaseApp *firebase.App
var FirebaseAuth *auth.Client
var Fiber *fiber.App
