package server

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/vince-II/auth-post-api/db/sqlc"
	"github.com/vince-II/auth-post-api/server/handlers"
)

func middleware(c *fiber.Ctx) error {
	// validations for input
	log.Println("Middleware executed")
	return c.Next()
}

// passing the database connection to the server
func NewServer(ctx context.Context, conn *sqlc.Queries) *fiber.App {
	app := fiber.New()

	// prefix all routes with /api -> it will run on a middleware first
	api := app.Group("/api", middleware)

	api.Get("/health", handlers.HealthCheck())
	api.Get("/status", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})
	// api/v1
	// v1 := api.Group("/v1", middleware)
	// v1.Get("/health", handlers.HealthCheck())

	return app
}
