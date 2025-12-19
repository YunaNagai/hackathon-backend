package usecase

import (
	"database/sql"
	"encoding/json"
	"hackathon-backend/dao"
	"hackathon-backend/model"
	"net/http"
)

func CreateTransaction(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var req model.Transaction

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid json"})
		return
	}

	if req.ID == "" || req.ProductID == "" || req.BuyerID == "" || req.SellerID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid fields"})
		return
	}

	if err := dao.InsertTransaction(db, req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "db error"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(req)
}
