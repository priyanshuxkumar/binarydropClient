package main

import (
	"binarydropclient/indexer"
	"binarydropclient/watcher"
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
	db, err := indexer.InitDB(os.Getenv("DATABASE_NAME"))
	if err != nil {
		log.Fatal("Failed to init DB:", err)
	}

	handler := &watcher.SyncHandler{
		DB:     db,
		Server: os.Getenv("SERVER_URL"),
	}
	watcher.Watcher(os.Getenv("BINARYDROP_SYNC_DIR"), handler)
}
