package main

import (
	"context"
	"log"

	"github.com/vince-II/auth-post-api/app/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app := server.NewServer(ctx)
	log.Fatal(app.Listen(":3000"))

	// Start the server on port 3000
}
