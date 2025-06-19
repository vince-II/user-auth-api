package server

import (
	"context"

	"github.com/vince-II/auth-post-api/server/handlers"

	"github.com/gofiber/fiber/v2"
)

func NewServer(ctx context.Context) *fiber.App {
	app := fiber.New()

	api := app.Group("/api")
	auth := api.Group("/auth")
	auth.Post("/register", handlers.RegisterUser(ctx))
	auth.Post("/login", handlers.LoginUser(ctx))

	// TODO: cors middleware
	v1 := api.Group("/v1")
	v1.Get("/health", handlers.HealthCheck())

	return app
}
