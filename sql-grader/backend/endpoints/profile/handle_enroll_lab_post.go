package profile

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	"backend/functions/enroll"
	"backend/functions/profile"
	"backend/types/common"
	"backend/types/payload"
	"backend/types/response"
	"backend/utils/text"
)

func EnrollLabPostHandler(c *fiber.Ctx) error {
	// * Parse user
	u := c.Locals("u").(*jwt.Token).Claims.(*common.UserClaims)

	// * Parse body
	var body *payload.ProfileEnrollmentRequest
	if err := c.BodyParser(&body); err != nil {
		return response.Error(c, false, "Unable to parse body", err)
	}

	// * Validate body
	if err := text.Validator.Struct(body); err != nil {
		return response.Error(c, false, "Unable to validate body", err)
	}

	// * Get user record
	user, err := profile.GetUser(u.UserId)
	if err != nil {
		return response.Error(c, false, "Unable to get user", err)
	}

	// * Get lab record
	lab, err := profile.GetLab(body.LabId)
	if err != nil {
		return response.Error(c, false, "Unable to get lab", err)
	}

	_, err = enroll.ActEnrollLab(user, lab)
	if err != nil {
		return response.Error(c, false, "Unable to enroll lab", err)
	}

	return c.JSON(response.Success(c, "Successfully enrolled lab"))
}
