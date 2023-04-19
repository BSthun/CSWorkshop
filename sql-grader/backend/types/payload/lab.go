package payload

type Lab struct {
	Id          *uint64 `json:"id"`
	Code        *string `json:"code"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

type TaskClickQuery struct {
	EnrollmentId *uint64 `query:"enrollmentId" validate:"required"`
	TaskId       *uint64 `query:"taskId" validate:"required"`
}
