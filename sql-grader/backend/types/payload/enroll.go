package payload

type EnrollInfoGetQuery struct {
	EnrollmentId *uint64 `json:"enrollmentId" validate:"required"`
}

type EnrollInfo struct {
	EnrollmentId *uint64     `json:"enrollmentId"`
	DbName       *string     `json:"dbName"`
	DbValid      *bool       `json:"dbValid"`
	DbHost       *string     `json:"dbHost,omitempty"`
	DbPort       *string     `json:"dbPort,omitempty"`
	DbUsername   *string     `json:"DbUsername,omitempty"`
	DbPassword   *string     `json:"dbPassword,omitempty"`
	Tasks        []*TaskList `json:"tasks,omitempty"`
}

type EnrollMockResponse struct {
	Token *string `json:"token"`
}
