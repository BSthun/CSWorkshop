package profile

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	"backend/functions/enroll"
	"backend/functions/profile"
	"backend/types/common"
	"backend/types/model"
	"backend/types/payload"
	"backend/types/response"
	"backend/utils/value"
)

func EnrollmentsGetHandler(c *fiber.Ctx) error {
	// * Parse user
	u := c.Locals("u").(*jwt.Token).Claims.(*common.UserClaims)

	// * Get enrollments
	enrollments, err := profile.GetEnrollments(u.UserId)
	if err != nil {
		return response.Error(c, false, "Unable to get enrollments", err)
	}

	mappedEnrollments, _ := value.Iterate(enrollments, func(enrollment *model.Enrollment) (*payload.EnrollInfo, *response.ErrorInstance) {
		return enroll.MapEnrollment(enrollment), nil
	})

	return c.JSON(response.Success(c, mappedEnrollments))
}
