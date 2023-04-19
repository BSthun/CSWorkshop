package lab

import (
	"backend/modules"
	"backend/types/model"
)

func QueryTask(taskId *uint64) (*model.Task, error) {
	var task *model.Task
	if result := modules.DB.Preload("Lab").First(&task, taskId); result.Error != nil {
		return nil, result.Error
	}
	return task, nil
}
