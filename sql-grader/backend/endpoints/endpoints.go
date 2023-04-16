package endpoints

import (
	"github.com/gofiber/fiber/v2"

	accountEndpoint "backend/endpoints/account"
	adminEndpoint "backend/endpoints/admin"
	labEndpoint "backend/endpoints/lab"
	profileEndpoint "backend/endpoints/profile"
	"backend/modules/fiber/middleware"
)

func Init(router fiber.Router) {
	// * Account
	account := router.Group("account/")
	account.Post("callback", accountEndpoint.CallbackPostHandler)

	// * Profile
	profile := router.Group("profile/", middleware.Jwt())
	profile.Get("state", profileEndpoint.StateGetHandler)
	profile.Get("enrollments", profileEndpoint.EnrollmentsGetHandler)
	profile.Get("labs", profileEndpoint.LabsGetHandler)
	profile.Post("enroll", profileEndpoint.EnrollLabPostHandler)

	// * Lab
	lab := router.Group("lab/", middleware.Jwt())
	lab.Get("enroll/info", labEndpoint.EnrollGetHandler)
	lab.Get("enroll/mock", labEndpoint.EnrollMockGetHandler)

	// * Admin
	admin := router.Group("admin/", middleware.Jwt())
	admin.Post("import/lab", adminEndpoint.ImportLabPostHandler)
}
