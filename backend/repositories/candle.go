package repositories

import (
	"context"

	"github.com/PresiyanaBB/crypto-price-tracker/models/stock"
	"gorm.io/gorm"
)

type CandleRepository struct {
	db *gorm.DB
}

func (r *CandleRepository) StockCandles(ctx context.Context, symbol string) ([]*stock.Candle, error) {
	candles := []*stock.Candle{}

	res := r.db.Model(&stock.Candle{}).Where("symbol = ?", symbol).Order("timestamp asc").Find(&candles)

	if res.Error != nil {
		return nil, res.Error
	}

	return candles, nil
}

func (r *CandleRepository) StocksHistory(ctx context.Context) (map[string][]*stock.Candle, error) {
	candles := []*stock.Candle{}

	res := r.db.Model(&stock.Candle{}).Order("timestamp asc").Find(&candles)

	if res.Error != nil {
		return nil, res.Error
	}

	groupedData := make(map[string][]*stock.Candle)

	for _, candle := range candles {
		groupedData[candle.Symbol] = append(groupedData[candle.Symbol], candle)
	}

	return groupedData, nil
}

func NewCanleRepository(db *gorm.DB) stock.CandleRepository {
	return &CandleRepository{
		db: db,
	}
}
