package payload

type Lab struct {
	Id          *uint64 `json:"id"`
	Code        *string `json:"code"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
}
