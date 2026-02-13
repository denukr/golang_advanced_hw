package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	Secret string
}

func NewJWT(secret string) *JWT {
	return &JWT{Secret: secret}
}

func (j *JWT) Create(phone string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"phone": phone,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := t.SignedString([]byte(j.Secret))
	if err != nil {
		fmt.Println("Ошибка подписи:", err)
		return "", err
	}
	return tokenString, nil
}
