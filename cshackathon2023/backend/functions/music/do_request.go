package functions

import (
	"backend/utils/network"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"backend/types/payload"
	"backend/types/response"
)

func DoRequest(c *fiber.Ctx, method string, url string, body io.Reader, modifier func(r *http.Request), data any) *response.ErrorInstance {
	// * Construct request
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return response.Error(c, true, "Unable to construct request", err)
	}

	// * Modify request
	if modifier != nil {
		modifier(req)
	}

	// * Construct client
	cli := network.NewClient()
	resp, err := cli.Do(req)
	if err != nil {
		return response.Error(c, true, "Unable to send request", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	// * Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return response.Error(c, true, "Unable to read response body", err)
	}

	// * Check response
	if resp.StatusCode == 204 {
		return nil
	}

	if resp.StatusCode != 200 {
		var errResp *payload.SpotifyApiError
		if err := json.Unmarshal(respBody, &errResp); err != nil {
			if string(respBody) == "User not registered in the Developer Dashboard" {
				return response.Error(c, false, "Spotify Error: "+string(respBody), nil)
			}
			return response.Error(c, true, "Spotify Error: "+string(respBody), nil)
		} else {
			return response.Error(c, true, "Spotify Error: "+*errResp.Error.Message, nil)
		}
	}

	// * Parse response body
	if err := json.Unmarshal(respBody, &data); err != nil {
		return response.Error(c, true, "Unable to parse response body", err)
	}

	return nil
}
