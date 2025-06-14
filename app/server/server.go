package server

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/vince-II/auth-post-api/internal/sqlc"
	"github.com/vince-II/auth-post-api/server/handlers"
)

// func middleware(c *fiber.Ctx) error {
// 	// validations for input
// 	log.Println("Middleware executed")
// 	return c.Next()
// }

// passing the database connection to the server
func NewServer(ctx context.Context, conn *sqlc.Queries) *fiber.App {
	app := fiber.New()

	// prefix all routes with /api -> it will run on a middleware first
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Post("/register", handlers.RegisterUserHandler(conn))

	// v1.Get("/posts", handlers.GetPostsHandler(conn))
	v1.Get("/health", handlers.HealthCheck())
	return app
}
