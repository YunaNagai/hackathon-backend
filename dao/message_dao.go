package dao

import (
	"database/sql"
	"hackathon-backend/model"
	"log"
)

func GetAllMessages(db *sql.DB) ([]model.Message, error) {
	rows, err := db.Query(`
        SELECT id, transaction_id, user_name, message, created_at
        FROM messages
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	messages := []model.Message{}

	for rows.Next() {
		var m model.Message
		if err := rows.Scan(
			&m.ID,
			&m.TransactionID,
			&m.UserName,
			&m.Message,
			&m.CreatedAt,
		); err != nil {
			return nil, err
		}
		messages = append(messages, m)
	}

	return messages, nil
}

func InsertMessage(db *sql.DB, msg model.Message) error {
	_, err := db.Exec(`
        INSERT INTO messages (id, transaction_id, user_name, message, created_at)
        VALUES ($1, $2, $3, $4, $5)
    `,
		msg.ID,
		msg.TransactionID,
		msg.UserName,
		msg.Message,
		msg.CreatedAt,
	)

	if err != nil {
		log.Printf("SQL Insert error: %v", err) // ← ここが一番重要！
	}

	return err
}
