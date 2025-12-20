package dao

import (
	"database/sql"
	"hackathon-backend/model"
)

func GetAllTransactions(db *sql.DB) ([]model.Transaction, error) {
	rows, err := db.Query(`
        SELECT id, product_id, buyer_id, seller_id, status, created_at
        FROM transactions
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []model.Transaction

	for rows.Next() {
		var t model.Transaction
		if err := rows.Scan(
			&t.ID,
			&t.ProductID,
			&t.BuyerID,
			&t.SellerID,
			&t.Status,
			&t.CreatedAt,
		); err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}

	return transactions, nil
}

func InsertTransaction(db *sql.DB, t model.Transaction) error {
	_, err := db.Exec(`
        INSERT INTO transactions (id, product_id, buyer_id, seller_id, status, created_at)
        VALUES (?, ?, ?, ?, ?, ?)
    `,
		t.ID,
		t.ProductID,
		t.BuyerID,
		t.SellerID,
		t.Status,
		t.CreatedAt,
	)
	return err
}
func GetTransactionByID(db *sql.DB, id string) (model.Transaction, error) {
	var tx model.Transaction
	err := db.QueryRow(`
        SELECT id, product_id, buyer_id, seller_id, status, created_at
        FROM transactions
        WHERE id = ?
    `, id).Scan(
		&tx.ID,
		&tx.ProductID,
		&tx.BuyerID,
		&tx.SellerID,
		&tx.Status,
		&tx.CreatedAt,
	)
	return tx, err
}
