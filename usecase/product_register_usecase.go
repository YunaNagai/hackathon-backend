package usecase

import (
	"database/sql"
	"encoding/json"
	"hackathon-backend/dao"
	"hackathon-backend/model"
	"hackathon-backend/utils"
	"net/http"
)

func RegisterProducts(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var req model.Product
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// NullString / NullInt64 に合わせたバリデーション
	if !req.SellerID.Valid || req.SellerID.String == "" ||
		!req.Title.Valid || req.Title.String == "" ||
		!req.Price.Valid || req.Price.Int64 <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	req.ID = utils.NewULID()
	req.Status = "selling"

	if err := dao.InsertProduct(db, req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": req.ID})
}
