package profile

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	"backend/functions/profile"
	"backend/types/common"
	"backend/types/payload"
	"backend/types/response"
)

func EnrollLabPostHandler(c *fiber.Ctx) error {
	// * Parse user
	u := c.Locals("u").(*jwt.Token).Claims.(*common.UserClaims)

	// * Parse request body
	var body *payload.ProfileEnrollmentRequest
	if err := c.BodyParser(&body); err != nil {
		return response.Error(c, false, "Unable to parse request body", err)
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

	if err := profile.ActEnrollLab(user, lab, body.Dump); err != nil {
		return response.Error(c, false, "Unable to enroll lab", err)
	}

	return c.JSON(response.Success(c, "Successfully enrolled the lab"))
}
