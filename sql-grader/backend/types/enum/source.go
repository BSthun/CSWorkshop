package enum

import (
	"encoding/json"
	"fmt"
)

type Source string

const (
	SourceSqlDump      Source = "sql_dump"
	SourceAutoGenerate Source = "auto_generate"
	SourceBlank        Source = "blank"
)

func (s *Source) UnmarshalJSON(data []byte) error {
	var val string
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}

	source := Source(val)
	if source != SourceSqlDump && source != SourceAutoGenerate && source != SourceBlank {
		return fmt.Errorf("invalid source enum value: %s", source)
	}

	*s = source

	return nil
}
