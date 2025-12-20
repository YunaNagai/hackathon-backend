package usecase

import (
	"database/sql"
	"hackathon-backend/dao"
	"hackathon-backend/model"
	"log"
)

func GetMessages(db *sql.DB) ([]model.Message, error) {
	return dao.GetAllMessages(db)
}

func CreateMessage(db *sql.DB, msg model.Message) error {
	err := dao.InsertMessage(db, msg)
	if err != nil {
		log.Printf("CreateMessage (usecase) error: %v", err) // ← これが必要！
		return err
	}
	return nil
}
