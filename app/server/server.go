package server

import (
	"context"

	"github.com/vince-II/auth-post-api/server/handlers"
	"github.com/vince-II/auth-post-api/server/middleware"

	"github.com/gofiber/fiber/v2"
)

func NewServer(ctx context.Context) *fiber.App {
	app := fiber.New()

	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/register", handlers.RegisterUser(ctx))
	auth.Post("/login", handlers.LoginUser(ctx))
	auth.Post("/logout", middleware.AuthenticateToken(), handlers.LogoutUser(ctx))

	v1 := api.Group("/v1")
	v1.Get("/health", handlers.HealthCheck())
	v1.Post("/post", middleware.AuthenticateToken(), handlers.CreatePost(ctx))

	// v1.Get("/posts/:id", middleware.AuthenticateToken(), handlers.CreatePost(ctx))
	v1.Put("/posts/:id", middleware.AuthenticateToken(), handlers.UpdatePost(ctx))
	v1.Delete("/posts/:id", middleware.AuthenticateToken(), handlers.DeletePost(ctx))

	return app
}
