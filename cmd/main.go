package main

import (
	"log"

	"github.com/Dzikuri/shopifyx/internal/config"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func main() {

	// Setup Database
	dbConn, err := config.NewDatabase()
	if err != nil {
		log.Fatal("Failed to open connection to database", err)
	}

	echo := config.NewEcho()

	config.Bootstrap(&config.BootstrapConfig{
		DB:  dbConn,
		App: echo,
		// Validate: ,
	})

	echo.Logger.Fatal(echo.Start(":8000"))
}
