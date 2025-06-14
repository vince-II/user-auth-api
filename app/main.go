package main

import (
	"context"
	"log"

	database "github.com/vince-II/auth-post-api/database"
	"github.com/vince-II/auth-post-api/internal/sqlc"
	"github.com/vince-II/auth-post-api/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// set up config for db
	dbPool := database.Config(ctx)
	conn := sqlc.New(dbPool)

	app := server.NewServer(ctx, conn)
	log.Fatal(app.Listen(":3000"))

	// Start the server on port 3000
}
