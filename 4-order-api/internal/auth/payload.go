package auth

type AuthRequestPayload struct {
	Phone string `json:"phone" validate:"required,e164"`
}

type AuthResponsePayload struct {
	SessionId string `json:"session_id"`
}

type VerifyRequest struct {
	Code      string `json:"code" validate:"required,min=4"`
	SessionId string `json:"session_id"`
}

type VerifyResponse struct {
	Token string `json:"token"`
}
