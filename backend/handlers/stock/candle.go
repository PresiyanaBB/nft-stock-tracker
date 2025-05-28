package stock

import (
	"encoding/json"
	"net/http"

	"github.com/PresiyanaBB/nft-stock-tracker/models/stock"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CandlesHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	symbol := r.URL.Query().Get("symbol")

	var candles []stock.Candle
	db.Where("symbol = ?", symbol).Order("timestamp asc").Find(&candles)

	jsonCandles, _ := json.Marshal(candles)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonCandles)
}

func CandlesFiberHandler(c *fiber.Ctx, db *gorm.DB) error {
	symbol := c.Query("symbol")

	var candles []stock.Candle
	db.Where("symbol = ?", symbol).Order("timestamp asc").Find(&candles)

	return c.JSON(candles)
}
