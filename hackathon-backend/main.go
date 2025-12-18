package main

import (
	"hackathon-backend/controller"
	"hackathon-backend/db"
	"hackathon-backend/middleware"
	"log"
	"net/http"
)

// deploy test
// deploy test
// deploy test

func main() {
	database := db.Connect()
	defer database.Close()
	router := http.NewServeMux()
	router.HandleFunc("/user", controller.RegisterUserHandler(database))
	router.HandleFunc("/products", controller.ProductsHandler(database))
	handler := middleware.CORS(router)

	log.Println("Listening...")
	if err := http.ListenAndServe(":8000", handler); err != nil {
		log.Fatal(err)
	}
}
