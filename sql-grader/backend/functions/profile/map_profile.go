package profile

import (
	"backend/types/model"
	"backend/types/payload"
)

func MapProfile(user *model.User) *payload.UserProfile {
	return &payload.UserProfile{
		Id:     user.Id,
		Email:  user.Email,
		Name:   user.Name,
		Avatar: user.Avatar,
	}
}
