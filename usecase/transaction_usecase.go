package usecase

import (
	"database/sql"
	"hackathon-backend/dao"
	"hackathon-backend/model"
)

func GetTransactions(db *sql.DB) ([]model.Transaction, error) {
	return dao.GetAllTransactions(db)
}
