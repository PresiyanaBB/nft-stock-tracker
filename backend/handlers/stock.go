package handlers

import (
	"context"
	"time"

	"github.com/PresiyanaBB/crypto-price-tracker/models/stock"
	"github.com/gofiber/fiber/v2"
)

type CandleHandler struct {
	repo stock.CandleRepository
}

func (h *CandleHandler) StockCandles(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()
	symbol := ctx.Params("symbol")
	candles, err := h.repo.StockCandles(c, symbol)

	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Stock candles retrieved successfully",
		"data":    candles,
	})
}

func (h *CandleHandler) StocksHistory(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()
	history, err := h.repo.StocksHistory(c)

	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Stocks history retrieved successfully",
		"data":    history,
	})
}

func NewCandleHandler(router fiber.Router, repo stock.CandleRepository) {
	handler := &CandleHandler{
		repo: repo,
	}

	router.Get("/stocks-history", handler.StocksHistory)
	router.Get("/stock-candles/:symbol", handler.StockCandles)
}
