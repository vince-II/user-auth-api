package server

import (
	"context"

	"github.com/vince-II/auth-post-api/app/server/handlers"

	"github.com/gofiber/fiber/v2"
)

func NewServer(ctx context.Context) *fiber.App {
	app := fiber.New()

	api := app.Group("/api")

	v1 := api.Group("/v1")
	// TODO: cors middleware
	v1.Post("/register", handlers.RegisterUser(ctx))
	v1.Get("/health", handlers.HealthCheck())

	return app
}
