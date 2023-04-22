package lab

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	"backend/functions/enroll"
	"backend/types/common"
	"backend/types/payload"
	"backend/types/response"
	"backend/utils/text"
)

func InfoGetHandler(c *fiber.Ctx) error {
	// * Parse user
	u := c.Locals("u").(*jwt.Token).Claims.(*common.UserClaims)

	// * Parse query
	query := new(payload.EnrollInfoGetQuery)
	if err := c.QueryParser(query); err != nil {
		return response.Error(c, false, "Unable to parse query", err)
	}

	// * Validate query
	if err := text.Validator.Struct(query); err != nil {
		return response.Error(c, false, "Unable to validate query", err)
	}

	// * Get enrollment info
	enrollment, err := enroll.GetEnrollment(query.EnrollmentId, u.UserId)
	if err != nil {
		return response.Error(c, false, "Unable to get enrollment info", err)
	}

	// * Get tasks info
	tasks, err := enroll.QueryTasks(enrollment.Id)
	if err != nil {
		return response.Error(c, false, "Unable to get tasks info", err)
	}

	// * Get enrollment session
	session, err := enroll.ActLoadEnrollmentSession(enrollment)
	if err != nil {
		return response.Error(c, false, "Unable to get enrollment session", err)
	}

	return c.JSON(response.Success(c, enroll.MapEnrollmentTask(enrollment, tasks, session)))
}
