package helpers

import (
	"encoding/json"
	"net/http"
)

func init() {}

// Message general message
type Message struct {
	Text string
}

// CreateJSONResponse return a standard JSON response
func CreateJSONResponse(data interface{}, status int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
