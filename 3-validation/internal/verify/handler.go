package verify

import (
	"fmt"
	"go/email-verify/config"
	"go/email-verify/pkg/json"
	"go/email-verify/pkg/req"
	"go/email-verify/pkg/resp"
	"net/http"
	"time"
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
	r.HandleFunc("GET /verify/{hash}", handler.Verify())
}

func (h *VerifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		sendReq, err := req.HandleBody[SendRequest](w, r)
		if err != nil {
			resp.JsonResp(w, err.Error(), 400)
			return
		}

		JsonStorage := json.NewJsonStorage("storage.json")
		generatedHash := Hash(time.Now().String())
		JsonStorage.Write(json.Data{
			Email: sendReq.Email,
			Hash:  generatedHash,
		})

		SendEmail(generatedHash)

		data := SendResponse{
			Email:   h.Config.Email,
			Address: h.Config.Address,
		}
		resp.JsonResp(w, data, 201)
	}
}

func (h *VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := req.HandleBody[VerifyRequest](w, r)
		if err != nil {
			resp.JsonResp(w, err.Error(), 400)
			return
		}
		JsonStorage := json.NewJsonStorage("storage.json")
		if err := JsonStorage.Read(); err != nil {
			resp.JsonResp(w, "Ошибка чтения хранилища: "+err.Error(), 500)
			return
		}
		fmt.Println(JsonStorage)
		for i, v := range JsonStorage.Data {
			fmt.Println(i, v)
		}
		resp.JsonResp(w, "", 201)
	}
}
