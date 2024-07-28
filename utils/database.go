package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() *sql.DB {
	if DB == nil {
		log.Fatal("Database connection is not initialized.")
	}
	return DB
}

// func InitDB() {
// 	var err error
// 	DB, err = sql.Open("mysql", "root:@tcp(localhost:3306)/merchant_pockets")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = DB.Ping()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func GetDB() *sql.DB {
// 	if DB == nil {
// 		log.Fatal("Database connection is not initialized.")
// 	}
// 	return DB
// }
