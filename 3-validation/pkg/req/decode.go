package req

import (
	"encoding/json"
	"io"
)

func Decode[T any](body io.ReadCloser, data *T) error {
	err := json.NewDecoder(body).Decode(data)
	if err != nil {
		return err
	}
	return nil
}
