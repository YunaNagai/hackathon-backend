package model

type Message struct {
	ID            int64  `json:"id"`
	TransactionID string `json:"transactionId"`
	UserName      string `json:"userName"`
	Message       string `json:"message"`
	CreatedAt     string `json:"createdAt"`
}
