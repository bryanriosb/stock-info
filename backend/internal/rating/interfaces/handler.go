package interfaces

import (
	"context"
	"log"
	"net/http"

	"github.com/bryanriosb/stock-info/internal/rating/domain"
	"github.com/bryanriosb/stock-info/shared/response"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	repo domain.RatingOptionRepository
}

func NewHandler(repo domain.RatingOptionRepository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) GetAllRatingOptions(c *fiber.Ctx) error {
	ctx := context.Background()

	options, err := h.repo.FindAll(ctx)
	if err != nil {
		log.Printf("Error fetching rating options: %v", err)
		return response.Error(c, http.StatusInternalServerError, "Failed to get rating options")
	}
	return response.Success(c, options)
}
