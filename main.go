package main

import (
	"hackathon-backend/controller"
	"hackathon-backend/db"
	"hackathon-backend/middleware"
	"os"

	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	database := db.Connect()
	defer database.Close()

	r := chi.NewRouter()
	r.Use(middleware.CORS)

	r.Post("/user", controller.RegisterUserHandler(database))
	r.Get("/products", controller.ProductsHandler(database))
	r.Get("/products/{id}", controller.GetProductByID(database))

	log.Println("Listening...")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
