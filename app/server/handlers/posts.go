package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/vince-II/auth-post-api/server/util"
)

func CreatePost(ctx context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {

		return util.SendResponse(c, fiber.StatusOK, "", "Post has been created")
	}
}
