package handler

import (
	"strconv"

	"github.com/bryanriosb/stock-info/internal/application"
	"github.com/bryanriosb/stock-info/internal/domain/repository"
	"github.com/bryanriosb/stock-info/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type StockHandler struct {
	useCase application.StockUseCase
}

func NewStockHandler(useCase application.StockUseCase) *StockHandler {
	return &StockHandler{useCase: useCase}
}

func (h *StockHandler) GetStocks(c *fiber.Ctx) error {
	params := repository.QueryParams{
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

func (h *StockHandler) GetStockByID(c *fiber.Ctx) error {
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

func (h *StockHandler) GetStockByTicker(c *fiber.Ctx) error {
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

func (h *StockHandler) SyncStocks(c *fiber.Ctx) error {
	count, err := h.useCase.SyncStocks(c.Context())
	if err != nil {
		return response.InternalError(c, "Failed to sync stocks: "+err.Error())
	}

	return response.Success(c, fiber.Map{
		"message": "Stocks synced successfully",
		"count":   count,
	})
}
