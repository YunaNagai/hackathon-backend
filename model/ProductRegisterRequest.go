package model

type ProductRegisterRequest struct {
	SellerID    string `json:"sellerId"`
	Title       string `json:"title"`
	Price       int64  `json:"price"`
	Description string `json:"description"`
	ImageURL    string `json:"imageUrl"`
}
