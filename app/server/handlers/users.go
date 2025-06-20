package handlers

import (
	"context"

	"github.com/vince-II/auth-post-api/server/dto"
	"github.com/vince-II/auth-post-api/server/services"
	"github.com/vince-II/auth-post-api/server/util"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func RegisterUser(ctx context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user dto.RegisterUser

		if err := c.BodyParser(&user); err != nil {
			log.Errorf("Failed to parse data", err)
			return util.SendError(c, fiber.ErrBadRequest.Code, "Invalid request body")
		}

		d, err := services.RegisterUser(user, ctx)
		if err != nil {
			log.Errorf("Failed to register user due to%v", err)
			return util.SendError(c, fiber.StatusInternalServerError, "Failed to register user")
		}

		return util.SendResponse(c, 200, d, "User registered successfully")
	}
}

func LoginUser(ctx context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user dto.LoginUser

		if err := c.BodyParser(&user); err != nil {
			log.Errorf("failed to parse data %v", err)
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

		return util.SendResponse(c, fiber.StatusOK, d, "Logged in successfully")
	}
}

func LogoutUser(ctx context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {

		// type assertion
		userID := c.Locals("user_id").(int32)

		if err := services.LogoutUser(userID, ctx); err != nil {

		}

		return util.SendResponse(c, fiber.StatusOK, map[string]interface{}{}, "Logged out successfully")

	}
}
