package req

import "github.com/go-playground/validator"

func isValid(s any) error {
	validator := validator.New()
	err := validator.Struct(s)
	if err != nil {
		return err
	}
	return nil
}
