package music

import (
	"backend/types/response"
	"github.com/gofiber/fiber/v2"
)

func StateGetHandler(c *fiber.Ctx) error {
	return c.JSON(response.Success(c, 1))
}
