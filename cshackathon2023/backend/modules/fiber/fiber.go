package ifiber

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"backend/endpoints"
	"backend/modules/fiber/middleware"
	"backend/modules/fiber/websocket"
	"backend/utils/text"
)

func Init() *fiber.App {
	// Initialize fiber instance
	app := fiber.New(fiber.Config{
		ErrorHandler:  ErrorHandler,
		AppName:       "CS Hackathon Panel",
		Prefork:       false,
		StrictRouting: true,
		ReadTimeout:   5 * time.Second,
		WriteTimeout:  5 * time.Second,
	})

	// Init middlewares
	// app.Use(middleware.Limiter())
	app.Use(middleware.Cors())
	app.Use(middleware.Sentry())

	// Init API endpoints
	apiGroup := app.Group("api/")
	endpoints.Init(apiGroup)

	// Init WebSocket endpoints
	websocketGroup := app.Group("ws/")
	websocket.Register(websocketGroup)

	// Init static files
	app.Static("/", text.RelativePath("resources/web"))
	app.Static("/image", text.RelativePath("resources/image"))

	// Init not found handler
	app.Use(NotFoundHandler)

	return app
}
