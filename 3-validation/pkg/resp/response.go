package resp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JsonResp(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Fprintln(w, err.Error(), http.StatusInternalServerError)
	}
}
