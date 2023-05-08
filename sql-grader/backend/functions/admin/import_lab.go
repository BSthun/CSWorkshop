package admin

import (
	"gorm.io/gorm"

	"backend/modules"
	"backend/types/model"
	"backend/types/payload"
	"backend/utils/text"
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
		// * Update lab
		lab.Id = currentLab.Id
		lab.CreatedAt = currentLab.CreatedAt
		if result := modules.DB.Save(lab); result.Error != nil {
			return result.Error
		}
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
		if err := ImportTask(lab.Id, task); err != nil {
			return err
		}
	}

	// * Create lab
	return nil
}
