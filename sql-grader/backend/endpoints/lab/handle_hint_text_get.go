package lab

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	"backend/functions/lab"
	"backend/modules"
	"backend/types/common"
	"backend/types/model"
	"backend/types/payload"
	"backend/types/response"
	"backend/utils/text"
	"backend/utils/value"
)

func HandleHintTextGet(c *fiber.Ctx) error {
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
		return response.Error(c, false, "Unable to fetch task info", err)
	}

	// * Get enrollment session
	session, ok := modules.Hub.Sessions[*query.EnrollmentId]
	if !ok {
		return response.Error(c, false, "Unable to find enrollment session", nil)
	}

	// * Check session ownership
	if *session.UserId != *u.UserId {
		return response.Error(c, false, "Unknown session user", nil)
	}

	// * Match enrollment session
	if *task.Lab.Id != *session.LabId {
		return response.Error(c, false, "Unable to match enrollment session", nil)
	}

	// * Add log
	log := &model.Log{
		UserId:    u.UserId,
		User:      nil,
		Type:      value.Ptr("hint_text"),
		Attrs:     value.Ptr(fmt.Sprintf("task_id=%d; enrollment_id=%d; lab_id=%d", *task.Id, *session.Id, *session.LabId)),
		CreatedAt: nil,
		UpdatedAt: nil,
	}
	if result := modules.DB.Create(log); result.Error != nil {
		return response.Error(c, false, "Unable to create log", result.Error)
	}

	// * Return response
	return c.JSON(response.Success(c, map[string]any{
		"hint_text": task.Hint,
	}))
}
