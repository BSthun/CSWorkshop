package profile

import (
	"backend/modules"
	"backend/types/model"
)

func GetEnrollments(userId *uint64) ([]*model.Enrollment, error) {
	var enrollments []*model.Enrollment
	if result := modules.DB.Where("user_id = ?", userId).Find(&enrollments); result.Error != nil {
		return nil, result.Error
	}

	return enrollments, nil
}
