package enroll

import (
	"fmt"
	"strconv"
	"strings"

	"backend/modules"
	idb "backend/modules/db"
	"backend/types/embed"
	"backend/types/model"
	"backend/utils/text"
	"backend/utils/value"
)

func NewEnrollment(user *model.User, lab *model.Lab) (*model.Enrollment, error) {
	// * Get user credentials
	cred, err := ExtUserCredential(user)
	if err != nil {
		return nil, err
	}

	// * Create database
	name, err := HelpCreateEnrollmentDatabase(user, lab, cred)
	if err != nil {
		return nil, err
	}

	// * Construct a new enrollment
	enrollment := &model.Enrollment{
		Id:        nil,
		UserId:    user.Id,
		User:      nil,
		LabId:     lab.Id,
		Lab:       nil,
		DbName:    name,
		DbValid:   value.Ptr(false),
		CreatedAt: nil,
		UpdatedAt: nil,
	}

	// * Create a new enrollment
	if result := modules.DB.Create(enrollment); result.Error != nil {
		return nil, result.Error
	}

	return enrollment, nil
}

func HelpCreateEnrollmentDatabase(user *model.User, lab *model.Lab, cred *embed.Credential) (*string, error) {
	// * Count user enrollment
	var count int64
	if result := modules.DB.
		Model(new(model.Enrollment)).
		Where(
			idb.Where(
				model.EnrollmentFieldUserId,
				"= ? AND",
				model.EnrollmentFieldLabId,
				"= ?",
			),
			user.Id,
			lab.Id,
		).
		Count(&count); result.Error != nil {
		return nil, result.Error
	}

	// * Construct database name
	formattedUserId := strconv.FormatUint(*user.Id, 32)
	formattedLabId := strconv.FormatUint(*lab.Id, 32)
	salt := strings.ToLower(*text.Random(text.RandomSet.UpperAlpha, 4))
	name := fmt.Sprintf("lab_%s_%s%s%s", strings.ToLower(*lab.Code), formattedUserId, formattedLabId, salt)

	// * Create database
	exec := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", name)
	if result := modules.DB.Exec(exec); result.Error != nil {
		return nil, result.Error
	}

	// * Grant privileges
	exec = fmt.Sprintf("GRANT ALL PRIVILEGES ON %s.* TO '%s'@'%%'", name, *cred.Username)
	if result := modules.DB.Exec(exec); result.Error != nil {
		return nil, result.Error
	}

	return &name, nil
}
