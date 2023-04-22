package ihub

import (
	"time"

	"github.com/sirupsen/logrus"

	"backend/types/model"
	"backend/utils/value"
)

func GradeSubmit(session *Session, startTime time.Time, sqlText string) *model.Submission {
	// * Construct submission
	submission := &model.Submission{
		EnrollmentId: session.Id,
		Enrollment:   nil,
		TaskId:       session.CurrentTask,
		Task:         nil,
		Query:        &sqlText,
		Passed:       value.Ptr(false),
		EventTime:    &startTime,
		CreatedAt:    nil,
		UpdatedAt:    nil,
	}

	if result := b.DB.Create(submission); result.Error != nil {
		logrus.Warn("CREATE SUBMISSION", result.Error)
	}

	return submission
}
