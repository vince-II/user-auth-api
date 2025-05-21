package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// returns a handler function
func HealthCheck() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	}
}
