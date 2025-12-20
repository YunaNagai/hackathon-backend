package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"hackathon-backend/model"
	"hackathon-backend/usecase"
)

func RegisterUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if user.Age < 20 || user.Age > 80 {
			http.Error(w, "invalid age", http.StatusBadRequest)
			return
		}

		if err := usecase.RegisterUser(db, user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}
