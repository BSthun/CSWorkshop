package profile

import (
	"net/http"

	"backend/tests/helpers"
	"backend/types/payload"
	"backend/types/response"
)

func RequestStateGet(u *helpers.User) (*http.Response, *payload.ProfileStateGetResponse, *response.ErrorResponse) {
	// * Construct request
	req, _ := http.NewRequest("GET", "/api/profile/state", nil)
	helpers.SetBasicHeader(req, u)

	// * Send request
	return helpers.Request[payload.ProfileStateGetResponse](req)
}
