package handlers

import (
	"github.com/gofiber/fiber"
)

func GetUserPosts(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// TODO: Extract user ID from JWT
		// TODO: Query posts for that user
		return c.JSON(fiber.Map{
			"posts": []string{}, // stub
		})
	}
}

//
