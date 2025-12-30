package handler

import (
	"github.com/bryanriosb/stock-info/internal/application"
	"github.com/bryanriosb/stock-info/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type RecommendationHandler struct {
	useCase application.RecommendationUseCase
}

func NewRecommendationHandler(useCase application.RecommendationUseCase) *RecommendationHandler {
	return &RecommendationHandler{useCase: useCase}
}

func (h *RecommendationHandler) GetRecommendations(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 10)

	recommendations, err := h.useCase.GetRecommendations(c.Context(), limit)
	if err != nil {
		return response.InternalError(c, "Failed to fetch recommendations")
	}

	return response.Success(c, recommendations)
}
