package main

import (
	"hackathon-backend/controller"
	"hackathon-backend/db"
	"hackathon-backend/middleware"

	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
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
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}
}
