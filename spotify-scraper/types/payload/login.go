package payload

type LoginSpotifyRedirect struct {
	Email string `json:"email"`
}

type LoginSpotifyCallback struct {
	Code  string `json:"code"`
	State string `json:"state"`
}

type LoginSpotifyCredentials struct {
	AccessToken  *string `json:"access_token"`
	RefreshToken *string `json:"refresh_token"`
	TokenType    *string `json:"token_type"`
	ExpiresIn    *int64  `json:"expires_in"`
	Scope        *string `json:"scope"`
	Error        *string `json:"error"`
}
