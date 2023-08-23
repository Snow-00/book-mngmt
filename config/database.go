package config

import (
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var DB *sql.DB

func ConnectDB() {
	connection := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?parseTime=true",
		ENV.DB_USER, ENV.DB_PASSWORD, ENV.DB_HOST, ENV.DB_PORT, ENV.DB_NAME,
	)

	db, err := sql.Open("mysql", connection)
	if err != nil {
		log.Fatal(err)
	}

	DB = db
	log.Println("Database Connected")
}