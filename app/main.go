package main

import (
	"context"
	"log"

	"github.com/vince-II/auth-post-api/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Uncomment the following lines to set up the database connection
	// and handle any errors that may occur.
	// set up config for db
	// conf := NewConfiguration()
	// db, err := database.NewDatabase(ctx, conf.DatabaseURL)

	// if err != nil {
	// 	log.Fatalf("failed to create database: %v", err)
	// }

	app := server.NewServer(ctx)
	log.Fatal(app.Listen(":3000"))

	// Start the server on port 3000
}
