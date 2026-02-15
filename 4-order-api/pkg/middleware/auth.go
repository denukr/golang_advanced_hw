package middleware

import (
	"context"
	"golang-adv/4-order-api/configs"
	"golang-adv/4-order-api/pkg/jwt"
	"net/http"
	"strings"
)

type PhoneKey string

const (
	Phone PhoneKey = "phone"
)

func Unauthed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}

func IsAuthed(next http.HandlerFunc, config *configs.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if !strings.HasPrefix(authHeader, "Bearer ") {
			Unauthed(w)
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			Unauthed(w)
			return
		}
		j := jwt.NewJWT(config.Key)
		ok, data := j.Parse(token)
		if !ok {
			Unauthed(w)
			return
		}
		ctx := context.WithValue(r.Context(), Phone, data.Phone)

		newReq := r.WithContext(ctx)

		next.ServeHTTP(w, newReq)
	}
}
