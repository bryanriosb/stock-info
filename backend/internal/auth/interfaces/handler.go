package interfaces

import (
	"time"

	"github.com/bryanriosb/stock-info/shared"
	"github.com/bryanriosb/stock-info/shared/response"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Handler struct {
	jwtConfig shared.JWTConfig
}

func NewHandler(jwtConfig shared.JWTConfig) *Handler {
	return &Handler{jwtConfig: jwtConfig}
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	if req.Username == "" || req.Password == "" {
		return response.BadRequest(c, "Username and password are required")
	}

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

func (h *Handler) generateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"sub": username,
		"exp": time.Now().Add(h.jwtConfig.Expiration).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(h.jwtConfig.Secret))
}
