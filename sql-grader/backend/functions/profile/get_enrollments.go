package profile

import (
	"backend/modules"
	"backend/types/model"
)

func GetEnrollments(userId *uint64, preloadUser bool, preloadLab bool) ([]*model.Enrollment, error) {
	var enrollments []*model.Enrollment
	tx := modules.DB
	if preloadLab {
		tx = tx.Preload("Lab")
	}
	if preloadUser {
		tx = tx.Preload("User")
	}
	if result := tx.Where("user_id = ?", userId).Find(&enrollments); result.Error != nil {
		return nil, result.Error
	}

	return enrollments, nil
}
