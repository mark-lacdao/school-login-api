package utils

import (
	"encoding/json"
	"net/http"
)

// WriteJSON is a utility function to write JSON responses
func WriteJSON(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

// ReadJSON is a utility function to read JSON requests
func ReadJSON(r *http.Request, dst interface{}) error {
	return json.NewDecoder(r.Body).Decode(dst)
}
