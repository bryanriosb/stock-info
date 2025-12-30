package domain

import stockDomain "github.com/bryanriosb/stock-info/internal/stock/domain"

type StockRecommendation struct {
	Stock         *stockDomain.Stock `json:"stock"`
	Score         float64            `json:"score"`
	Reason        string             `json:"reason"`
	PotentialGain float64            `json:"potential_gain_percent"`
}
