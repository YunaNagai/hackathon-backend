package model

type Product struct {
	ID          string `json:"id"`
	SellerID    string `json:"sellerId"`
	Title       string `json:"title"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	ImageURL    string `json:"imageUrl"`
}
