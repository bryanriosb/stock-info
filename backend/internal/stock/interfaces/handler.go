package interfaces

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/bryanriosb/stock-info/internal/stock/application"
	"github.com/bryanriosb/stock-info/internal/stock/domain"
	"github.com/bryanriosb/stock-info/internal/stock/infrastructure"
	"github.com/bryanriosb/stock-info/shared/response"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

type Handler struct {
	useCase application.StockUseCase
}

func NewHandler(useCase application.StockUseCase) *Handler {
	return &Handler{useCase: useCase}
}

func (h *Handler) GetStocks(c *fiber.Ctx) error {
	params := domain.QueryParams{
		Page:    c.QueryInt("page", 1),
		Limit:   c.QueryInt("limit", 20),
		SortBy:  c.Query("sort_by", "id"),
		SortDir: c.Query("sort_dir", "asc"),
	}

	// Only set filters if they have values
	if search := c.Query("search"); search != "" {
		params.Search = search
	}
	if ratingFrom := c.Query("rating_from"); ratingFrom != "" {
		params.RatingFrom = ratingFrom
	}
	if ratingTo := c.Query("rating_to"); ratingTo != "" {
		params.RatingTo = ratingTo
	}

	stocks, total, err := h.useCase.GetStocks(c.Context(), params)
	if err != nil {
		return response.InternalError(c, "Failed to fetch stocks")
	}

	totalPages := int(total) / params.Limit
	if int(total)%params.Limit > 0 {
		totalPages++
	}

	return response.SuccessWithMeta(c, stocks, &response.Meta{
		Page:       params.Page,
		Limit:      params.Limit,
		Total:      total,
		TotalPages: totalPages,
	})
}

func (h *Handler) GetStockByID(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return response.BadRequest(c, "Invalid stock ID")
	}

	stock, err := h.useCase.GetStockByID(c.Context(), id)
	if err != nil {
		return response.InternalError(c, "Failed to fetch stock")
	}

	if stock == nil {
		return response.NotFound(c, "Stock not found")
	}

	return response.Success(c, stock)
}

// SyncStocksStream handles SSE streaming for stock sync with progress
func (h *Handler) SyncStocksStream(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")
	c.Set("X-Accel-Buffering", "no")
	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("Access-Control-Allow-Headers", "Authorization, Content-Type")

	c.Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
		defer cancel()

		// Channel for progress updates
		progressCh := make(chan infrastructure.SyncProgress, 100)

		// Start sync in background
		go func() {
			defer close(progressCh)
			count, err := h.useCase.SyncStocksWithProgress(ctx, func(p infrastructure.SyncProgress) {
				select {
				case progressCh <- p:
				case <-ctx.Done():
				}
			})
			if err != nil {
				select {
				case progressCh <- infrastructure.SyncProgress{Status: "error", Message: err.Error()}:
				case <-ctx.Done():
				}
				return
			}
			select {
			case progressCh <- infrastructure.SyncProgress{
				Percent: 100,
				Status:  "completed",
				Message: fmt.Sprintf("Synced %d stocks", count),
			}:
			case <-ctx.Done():
			}
		}()

		// Send initial event
		fmt.Fprintf(w, "data: %s\n\n", mustJSON(infrastructure.SyncProgress{
			Status:  "starting",
			Message: "Starting sync...",
		}))
		w.Flush()

		heartbeat := time.NewTicker(2 * time.Second)
		defer heartbeat.Stop()

		for {
			select {
			case p, ok := <-progressCh:
				if !ok {
					return
				}
				fmt.Fprintf(w, "data: %s\n\n", mustJSON(p))
				w.Flush()

				if p.Status == "completed" || p.Status == "error" {
					return
				}

			case <-heartbeat.C:
				fmt.Fprint(w, ": ping\n\n")
				w.Flush()

			case <-ctx.Done():
				fmt.Fprintf(w, "data: %s\n\n", mustJSON(infrastructure.SyncProgress{
					Status:  "error",
					Message: "Request timeout",
				}))
				w.Flush()
				return
			}
		}
	}))

	return nil
}

func mustJSON(v interface{}) string {
	data, _ := json.Marshal(v)
	return string(data)
}
