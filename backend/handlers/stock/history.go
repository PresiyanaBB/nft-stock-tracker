package stock

import (
	"encoding/json"
	"github.com/PresiyanaBB/crypto-price-tracker/models/stock"
	"gorm.io/gorm"
	"net/http"
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
