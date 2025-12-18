package model

type User struct {
	ID    string `json:"uid"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}
