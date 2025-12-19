package controller

import (
	"database/sql"
	"encoding/json"
	"hackathon-backend/usecase"
	"net/http"
)

func GetTransactionsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		transactions, err := usecase.GetTransactions(db)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(transactions)
	}
}
func CreateTransactionHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		usecase.CreateTransaction(db, w, r)
	}
}
