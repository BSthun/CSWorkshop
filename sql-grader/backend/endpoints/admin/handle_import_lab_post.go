package admin

import (
	"github.com/gofiber/fiber/v2"

	"backend/functions/admin"
	"backend/types/payload"
	"backend/types/response"
)

func ImportLabPostHandler(c *fiber.Ctx) error {
	var body *payload.AdminLabImport
	if err := c.BodyParser(&body); err != nil {
		return response.Error(c, false, "Unable to parse body", err)
	}

	if err := admin.ImportLab(body); err != nil {
		return response.Error(c, false, "Unable to import lab", err)
	}

	return c.JSON(response.Success(c, "Successfully imported lab"))
}
