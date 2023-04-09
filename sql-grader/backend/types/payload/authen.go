package payload

type AuthenCallbackBody struct {
	IdToken *string `json:"idToken" validate:"required"`
}
