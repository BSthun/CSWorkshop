package modules

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"mock/modules/config"
)

var Conf *iconfig.Config
var DB *gorm.DB
var Fiber *fiber.App
