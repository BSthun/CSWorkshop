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

	if err := b.DB.Create(submission); err != nil {
		logrus.Warn(err)
	}

	return submission
}
