package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"backend/modules"
)

func Cors() fiber.Handler {
	// origins is the value of allowed CORS addresses, separated by comma (,).
	// Example: "https://www.google.com, https://www.bsthun.com, http://localhost:8080"
	origins := ""
	for i, s := range modules.Conf.Cors {
		origins += s
		if i < len(modules.Conf.Cors)-1 {
			origins += ", "
		}
	}

	c := cors.Config{
		AllowOrigins:     origins,
		AllowCredentials: true,
	}

	return cors.New(c)
}
