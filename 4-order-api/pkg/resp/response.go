package resp

import (
	"encoding/json"
	"net/http"
)

func JsonResp(w http.ResponseWriter, data any, statusCode int) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-ONE", "Test")
	json.NewEncoder(w).Encode(data)
}
