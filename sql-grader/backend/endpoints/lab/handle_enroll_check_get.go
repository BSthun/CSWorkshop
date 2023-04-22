package lab

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	"backend/modules"
	"backend/types/common"
	"backend/types/model"
	"backend/types/payload"
	"backend/types/response"
	"backend/utils/text"
)

func EnrollCheckGetHandler(c *fiber.Ctx) error {
	// * Parse user
	u := c.Locals("u").(*jwt.Token).Claims.(*common.UserClaims)

	// * Parse query
	query := new(payload.EnrollInfoGetQuery)
	if err := c.QueryParser(query); err != nil {
		return response.Error(c, false, "Unable to parse query", err)
	}

	// * Validate query
	if err := text.Validator.Struct(query); err != nil {
		return response.Error(c, false, "Unable to validate query", err)
	}

	// * Update enrollment info
	if result := modules.DB.Model(new(model.Enrollment)).
		Where("id = ? AND user_id = ?", query.EnrollmentId, u.UserId).
		Update("db_valid", true); result.Error != nil {
		return response.Error(c, false, "Unable to update enrollment info", result.Error)
	}

	return nil
}
