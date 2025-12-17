package controller

import (
	"database/sql"
	"net/http"
	"uttc/usecase"
)

func RegisterUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		usecase.RegisterUser(db, w, r)
	}
}
