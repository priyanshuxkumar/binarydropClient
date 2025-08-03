package indexer

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	schemaBytes, err := os.ReadFile("schema.sql")
	if err != nil {
		return nil, fmt.Errorf("failed to read schema file: %w", err)
	}

	schema := string(schemaBytes)

	_, err = db.Exec(schema)
	if err != nil {
		return nil, err
	}

	return db, nil
}
