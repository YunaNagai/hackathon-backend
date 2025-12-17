package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/oklog/ulid/v2"
)

type UserResForHTTPGet struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// ① GoプログラムからMySQLへ接続
var db *sql.DB

func init() {
	// ①-1
	mysqlUser := os.Getenv("DB_USER")
	mysqlPwd := os.Getenv("DB_USERPWD")
	instanceConn := os.Getenv("INSTANCE_CONNECTION_NAME")
	mysqlDatabase := os.Getenv("DB_DATABASE")
	connStr := fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/%s", mysqlUser, mysqlPwd, instanceConn, mysqlDatabase)
	_db, err := sql.Open("mysql", connStr)
	// ①-2
	// ①-3
	if err != nil {
		log.Fatalf("fail: sql.Open, %v\n", err)
	}
	if err := _db.Ping(); err != nil {
		log.Fatalf("fail: _db.Ping, %v\n", err)
	}
	db = _db
}

func newULID() string {
	t := time.Now().UTC()
	entropy := ulid.Monotonic(rand.Reader, 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}

// ② /userでリクエストされたらnameパラメーターと一致する名前を持つレコードをJSON形式で返す
func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// ②-1
		name := r.URL.Query().Get("name") // To be filled
		if name == "" {
			log.Println("fail: name is empty")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// ②-2
		rows, err := db.Query("SELECT id, name, age FROM `user` WHERE name = ?", name)
		if err != nil {
			log.Printf("fail: db.Query, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		// ②-3
		users := make([]UserResForHTTPGet, 0)
		for rows.Next() {
			var u UserResForHTTPGet
			if err := rows.Scan(&u.Id, &u.Name, &u.Age); err != nil {
				log.Printf("fail: rows.Scan, %v\n", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			users = append(users, u)
		}

		// ②-4
		bytes, err := json.Marshal(users)
		if err != nil {
			log.Printf("fail: json.Marshal, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)
	case http.MethodPost:
		var req struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Printf("fail: json.NewDecoder, %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if req.Name == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if len(req.Name) > 50 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if req.Age < 20 || req.Age > 80 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		id := newULID()
		tx, err := db.Begin()
		if err != nil {
			log.Printf("fail: db.Begin, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = tx.Exec("INSERT INTO `user` (id,name, age) VALUES (?,?,?)", id, req.Name, req.Age)
		if err != nil {
			log.Printf("fail: db.Exec, %v\n", err)
			if rbErr := tx.Rollback(); rbErr != nil {
				log.Printf("fail: tx.Rollback, %v\n", rbErr)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err := tx.Commit(); err != nil {
			log.Printf("fail: tx.Commit, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"id": id})

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
func handleProductPost(w http.ResponseWriter, r *http.Request) {
	var req struct {
		SellerID    string `json:"seller_id"`
		Title       string `json:"title"`
		Price       int    `json:"price"`
		Description string `json:"description"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("fail: json.Decode, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.SellerID == "" || req.Title == "" || req.Price <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	id := newULID()

	_, err = tx.Exec(`
        INSERT INTO products (id, seller_id, title, price, description, status, created_at)
        VALUES (?, ?, ?, ?, ?, 'selling', NOW())
    `, id, req.SellerID, req.Title, req.Price, req.Description)

	if err != nil {
		log.Printf("fail: db.Exec, %v\n", err)
		tx.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := tx.Commit(); err != nil {
		log.Printf("fail: tx.Commit, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}
func handleProductGet(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`
        SELECT id, seller_id, title, price, description, status, created_at
        FROM products
        ORDER BY created_at DESC
    `)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type Product struct {
		ID          string `json:"id"`
		SellerID    string `json:"seller_id"`
		Title       string `json:"title"`
		Price       int    `json:"price"`
		Description string `json:"description"`
		Status      string `json:"status"`
		CreatedAt   string `json:"created_at"`
	}

	products := make([]Product, 0)

	for rows.Next() {
		var p Product
		if err := rows.Scan(
			&p.ID, &p.SellerID, &p.Title, &p.Price,
			&p.Description, &p.Status, &p.CreatedAt,
		); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		products = append(products, p)
	}

	bytes, err := json.Marshal(products)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}
func productsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleProductGet(w, r)
	case http.MethodPost:
		handleProductPost(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func main() {
	// ② /userでリクエストされたらnameパラメーターと一致する名前を持つレコードをJSON形式で返す
	http.HandleFunc("/user", handler)
	http.HandleFunc("/products", productsHandler)

	// ③ Ctrl+CでHTTPサーバー停止時にDBをクローズする
	closeDBWithSysCall()

	// 8000番ポートでリクエストを待ち受ける
	log.Println("Listening...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

// ③ Ctrl+CでHTTPサーバー停止時にDBをクローズする
func closeDBWithSysCall() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		s := <-sig
		log.Printf("received syscall, %v", s)

		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
		log.Printf("success: db.Close()")
		os.Exit(0)
	}()
}
