package usecase

import (
	"database/sql"
	"hackathon-backend/dao"
	"hackathon-backend/model"
)

func RegisterUser(db *sql.DB, user model.User) error {
	return dao.InsertUser(db, user)
}
