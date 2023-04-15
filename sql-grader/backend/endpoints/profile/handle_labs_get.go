package profile

import (
	"github.com/gofiber/fiber/v2"

	"backend/functions/profile"
	"backend/types/model"
	"backend/types/payload"
	"backend/types/response"
	"backend/utils/value"
)

func LabsGetHandler(c *fiber.Ctx) error {
	labs, err := profile.GetLabs()
	if err != nil {
		return response.Error(c, false, "Unable to get labs", err)
	}

	mappedLabs, err := value.Iterate(labs, func(lab *model.Lab) (*payload.Lab, *response.ErrorInstance) {
		return profile.MapLab(lab), nil
	})

	return c.JSON(response.Success(c, &payload.ProfileLabGetResponse{
		Labs: mappedLabs,
	}))
}
