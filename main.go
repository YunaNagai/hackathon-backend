package main

import (
	"hackathon-backend/controller"
	"hackathon-backend/db"
	"hackathon-backend/middleware"
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
	r.Use(middleware.CORS)

	r.Post("/user", controller.RegisterUserHandler(database))
	r.Get("/user/{id}", controller.GetUserHandler(database))

	r.Route("/products", func(r chi.Router) {
		r.Get("/", controller.ProductsHandler(database))
		r.Post("/", controller.ProductsHandler(database))
	})
	r.Get("/products/{id}", controller.GetProductByID(database))

	// 🔥 重複を削除して1つに統一
	r.Route("/transactions", func(r chi.Router) {
		r.Get("/", withCORS(controller.GetTransactionsHandler(database)))
		r.Post("/", withCORS(controller.CreateTransactionHandler(database)))
	})

	r.Get("/transactions/{id}", withCORS(controller.GetTransactionByIDHandler(database)))

	r.Options("/transactions/{id}", withCORS(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	r.Put("/transactions/{id}", withCORS(func(w http.ResponseWriter, r *http.Request) {
		usecase.UpdateTransaction(database, w, r)
	}))
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
