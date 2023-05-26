package endpoints

import (
	"github.com/gofiber/fiber/v2"

	loginEndpoint "backend/endpoints/login"
	profileEndpoint "backend/endpoints/profile"
	"backend/modules/fiber/middlewares"
)

func Init(router fiber.Router) {
	// * Login
	login := router.Group("login/", middlewares.Sentry())
	login.Get("redirect", loginEndpoint.RedirectHandler)
	login.Post("callback", loginEndpoint.HandleCallback)

	// * Account
	profile := router.Group("profile/", middlewares.Jwt(), middlewares.Sentry())
	profile.Get("info", profileEndpoint.InfoHandler)
	profile.Get("download", profileEndpoint.DownloadHandler)
}
