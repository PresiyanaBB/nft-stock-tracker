package handlers

import (
	"context"
	"fmt"
	"time"

	hs "github.com/PresiyanaBB/nft-stock-tracker/handlers/stock"
	"github.com/PresiyanaBB/nft-stock-tracker/models/stock"
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

func WSHandler(ctx *fiber.Ctx) error {
	wsconn := hs.GetWSConn()

	defer func() {
		fmt.Println("Client disconnected:", wsconn.RemoteAddr())
		wsconn.Close()
	}()

	for {
		messageType, msg, err := wsconn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading WebSocket message:", err)
			break
		}

		fmt.Printf("Received message: %s", msg)

		err = wsconn.WriteMessage(messageType, msg)
		if err != nil {
			fmt.Println("Error writing WebSocket response:", err)
			break
		}
	}
	return nil
}

func NewCandleHandler(router fiber.Router, repo stock.CandleRepository) {
	handler := &CandleHandler{
		repo: repo,
	}

	router.Get("/stocks-history", handler.StocksHistory)
	router.Get("/stock-candles/:symbol", handler.StockCandles)
}
