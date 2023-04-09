package endpoints

import (
	"github.com/gofiber/fiber/v2"

	accountEndpoint "backend/endpoints/account"
)

func Init(router fiber.Router) {
	account := router.Group("account/")
	account.Post("callback", accountEndpoint.HandleCallbackPost)
}
