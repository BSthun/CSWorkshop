package profile

import (
	"backend/modules"
	"backend/types/model"
)

func GetProfile(userId *uint64) (*model.User, error) {
	var user *model.User
	if result := modules.DB.First(&user, userId); result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
