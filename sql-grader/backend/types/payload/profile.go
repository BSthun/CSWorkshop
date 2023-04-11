package payload

type UserProfile struct {
	Id     *uint64 `json:"id"`
	Email  *string `json:"email"`
	Name   *string `json:"name"`
	Avatar *string `json:"avatar"`
}
