package usecase

import (
	"database/sql"
	"encoding/json"
	"hackathon-backend/dao"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func UpdateTransaction(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var body struct {
		Status string `json:"status"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid json"})
		return
	}

	if body.Status == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "missing status"})
		return
	}

	if err := dao.UpdateTransactionStatus(db, id, body.Status); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "db error"})
		return
	}

	tx, err := dao.GetTransactionByID(db, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "fetch error"})
		return
	}

	json.NewEncoder(w).Encode(tx)
}
