package ihub

import (
	"reflect"

	"github.com/sirupsen/logrus"

	"backend/types/model"
	"backend/utils/value"
)

func GradePasser(session *Session, submission *model.Submission) {
	// * Query task
	var task *model.Task
	if result := b.DB.First(&task, submission.TaskId); result.Error != nil {
		logrus.Warn(result.Error)
		return
	}

	// * Query actual results
	actualRows, err := session.Db.Query(*submission.Query)
	if err != nil {
		logrus.Warn(err)
		return
	}

	actualColumnNames, err := actualRows.Columns()
	if err != nil {
		logrus.Warn(err)
		return
	}

	actualResults := make([][]string, 0)
	actualPointers := make([]any, len(actualColumnNames))
	actualContainer := make([]string, len(actualColumnNames))
	for i, _ := range actualPointers {
		actualPointers[i] = &actualContainer[i]
	}
	for actualRows.Next() {
		if err := actualRows.Scan(actualPointers...); err != nil {
			logrus.Warn(err)
			return
		}
		actualResults = append(actualResults, actualContainer)
	}

	// * Query expected results
	expectedRows, err := session.Db.Query(*task.Query)
	if err != nil {
		logrus.Warn(err)
		return
	}

	expectedColumnNames, err := expectedRows.Columns()
	if err != nil {
		logrus.Warn(err)
		return
	}

	expectedResults := make([][]string, 0)
	expectedPointers := make([]any, len(expectedColumnNames))
	expectedContainer := make([]string, len(expectedColumnNames))
	for i, _ := range expectedPointers {
		expectedPointers[i] = &expectedContainer[i]
	}
	for expectedRows.Next() {
		if err := expectedRows.Scan(expectedPointers...); err != nil {
			logrus.Warn(err)
			return
		}
		expectedResults = append(expectedResults, expectedContainer)
	}

	// * Compare results
	header := reflect.DeepEqual(actualColumnNames, expectedColumnNames)
	body := reflect.DeepEqual(actualResults, expectedResults)

	// * Update submission
	submission.Passed = value.Ptr(header && body)
	if *submission.Passed {
		if err := b.DB.Model(submission).Where("enrollment_id = ? AND task_id = ? AND event_time = ?", submission.EnrollmentId, submission.TaskId, submission.EventTime).Update("passed", true).Error; err != nil {
			logrus.Warn(err)
		}
	}
}
