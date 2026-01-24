package verify

import (
	"go/email-verify/config"
	"go/email-verify/pkg/resp"
	"net/http"
)

type VerifyHandler struct {
	Config *config.VerifyConfig
}

type VerifyHandlerDeps struct {
	Config *config.VerifyConfig
}

func NewVerifyHandler(r *http.ServeMux, deps *VerifyHandlerDeps) {
	handler := VerifyHandler{
		Config: deps.Config,
	}
	r.HandleFunc("POST /send", handler.Send())
	r.HandleFunc("GET /verify/{hash}", handler.Send())
}

func (h *VerifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := SendResponse{
			Email:   h.Config.Email,
			Address: h.Config.Address,
		}
		resp.JsonResp(w, data, 201)
	}
}

func (h *VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := SendResponse{
			Email:   h.Config.Email,
			Address: h.Config.Address,
		}
		resp.JsonResp(w, data, 201)
	}
}
