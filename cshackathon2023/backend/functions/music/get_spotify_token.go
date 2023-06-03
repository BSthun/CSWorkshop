package functions

import (
	"backend/modules"
	"time"
)

func GetSpotifyToken() string {
	if modules.Hub.SpotifyCredential.ExpiresIn.Before(time.Now()) {
		credential, err := SpotifyRefreshToken()
		if err != nil {
			panic(err)
		}

		modules.Hub.SpotifyCredential.AccessToken = *credential.AccessToken
		modules.Hub.SpotifyCredential.ExpiresIn = time.Now().Add(time.Duration(*credential.ExpiresIn) * time.Second)
	}

	return modules.Hub.SpotifyCredential.AccessToken
}
