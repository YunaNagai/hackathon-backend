package dao

import (
	"database/sql"
	"hackathon-backend/model"
)

func InsertUser(db *sql.DB, user model.User) error {
	_, err := db.Exec("INSERT IGNORE INTO users (id, name, email, age) VALUES (?, ?, ?, ?)", user.ID, user.Name, user.Email, user.Age)
	return err
}

func GetUserByID(db *sql.DB, id string) (model.User, error) {
	var user model.User
	err := db.QueryRow(
		"SELECT id, name, email, age FROM users WHERE id = ?",
		id,
	).Scan(&user.ID, &user.Name, &user.Email, &user.Age)

	return user, err
}
