package enroll

import (
	"backend/modules"
	"backend/types/model"
)

func GetEnrollment(enrollId *uint64, userId *uint64) (*model.Enrollment, error) {
	var enrollment *model.Enrollment
	if result := modules.DB.
		Preload("User").
		Preload("Lab").
		Where(
			"id = ? AND user_id = ?",
			enrollId,
			userId,
		).
		First(&enrollment); result.Error != nil {
		return nil, result.Error
	}
	return enrollment, nil
}
