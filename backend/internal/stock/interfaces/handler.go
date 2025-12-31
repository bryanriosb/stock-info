package interfaces

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
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
		Ticker:  c.Query("ticker"),
		Company: c.Query("company"),
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

func (h *Handler) GetStockByTicker(c *fiber.Ctx) error {
	ticker := c.Params("ticker")
	if ticker == "" {
		return response.BadRequest(c, "Ticker is required")
	}

	stocks, err := h.useCase.GetStockByTicker(c.Context(), ticker)
	if err != nil {
		return response.InternalError(c, "Failed to fetch stocks")
	}

	return response.Success(c, stocks)
}

// SyncStocksStream handles SSE streaming for stock sync with progress
func (h *Handler) SyncStocksStream(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Set("Connection", "keep-alive")
	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("X-Accel-Buffering", "no")
	// Don't set Transfer-Encoding manually - let Fiber handle it

	c.Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
		// Create a new context for the sync operation since the fiber context
		// becomes invalid inside the StreamWriter goroutine
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
		defer cancel()

		log.Println("Starting stock sync from external API...")

		// Send initial event
		fmt.Fprintf(w, "data: %s\n\n", mustJSON(infrastructure.SyncProgress{
			Percent: 0,
			Status:  "starting",
			Message: "Starting sync...",
		}))
		w.Flush()

		// Channel to receive progress updates
		progressChan := make(chan infrastructure.SyncProgress, 100)
		doneChan := make(chan error, 1)

		// Create progress callback that sends to channel
		onProgress := func(progress infrastructure.SyncProgress) {
			select {
			case progressChan <- progress:
			default:
				// Channel full, skip this update
			}
		}

		// Run sync in goroutine
		go func() {
			_, err := h.useCase.SyncStocksWithProgress(ctx, onProgress)
			doneChan <- err
		}()

		// Heartbeat ticker to keep connection alive
		heartbeat := time.NewTicker(10 * time.Second)
		defer heartbeat.Stop()

		var lastProgress infrastructure.SyncProgress

		for {
			select {
			case progress := <-progressChan:
				lastProgress = progress
				fmt.Fprintf(w, "data: %s\n\n", mustJSON(progress))
				w.Flush()

			case <-heartbeat.C:
				// Send heartbeat comment to keep connection alive
				fmt.Fprintf(w, ": heartbeat\n\n")
				w.Flush()

			case err := <-doneChan:
				if err != nil {
					fmt.Fprintf(w, "data: %s\n\n", mustJSON(infrastructure.SyncProgress{
						Status:  "error",
						Message: err.Error(),
					}))
				} else if lastProgress.Status != "completed" {
					// Ensure completed message is sent
					fmt.Fprintf(w, "data: %s\n\n", mustJSON(infrastructure.SyncProgress{
						Percent: 100,
						Status:  "completed",
						Message: "Sync completed successfully",
					}))
				}
				w.Flush()
				return

			case <-ctx.Done():
				fmt.Fprintf(w, "data: %s\n\n", mustJSON(infrastructure.SyncProgress{
					Status:  "error",
					Message: "Sync timed out",
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
