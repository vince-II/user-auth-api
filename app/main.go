package main

import (
	"context"
	"log"

	database "github.com/vince-II/auth-post-api/db"
	"github.com/vince-II/auth-post-api/db/sqlc"
	"github.com/vince-II/auth-post-api/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// set up config for db
	dbPool := database.Config(ctx)
	conn := sqlc.New(dbPool)

	// if err != nil {
	// 	log.Fatalf("failed to create database: %v", err)
	// }

	app := server.NewServer(ctx, conn)
	log.Fatal(app.Listen(":3000"))

	// Start the server on port 3000
}
