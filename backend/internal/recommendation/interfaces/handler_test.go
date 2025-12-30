package interfaces

import (
	"context"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/bryanriosb/stock-info/internal/recommendation/domain"
	stockDomain "github.com/bryanriosb/stock-info/internal/stock/domain"
	"github.com/bryanriosb/stock-info/shared/response"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock RecommendationUseCase
type MockRecommendationUseCase struct {
	mock.Mock
}

func (m *MockRecommendationUseCase) GetRecommendations(ctx context.Context, limit int) ([]*domain.StockRecommendation, error) {
	args := m.Called(ctx, limit)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.StockRecommendation), args.Error(1)
}

func setupTestApp(handler *Handler) *fiber.App {
	app := fiber.New()
	app.Get("/recommendations", handler.GetRecommendations)
	return app
}

func TestGetRecommendations_Success(t *testing.T) {
	mockUC := new(MockRecommendationUseCase)
	handler := NewHandler(mockUC)
	app := setupTestApp(handler)

	recommendations := []*domain.StockRecommendation{
		{
			Stock:         &stockDomain.Stock{ID: 1, Ticker: "AAPL", Company: "Apple Inc."},
			Score:         0.85,
			Reason:        "Positive rating, Target price increased",
			PotentialGain: 20.0,
		},
		{
			Stock:         &stockDomain.Stock{ID: 2, Ticker: "GOOGL", Company: "Alphabet Inc."},
			Score:         0.65,
			Reason:        "Target price increased",
			PotentialGain: 15.0,
		},
	}

	mockUC.On("GetRecommendations", mock.Anything, 10).Return(recommendations, nil)

	req := httptest.NewRequest("GET", "/recommendations?limit=10", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var result response.Response
	json.NewDecoder(resp.Body).Decode(&result)

	assert.True(t, result.Success)
	assert.NotNil(t, result.Data)
	mockUC.AssertExpectations(t)
}

func TestGetRecommendations_DefaultLimit(t *testing.T) {
	mockUC := new(MockRecommendationUseCase)
	handler := NewHandler(mockUC)
	app := setupTestApp(handler)

	recommendations := []*domain.StockRecommendation{}
	mockUC.On("GetRecommendations", mock.Anything, 10).Return(recommendations, nil)

	req := httptest.NewRequest("GET", "/recommendations", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var result response.Response
	json.NewDecoder(resp.Body).Decode(&result)

	assert.True(t, result.Success)
	mockUC.AssertExpectations(t)
}

func TestGetRecommendations_CustomLimit(t *testing.T) {
	mockUC := new(MockRecommendationUseCase)
	handler := NewHandler(mockUC)
	app := setupTestApp(handler)

	recommendations := []*domain.StockRecommendation{
		{
			Stock:         &stockDomain.Stock{ID: 1, Ticker: "AAPL"},
			Score:         0.85,
			Reason:        "Positive",
			PotentialGain: 20.0,
		},
	}
	mockUC.On("GetRecommendations", mock.Anything, 5).Return(recommendations, nil)

	req := httptest.NewRequest("GET", "/recommendations?limit=5", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUC.AssertExpectations(t)
}

func TestGetRecommendations_Error(t *testing.T) {
	mockUC := new(MockRecommendationUseCase)
	handler := NewHandler(mockUC)
	app := setupTestApp(handler)

	mockUC.On("GetRecommendations", mock.Anything, 10).Return(nil, errors.New("database error"))

	req := httptest.NewRequest("GET", "/recommendations?limit=10", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

	var result response.Response
	json.NewDecoder(resp.Body).Decode(&result)

	assert.False(t, result.Success)
	assert.Contains(t, result.Error, "Failed to fetch recommendations")
	mockUC.AssertExpectations(t)
}

func TestGetRecommendations_Empty(t *testing.T) {
	mockUC := new(MockRecommendationUseCase)
	handler := NewHandler(mockUC)
	app := setupTestApp(handler)

	mockUC.On("GetRecommendations", mock.Anything, 10).Return([]*domain.StockRecommendation{}, nil)

	req := httptest.NewRequest("GET", "/recommendations?limit=10", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var result response.Response
	json.NewDecoder(resp.Body).Decode(&result)

	assert.True(t, result.Success)
	mockUC.AssertExpectations(t)
}
