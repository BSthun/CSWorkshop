package fiber

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"backend/endpoints"
	"backend/modules"

	"backend/modules/fiber/middlewares"
)

func Init(b *modules.Base) {
	// Initialize fiber instance
	app := fiber.New(fiber.Config{
		ErrorHandler:  ErrorHandler,
		Prefork:       false,
		StrictRouting: true,
		ReadTimeout:   5 * time.Second,
		WriteTimeout:  5 * time.Second,
	})

	// Init middlewares
	app.Use(middlewares.Limiter(b))
	app.Use(middlewares.Cors(b))
	app.Use(middlewares.Sentry(b))

	// Init API endpoints
	apiGroup := app.Group("api/")
	endpoints.Init(apiGroup)

	// Init static files
	app.Static("/", "./web/")

	// Init not found handler
	app.Use(NotFoundHandler)

	// Startup
	err := app.Listen(b.Conf.Address)
	if err != nil {
		logrus.Fatal(err.Error())
	}
}
