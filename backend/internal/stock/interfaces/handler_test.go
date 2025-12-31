package interfaces

import (
	"context"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/bryanriosb/stock-info/internal/stock/domain"
	"github.com/bryanriosb/stock-info/internal/stock/infrastructure"
	"github.com/bryanriosb/stock-info/shared/response"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock StockUseCase
type MockStockUseCase struct {
	mock.Mock
}

func (m *MockStockUseCase) SyncStocks(ctx context.Context) (int, error) {
	args := m.Called(ctx)
	return args.Int(0), args.Error(1)
}

func (m *MockStockUseCase) SyncStocksWithProgress(ctx context.Context, onProgress infrastructure.ProgressCallback) (int, error) {
	args := m.Called(ctx, onProgress)
	return args.Int(0), args.Error(1)
}

func (m *MockStockUseCase) GetStocks(ctx context.Context, params domain.QueryParams) ([]*domain.Stock, int64, error) {
	args := m.Called(ctx, params)
	if args.Get(0) == nil {
		return nil, args.Get(1).(int64), args.Error(2)
	}
	return args.Get(0).([]*domain.Stock), args.Get(1).(int64), args.Error(2)
}

func (m *MockStockUseCase) GetStockByTicker(ctx context.Context, ticker string) ([]*domain.Stock, error) {
	args := m.Called(ctx, ticker)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Stock), args.Error(1)
}

func (m *MockStockUseCase) GetStockByID(ctx context.Context, id int64) (*domain.Stock, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Stock), args.Error(1)
}

func setupTestApp(handler *Handler) *fiber.App {
	app := fiber.New()
	app.Get("/stocks", handler.GetStocks)
	app.Get("/stocks/:id", handler.GetStockByID)
	app.Get("/stocks/ticker/:ticker", handler.GetStockByTicker)
	// Note: SyncStocksStream is SSE and tested separately
	return app
}

func TestGetStocks_Success(t *testing.T) {
	mockUC := new(MockStockUseCase)
	handler := NewHandler(mockUC)
	app := setupTestApp(handler)

	stocks := []*domain.Stock{
		{ID: 1, Ticker: "AAPL", Company: "Apple Inc."},
		{ID: 2, Ticker: "GOOGL", Company: "Alphabet Inc."},
	}

	mockUC.On("GetStocks", mock.Anything, mock.AnythingOfType("domain.QueryParams")).
		Return(stocks, int64(2), nil)

	req := httptest.NewRequest("GET", "/stocks?page=1&limit=10", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var result response.Response
	json.NewDecoder(resp.Body).Decode(&result)

	assert.True(t, result.Success)
	assert.NotNil(t, result.Data)
	assert.NotNil(t, result.Meta)
	mockUC.AssertExpectations(t)
}

func TestGetStocks_Error(t *testing.T) {
	mockUC := new(MockStockUseCase)
	handler := NewHandler(mockUC)
	app := setupTestApp(handler)

	mockUC.On("GetStocks", mock.Anything, mock.AnythingOfType("domain.QueryParams")).
		Return(nil, int64(0), errors.New("database error"))

	req := httptest.NewRequest("GET", "/stocks", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

	var result response.Response
	json.NewDecoder(resp.Body).Decode(&result)

	assert.False(t, result.Success)
	assert.Contains(t, result.Error, "Failed to fetch stocks")
	mockUC.AssertExpectations(t)
}

func TestGetStockByID_Success(t *testing.T) {
	mockUC := new(MockStockUseCase)
	handler := NewHandler(mockUC)
	app := setupTestApp(handler)

	stock := &domain.Stock{ID: 1, Ticker: "AAPL", Company: "Apple Inc."}
	mockUC.On("GetStockByID", mock.Anything, int64(1)).Return(stock, nil)

	req := httptest.NewRequest("GET", "/stocks/1", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var result response.Response
	json.NewDecoder(resp.Body).Decode(&result)

	assert.True(t, result.Success)
	assert.NotNil(t, result.Data)
	mockUC.AssertExpectations(t)
}

func TestGetStockByID_InvalidID(t *testing.T) {
	mockUC := new(MockStockUseCase)
	handler := NewHandler(mockUC)
	app := setupTestApp(handler)

	req := httptest.NewRequest("GET", "/stocks/invalid", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

	var result response.Response
	json.NewDecoder(resp.Body).Decode(&result)

	assert.False(t, result.Success)
	assert.Contains(t, result.Error, "Invalid stock ID")
}

func TestGetStockByID_NotFound(t *testing.T) {
	mockUC := new(MockStockUseCase)
	handler := NewHandler(mockUC)
	app := setupTestApp(handler)

	mockUC.On("GetStockByID", mock.Anything, int64(999)).Return(nil, nil)

	req := httptest.NewRequest("GET", "/stocks/999", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)

	var result response.Response
	json.NewDecoder(resp.Body).Decode(&result)

	assert.False(t, result.Success)
	assert.Contains(t, result.Error, "Stock not found")
	mockUC.AssertExpectations(t)
}

func TestGetStockByTicker_Success(t *testing.T) {
	mockUC := new(MockStockUseCase)
	handler := NewHandler(mockUC)
	app := setupTestApp(handler)

	stocks := []*domain.Stock{
		{ID: 1, Ticker: "AAPL", Company: "Apple Inc."},
	}
	mockUC.On("GetStockByTicker", mock.Anything, "AAPL").Return(stocks, nil)

	req := httptest.NewRequest("GET", "/stocks/ticker/AAPL", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var result response.Response
	json.NewDecoder(resp.Body).Decode(&result)

	assert.True(t, result.Success)
	mockUC.AssertExpectations(t)
}

func TestGetStockByTicker_Error(t *testing.T) {
	mockUC := new(MockStockUseCase)
	handler := NewHandler(mockUC)
	app := setupTestApp(handler)

	mockUC.On("GetStockByTicker", mock.Anything, "AAPL").Return(nil, errors.New("database error"))

	req := httptest.NewRequest("GET", "/stocks/ticker/AAPL", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

	var result response.Response
	json.NewDecoder(resp.Body).Decode(&result)

	assert.False(t, result.Success)
	mockUC.AssertExpectations(t)
}

// Note: SyncStocksStream uses SSE (Server-Sent Events) which requires
// integration tests rather than unit tests. The streaming nature of SSE
// makes it difficult to test with httptest.
