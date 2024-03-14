package config

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"

	_ "github.com/lib/pq"
)

var (
	// DB variable for connection DB postgresql
	DB *sql.DB
)

func NewDatabase() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")

	// connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)
	connection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, database)

	val := url.Values{}
	// val.Add("parseTime", "1")
	// val.Add("loc", "Asia/Jakarta")
	val.Add("sslmode", "disable")

	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())

	dbConn, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal("failed to open connection to database", err)
	}

	err = dbConn.Ping()
	if err != nil {
		log.Fatal("failed to ping database ", err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal("got error when closing the DB connection", err)
		}
	}()

	return dbConn, err
}

// func NewDatabase() (*sql.DB, error) {
// 	host := os.Getenv("DB_HOST")
// 	port := os.Getenv("DB_PORT")
// 	username := os.Getenv("DB_USERNAME")
// 	password := os.Getenv("DB_PASSWORD")
// 	database := os.Getenv("DB_NAME")

// 	// connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)
// 	connection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, database)
// 	// connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
// 	val := url.Values{}
// 	// val.Add("parseTime", "1")
// 	// val.Add("loc", "Asia/Jakarta")
// 	val.Add("sslmode", "disable")

// 	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
// 	dbConn, err := sql.Open(`postgres`, dsn)
// 	if err != nil {
// 		log.Fatal("failed to open connection to database", err)
// 	}
// 	err = dbConn.Ping()
// 	if err != nil {
// 		log.Fatal("failed to ping database ", err)
// 	}

// 	// defer func() {
// 	// 	err := dbConn.Close()
// 	// 	if err != nil {
// 	// 		log.Fatal("got error when closing the DB connection", err)
// 	// 	}
// 	// }()

// 	return dbConn, err
// }
