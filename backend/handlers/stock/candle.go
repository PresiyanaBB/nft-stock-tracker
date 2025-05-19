package stock

import (
	"encoding/json"
	"github.com/PresiyanaBB/crypto-price-tracker/models/stock"
	"gorm.io/gorm"
	"net/http"
)

func CandlesHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	symbol := r.URL.Query().Get("symbol")

	var candles []stock.Candle
	db.Where("symbol = ?", symbol).Order("timestamp asc").Find(&candles)

	jsonCandles, _ := json.Marshal(candles)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonCandles)
}
