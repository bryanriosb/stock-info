package interfaces

import (
	"errors"
	"time"

	"github.com/bryanriosb/stock-info/internal/user/application"
	"github.com/bryanriosb/stock-info/shared"
	"github.com/bryanriosb/stock-info/shared/response"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Handler struct {
	jwtConfig   shared.JWTConfig
	userUseCase application.UserUseCase
}

func NewHandler(jwtConfig shared.JWTConfig, userUseCase application.UserUseCase) *Handler {
	return &Handler{
		jwtConfig:   jwtConfig,
		userUseCase: userUseCase,
	}
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

	user, err := h.userUseCase.Authenticate(c.Context(), req.Username, req.Password)
	if err != nil {
		if errors.Is(err, application.ErrInvalidCredentials) {
			return response.Unauthorized(c, "Invalid credentials")
		}
		return response.InternalError(c, "Authentication failed")
	}

	token, err := h.generateToken(user.Username)
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
