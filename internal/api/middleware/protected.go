package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"strings"
	"time"
)

func Protected(viper *viper.Viper) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return fiber.ErrUnauthorized
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 && parts[0] != "Bearer" {
			return fiber.ErrUnauthorized
		}

		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(viper.GetString("jwt.secret")), nil
		})
		if err != nil || !token.Valid {
			return fiber.ErrUnauthorized
		}

		claims := token.Claims.(jwt.MapClaims)
		expTime := time.Unix(int64(claims["exp"].(float64)), 0)
		currentTime := time.Now()
		if currentTime.After(expTime) {
			return fiber.ErrUnauthorized
		}

		c.Locals("user", token)
		return c.Next()
	}
}
