package main

import (
	"binarydropclient/indexer"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() error {
	return godotenv.Load(".env")
}

func main() {
	err := loadEnv()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	//init db
	_, err = indexer.InitDB(os.Getenv("DATABASE_NAME"))
	if err != nil {
		log.Fatal("Failed to init DB:", err)
	}
}
