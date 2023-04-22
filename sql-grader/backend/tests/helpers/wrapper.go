package helpers

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"backend/modules"
	"backend/types/response"
)

func ReadBody(v interface{}) *strings.Reader {
	body, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return strings.NewReader(string(body))
}

func SetBasicHeader(req *http.Request, user ...*User) {
	req.Header.Set("Content-Type", "application/json")
	if len(user) > 0 {
		req.AddCookie(&http.Cookie{
			Name:  "user",
			Value: *user[0].Token,
		})
	}
}

func Request[T any](req *http.Request) (*http.Response, *T, *response.ErrorResponse) {
	res, err := modules.Fiber.Test(req, -1)
	if err != nil {
		panic(err)
	}

	raw, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	if res.StatusCode == 200 {
		var body *response.GenericSuccessResponse[*T]
		if err := json.Unmarshal(raw, &body); err != nil {
			panic(err)
		}
		return res, body.Data, nil
	} else {
		var body *response.ErrorResponse
		if err := json.Unmarshal(raw, &body); err != nil {
			panic(err)
		}
		return res, nil, body
	}
}
