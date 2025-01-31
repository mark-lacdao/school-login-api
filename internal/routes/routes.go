package routes

import (
	"net/http"
	"school-login-api/internal/handlers"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
)

func InitializeRoutes(db *pgxpool.Pool) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/login", handlers.LoginHandler(db)).Methods("POST")
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	return r
}
