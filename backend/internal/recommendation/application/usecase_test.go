package application

import (
	"context"
	"errors"
	"testing"

	stockDomain "github.com/bryanriosb/stock-info/internal/stock/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock StockRepository
type MockStockRepository struct {
	mock.Mock
}

func (m *MockStockRepository) Create(ctx context.Context, stock *stockDomain.Stock) error {
	args := m.Called(ctx, stock)
	return args.Error(0)
}

func (m *MockStockRepository) CreateBatch(ctx context.Context, stocks []*stockDomain.Stock) error {
	args := m.Called(ctx, stocks)
	return args.Error(0)
}

func (m *MockStockRepository) FindAll(ctx context.Context, params stockDomain.QueryParams) ([]*stockDomain.Stock, int64, error) {
	args := m.Called(ctx, params)
	if args.Get(0) == nil {
		return nil, args.Get(1).(int64), args.Error(2)
	}
	return args.Get(0).([]*stockDomain.Stock), args.Get(1).(int64), args.Error(2)
}

func (m *MockStockRepository) FindByTicker(ctx context.Context, ticker string) ([]*stockDomain.Stock, error) {
	args := m.Called(ctx, ticker)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*stockDomain.Stock), args.Error(1)
}

func (m *MockStockRepository) FindByID(ctx context.Context, id int64) (*stockDomain.Stock, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*stockDomain.Stock), args.Error(1)
}

func TestGetRecommendations_Success(t *testing.T) {
	mockRepo := new(MockStockRepository)

	stocks := []*stockDomain.Stock{
		{
			ID:         1,
			Ticker:     "AAPL",
			Company:    "Apple Inc.",
			RatingFrom: "Hold",
			RatingTo:   "Buy",
			TargetFrom: 150.0,
			TargetTo:   180.0,
			Action:     "target raised by",
		},
		{
			ID:         2,
			Ticker:     "GOOGL",
			Company:    "Alphabet Inc.",
			RatingFrom: "Neutral",
			RatingTo:   "Neutral",
			TargetFrom: 100.0,
			TargetTo:   100.0,
			Action:     "maintained",
		},
	}

	mockRepo.On("FindAll", mock.Anything, mock.Anything).Return(stocks, int64(2), nil)

	uc := NewRecommendationUseCase(mockRepo)
	recommendations, err := uc.GetRecommendations(context.Background(), 10)

	assert.NoError(t, err)
	assert.Len(t, recommendations, 2)
	// First recommendation should be AAPL (higher score due to rating upgrade and target raise)
	assert.Equal(t, "AAPL", recommendations[0].Stock.Ticker)
	assert.Greater(t, recommendations[0].Score, recommendations[1].Score)
	mockRepo.AssertExpectations(t)
}

func TestGetRecommendations_WithLimit(t *testing.T) {
	mockRepo := new(MockStockRepository)

	stocks := []*stockDomain.Stock{
		{ID: 1, Ticker: "AAPL", RatingFrom: "Hold", RatingTo: "Buy", TargetFrom: 100, TargetTo: 150, Action: "upgraded"},
		{ID: 2, Ticker: "GOOGL", RatingFrom: "Neutral", RatingTo: "Neutral", TargetFrom: 100, TargetTo: 100},
		{ID: 3, Ticker: "MSFT", RatingFrom: "Sell", RatingTo: "Hold", TargetFrom: 80, TargetTo: 100, Action: "raised"},
	}

	mockRepo.On("FindAll", mock.Anything, mock.Anything).Return(stocks, int64(3), nil)

	uc := NewRecommendationUseCase(mockRepo)
	recommendations, err := uc.GetRecommendations(context.Background(), 2)

	assert.NoError(t, err)
	assert.Len(t, recommendations, 2)
	mockRepo.AssertExpectations(t)
}

func TestGetRecommendations_DefaultLimit(t *testing.T) {
	mockRepo := new(MockStockRepository)

	mockRepo.On("FindAll", mock.Anything, mock.Anything).Return([]*stockDomain.Stock{}, int64(0), nil)

	uc := NewRecommendationUseCase(mockRepo)

	// Test with invalid limit (0)
	_, err := uc.GetRecommendations(context.Background(), 0)
	assert.NoError(t, err)

	// Test with negative limit
	_, err = uc.GetRecommendations(context.Background(), -5)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestGetRecommendations_RepoError(t *testing.T) {
	mockRepo := new(MockStockRepository)

	mockRepo.On("FindAll", mock.Anything, mock.Anything).Return(nil, int64(0), errors.New("database error"))

	uc := NewRecommendationUseCase(mockRepo)
	recommendations, err := uc.GetRecommendations(context.Background(), 10)

	assert.Error(t, err)
	assert.Nil(t, recommendations)
	mockRepo.AssertExpectations(t)
}

func TestGetRecommendations_Empty(t *testing.T) {
	mockRepo := new(MockStockRepository)

	mockRepo.On("FindAll", mock.Anything, mock.Anything).Return([]*stockDomain.Stock{}, int64(0), nil)

	uc := NewRecommendationUseCase(mockRepo)
	recommendations, err := uc.GetRecommendations(context.Background(), 10)

	assert.NoError(t, err)
	assert.Empty(t, recommendations)
	mockRepo.AssertExpectations(t)
}

func TestCalculateScore_PositiveRating(t *testing.T) {
	stock := &stockDomain.Stock{
		RatingFrom: "Hold",
		RatingTo:   "Buy",
		TargetFrom: 100.0,
		TargetTo:   100.0,
		Action:     "",
	}

	score, reason := calculateScore(stock)

	assert.Greater(t, score, 0.0)
	assert.Contains(t, reason, "Positive rating")
}

func TestCalculateScore_TargetIncrease(t *testing.T) {
	stock := &stockDomain.Stock{
		RatingFrom: "Hold",
		RatingTo:   "Hold",
		TargetFrom: 100.0,
		TargetTo:   150.0,
		Action:     "",
	}

	score, reason := calculateScore(stock)

	assert.Greater(t, score, 0.0)
	assert.Contains(t, reason, "Target price increased")
}

func TestCalculateScore_PositiveAction(t *testing.T) {
	stock := &stockDomain.Stock{
		RatingFrom: "Hold",
		RatingTo:   "Hold",
		TargetFrom: 100.0,
		TargetTo:   100.0,
		Action:     "target raised by analyst",
	}

	score, reason := calculateScore(stock)

	assert.Greater(t, score, 0.0)
	assert.Contains(t, reason, "Positive action")
}

func TestCalculateScore_NoSignals(t *testing.T) {
	stock := &stockDomain.Stock{
		RatingFrom: "",
		RatingTo:   "",
		TargetFrom: 0,
		TargetTo:   0,
		Action:     "",
	}

	score, reason := calculateScore(stock)

	assert.Equal(t, 0.0, score)
	assert.Equal(t, "No strong signals", reason)
}

func TestGetRatingScore(t *testing.T) {
	tests := []struct {
		from     string
		to       string
		positive bool
	}{
		{"sell", "buy", true},
		{"buy", "sell", false},
		{"hold", "strong buy", true},
		{"neutral", "neutral", false},
		{"", "buy", false},
	}

	for _, tt := range tests {
		score := getRatingScore(tt.from, tt.to)
		if tt.positive {
			assert.Greater(t, score, 0.0, "expected positive score for %s -> %s", tt.from, tt.to)
		} else {
			assert.LessOrEqual(t, score, 0.0, "expected non-positive score for %s -> %s", tt.from, tt.to)
		}
	}
}

func TestGetActionScore(t *testing.T) {
	tests := []struct {
		action   string
		expected float64
	}{
		{"target raised by analyst", 1.0},
		{"upgraded to buy", 1.0},
		{"maintained at hold", 0.5},
		{"reiterated buy", 0.5},
		{"target lowered", -0.5},
		{"downgraded to sell", -0.5},
		{"unknown action", 0.0},
	}

	for _, tt := range tests {
		score := getActionScore(tt.action)
		assert.Equal(t, tt.expected, score, "unexpected score for action: %s", tt.action)
	}
}

func TestCalculatePotentialGain(t *testing.T) {
	tests := []struct {
		from     float64
		to       float64
		expected float64
	}{
		{100.0, 150.0, 50.0},
		{100.0, 100.0, 0.0},
		{100.0, 80.0, -20.0},
		{0.0, 100.0, 0.0}, // Division by zero protection
	}

	for _, tt := range tests {
		stock := &stockDomain.Stock{TargetFrom: tt.from, TargetTo: tt.to}
		gain := calculatePotentialGain(stock)
		assert.Equal(t, tt.expected, gain, "unexpected gain for %v -> %v", tt.from, tt.to)
	}
}
