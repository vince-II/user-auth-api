package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/vince-II/auth-post-api/internal/handlers"
)

func main() {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// initialiaze the fiber instance
	app := fiber.New()

	// connect to the database
	// db, err := database.Connect()
	// if err != nil {
	// 	log.Fatal("Error connecting to the database")
	// }

	// prefix
	api := app.Group("/api")

	// register routes
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/posts", handlers.GetUserPosts)

	app.Listen(":3000")
}
