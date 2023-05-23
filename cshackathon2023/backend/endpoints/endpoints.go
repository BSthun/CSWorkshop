package endpoints

import (
	accountEndpoint "backend/endpoints/account"
	"github.com/gofiber/fiber/v2"
)

func Init(router fiber.Router) {
	// * Account
	account := router.Group("account/")
	account.Post("callback", accountEndpoint.CallbackPostHandler)
}
