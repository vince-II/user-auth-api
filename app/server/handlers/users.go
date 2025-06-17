package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/vince-II/auth-post-api/internal/sqlc"
	"github.com/vince-II/auth-post-api/server/dto"
	"github.com/vince-II/auth-post-api/server/services"
	"github.com/vince-II/auth-post-api/server/util"
)

func RegisterUser(conn *sqlc.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user dto.RegisterUser

		if err := c.BodyParser(&user); err != nil {
			log.Println("Error parsing request body:", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		d, err := services.RegisterUser(user, conn)
		if err != nil {
			util.LogError(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to register user",
			})
		}

		return util.SendResponse(c, 200, d, "User registered successfully")
	}
}
