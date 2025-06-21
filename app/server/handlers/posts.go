package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/vince-II/auth-post-api/server/dto"
	"github.com/vince-II/auth-post-api/server/services"
	"github.com/vince-II/auth-post-api/server/util"
)

func CreatePost(ctx context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var data dto.PostParams

		if err := c.BodyParser(&data); err != nil {
			log.Errorf("Failed to parse data", err)
			return util.SendError(c, fiber.ErrBadRequest.Code, "Invalid request body")
		}

		userID := c.Locals("user_id").(int32)

		d, err := services.CreatePost(userID, data, ctx)
		if err != nil {
			log.Errorf("Failed to create post %v", err.Error())
			return util.SendError(c, fiber.StatusInternalServerError, "Failed to register user")
		}

		return util.SendResponse(c, fiber.StatusOK, d, "Post has been created")
	}
}

func UpdatePost(ctx context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var data dto.PostParams

		if err := c.BodyParser(&data); err != nil {
			log.Errorf("Failed to parse data", err)
			return util.SendError(c, fiber.ErrBadRequest.Code, "Invalid request body")
		}

		userID := c.Locals("user_id").(int32)

		d, err := services.UpdatePost(userID, data, ctx)
		if err != nil {
			log.Errorf("Failed to create post %v", err.Error())
			return util.SendError(c, fiber.StatusInternalServerError, "Failed to register user")
		}

		return util.SendResponse(c, fiber.StatusOK, d, "Post has been updated")
	}
}

func DeletePost(ctx context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var data dto.PostParams

		if err := c.BodyParser(&data); err != nil {
			log.Errorf("Failed to parse data", err)
			return util.SendError(c, fiber.ErrBadRequest.Code, "Invalid request body")
		}

		userID := c.Locals("user_id").(int32)

		d, err := services.DeletePost(userID, data, ctx)
		if err != nil {
			log.Errorf("Failed to create post %v", err.Error())
			return util.SendError(c, fiber.StatusInternalServerError, "Failed to register user")
		}

		return util.SendResponse(c, fiber.StatusOK, d, "Post has been updated")
	}
}
