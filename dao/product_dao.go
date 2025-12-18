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
        SELECT
            id,
            COALESCE(seller_id, ''),
            COALESCE(title, ''),
            COALESCE(price, 0),
            COALESCE(description, ''),
            status,
            COALESCE(created_at, NOW()),
            COALESCE(image_url, '')
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
			&p.ImageURL,
		); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}
func GetProductByID(db *sql.DB, id string) (*model.Product, error) {
	var p model.Product
	query := `
		SELECT id, seller_id, title, price, description, status, created_at, image_url
		FROM products
		WHERE id = ?
`
	err := db.QueryRow(query, id).Scan(&p.ID, &p.SellerID, &p.Title, &p.Price, &p.Description,
		&p.Status, &p.CreatedAt, &p.ImageURL)

	if err != nil {
		return nil, err
	}

	return &p, nil
}
