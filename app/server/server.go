package server

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/vince-II/auth-post-api/internal/sqlc"
	"github.com/vince-II/auth-post-api/server/handlers"
)

func NewServer(ctx context.Context, conn *sqlc.Queries) *fiber.App {
	app := fiber.New()

	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Post("/register", handlers.RegisterUserHandler(conn))

	v1.Get("/health", handlers.HealthCheck())
	return app
}
