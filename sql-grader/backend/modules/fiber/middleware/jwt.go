package middleware

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	"backend/types/common"
)

func Jwt() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		user := ctx.Cookies("user")
		userId, err := strconv.ParseUint(user, 10, 64)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  false,
				"message": "Unauthorized",
			})
		}
		ctx.Locals("u", &jwt.Token{
			Claims: &common.UserClaims{
				UserId: &userId,
			},
		})
		return ctx.Next()
	}
}
