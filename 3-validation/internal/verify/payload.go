package verify

type SendRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type SendResponse struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
}
