package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTData struct {
	Phone string
}

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

func (j *JWT) Parse(token string) (bool, *JWTData) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		return false, nil
	}
	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		// Если внутри не MapClaims, ok будет false
		fmt.Println("Ошибка: это не те клеймы, которые мы ждали")
		return false, nil
	}
	phone, ok := claims["phone"]
	if !ok {
		return false, nil
	}
	return t.Valid, &JWTData{
		Phone: phone.(string),
	}
}
