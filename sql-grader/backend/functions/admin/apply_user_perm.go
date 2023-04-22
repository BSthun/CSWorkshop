package admin

import (
	"fmt"

	"backend/modules"
	"backend/types/model"
)

func ApplyUserPerm() error {
	// * Flush all lab users
	var usernames []string
	if result := modules.DB.Raw("SELECT user FROM mysql.user WHERE user LIKE 'lab_%'").Find(&usernames); result.Error != nil {
		return result.Error
	}

	for _, username := range usernames {
		if result := modules.DB.Exec(fmt.Sprintf("DROP USER '%s'@'%%'", username)); result.Error != nil {
			return result.Error
		}
	}

	// * Query all users
	var users []*model.User
	if result := modules.DB.Find(&users); result.Error != nil {
		return result.Error
	}

	// * Create users
	for _, user := range users {
		// * Skip if user has no credential
		if user.Credential == nil {
			continue
		}

		// * Create user
		exec := fmt.Sprintf("CREATE USER '%s'@'%%' IDENTIFIED BY '%s'", *user.Credential.Username, *user.Credential.Password)
		if result := modules.DB.Exec(exec); result.Error != nil {
			return result.Error
		}

		// * Get all enrollments
		var enrollments []*model.Enrollment
		if result := modules.DB.Find(&enrollments, "user_id = ?", user.Id); result.Error != nil {
			return result.Error
		}

		// * Grant privileges
		for _, enrollment := range enrollments {
			exec := fmt.Sprintf("GRANT ALL ON %s.* TO '%s'@'%%'", *enrollment.DbName, *user.Credential.Username)
			if result := modules.DB.Exec(exec); result.Error != nil {
				return result.Error
			}
		}
	}

	return nil
}
