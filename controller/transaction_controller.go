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
			http.Error(w, err.Error(), http.StatusInternalServerError)
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
