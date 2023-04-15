package admin

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"

	"backend/modules"
)

func CheckTemplateDb(templateDb string) error {
	// * Validate template db is exist
	rows, err := modules.SqlDB.Query(fmt.Sprintf("SHOW DATABASES LIKE '%s'", templateDb))
	if err != nil {
		spew.Dump(err)
		return err
	}

	for rows.Next() {
		var dbName string
		if err := rows.Scan(&dbName); err != nil {
			return err
		}
		if dbName == templateDb {
			return nil
		}
	}
	return fmt.Errorf("template db not found")
}
