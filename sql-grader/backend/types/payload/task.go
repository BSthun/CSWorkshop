package payload

type TaskList struct {
	Id    *uint64 `json:"id"`
	Title *string `json:"title"`
}
