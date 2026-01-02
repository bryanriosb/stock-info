package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/bryanriosb/stock-info/internal/stock/domain"
	"github.com/bryanriosb/stock-info/shared"
)

// SyncProgress represents the current sync progress
type SyncProgress struct {
	Current int    `json:"current"`
	Total   int    `json:"total"`
	Percent int    `json:"percent"`
	Status  string `json:"status"` // "fetching", "saving", "completed", "error"
	Message string `json:"message,omitempty"`
}

// ProgressCallback is called during sync to report progress
type ProgressCallback func(progress SyncProgress)

type StockAPIClient interface {
	FetchStocks(ctx context.Context, nextPage string) (*StockAPIResponse, error)
	FetchAllStocks(ctx context.Context) ([]*domain.Stock, error)
	FetchAllStocksWithProgress(ctx context.Context, onProgress ProgressCallback) ([]*domain.Stock, error)
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
	Time       string `json:"time"`
}

type stockAPIClient struct {
	baseURL    string
	token      string
	httpClient *http.Client
}

func NewStockAPIClient(cfg shared.StockAPIConfig) StockAPIClient {
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

func (c *stockAPIClient) FetchAllStocks(ctx context.Context) ([]*domain.Stock, error) {
	return c.FetchAllStocksWithProgress(ctx, nil)
}

func (c *stockAPIClient) FetchAllStocksWithProgress(ctx context.Context, onProgress ProgressCallback) ([]*domain.Stock, error) {
	var allStocks []*domain.Stock
	nextPage := ""
	pageCount := 0

	// Known total pages from API
	const totalPages = 2184

	for {
		pageCount++

		if onProgress != nil {
			// Calculate percent based on actual progress (reserve 5% for saving)
			percent := (pageCount * 95) / totalPages
			if percent > 95 {
				percent = 95
			}
			onProgress(SyncProgress{
				Current: pageCount,
				Total:   totalPages,
				Percent: percent,
				Status:  "fetching",
				Message: fmt.Sprintf("Fetching page %d of %d...", pageCount, totalPages),
			})
		}

		response, err := c.FetchStocks(ctx, nextPage)
		if err != nil {
			if onProgress != nil {
				onProgress(SyncProgress{
					Status:  "error",
					Message: err.Error(),
				})
			}
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

func itemToEntity(item StockItem) *domain.Stock {
	return &domain.Stock{
		Ticker:     item.Ticker,
		Company:    item.Company,
		Brokerage:  item.Brokerage,
		Action:     item.Action,
		RatingFrom: item.RatingFrom,
		RatingTo:   item.RatingTo,
		TargetFrom: parsePrice(item.TargetFrom),
		TargetTo:   parsePrice(item.TargetTo),
		Time:       parseTime(item.Time),
	}
}

func parseTime(timeStr string) time.Time {
	// Try common formats
	formats := []string{
		"2006-01-02T15:04:05Z",
		"2006-01-02 15:04:05",
		"2006-01-02",
		time.RFC3339,
	}
	for _, format := range formats {
		if t, err := time.Parse(format, timeStr); err == nil {
			return t
		}
	}
	return time.Time{}
}

func parsePrice(price string) float64 {
	cleaned := strings.ReplaceAll(price, "$", "")
	cleaned = strings.ReplaceAll(cleaned, ",", "")
	cleaned = strings.TrimSpace(cleaned)

	value, _ := strconv.ParseFloat(cleaned, 64)
	return value
}
