package usecase

import (
	"database/sql"
	"encoding/json"
	"hackathon-backend/dao"
	"hackathon-backend/model"
	"hackathon-backend/utils"
	"log"
	"net/http"
)

func RegisterUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	log.Println("usecase.RegisterUser called")

	var req model.User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if req.Age < 20 || req.Age > 80 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	req.ID = utils.NewULID()
	if err := dao.InsertUser(db, req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": req.ID})
}
