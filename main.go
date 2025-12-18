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
	r.Get("/transactions", withCORS(controller.GetTransactionsHandler(database)))
	r.Get("/messages", withCORS(controller.GetMessagesHandler(database)))
	r.Post("/messages", withCORS(controller.CreateMessageHandler(database)))

	log.Println("Listening...")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}

func withCORS(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		h(w, r)
	}
}
