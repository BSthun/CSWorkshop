package ihub

import (
	"reflect"

	"github.com/sirupsen/logrus"

	"backend/types/model"
	"backend/utils/value"
)

func GradePasser(session *Session, submission *model.Submission) {
	// * Query task
	if result := b.DB.First(&submission.Task, submission.TaskId); result.Error != nil {
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
	for actualRows.Next() {
		actualPointers := make([]any, len(actualColumnNames))
		actualContainer := make([]string, len(actualColumnNames))
		for i, _ := range actualPointers {
			actualPointers[i] = &actualContainer[i]
		}
		if err := actualRows.Scan(actualPointers...); err != nil {
			logrus.Warn(err)
			return
		}
		actualResults = append(actualResults, actualContainer)
	}

	// * Query expected results
	expectedRows, err := session.Db.Query(*submission.Task.Query)
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
	for expectedRows.Next() {
		expectedPointers := make([]any, len(expectedColumnNames))
		expectedContainer := make([]string, len(expectedColumnNames))
		for i, _ := range expectedPointers {
			expectedPointers[i] = &expectedContainer[i]
		}
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
		if err := b.DB.Model(submission).Where("enrollment_id = ? AND task_id = ? AND query = ?", submission.EnrollmentId, submission.TaskId, submission.Query).Update("passed", true).Error; err != nil {
			logrus.Warn(err)
		}
	}

	// * Notify user
	GradePush(session, submission, actualColumnNames, actualResults, expectedColumnNames, expectedResults)
}
