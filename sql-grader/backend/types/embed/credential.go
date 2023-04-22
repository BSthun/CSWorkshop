package embed

import (
	"database/sql/driver"
	"encoding/json"

	"backend/utils/text"
)

type Credential struct {
	Username *string `json:"username" validate:"required"`
	Password *string `json:"password" validate:"required"`
}

func (c *Credential) Value() (driver.Value, error) {
	if c == nil {
		return nil, nil
	}
	if err := text.Validator.Struct(c); err != nil {
		return nil, err
	}
	return json.Marshal(c)
}

func (c *Credential) Scan(src interface{}) error {
	if err := json.Unmarshal(src.([]byte), c); err != nil {
		return err
	}
	return text.Validator.Struct(c)
}
