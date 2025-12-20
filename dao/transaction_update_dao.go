package dao

import "database/sql"

func UpdateTransactionStatus(db *sql.DB, id string, status string) error {
	_, err := db.Exec(`
        UPDATE transactions
        SET status = ?
        WHERE id = ?
    `, status, id)
	return err
}
