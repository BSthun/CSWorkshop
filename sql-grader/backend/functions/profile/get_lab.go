package profile

import (
	"backend/modules"
	"backend/types/model"
)

func GetLab(labId *uint64) (*model.Lab, error) {
	var lab *model.Lab
	if result := modules.DB.First(&lab, labId); result.Error != nil {
		return nil, result.Error
	}

	return lab, nil
}

func GetLabs() ([]*model.Lab, error) {
	var labs []*model.Lab
	if result := modules.DB.Find(&labs); result.Error != nil {
		return nil, result.Error
	}

	return labs, nil
}
