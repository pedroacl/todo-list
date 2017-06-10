package helpers

import "net/http"

func init() {}

// Message general message
type Message struct {
	Text string
}

// CreateJSONResponse return a standard JSON response
func CreateJSONResponse(data []byte, status int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	w.Write(data)
}
