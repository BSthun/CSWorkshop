package enroll

import (
	"backend/types/model"
)

func ActEnrollLab(user *model.User, lab *model.Lab, dump *string) error {
	// * Create a new enrollment
	enrollment, err := NewEnrollment(user, lab)
	if err != nil {
		return err
	}

	_ = enrollment

	return nil
}
