package model

import "database/sql"

type Product struct {
	ID          string
	SellerID    sql.NullString
	Title       sql.NullString
	Price       sql.NullInt64
	Description sql.NullString
	Status      string
	CreatedAt   sql.NullTime
	ImageURL    sql.NullString
}
