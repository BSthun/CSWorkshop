package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/jwt/v3"

	"backend/types/common"
	"backend/types/response"

	"backend/modules"
)

func Jwt(b *modules.Base) fiber.Handler {
	conf := jwtware.Config{
		SigningKey:  []byte(b.Conf.JwtSecret),
		TokenLookup: "cookie:user",
		ContextKey:  "u",
		Claims:      &common.UserClaims{},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return response.Error(c, false, "JWT validation failure", err)
		},
	}

	return jwtware.New(conf)
}
