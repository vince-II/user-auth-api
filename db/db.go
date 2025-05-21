package database

import (
	"context"
	"database/sql"
	"fmt" // for formatting strings
	"log" // for logging strings
)

func NewDatabase(ctx context.Context, dbURL string) (*sql.DB, error) {
	if dbURL == "" {
		log.Fatal("missing database URL")
		return nil, fmt.Errorf("missing database URL")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
