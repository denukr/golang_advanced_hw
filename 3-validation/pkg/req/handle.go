package req

import (
	"log"
	"net/http"
)

func HandleBody[T any](w http.ResponseWriter, r *http.Request) (*T, error) {
	var payload T
	err := Decode[T](r.Body, &payload)
	log.Println(payload)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err = IsValidate(payload); err != nil {
		return nil, err
	}

	return &payload, nil
}
