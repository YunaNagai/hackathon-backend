package controller

import (
	"database/sql"
	"hackathon-backend/usecase"
	"net/http"
)

func ProductsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodPost:
			usecase.RegisterProducts(db, w, r)
		case http.MethodGet:
			usecase.GetProducts(db, w, r)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}
