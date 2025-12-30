package interfaces

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"time"

	"github.com/bryanriosb/stock-info/internal/auth/domain"
	"github.com/bryanriosb/stock-info/internal/user/application"
	"github.com/bryanriosb/stock-info/shared"
	"github.com/bryanriosb/stock-info/shared/response"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type Handler struct {
	db          *gorm.DB
	jwtConfig   shared.JWTConfig
	userUseCase application.UserUseCase
}

func NewHandler(db *gorm.DB, jwtConfig shared.JWTConfig, userUseCase application.UserUseCase) *Handler {
	return &Handler{
		db:          db,
		jwtConfig:   jwtConfig,
		userUseCase: userUseCase,
	}
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
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

	accessToken, err := h.generateAccessToken(user.Username, user.Email, string(user.Role))
	if err != nil {
		return response.InternalError(c, "Failed to generate access token")
	}

	refreshToken, err := h.createRefreshToken(user.ID)
	if err != nil {
		return response.InternalError(c, "Failed to generate refresh token")
	}

	return response.Success(c, fiber.Map{
		"access_token":       accessToken,
		"refresh_token":      refreshToken,
		"expires_in":         h.jwtConfig.Expiration.Seconds(),
		"refresh_expires_in": h.jwtConfig.RefreshExpiration.Seconds(),
	})
}

func (h *Handler) Refresh(c *fiber.Ctx) error {
	var req RefreshRequest
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	if req.RefreshToken == "" {
		return response.BadRequest(c, "Refresh token is required")
	}

	var storedToken domain.RefreshToken
	if err := h.db.Where("token = ?", req.RefreshToken).First(&storedToken).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.Unauthorized(c, "Invalid refresh token")
		}
		return response.InternalError(c, "Failed to validate refresh token")
	}

	if !storedToken.IsValid() {
		return response.Unauthorized(c, "Refresh token expired or revoked")
	}

	user, err := h.userUseCase.GetByID(c.Context(), storedToken.UserID)
	if err != nil {
		return response.Unauthorized(c, "User not found")
	}

	// Revoke old refresh token
	h.db.Model(&storedToken).Update("revoked", true)

	accessToken, err := h.generateAccessToken(user.Username, user.Email, string(user.Role))
	if err != nil {
		return response.InternalError(c, "Failed to generate access token")
	}

	newRefreshToken, err := h.createRefreshToken(user.ID)
	if err != nil {
		return response.InternalError(c, "Failed to generate refresh token")
	}

	return response.Success(c, fiber.Map{
		"access_token":       accessToken,
		"refresh_token":      newRefreshToken,
		"expires_in":         h.jwtConfig.Expiration.Seconds(),
		"refresh_expires_in": h.jwtConfig.RefreshExpiration.Seconds(),
	})
}

func (h *Handler) Logout(c *fiber.Ctx) error {
	var req RefreshRequest
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	if req.RefreshToken != "" {
		h.db.Model(&domain.RefreshToken{}).Where("token = ?", req.RefreshToken).Update("revoked", true)
	}

	return response.Success(c, fiber.Map{
		"message": "Logged out successfully",
	})
}

func (h *Handler) generateAccessToken(username, email, role string) (string, error) {
	claims := jwt.MapClaims{
		"sub":   username,
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(h.jwtConfig.Expiration).Unix(),
		"iat":   time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(h.jwtConfig.Secret))
}

func (h *Handler) createRefreshToken(userID int64) (string, error) {
	tokenBytes := make([]byte, 32)
	if _, err := rand.Read(tokenBytes); err != nil {
		return "", err
	}
	tokenString := base64.URLEncoding.EncodeToString(tokenBytes)

	refreshToken := domain.RefreshToken{
		UserID:    userID,
		Token:     tokenString,
		ExpiresAt: time.Now().Add(h.jwtConfig.RefreshExpiration),
		Revoked:   false,
	}

	if err := h.db.Create(&refreshToken).Error; err != nil {
		return "", err
	}

	return tokenString, nil
}
