package endpoints

import (
	accountEndpoint "backend/endpoints/account"
	backdropEndpoint "backend/endpoints/backdrop"
	imageEndpoint "backend/endpoints/image"
	musicEndpoint "backend/endpoints/music"
	"backend/modules/fiber/middleware"
	"github.com/gofiber/fiber/v2"
)

func Init(router fiber.Router) {
	// * Account
	account := router.Group("account/")
	account.Post("callback", accountEndpoint.CallbackPostHandler)

	// * Backdrop
	backdrop := router.Group("backdrop/")
	backdrop.Get("state", backdropEndpoint.StateGetHandler)

	// * Music
	music := router.Group("music/", middleware.Jwt())
	music.Get("state", musicEndpoint.StateGetHandler)
	music.Get("search", musicEndpoint.SearchGetHandler)
	music.Post("queue", musicEndpoint.QueuePostHandler)

	// * Image
	image := router.Group("image/")
	image.Get("rand", imageEndpoint.RandGetHandler)
}
