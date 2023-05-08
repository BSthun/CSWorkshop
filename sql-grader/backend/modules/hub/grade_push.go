package ihub

import (
	"github.com/sirupsen/logrus"

	"backend/functions/enroll"
	"backend/types/extern"
	"backend/types/model"
	"backend/types/payload"
	"backend/types/response"
	"backend/utils/value"
)

func GradePush(session *extern.Session, submission *model.Submission, actualColumnNames []string, actualResults [][]string, expectedColumnNames []string, expectedResult [][]string) {
	// * Check session taskId
	if *session.CurrentTask != *submission.TaskId {
		return
	}

	// * Query tasks
	tasks, err := enroll.QueryTasks(submission.EnrollmentId)
	if err != nil {
		logrus.Warn(err)
		return
	}

	// * Map tasks
	mappedTasks, _ := value.Iterate(tasks, func(task *model.TaskPassed) (*payload.TaskList, *response.ErrorInstance) {
		return &payload.TaskList{
			Id:     task.Task.Id,
			Title:  task.Task.Title,
			Passed: task.Passed,
		}, nil
	})

	// * Construct payload
	var tags []map[string]any
	for k, v := range *submission.Task.Tags {
		tags = append(tags, map[string]any{
			"key":   k,
			"value": v,
		})
	}

	// * Construct result
	result := &payload.LabStateResult{
		ExpectedHeader: expectedColumnNames,
		ExpectedRows:   expectedResult,
		ActualHeader:   actualColumnNames,
		ActualRows:     actualResults,
	}
	session.TaskResults[*submission.TaskId] = result

	session.Emit(&extern.OutboundMessage{
		Event: extern.LabStateEvent,
		Payload: &payload.LabState{
			DbValid:         session.DbValid,
			TaskTitle:       submission.Task.Title,
			TaskDescription: submission.Task.Description,
			TaskTags:        tags,
			Query:           submission.Query,
			QueryPassed:     submission.Passed,
			QueryError:      nil,
			Result:          result,
			Tasks:           mappedTasks,
		},
	})

}
