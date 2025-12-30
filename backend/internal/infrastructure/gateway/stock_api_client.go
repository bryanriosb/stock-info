package gateway

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/bryanriosb/stock-info/internal/domain/entity"
	"github.com/bryanriosb/stock-info/pkg/config"
)

type StockAPIClient interface {
	FetchStocks(ctx context.Context, nextPage string) (*StockAPIResponse, error)
	FetchAllStocks(ctx context.Context) ([]*entity.Stock, error)
}

type StockAPIResponse struct {
	Items    []StockItem `json:"items"`
	NextPage string      `json:"next_page"`
}

type StockItem struct {
	Ticker     string `json:"ticker"`
	Company    string `json:"company"`
	Brokerage  string `json:"brokerage"`
	Action     string `json:"action"`
	RatingFrom string `json:"rating_from"`
	RatingTo   string `json:"rating_to"`
	TargetFrom string `json:"target_from"`
	TargetTo   string `json:"target_to"`
}

type stockAPIClient struct {
	baseURL    string
	token      string
	httpClient *http.Client
}

func NewStockAPIClient(cfg config.StockAPIConfig) StockAPIClient {
	return &stockAPIClient{
		baseURL: cfg.URL,
		token:   cfg.Token,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *stockAPIClient) FetchStocks(ctx context.Context, nextPage string) (*StockAPIResponse, error) {
	url := c.baseURL
	if nextPage != "" {
		url = fmt.Sprintf("%s?next_page=%s", c.baseURL, nextPage)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var response StockAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &response, nil
}

func (c *stockAPIClient) FetchAllStocks(ctx context.Context) ([]*entity.Stock, error) {
	var allStocks []*entity.Stock
	nextPage := ""

	for {
		response, err := c.FetchStocks(ctx, nextPage)
		if err != nil {
			return nil, err
		}

		for _, item := range response.Items {
			stock := itemToEntity(item)
			allStocks = append(allStocks, stock)
		}

		if response.NextPage == "" {
			break
		}
		nextPage = response.NextPage
	}

	return allStocks, nil
}

func itemToEntity(item StockItem) *entity.Stock {
	return &entity.Stock{
		Ticker:     item.Ticker,
		Company:    item.Company,
		Brokerage:  item.Brokerage,
		Action:     item.Action,
		RatingFrom: item.RatingFrom,
		RatingTo:   item.RatingTo,
		TargetFrom: parsePrice(item.TargetFrom),
		TargetTo:   parsePrice(item.TargetTo),
	}
}

func parsePrice(price string) float64 {
	cleaned := strings.ReplaceAll(price, "$", "")
	cleaned = strings.ReplaceAll(cleaned, ",", "")
	cleaned = strings.TrimSpace(cleaned)

	value, _ := strconv.ParseFloat(cleaned, 64)
	return value
}
