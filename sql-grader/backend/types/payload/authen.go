package payload

type AuthCallbackBody struct {
	IdToken *string `json:"idToken" validate:"required"`
}

type AuthCallbackResponse struct {
	Token *string `json:"token"`
}

type FirebaseIdTokenResponse struct {
	Kind         *string `json:"kind"`
	IsNewUser    *bool   `json:"isNewUser"`
	IdToken      *string `json:"idToken"`
	RefreshToken *string `json:"refreshToken"`
	ExpiresIn    *string `json:"expiresIn"`
}
