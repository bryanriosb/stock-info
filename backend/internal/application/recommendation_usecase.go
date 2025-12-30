package application

import (
	"context"
	"sort"
	"strings"

	"github.com/bryanriosb/stock-info/internal/domain/entity"
	"github.com/bryanriosb/stock-info/internal/domain/repository"
)

type StockRecommendation struct {
	Stock         *entity.Stock `json:"stock"`
	Score         float64       `json:"score"`
	Reason        string        `json:"reason"`
	PotentialGain float64       `json:"potential_gain_percent"`
}

type RecommendationUseCase interface {
	GetRecommendations(ctx context.Context, limit int) ([]*StockRecommendation, error)
}

type recommendationUseCase struct {
	repo repository.StockRepository
}

func NewRecommendationUseCase(repo repository.StockRepository) RecommendationUseCase {
	return &recommendationUseCase{repo: repo}
}

func (uc *recommendationUseCase) GetRecommendations(ctx context.Context, limit int) ([]*StockRecommendation, error) {
	if limit <= 0 || limit > 50 {
		limit = 10
	}

	stocks, _, err := uc.repo.FindAll(ctx, repository.QueryParams{
		Page:  1,
		Limit: 100,
	})
	if err != nil {
		return nil, err
	}

	var recommendations []*StockRecommendation
	for _, stock := range stocks {
		score, reason := calculateScore(stock)
		potentialGain := calculatePotentialGain(stock)

		recommendations = append(recommendations, &StockRecommendation{
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

func calculateScore(stock *entity.Stock) (float64, string) {
	score := 0.0
	reasons := []string{}

	// Rating improvement score (0.3 weight)
	ratingScore := getRatingScore(stock.RatingFrom, stock.RatingTo)
	score += ratingScore * 0.3
	if ratingScore > 0 {
		reasons = append(reasons, "Positive rating")
	}

	// Target price change score (0.4 weight)
	if stock.TargetFrom > 0 {
		targetChange := (stock.TargetTo - stock.TargetFrom) / stock.TargetFrom
		score += targetChange * 0.4
		if targetChange > 0 {
			reasons = append(reasons, "Target price increased")
		}
	}

	// Action type score (0.3 weight)
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
		"strong sell":     1,
		"sell":            2,
		"underperform":    3,
		"market perform":  4,
		"neutral":         5,
		"hold":            5,
		"buy":             7,
		"outperform":      8,
		"strong buy":      9,
		"speculative buy": 8,
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

func calculatePotentialGain(stock *entity.Stock) float64 {
	if stock.TargetFrom <= 0 {
		return 0
	}
	return ((stock.TargetTo - stock.TargetFrom) / stock.TargetFrom) * 100
}
