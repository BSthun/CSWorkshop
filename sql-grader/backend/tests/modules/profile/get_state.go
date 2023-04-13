package profile

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"backend/tests/endpoints/profile"
	"backend/tests/helpers"
)

func GetState(t *testing.T) {
	for _, user := range helpers.D.Users {
		t.Run("get state for existing user", func(t *testing.T) {
			res, body, err := profile.RequestStateGet(user)
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, *user.Name, *body.Profile.Name)
			assert.Equal(t, 200, res.StatusCode)
		})
	}
}
