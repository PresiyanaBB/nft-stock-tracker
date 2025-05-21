package stock

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/PresiyanaBB/crypto-price-tracker/config"
	"github.com/PresiyanaBB/crypto-price-tracker/models/stock"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
)

var (
	symbols = []string{"AAPL", "AMZN", "GOOGL", "MSFT", "TSLA", "NFLX", "FB", "NVDA", "AMD", "INTC"}

	broadcast = make(chan *stock.BroadcastMessage)

	// clients connections
	clientConns = make(map[*websocket.Conn]string)

	tempCandles = make(map[string]*stock.TempCandle)

	mutex sync.Mutex
)

func ConnectToFinnhub(env *config.EnvConfig) *websocket.Conn {
	ws, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("wss://ws.finnhub.io?token=%s", env.APIKey), nil)
	if err != nil {
		panic(err)
	}

	for _, s := range symbols {
		msg, _ := json.Marshal(map[string]interface{}{"type": "subscribe", "symbol": s})
		ws.WriteMessage(websocket.TextMessage, msg)
	}

	return ws
}

func HandleFinnhubMessages(ws *websocket.Conn, db *gorm.DB) {
	finnhubMessage := &stock.FinnhubMessage{}

	for {
		if err := ws.ReadJSON(finnhubMessage); err != nil {
			fmt.Println("Error reading the message: ", err)
			continue
		}

		// As the type can be either ping or trade, we are processing only trade operations
		if finnhubMessage.Type == "trade" {
			for _, trade := range finnhubMessage.Data {
				processTradeData(&trade, db)
			}
		}
	}
}

// Update ot Create temporary Candles
func processTradeData(trade *stock.TradeData, db *gorm.DB) {
	// protect the goroutine from data racing via mutex
	mutex.Lock()
	defer mutex.Unlock()

	// extract data
	price := trade.Price
	symbol := trade.Symbol
	timestamp := time.UnixMilli(trade.Timestamp)
	volume := float64(trade.Volume)

	tempCandle, exists := tempCandles[symbol]

	if !exists || timestamp.After(tempCandle.CloseTime) {
		if exists {
			// convert the temporary candle to candle
			candle := tempCandle.ToCandle()

			// save to db
			if err := db.Create(candle).Error; err != nil {
				fmt.Println("Error saving the candle to db: ", err)
			}

			broadcast <- &stock.BroadcastMessage{
				UpdateType: stock.Closed,
				Candle:     candle,
			}
		}

		tempCandle = &stock.TempCandle{
			Symbol:     symbol,
			OpenTime:   timestamp,
			CloseTime:  timestamp.Add(time.Minute),
			OpenPrice:  price,
			ClosePrice: price,
			HighPrice:  price,
			LowPrice:   price,
			Volume:     volume,
		}
	}

	tempCandle.ClosePrice = price
	tempCandle.Volume += volume
	if price > tempCandle.HighPrice {
		tempCandle.HighPrice = price
	}
	if price < tempCandle.LowPrice {
		tempCandle.LowPrice = price
	}

	tempCandles[symbol] = tempCandle

	broadcast <- &stock.BroadcastMessage{
		UpdateType: stock.Live,
		Candle:     tempCandle.ToCandle(),
	}
}
