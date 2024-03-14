package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/Dzikuri/shopifyx/internal/config"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func main() {

	// // Setup Database
	// dbConn, err := config.NewDatabase()
	// if err != nil {
	// 	log.Fatal("Failed to open connection to database", err)
	// }

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

	echo := config.NewEcho()

	// Validator
	validator := config.NewValidator()

	config.Bootstrap(&config.BootstrapConfig{
		DB:       dbConn,
		App:      echo,
		Validate: validator,
	})

	echo.Logger.Fatal(echo.Start(":8000"))
}
