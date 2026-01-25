package verify

type SendResponse struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
}
