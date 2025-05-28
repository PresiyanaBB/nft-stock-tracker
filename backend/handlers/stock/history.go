package stock

import (
	"encoding/json"
	"net/http"

	"github.com/PresiyanaBB/nft-stock-tracker/models/stock"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func StocksHistoryHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var candles []stock.Candle
	db.Order("timestamp asc").Find(&candles)

	groupedData := make(map[string][]stock.Candle)

	for _, candle := range candles {
		symbol := candle.Symbol
		groupedData[symbol] = append(groupedData[symbol], candle)
	}

	jsonResponse, _ := json.Marshal(groupedData)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func StocksHistoryFiberHandler(c *fiber.Ctx, db *gorm.DB) error {
	var candles []stock.Candle
	db.Order("timestamp asc").Find(&candles)

	groupedData := make(map[string][]stock.Candle)
	for _, candle := range candles {
		groupedData[candle.Symbol] = append(groupedData[candle.Symbol], candle)
	}

	return c.JSON(groupedData)
}
