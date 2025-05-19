package stock

import "time"

// Candle Represents single OHLC(Open, High, Low, Close)
type Candle struct {
	Symbol    string    `json:"symbol"`
	Open      float64   `json:"open"`
	High      float64   `json:"high"`
	Low       float64   `json:"low"`
	Close     float64   `json:"close"`
	Timestamp time.Time `json:"timestamp"`
}

// TempCandle Temporary candle is an item from the temporary candle slice building the candles
type TempCandle struct {
	Symbol     string
	OpenTime   time.Time
	CloseTime  time.Time
	OpenPrice  float64
	HighPrice  float64
	LowPrice   float64
	ClosePrice float64
	Volume     float64
}

func (tc *TempCandle) ToCandle() *Candle {
	return &Candle{
		Symbol:    tc.Symbol,
		Open:      tc.OpenPrice,
		High:      tc.HighPrice,
		Low:       tc.LowPrice,
		Close:     tc.ClosePrice,
		Timestamp: tc.CloseTime,
	}
}
