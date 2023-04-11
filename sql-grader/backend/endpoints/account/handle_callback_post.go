package account

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	"backend/modules"
	"backend/types/common"
	"backend/types/model"
	"backend/types/payload"
	"backend/types/response"
)

func CallbackPostHandler(c *fiber.Ctx) error {
	// * Parse body
	var body *payload.AuthenCallbackBody
	if err := c.BodyParser(&body); err != nil {
		return response.Error(c, false, "Unable to parse body", err)
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

	// * Apply user record
	var user *model.User
	if result := modules.DB.FirstOrCreate(&user, &model.User{
		Id:          nil,
		FirebaseUid: &token.UID,
		Email:       &profile.Email,
		Name:        &profile.DisplayName,
		Avatar:      &profile.PhotoURL,
		CreatedAt:   nil,
		UpdatedAt:   nil,
	}); result.Error != nil {
		return response.Error(c, true, "Unable to apply user record", result.Error)
	}

	// Create JWT claims
	claims := &common.UserClaims{
		UserId: user.Id,
		Name:   user.Name,
	}

	// Sign JWT token
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedJwtToken, err := jwtToken.SignedString([]byte(modules.Conf.JwtSecret))
	if err != nil {
		return response.Error(c, true, "Unable to sign JWT token", err)
	}

	// Set cookie
	c.Cookie(&fiber.Cookie{
		Name:    "user",
		Value:   signedJwtToken,
		Expires: time.Now().Add(time.Hour * 24 * 7),
	})

	return c.JSON(response.Success(c, map[string]interface{}{
		"token": signedJwtToken,
	}))
}
