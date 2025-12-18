package usecase

import (
	"database/sql"
	"hackathon-backend/dao"
	"hackathon-backend/model"
)

func GetMessages(db *sql.DB) ([]model.Message, error) {
	return dao.GetAllMessages(db)
}

func CreateMessage(db *sql.DB, msg model.Message) error {
	return dao.InsertMessage(db, msg)
}
