package account

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	"backend/modules"
	"backend/types/common"
	"backend/types/model"
	"backend/types/payload"
	"backend/types/response"
)

func HandleCallbackPost(c *fiber.Ctx) error {
	// * Parse base
	b := modules.B

	// * Parse body
	var body *payload.AuthenCallbackBody
	if err := c.BodyParser(&body); err != nil {
		return response.Error(c, false, "Unable to parse body", err)
	}

	// * Verify token
	token, err := b.FirebaseAuth.VerifyIDToken(context.Background(), *body.IdToken)
	if err != nil {
		return response.Error(c, false, "Unable to verify token", err)
	}

	// * Get profile
	profile, err := b.FirebaseAuth.GetUser(context.Background(), token.UID)
	if err != nil {
		return response.Error(c, false, "Unable to get profile", err)
	}

	// * Apply user record
	var user *model.User
	if result := b.DB.FirstOrCreate(&user, &model.User{
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
	signedJwtToken, err := jwtToken.SignedString([]byte(b.Conf.JwtSecret))
	if err != nil {
		return response.Error(c, true, "Unable to sign JWT token", err)
	}

	return c.JSON(response.Info(c, map[string]interface{}{
		"token": signedJwtToken,
	}))
}
