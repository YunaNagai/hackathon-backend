package controller

import (
	"database/sql"
	"hackathon-backend/usecase"
	"log"
	"net/http"
)

func RegisterUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("RegisterUserHandler called")

		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Println("Calling usecase.RegisterUser")

		usecase.RegisterUser(db, w, r)
	}
}
