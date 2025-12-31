package application

import (
	"context"
	"errors"
	"testing"

	"github.com/bryanriosb/stock-info/internal/stock/domain"
	"github.com/bryanriosb/stock-info/internal/stock/infrastructure"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock StockRepository
type MockStockRepository struct {
	mock.Mock
}

func (m *MockStockRepository) Create(ctx context.Context, stock *domain.Stock) error {
	args := m.Called(ctx, stock)
	return args.Error(0)
}

func (m *MockStockRepository) CreateBatch(ctx context.Context, stocks []*domain.Stock) error {
	args := m.Called(ctx, stocks)
	return args.Error(0)
}

func (m *MockStockRepository) FindAll(ctx context.Context, params domain.QueryParams) ([]*domain.Stock, int64, error) {
	args := m.Called(ctx, params)
	if args.Get(0) == nil {
		return nil, args.Get(1).(int64), args.Error(2)
	}
	return args.Get(0).([]*domain.Stock), args.Get(1).(int64), args.Error(2)
}

func (m *MockStockRepository) FindByTicker(ctx context.Context, ticker string) ([]*domain.Stock, error) {
	args := m.Called(ctx, ticker)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Stock), args.Error(1)
}

func (m *MockStockRepository) FindByID(ctx context.Context, id int64) (*domain.Stock, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Stock), args.Error(1)
}

// Mock StockAPIClient
type MockStockAPIClient struct {
	mock.Mock
}

func (m *MockStockAPIClient) FetchStocks(ctx context.Context, nextPage string) (*infrastructure.StockAPIResponse, error) {
	args := m.Called(ctx, nextPage)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*infrastructure.StockAPIResponse), args.Error(1)
}

func (m *MockStockAPIClient) FetchAllStocks(ctx context.Context) ([]*domain.Stock, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Stock), args.Error(1)
}

func (m *MockStockAPIClient) FetchAllStocksWithProgress(ctx context.Context, onProgress infrastructure.ProgressCallback) ([]*domain.Stock, error) {
	args := m.Called(ctx, onProgress)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Stock), args.Error(1)
}

func TestSyncStocks_Success(t *testing.T) {
	mockRepo := new(MockStockRepository)
	mockAPI := new(MockStockAPIClient)

	stocks := []*domain.Stock{
		{Ticker: "AAPL", Company: "Apple Inc."},
		{Ticker: "GOOGL", Company: "Alphabet Inc."},
	}

	mockAPI.On("FetchAllStocksWithProgress", mock.Anything, mock.Anything).Return(stocks, nil)
	mockRepo.On("CreateBatch", mock.Anything, stocks).Return(nil)

	uc := NewStockUseCase(mockRepo, mockAPI)
	count, err := uc.SyncStocks(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, 2, count)
	mockAPI.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestSyncStocks_APIError(t *testing.T) {
	mockRepo := new(MockStockRepository)
	mockAPI := new(MockStockAPIClient)

	mockAPI.On("FetchAllStocksWithProgress", mock.Anything, mock.Anything).Return(nil, errors.New("API error"))

	uc := NewStockUseCase(mockRepo, mockAPI)
	count, err := uc.SyncStocks(context.Background())

	assert.Error(t, err)
	assert.Equal(t, 0, count)
	mockAPI.AssertExpectations(t)
}

func TestSyncStocks_RepoError(t *testing.T) {
	mockRepo := new(MockStockRepository)
	mockAPI := new(MockStockAPIClient)

	stocks := []*domain.Stock{
		{Ticker: "AAPL", Company: "Apple Inc."},
	}

	mockAPI.On("FetchAllStocksWithProgress", mock.Anything, mock.Anything).Return(stocks, nil)
	mockRepo.On("CreateBatch", mock.Anything, stocks).Return(errors.New("DB error"))

	uc := NewStockUseCase(mockRepo, mockAPI)
	count, err := uc.SyncStocks(context.Background())

	assert.Error(t, err)
	assert.Equal(t, 0, count)
	mockAPI.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestGetStocks_Success(t *testing.T) {
	mockRepo := new(MockStockRepository)
	mockAPI := new(MockStockAPIClient)

	stocks := []*domain.Stock{
		{ID: 1, Ticker: "AAPL", Company: "Apple Inc."},
		{ID: 2, Ticker: "GOOGL", Company: "Alphabet Inc."},
	}
	params := domain.QueryParams{Page: 1, Limit: 10}

	mockRepo.On("FindAll", mock.Anything, params).Return(stocks, int64(2), nil)

	uc := NewStockUseCase(mockRepo, mockAPI)
	result, total, err := uc.GetStocks(context.Background(), params)

	assert.NoError(t, err)
	assert.Equal(t, int64(2), total)
	assert.Len(t, result, 2)
	mockRepo.AssertExpectations(t)
}

func TestGetStocks_Empty(t *testing.T) {
	mockRepo := new(MockStockRepository)
	mockAPI := new(MockStockAPIClient)

	params := domain.QueryParams{Page: 1, Limit: 10}
	mockRepo.On("FindAll", mock.Anything, params).Return([]*domain.Stock{}, int64(0), nil)

	uc := NewStockUseCase(mockRepo, mockAPI)
	result, total, err := uc.GetStocks(context.Background(), params)

	assert.NoError(t, err)
	assert.Equal(t, int64(0), total)
	assert.Empty(t, result)
	mockRepo.AssertExpectations(t)
}

func TestGetStockByTicker_Success(t *testing.T) {
	mockRepo := new(MockStockRepository)
	mockAPI := new(MockStockAPIClient)

	stocks := []*domain.Stock{
		{ID: 1, Ticker: "AAPL", Company: "Apple Inc."},
	}

	mockRepo.On("FindByTicker", mock.Anything, "AAPL").Return(stocks, nil)

	uc := NewStockUseCase(mockRepo, mockAPI)
	result, err := uc.GetStockByTicker(context.Background(), "AAPL")

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "AAPL", result[0].Ticker)
	mockRepo.AssertExpectations(t)
}

func TestGetStockByTicker_NotFound(t *testing.T) {
	mockRepo := new(MockStockRepository)
	mockAPI := new(MockStockAPIClient)

	mockRepo.On("FindByTicker", mock.Anything, "INVALID").Return([]*domain.Stock{}, nil)

	uc := NewStockUseCase(mockRepo, mockAPI)
	result, err := uc.GetStockByTicker(context.Background(), "INVALID")

	assert.NoError(t, err)
	assert.Empty(t, result)
	mockRepo.AssertExpectations(t)
}

func TestGetStockByID_Success(t *testing.T) {
	mockRepo := new(MockStockRepository)
	mockAPI := new(MockStockAPIClient)

	stock := &domain.Stock{ID: 1, Ticker: "AAPL", Company: "Apple Inc."}

	mockRepo.On("FindByID", mock.Anything, int64(1)).Return(stock, nil)

	uc := NewStockUseCase(mockRepo, mockAPI)
	result, err := uc.GetStockByID(context.Background(), 1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, int64(1), result.ID)
	mockRepo.AssertExpectations(t)
}

func TestGetStockByID_NotFound(t *testing.T) {
	mockRepo := new(MockStockRepository)
	mockAPI := new(MockStockAPIClient)

	mockRepo.On("FindByID", mock.Anything, int64(999)).Return(nil, errors.New("not found"))

	uc := NewStockUseCase(mockRepo, mockAPI)
	result, err := uc.GetStockByID(context.Background(), 999)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}
