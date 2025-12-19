package model

type Message struct {
	ID            int64  `json:"id"`
	TransactionID string `json:"transactionId"`
	UserName      string `json:"userName"`
	Text          string `json:"text"`
	CreatedAt     string `json:"createdAt"`
}
