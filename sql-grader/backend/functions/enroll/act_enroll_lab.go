package enroll

import (
	"backend/types/model"
)

func ActEnrollLab(user *model.User, lab *model.Lab) (*model.Enrollment, error) {
	// * Create a new enrollment
	enrollment, err := NewEnrollment(user, lab)
	if err != nil {
		return nil, err
	}

	return enrollment, nil
}
