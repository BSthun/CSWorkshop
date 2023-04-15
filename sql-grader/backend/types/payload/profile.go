package payload

import "backend/types/enum"

type UserProfile struct {
	Id     *uint64 `json:"id"`
	Email  *string `json:"email"`
	Name   *string `json:"name"`
	Avatar *string `json:"avatar"`
}

type ProfileStateGetResponse struct {
	Profile *UserProfile `json:"profile"`
}

type ProfileLabGetResponse struct {
	Labs []*Lab `json:"labs"`
}

type ProfileEnrollmentRequest struct {
	LabId  *uint64      `json:"lab_id" validate:"required"`
	Source *enum.Source `json:"source" validate:"required"`
	Dump   *string      `json:"dump" validate:"required_if=Source sql_dump"`
}
