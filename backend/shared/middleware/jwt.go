package middleware

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

const defaultRole = "user"

func JWTProtected(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(secret)},
		// Support token from query param for SSE (EventSource doesn't support headers)
		Filter: func(c *fiber.Ctx) bool {
			if token := c.Query("token"); token != "" {
				c.Request().Header.Set("Authorization", "Bearer "+token)
			}
			return false // never skip validation
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"error":   "Invalid or expired token",
			})
		},
	})
}

func GetUserFromToken(c *fiber.Ctx) string {
	token, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return ""
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return ""
	}
	sub, ok := claims["sub"].(string)
	if !ok {
		return ""
	}
	return sub
}

func GetRoleFromToken(c *fiber.Ctx) string {
	token, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return defaultRole
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return defaultRole
	}
	if role, ok := claims["role"].(string); ok {
		return role
	}
	return defaultRole
}

func RequireAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		role := GetRoleFromToken(c)
		if role != "admin" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"success": false,
				"error":   "Admin access required",
			})
		}
		return c.Next()
	}
}
