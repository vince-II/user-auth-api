package handlers

import (
	"context"

	"github.com/vince-II/auth-post-api/app/server/dto"
	"github.com/vince-II/auth-post-api/app/server/services"
	"github.com/vince-II/auth-post-api/app/server/util"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func RegisterUser(ctx context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user dto.RegisterUser

		if err := c.BodyParser(&user); err != nil {
			log.Errorf(err.Error())
			return util.SendError(c, fiber.ErrBadRequest.Code, "Invalid request body")
		}

		d, err := services.RegisterUser(user, ctx)
		if err != nil {
			log.Errorf(err.Error())
			return util.SendError(c, fiber.StatusInternalServerError, "Failed to register user")
		}

		return util.SendResponse(c, 200, d, "User registered successfully")
	}
}

func LoginUser(ctx context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user dto.LoginUser

		if err := c.BodyParser(&user); err != nil {
			log.Errorf(err.Error())
			return util.SendError(c, fiber.ErrBadRequest.Code, "Invalid request body")
		}

		d, err := services.LoginUser(user, ctx)

		if err != nil {
			log.Errorf(err.Error())
			if err.Error() == "Username not found" {
				return util.SendError(c, fiber.StatusNotFound, "Username not found")
			}
			return util.SendError(c, fiber.StatusInternalServerError, "Failed to login user")
		}

		// todo: prep the response data
		// todo: prep the data type

		return util.SendResponse(c, fiber.StatusOK, d, "User logged in successfully")
	}
}
