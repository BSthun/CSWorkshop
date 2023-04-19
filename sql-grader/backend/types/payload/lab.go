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

type LabState struct {
	TaskTitle       *string          `json:"taskTitle"`
	TaskDescription *string          `json:"taskDescription"`
	TaskTags        []map[string]any `json:"taskTags"`
	Query           *string          `json:"query"`
	QueryPassed     *bool            `json:"queryPassed"`
	QueryError      *string          `json:"queryError"`
	Result          *LabStateResult  `json:"result"`
}

type LabStateResult struct {
	ExpectedHeader []string   `json:"expectedHeader"`
	ExpectedRows   [][]string `json:"expectedRows"`
	ActualHeader   []string   `json:"actualHeader"`
	ActualRows     [][]string `json:"actualRows"`
}
