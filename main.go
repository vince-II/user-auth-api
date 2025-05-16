package main

import (
	"context"
	"log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// set up config for db
	conf := NewConfiguration()
	db, err := database.NewDatabase(ctx, conf.DatabaseURL)

	if err != nil {
		log.Fatalf("failed to create database: %v", err)
	}

}
