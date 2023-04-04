package crypto

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	"backend/modules/config"
	"backend/types/response"
)

func SignJwt(c *fiber.Ctx, claim jwt.Claims) (string, *response.ErrorInstance) {
	// * Generate signed token with HMAC
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	if signedToken, err := token.SignedString([]byte(config.C.JwtSecret)); err != nil {
		return "", response.Error(c, true, "Unable to sign JWT token", err)
	} else {
		return signedToken, nil
	}
}
