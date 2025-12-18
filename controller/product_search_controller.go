package controller

import (
	"encoding/json"
	"net/http"

	"database/sql"
	"hackathon-backend/dao"

	"github.com/go-chi/chi/v5"
)

func GetProductByID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		product, err := dao.GetProductByID(db, id)
		if err != nil {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(product)
	}
}
