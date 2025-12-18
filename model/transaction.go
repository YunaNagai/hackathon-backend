package model

type Transaction struct {
	ID        string `json:"id"`
	ProductID string `json:"productId"`
	BuyerID   string `json:"buyerId"`
	SellerID  string `json:"sellerId"`
	Status    string `json:"status"`
	CreatedAt string `json:"createdAt"`
}
