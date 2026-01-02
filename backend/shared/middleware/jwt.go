package middleware

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

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
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	return claims["sub"].(string)
}

func GetRoleFromToken(c *fiber.Ctx) string {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	if role, ok := claims["role"].(string); ok {
		return role
	}
	return "user"
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
