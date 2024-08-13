package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Gateway(permission ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user")
		if user == nil {
			return fiber.ErrUnauthorized
		}

		token, ok := user.(*jwt.Token)
		if !ok {
			return fiber.ErrUnauthorized
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return fiber.ErrUnauthorized
		}

		permissions, ok := claims["permissions"].([]interface{})
		if !ok {
			return fiber.ErrUnauthorized
		}

		if len(permission) == 0 {
			return c.Next()
		}

		// Has Permission
		for _, p := range permissions {
			if p == permission[0] {
				return c.Next()
			}
		}

		return fiber.ErrForbidden
	}
}
