package dao

import (
	"database/sql"
	"hackathon-backend/model"
)

// 商品登録（POST）
func InsertProduct(db *sql.DB, p model.Product) error {
	_, err := db.Exec(`
        INSERT INTO products (id, seller_id, title, price, description, status, created_at)
        VALUES (?, ?, ?, ?, ?, ?, NOW())
    `,
		p.ID, p.SellerID, p.Title, p.Price, p.Description, p.Status,
	)
	return err
}

// 商品一覧取得（GET）
func SelectAllProducts(db *sql.DB) ([]model.Product, error) {
	rows, err := db.Query(`
        SELECT id, seller_id, title, price, description, status, created_at
        FROM products
        ORDER BY created_at DESC
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]model.Product, 0)

	for rows.Next() {
		var p model.Product
		if err := rows.Scan(
			&p.ID,
			&p.SellerID,
			&p.Title,
			&p.Price,
			&p.Description,
			&p.Status,
			&p.CreatedAt,
		); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}
