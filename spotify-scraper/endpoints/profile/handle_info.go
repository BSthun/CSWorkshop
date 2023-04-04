package profile

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	"backend/modules/db/branchModel"
	"backend/modules/hub"
	"backend/types/common"
	"backend/types/response"
)

func InfoHandler(c *fiber.Ctx) error {
	// * Parse user claims
	u := c.Locals("u").(*jwt.Token).Claims.(*common.UserClaims)

	// * Get branch
	branch, ok := hub.Hub.Branches[*u.UserId]
	if !ok {
		return response.Error(c, true, "Unable to get branch")
	}

	// * Count minutes
	var minuteCount int64
	if result := branch.DB.Model(new(branchModel.Lv3Activity)).Count(&minuteCount); result.Error != nil {
		return response.Error(c, true, "Unable to count minutes", result.Error)
	}

	var profileImg *string
	if len(branch.Profile.Profile.Images) > 0 {
		profileImg = branch.Profile.Profile.Images[0].Url
	}

	return c.JSON(response.Info(c, map[string]any{
		"profile_img": profileImg,
		"username":    branch.Profile.Profile.DisplayName,
		"minutes":     minuteCount,
	}))
}
