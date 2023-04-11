package profile

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	"backend/functions/profile"
	"backend/types/common"
	"backend/types/response"
)

func StateGetHandler(c *fiber.Ctx) error {
	// * Parse user
	u := c.Locals("u").(*jwt.Token).Claims.(*common.UserClaims)

	// * Fetch user information
	user, err := profile.GetProfile(u.UserId)
	if err != nil {
		return response.Error(c, true, "Unable to get profile", err)
	}

	return c.JSON(response.Success(c, map[string]any{
		"profile": profile.MapProfile(user),
	}))
}
