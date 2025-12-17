package controller

import (
	"database/sql"
	"hackathon-backend/usecase"
	"net/http"
)

func RegisterUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		usecase.RegisterUser(db, w, r)
	}
}
