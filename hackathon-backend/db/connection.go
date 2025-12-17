package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	mysqlUser := os.Getenv("DB_USER")
	mysqlPwd := os.Getenv("DB_USERPWD")
	instanceConn := os.Getenv("INSTANCE_CONNECTION_NAME")
	mysqlDatabase := os.Getenv("DB_DATABASE")

	connStr := fmt.Sprintf(
		"%s:%s@unix(/cloudsql/%s)/%s",
		mysqlUser, mysqlPwd, instanceConn, mysqlDatabase,
	)

	database, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("fail: sql.Open, %v\n", err)
	}

	if err := database.Ping(); err != nil {
		log.Fatalf("fail: database.Ping, %v\n", err)
	}

	return database
}
