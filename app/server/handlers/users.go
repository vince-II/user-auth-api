package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/vince-II/auth-post-api/internal/sqlc"
	"github.com/vince-II/auth-post-api/server/dto"
	"github.com/vince-II/auth-post-api/server/services"
	"github.com/vince-II/auth-post-api/server/util"
)

func RegisterUser(conn *sqlc.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user dto.RegisterUser

		if err := c.BodyParser(&user); err != nil {
			log.Errorf(err.Error())
			return util.SendError(c, fiber.ErrBadRequest.Code, "Invalid request body")
		}

		d, err := services.RegisterUser(user, conn)
		if err != nil {
			log.Errorf(err.Error())
			return util.SendError(c, fiber.StatusInternalServerError, "Failed to register user")
		}

		return util.SendResponse(c, 200, d, "User registered successfully")
	}
}

func LoginUser(conn *sqlc.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user dto.LoginUser

		if err := c.BodyParser(&user); err != nil {
			log.Errorf(err.Error())
			return util.SendError(c, fiber.ErrBadRequest.Code, "Invalid request body")
		}

		d, err := services.LoginUser(user, conn)
		if err != nil {
			log.Errorf(err.Error())
			if err.Error() == "Invalid credentials" {
				return util.SendError(c, fiber.StatusUnauthorized, "Invalid credentials")
			}
			return util.SendError(c, fiber.StatusInternalServerError, "Failed to login user")
		}

		if d == nil {
			return util.SendResponse(c, fiber.StatusOK, d, "User logged in successfully")
		}

		// todo: prep the response data
		// todo: prep the data type

		return util.SendResponse(c, fiber.StatusOK, d, "User logged in successfully")
	}
}
