package account

import (
	"net/http"

	"backend/tests/helpers"
	"backend/types/payload"
	"backend/types/response"
)

func RequestCallbackPost(idToken string) (*http.Response, *payload.AuthCallbackResponse, *response.ErrorResponse) {
	// * Construct body
	body := helpers.ReadBody(map[string]interface{}{
		"idToken": idToken,
	})

	// * Construct request
	req, _ := http.NewRequest("POST", "/api/account/callback", body)
	helpers.SetBasicHeader(req)

	// * Send request
	return helpers.Request[payload.AuthCallbackResponse](req)
}
