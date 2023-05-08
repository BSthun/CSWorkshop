package lab

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	"backend/functions/lab"
	"backend/modules"
	"backend/types/common"
	"backend/types/extern"
	"backend/types/model"
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

	// * Check session ownership
	if *session.UserId != *u.UserId {
		return response.Error(c, false, "Unknown session user", nil)
	}

	// * Match enrollment session
	if *task.Lab.Id != *session.LabId {
		return response.Error(c, false, "Unable to match enrollment session", nil)
	}

	// * Check db validity
	if !*session.DbValid {
		return response.Error(c, false, "Database is invalid", nil)
	}

	// * Set current task
	var previousTask uint64
	if session.CurrentTask != nil {
		previousTask = *session.CurrentTask
	}
	session.CurrentTask = query.TaskId

	// * Query submission
	var submission *model.Submission
	if result := modules.DB.Where("enrollment_id = ? AND task_id = ?", *session.Id, *task.Id).Order("event_time DESC").First(&submission); result.Error != nil {
		submission = new(model.Submission)
	}

	var tags []map[string]any
	for k, v := range *task.Tags {
		tags = append(tags, map[string]any{
			"key":   k,
			"value": v,
		})
	}

	// * Get cached result
	var result *payload.LabStateResult
	if e, ok := session.TaskResults[*task.Id]; ok {
		result = e
	}

	if previousTask != *session.CurrentTask {
		session.Emit(&extern.OutboundMessage{
			Event: extern.LabStateEvent,
			Payload: &payload.LabState{
				DbValid:         session.DbValid,
				TaskTitle:       task.Title,
				TaskDescription: task.Description,
				TaskTags:        tags,
				Query:           submission.Query,
				QueryPassed:     submission.Passed,
				QueryError:      nil,
				Result:          result,
				Tasks:           nil,
			},
		})
	}

	return c.JSON(response.Success(c, map[string]any{}))
}
