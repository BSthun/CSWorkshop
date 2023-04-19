package lab

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	"backend/functions/lab"
	"backend/modules"
	"backend/types/common"
	"backend/types/payload"
	"backend/types/response"
	"backend/utils/text"
)

func TaskClickGetHandler(c *fiber.Ctx) error {
	// * Parse user
	u := c.Locals("u").(*jwt.Token).Claims.(*common.UserClaims)

	// * Parse query
	query := new(payload.TaskClickQuery)
	if err := c.QueryParser(query); err != nil {
		return response.Error(c, false, "Unable to parse query", err)
	}

	// * Validate query
	if err := text.Validator.Struct(query); err != nil {
		return response.Error(c, false, "Unable to validate query", err)
	}

	// * Fetch task info
	task, err := lab.QueryTask(query.TaskId)
	if err != nil {
		return err
	}

	// * Get enrollment session
	session, ok := modules.Hub.Sessions[*query.EnrollmentId]
	if !ok {
		return response.Error(c, false, "Unable to find enrollment session", nil)
	}

	// * Match enrollment session
	if *task.Lab.Id != *session.LabId {
		return response.Error(c, false, "Unable to match enrollment session", nil)
	}

	// * Set current task
	session.CurrentTask = query.TaskId

	_ = u

	return c.JSON(response.Success(c, map[string]any{}))
}
