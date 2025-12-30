package interfaces

import (
	"github.com/bryanriosb/stock-info/internal/recommendation/application"
	"github.com/bryanriosb/stock-info/shared/response"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	useCase application.RecommendationUseCase
}

func NewHandler(useCase application.RecommendationUseCase) *Handler {
	return &Handler{useCase: useCase}
}

func (h *Handler) GetRecommendations(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 10)

	recommendations, err := h.useCase.GetRecommendations(c.Context(), limit)
	if err != nil {
		return response.InternalError(c, "Failed to fetch recommendations")
	}

	return response.Success(c, recommendations)
}
