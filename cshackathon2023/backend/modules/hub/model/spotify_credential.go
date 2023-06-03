package hubModel

import "time"

type SpotifyCredential struct {
	AccessToken string
	ExpiresIn   time.Time
}

func NewSpotifyCredential() *SpotifyCredential {
	return &SpotifyCredential{
		AccessToken: "",
		ExpiresIn:   time.Now(),
	}
}
