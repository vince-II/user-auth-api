package util

import "github.com/gofiber/fiber/v2"

func SendResponse(c *fiber.Ctx, status int, data interface{}, message string) error {
	if data == nil {
		data = fiber.Map{}
	}

	response := fiber.Map{
		"status":  status,
		"message": message,
		"data":    data,
	}

	return c.Status(status).JSON(response)
}

func SendError(c *fiber.Ctx, status int, message string) error {
	response := fiber.Map{
		"status":  status,
		"message": message,
	}

	return c.Status(status).JSON(response)
}
