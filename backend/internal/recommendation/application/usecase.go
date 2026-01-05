package application

import (
	"context"
	"sort"
	"strings"

	"github.com/bryanriosb/stock-info/internal/recommendation/domain"
	stockDomain "github.com/bryanriosb/stock-info/internal/stock/domain"
)

type RecommendationUseCase interface {
	GetRecommendations(ctx context.Context, limit int) ([]*domain.StockRecommendation, error)
}

type recommendationUseCase struct {
	repo stockDomain.StockRepository
}

func NewRecommendationUseCase(repo stockDomain.StockRepository) RecommendationUseCase {
	return &recommendationUseCase{repo: repo}
}

func (uc *recommendationUseCase) GetRecommendations(ctx context.Context, limit int) ([]*domain.StockRecommendation, error) {
	if limit <= 0 || limit > 50 {
		limit = 10
	}

	stocks, _, err := uc.repo.FindAll(ctx, stockDomain.QueryParams{
		Page:  1,
		Limit: 100,
	})
	if err != nil {
		return nil, err
	}

	recommendations := make([]*domain.StockRecommendation, 0, len(stocks))
	for _, stock := range stocks {
		score, reason := calculateScore(stock)
		potentialGain := calculatePotentialGain(stock)

		recommendations = append(recommendations, &domain.StockRecommendation{
			Stock:         stock,
			Score:         score,
			Reason:        reason,
			PotentialGain: potentialGain,
		})
	}

	sort.Slice(recommendations, func(i, j int) bool {
		return recommendations[i].Score > recommendations[j].Score
	})

	if len(recommendations) > limit {
		recommendations = recommendations[:limit]
	}

	return recommendations, nil
}

func calculateScore(stock *stockDomain.Stock) (float64, string) {
	score := 0.0
	reasons := []string{}

	ratingScore := getRatingScore(stock.RatingFrom, stock.RatingTo)
	score += ratingScore * 0.3
	if ratingScore > 0 {
		reasons = append(reasons, "Positive rating")
	}

	if stock.TargetFrom > 0 {
		targetChange := (stock.TargetTo - stock.TargetFrom) / stock.TargetFrom
		score += targetChange * 0.4
		if targetChange > 0 {
			reasons = append(reasons, "Target price increased")
		}
	}

	actionScore := getActionScore(stock.Action)
	score += actionScore * 0.3
	if actionScore > 0 {
		reasons = append(reasons, "Positive action")
	}

	reason := "No strong signals"
	if len(reasons) > 0 {
		reason = strings.Join(reasons, ", ")
	}

	return score, reason
}

func getRatingScore(from, to string) float64 {
	ratings := map[string]int{
		"sell":                1,
		"negative":            2,
		"underperform":        3,
		"sector underperform": 3,
		"cautious":            4,
		"market perform":      5,
		"sector perform":      5,
		"neutral":             5,
		"hold":                5,
		"equal weight":        5,
		"in-line":             5,
		"buy":                 7,
		"positive":            7,
		"overweight":          7,
		"outperform":          8,
		"outperformer":        8,
		"market outperform":   8,
		"sector outperform":   8,
		"speculative buy":     8,
		"strong-buy":          9,
	}

	fromScore := ratings[strings.ToLower(from)]
	toScore := ratings[strings.ToLower(to)]

	if fromScore == 0 || toScore == 0 {
		return 0
	}

	diff := float64(toScore - fromScore)
	return diff / 8.0
}

func getActionScore(action string) float64 {
	action = strings.ToLower(action)

	if strings.Contains(action, "raised") || strings.Contains(action, "upgraded") {
		return 1.0
	}
	if strings.Contains(action, "maintained") || strings.Contains(action, "reiterated") {
		return 0.5
	}
	if strings.Contains(action, "lowered") || strings.Contains(action, "downgraded") {
		return -0.5
	}

	return 0
}

func calculatePotentialGain(stock *stockDomain.Stock) float64 {
	if stock.TargetFrom <= 0 {
		return 0
	}
	return ((stock.TargetTo - stock.TargetFrom) / stock.TargetFrom) * 100
}
