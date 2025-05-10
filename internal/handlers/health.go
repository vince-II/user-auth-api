package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// HealthCheckHandler is a simple health check handler
func HealthCheckHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	}
}
