package profile

import (
	"backend/types/model"
	"backend/types/payload"
)

func MapLab(lab *model.Lab) *payload.Lab {
	return &payload.Lab{
		Id:          lab.Id,
		Code:        lab.Code,
		Name:        lab.Name,
		Description: lab.Description,
	}
}
