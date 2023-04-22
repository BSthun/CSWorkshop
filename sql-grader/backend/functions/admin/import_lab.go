package admin

import (
	"gorm.io/gorm"

	"backend/modules"
	"backend/types/model"
	"backend/types/payload"
	"backend/utils/text"
	"backend/utils/value"
)

func ImportLab(imp *payload.AdminLabImport) error {
	// * Validate lab import
	if err := text.Validator.Struct(imp); err != nil {
		return err
	}

	// * Check template db existence
	if err := CheckTemplateDb(*imp.TemplateDb); err != nil {
		return err
	}

	// * Construct lab record
	lab := &model.Lab{
		Id:          nil,
		Code:        imp.Code,
		Name:        imp.Name,
		Description: imp.Description,
		TemplateDb:  imp.TemplateDb,
		Generator:   imp.Generator,
		CreatedAt:   nil,
		UpdatedAt:   nil,
	}

	// * Check lab exist
	var currentLab *model.Lab
	if result := modules.DB.Where("code = ?", lab.Code).First(&currentLab); result.Error == nil {
		// # Case of lab already exist
		// * Update lab
		// lab.Id = currentLab.Id
		// if result := modules.DB.Save(lab); result.Error != nil {
		// 	return result.Error
		// }
	} else if result.Error == gorm.ErrRecordNotFound {
		// # Case of lab not exist
		// * Create lab
		if result := modules.DB.Create(lab); result.Error != nil {
			return result.Error
		}
	} else {
		return result.Error
	}

	// # Checkpoint: Lab record is up-to-dated, as `lab` variable

	// * Import tasks
	for _, task := range imp.Tasks {
		if err := ImportTask(value.Ptr[uint64](1), task); err != nil {
			return err
		}
	}

	// * Create lab
	return nil
}
