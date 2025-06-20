package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/vince-II/auth-post-api/server/util"
)

func AuthenticateToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.GetRespHeader("Authorization")
		if authHeader == "" {
			return util.SendError(c, fiber.StatusUnauthorized, "Missing Authorization Token")
		}

		claims, err := util.VerifyToken(authHeader)
		if err != nil {
			log.Errorf("Token verification failed: %v", err)

			return util.SendError(c, fiber.StatusForbidden, err.Error())
		}

		// saves in request context
		c.Locals("user_id", claims.UserID)
		c.Locals("username", claims.Username)

		return c.Next()
	}
}
