package admin

import (
	"net/http"

	"backend/tests/helpers"
	"backend/types/payload"
	"backend/types/response"
)

func RequestImportLabPost(imp *payload.AdminLabImport) (*http.Response, any, *response.ErrorResponse) {
	// * Construct body
	body := helpers.ReadBody(imp)

	// * Construct request
	req, _ := http.NewRequest("POST", "/api/admin/import/lab", body)
	helpers.SetBasicHeader(req, helpers.D.Users[0])

	// * Send request
	return helpers.Request[any](req)
}
