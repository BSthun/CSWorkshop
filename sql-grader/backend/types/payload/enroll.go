package payload

import "time"

type EnrollInfoGetQuery struct {
	EnrollmentId *uint64 `json:"enrollmentId" validate:"required"`
}

type EnrollmentInfo struct {
	EnrollmentId *uint64     `json:"enrollmentId"`
	EnrolledAt   *time.Time  `json:"enrolledAt"`
	LabName      *string     `json:"labName"`
	DbName       *string     `json:"dbName"`
	DbValid      *bool       `json:"dbValid"`
	DbHost       *string     `json:"dbHost,omitempty"`
	DbPort       *string     `json:"dbPort,omitempty"`
	DbUsername   *string     `json:"DbUsername,omitempty"`
	DbPassword   *string     `json:"dbPassword,omitempty"`
	Tasks        []*TaskList `json:"tasks,omitempty"`
	Token        *string     `json:"token,omitempty"`
}

type EnrollMockResponse struct {
	Token *string `json:"token"`
}
