package common

type UserClaims struct {
	UserId *uint64 `json:"user_id"`
	Name   *string `json:"name"`
	Exp    *int64  `json:"exp"`
}

func (v *UserClaims) Valid() error {
	return nil
}
