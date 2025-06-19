package main

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/vince-II/auth-post-api/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app := server.NewServer(ctx)
	log.Infof("Server starting...")
	log.Fatal(app.Listen(":3000"))

}
