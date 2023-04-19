package ihub

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"

	"backend/types/model"
)

func GradePasser(session *Session, submission *model.Submission) {
	// * Query task
	var task *model.Task
	if result := b.DB.First(&task, submission.TaskId); result.Error != nil {
		logrus.Warn(result.Error)
		return
	}

	// * Query task solution
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
		spew.Dump(222)
		if err := actualRows.Scan(actualPointers...); err != nil {
			logrus.Warn(err)
			return
		}
		actualResults = append(actualResults, actualContainer)
	}

	spew.Dump(actualResults)
}
