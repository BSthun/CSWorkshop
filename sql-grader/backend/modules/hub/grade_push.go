package ihub

import (
	"backend/types/extern"
	"backend/types/model"
	"backend/types/payload"
)

func GradePush(session *Session, submission *model.Submission, actualColumnNames []string, actualResults [][]string, expectedColumnNames []string, expectedResult [][]string) {
	// * Check session taskId
	if *session.CurrentTask != *submission.TaskId {
		return
	}

	// * Construct payload
	var tags []map[string]any
	for k, v := range *submission.Task.Tags {
		tags = append(tags, map[string]any{
			"key":   k,
			"value": v,
		})
	}
	session.Emit(&extern.OutboundMessage{
		Event: extern.LabStateEvent,
		Payload: &payload.LabState{
			TaskTitle:       submission.Task.Title,
			TaskDescription: submission.Task.Description,
			TaskTags:        tags,
			Query:           submission.Query,
			QueryPassed:     submission.Passed,
			QueryError:      nil,
			Result: &payload.LabStateResult{
				ExpectedHeader: expectedColumnNames,
				ExpectedRows:   expectedResult,
				ActualHeader:   actualColumnNames,
				ActualRows:     actualResults,
			},
		},
	})

}
