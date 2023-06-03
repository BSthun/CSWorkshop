package backdrop

import (
	"backend/functions/backdrop"
	"backend/types/response"
	"github.com/gofiber/fiber/v2"
)

func StateGetHandler(c *fiber.Ctx) error {
	state := backdrop.GetBackdrop()
	return c.JSON(response.Success(c, state))
}
