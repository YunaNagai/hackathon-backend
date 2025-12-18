package dao

import (
	"database/sql"
	"hackathon-backend/model"
)

func InsertUser(db *sql.DB, user model.User) error {
	_, err := db.Exec("INSERT INTO users (id, name, email, age) VALUES (?, ?, ?, ?)", user.ID, user.Name, user.Email, user.Age)
	return err
}
