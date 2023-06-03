package account

import (
	"backend/utils/text"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"

	"backend/functions/account"
	"backend/modules"
	"backend/types/model"
	"backend/types/payload"
	"backend/types/response"
)

func CallbackPostHandler(c *fiber.Ctx) error {
	// * Parse body
	var body *payload.AuthCallbackBody
	if err := c.BodyParser(&body); err != nil {
		return response.Error(c, false, "Unable to parse body", err)
	}

	// * Validate body
	if err := text.Validator.Struct(body); err != nil {
		return response.Error(c, false, "Unable to validate body", err)
	}

	// * Verify token
	token, err := modules.FirebaseAuth.VerifyIDToken(context.Background(), *body.IdToken)
	if err != nil {
		return response.Error(c, false, "Unable to verify token", err)
	}

	// * Get profile
	profile, err := modules.FirebaseAuth.GetUser(context.Background(), token.UID)
	if err != nil {
		return response.Error(c, false, "Unable to get profile", err)
	}

	// * Construct user record
	user := &model.User{
		Id:          nil,
		FirebaseUid: &token.UID,
		Email:       &profile.Email,
		Name:        &profile.DisplayName,
		Avatar:      &profile.PhotoURL,
		CreatedAt:   nil,
		UpdatedAt:   nil,
	}

	// * Affirm user record
	userToken, err := account.ConfirmUser(user)
	if err != nil {
		return response.Error(c, true, "Unable to affirm user", err)
	}

	// Set cookie
	c.Cookie(&fiber.Cookie{
		Name:    "user",
		Value:   *userToken,
		Expires: time.Now().Add(time.Hour * 24 * 7),
	})

	return c.JSON(response.Success(c, &payload.AuthCallbackResponse{
		Token: userToken,
	}))
}
