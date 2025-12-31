package interfaces

import (
	"bufio"
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
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")
	c.Set("Transfer-Encoding", "chunked")
	c.Set("Access-Control-Allow-Origin", "*")

	c.Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
		// Send initial event
		fmt.Fprintf(w, "data: %s\n\n", mustJSON(infrastructure.SyncProgress{
			Percent: 0,
			Status:  "starting",
			Message: "Starting sync...",
		}))
		w.Flush()

		// Create progress callback
		onProgress := func(progress infrastructure.SyncProgress) {
			data := mustJSON(progress)
			fmt.Fprintf(w, "data: %s\n\n", data)
			w.Flush()
		}

		// Run sync with progress
		_, err := h.useCase.SyncStocksWithProgress(c.Context(), onProgress)
		if err != nil {
			fmt.Fprintf(w, "data: %s\n\n", mustJSON(infrastructure.SyncProgress{
				Status:  "error",
				Message: err.Error(),
			}))
			w.Flush()
			return
		}

		// Small delay to ensure client receives final message
		time.Sleep(100 * time.Millisecond)
	}))

	return nil
}

func mustJSON(v interface{}) string {
	data, _ := json.Marshal(v)
	return string(data)
}
