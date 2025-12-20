package main

import (
	"hackathon-backend/controller"
	"hackathon-backend/db"
	"hackathon-backend/usecase"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func main() {
	database := db.Connect()
	defer database.Close()

	r := chi.NewRouter()

	r.Post("/user", withCORS(controller.RegisterUserHandler(database)))
	r.Get("/user/{id}", withCORS(controller.GetUserHandler(database)))

	r.Route("/products", func(r chi.Router) {
		r.Get("/", withCORS(controller.ProductsHandler(database)))
		r.Post("/", withCORS(controller.ProductsHandler(database)))
	})
	r.Get("/products/{id}", withCORS(controller.GetProductByID(database)))

	r.Route("/transactions", func(r chi.Router) {
		r.Get("/", withCORS(controller.GetTransactionsHandler(database)))
		r.Post("/", withCORS(controller.CreateTransactionHandler(database)))

		r.Options("/{id}", withCORS(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		r.Get("/{id}", withCORS(controller.GetTransactionByIDHandler(database)))
		r.Put("/{id}", withCORS(func(w http.ResponseWriter, r *http.Request) {
			usecase.UpdateTransaction(database, w, r)
		}))
	})

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
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS") // ← PUT を追加！
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		h(w, r)
	}
}
