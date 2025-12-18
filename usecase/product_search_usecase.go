package usecase

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"hackathon-backend/dao"
)

func GetProducts(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	products, err := dao.SelectAllProducts(db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
