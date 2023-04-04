package fiber

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"backend/endpoints"
	"backend/modules/config"
	"backend/modules/fiber/middlewares"
)

func Init() {
	// Initialize fiber instance
	app := fiber.New(fiber.Config{
		ErrorHandler:  ErrorHandler,
		Prefork:       false,
		StrictRouting: true,
		ReadTimeout:   5 * time.Second,
		WriteTimeout:  5 * time.Second,
	})

	// Init middlewares
	app.Use(middlewares.Limiter())
	app.Use(middlewares.Cors())
	app.Use(middlewares.Sentry())

	// Init API endpoints
	apiGroup := app.Group("api/")
	endpoints.Init(apiGroup)

	// Init static files
	app.Static("/", "./web/")

	// Init not found handler
	app.Use(NotFoundHandler)

	// Startup
	err := app.Listen(config.C.Address)
	if err != nil {
		logrus.Fatal(err.Error())
	}
}
