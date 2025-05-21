package server

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/vince-II/auth-post-api/server/handlers"
)

func middleware(c *fiber.Ctx) error {
	log.Println("Middleware executed")
	return c.Next()
}

func NewServer(ctx context.Context) *fiber.App {
	app := fiber.New()

	// prefix all routes with /api -> it will run on a middleware first
	api := app.Group("/api", middleware)

	// api.Get("/health", handlers.HealthCheck())

	// api/v1
	v1 := api.Group("/v1", middleware)
	v1.Get("/health", handlers.HealthCheck())

	return app
}
