package handlers

import (
	"context"
	"net/http"
	"school-login-api/internal/models"
	"school-login-api/pkg/utils"

	"github.com/jackc/pgx/v4/pgxpool"
)

func LoginHandler(db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := utils.ReadJSON(r, &user)
		if err != nil {
			utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
			return
		}

		var dbUser models.User
		err = db.QueryRow(context.Background(), "SELECT id, username, password FROM users WHERE username=$1", user.Username).Scan(&dbUser.ID, &dbUser.Username, &dbUser.Password)
		if err != nil {
			utils.WriteJSON(w, http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
			return
		}

		if user.Password != dbUser.Password {
			utils.WriteJSON(w, http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
			return
		}

		utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "Login successful"})
	}
}
