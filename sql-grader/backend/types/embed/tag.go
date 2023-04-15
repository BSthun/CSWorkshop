package embed

import (
	"database/sql/driver"
	"encoding/json"
)

type TaskTags map[string]string

func (t *TaskTags) Scan(src interface{}) error {
	return json.Unmarshal(src.([]byte), t)
}

func (t *TaskTags) Value() (driver.Value, error) {
	return json.Marshal(t)
}
