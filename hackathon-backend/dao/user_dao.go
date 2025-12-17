package dao

import (
	"database/sql"
	"hackathon-backend/model"
)

func InsertUser(db *sql.DB, user model.User) error {
	_, err := db.Exec("INSERT INTO user (id, name, age) VALUES (?, ?, ?)", user.ID, user.Name, user.Age)
	return err
}
