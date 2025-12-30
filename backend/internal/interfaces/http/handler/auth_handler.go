package handler

import (
	"time"

	"github.com/bryanriosb/stock-info/pkg/config"
	"github.com/bryanriosb/stock-info/pkg/response"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AuthHandler struct {
	jwtConfig config.JWTConfig
}

func NewAuthHandler(jwtConfig config.JWTConfig) *AuthHandler {
	return &AuthHandler{jwtConfig: jwtConfig}
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	// Simple auth for testing purposes
	// In production, validate against a user database
	if req.Username == "" || req.Password == "" {
		return response.BadRequest(c, "Username and password are required")
	}

	// For demo: accept any non-empty credentials
	// Replace with real authentication logic
	if req.Username != "admin" || req.Password != "admin" {
		return response.Unauthorized(c, "Invalid credentials")
	}

	token, err := h.generateToken(req.Username)
	if err != nil {
		return response.InternalError(c, "Failed to generate token")
	}

	return response.Success(c, fiber.Map{
		"token":      token,
		"expires_in": h.jwtConfig.Expiration.Seconds(),
	})
}

func (h *AuthHandler) generateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"sub": username,
		"exp": time.Now().Add(h.jwtConfig.Expiration).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(h.jwtConfig.Secret))
}
