package enroll

import (
	"backend/modules"
	"backend/types/model"
)

func QueryTasks(labId *uint64) ([]*model.Task, error) {
	var tasks []*model.Task
	if result := modules.DB.
		Where(
			"lab_id = ?",
			labId,
		).
		Find(&tasks); result.Error != nil {
		return nil, result.Error
	}

	return tasks, nil
}
