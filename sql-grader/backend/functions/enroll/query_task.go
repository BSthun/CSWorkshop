package enroll

import (
	"backend/modules"
	"backend/types/model"
)

func QueryTasks(enrollmentId *uint64) ([]*model.TaskPassed, error) {
	var tasks []*model.TaskPassed
	if result := modules.DB.
		Model(new(model.Task)).
		Joins("INNER JOIN enrollments ON enrollments.id = ? AND enrollments.lab_id = tasks.lab_id", enrollmentId).
		Select("tasks.*, COALESCE((SELECT 1 FROM submissions WHERE submissions.enrollment_id = enrollments.id AND submissions.task_id = tasks.id AND submissions.passed = 1 LIMIT 1), 0) AS passed").
		Find(&tasks); result.Error != nil {
		return nil, result.Error
	}

	return tasks, nil
}
