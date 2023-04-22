package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"

	"backend/types/response"
)

var limiterReached = &response.ErrorInstance{
	Message: "Rate limit reached, try again in a few seconds.",
}

func Limiter() fiber.Handler {
	config := limiter.Config{
		Next:       nil,
		Max:        20,
		Expiration: 60 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Get("X-Forwarded-For", c.IP())
		},
		LimitReached: func(c *fiber.Ctx) error {
			return limiterReached
		},
	}

	return limiter.New(config)
}
