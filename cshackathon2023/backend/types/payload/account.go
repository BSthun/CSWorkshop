package payload

type AuthCallbackBody struct {
	IdToken *string `json:"idToken" validate:"required"`
}

type AuthCallbackResponse struct {
	Token *string `json:"token"`
}
