package handlers

import (
	"net/http"
	"school/pkg/utils"
)

// HomeHandler handles the home page request
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{"message": "Welcome to the School Management App"}
	utils.WriteJSON(w, http.StatusOK, data)
}

// NotFoundHandler handles 404 errors
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{"error": "Page not found"}
	utils.WriteJSON(w, http.StatusNotFound, data)
}
