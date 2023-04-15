package account

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"backend/modules"
	"backend/tests/endpoints/account"
	"backend/tests/helpers"
	"backend/types/payload"
	"backend/utils/value"
)

func TestCallback(t *testing.T) {
	for _, user := range helpers.D.Users {
		t.Run("callback with valid token", func(t *testing.T) {
			token := GetFirebaseIdToken(*user.Uid)
			res1, body, errBody := account.RequestCallbackPost(&token)
			if errBody != nil {
				t.Error(errBody)
				return
			} else {
				user.Token = body.Token
			}
			assert.Equal(t, 200, res1.StatusCode)
		})
	}
}

func TestInvalidCallback(t *testing.T) {
	t.Run("callback with invalid token", func(t *testing.T) {
		res1, _, _ := account.RequestCallbackPost(value.Ptr("invalid_token"))
		assert.Equal(t, 400, res1.StatusCode)
	})
}

func GetFirebaseIdToken(uid string) string {
	// * Get firebase token
	token, err := modules.FirebaseAuth.CustomToken(context.Background(), uid)
	if err != nil {
		panic(err)
	}

	// * Request ID Token
	body := helpers.ReadBody(map[string]any{
		"token":             token,
		"returnSecureToken": true,
	})
	host, _ := os.LookupEnv("FIREBASE_AUTH_EMULATOR_HOST")
	path, _ := url.JoinPath("http://", host, "/www.googleapis.com/identitytoolkit/v3/relyingparty/verifyCustomToken")
	req, _ := http.NewRequest("POST", path+"?key="+os.Getenv("FIREBASE_API_KEY"), body)
	req.Header.Set("Content-Type", "application/json")
	res, _ := http.DefaultClient.Do(req)
	defer func() {
		_ = res.Body.Close()
	}()
	var tokenRes *payload.FirebaseIdTokenResponse
	_ = json.NewDecoder(res.Body).Decode(&tokenRes)
	return *tokenRes.IdToken
}
