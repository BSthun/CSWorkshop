package endpoints

import (
	"github.com/gofiber/fiber/v2"

	accountEndpoint "backend/endpoints/account"
	profileEndpoint "backend/endpoints/profile"
	"backend/modules/fiber/middlewares"
)

func Init(router fiber.Router) {
	// * Account
	account := router.Group("account/")
	account.Post("callback", accountEndpoint.CallbackPostHandler)

	// * Profile
	profile := router.Group("profile/", middlewares.Jwt())
	profile.Get("state", profileEndpoint.StateGetHandler)
}
