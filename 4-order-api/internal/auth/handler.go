package auth

import (
	"golang-adv/4-order-api/configs"
	"golang-adv/4-order-api/pkg/jwt"
	"golang-adv/4-order-api/pkg/req"
	"golang-adv/4-order-api/pkg/resp"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}

type AuthHandler struct {
	*configs.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps *AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}
	router.Handle("POST /api/auth", handler.Auth())
	router.Handle("POST /api/verify", handler.Verify())
}

func (h *AuthHandler) Auth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[AuthRequestPayload](w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		sessionId, err := h.AuthService.AuthByPhone(body.Phone)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		resp.JsonResp(w, AuthResponsePayload{SessionId: sessionId}, http.StatusOK)
	}
}

func (h *AuthHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[VerifyRequest](w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		user, err := h.AuthService.GetBySessionId(body.SessionId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		j := jwt.NewJWT(h.Config.AuthConf.Key)
		token, err := j.Create(user.Phone)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp.JsonResp(w, VerifyResponse{
			Token: token,
		}, http.StatusOK)
	}
}
