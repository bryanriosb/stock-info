package interfaces

import (
	"errors"
	"strconv"

	"github.com/bryanriosb/stock-info/internal/user/application"
	"github.com/bryanriosb/stock-info/shared/response"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	useCase application.UserUseCase
}

func NewHandler(useCase application.UserUseCase) *Handler {
	return &Handler{useCase: useCase}
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role,omitempty"`
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var req CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	if req.Username == "" || req.Email == "" || req.Password == "" {
		return response.BadRequest(c, "Username, email and password are required")
	}

	user, err := h.useCase.Create(c.Context(), application.CreateUserRequest{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		if errors.Is(err, application.ErrUserAlreadyExists) {
			return response.BadRequest(c, "User already exists")
		}
		return response.InternalError(c, "Failed to create user")
	}

	return response.Created(c, user)
}

func (h *Handler) GetAll(c *fiber.Ctx) error {
	users, err := h.useCase.GetAll(c.Context())
	if err != nil {
		return response.InternalError(c, "Failed to fetch users")
	}
	return response.Success(c, users)
}

func (h *Handler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return response.BadRequest(c, "Invalid user ID")
	}

	user, err := h.useCase.GetByID(c.Context(), id)
	if err != nil {
		if errors.Is(err, application.ErrUserNotFound) {
			return response.NotFound(c, "User not found")
		}
		return response.InternalError(c, "Failed to fetch user")
	}

	return response.Success(c, user)
}

func (h *Handler) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return response.BadRequest(c, "Invalid user ID")
	}

	var req UpdateUserRequest
	if parseErr := c.BodyParser(&req); parseErr != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	user, err := h.useCase.Update(c.Context(), id, application.UpdateUserRequest{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	})
	if err != nil {
		if errors.Is(err, application.ErrUserNotFound) {
			return response.NotFound(c, "User not found")
		}
		if errors.Is(err, application.ErrLastAdmin) {
			return response.BadRequest(c, "Cannot remove the last admin")
		}
		return response.InternalError(c, "Failed to update user")
	}

	return response.Success(c, user)
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return response.BadRequest(c, "Invalid user ID")
	}

	if err := h.useCase.Delete(c.Context(), id); err != nil {
		if errors.Is(err, application.ErrUserNotFound) {
			return response.NotFound(c, "User not found")
		}
		if errors.Is(err, application.ErrLastAdmin) {
			return response.BadRequest(c, "Cannot delete the last admin")
		}
		return response.InternalError(c, "Failed to delete user")
	}

	return response.Success(c, fiber.Map{"message": "User deleted"})
}
