package stock

type FinnhubMessage struct {
	Data []TradeData `json:"data"`
	Type string      `json:"type"` // ping | trade
}

type TradeData struct {
	Close     []string `json:"c"`
	Price     float64  `json:"p"`
	Symbol    string   `json:"s"`
	Timestamp int64    `json:"t"`
	Volume    float64  `json:"v"`
}

type BroadcastMessage struct {
	UpdateType UpdateType `json:"update_type"` // live | closed
	Candle     *Candle    `json:"candle"`
}

type UpdateType string

const (
	Live   UpdateType = "live"
	Closed UpdateType = "closed"
)
