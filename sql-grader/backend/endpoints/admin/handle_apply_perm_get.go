package admin

import (
	"github.com/gofiber/fiber/v2"

	"backend/functions/admin"
	"backend/types/response"
)

func ApplyPermGetHandler(c *fiber.Ctx) error {
	if err := admin.ApplyUserPerm(); err != nil {
		return response.Error(c, false, "Unable to apply user perm", err)
	}

	return c.JSON(response.Success(c, "Successfully applied user perm"))
}
