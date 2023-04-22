package admin

import (
	"gorm.io/gorm"

	"backend/modules"
	"backend/types/model"
	"backend/types/payload"
	"backend/utils/text"
)

func ImportTask(labId *uint64, imp *payload.AdminTaskImport) error {
	// * Validate task import
	if err := text.Validator.Struct(imp); err != nil {
		return err
	}

	// * Construct task record
	task := &model.Task{
		Id:          nil,
		Code:        imp.Code,
		LabId:       labId,
		Lab:         nil,
		Title:       imp.Title,
		Description: imp.Description,
		Tags:        imp.Tags,
		Query:       imp.Query,
		Hint:        imp.Hint,
		CreatedAt:   nil,
		UpdatedAt:   nil,
	}

	// * Check task exist
	var currentTask *model.Task
	if result := modules.DB.Where("code = ?", task.Code).First(&currentTask); result.Error == nil {
		// # Case of task already exist
		// * Update task
		// task.Id = currentTask.Id
		// if result := modules.DB.Save(task); result.Error != nil {
		// 	return result.Error
		// }
	} else if result.Error == gorm.ErrRecordNotFound {
		// # Case of task not exist
		// * Create task
		if result := modules.DB.Create(task); result.Error != nil {
			return result.Error
		}
	} else {
		return result.Error
	}

	return nil
}
