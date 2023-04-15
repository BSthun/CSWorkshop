package admin

import (
	"fmt"

	"backend/modules"
)

func CheckTemplateDb(templateDb string) error {
	// * Validate template db is exist
	var dbs []*string
	if result := modules.DB.Raw(fmt.Sprintf("SHOW DATABASES LIKE '%s'", templateDb)).Scan(&dbs); result.Error != nil {
		return result.Error
	}

	if len(dbs) == 0 {
		return fmt.Errorf("template database is not exist")
	}

	return nil
}
