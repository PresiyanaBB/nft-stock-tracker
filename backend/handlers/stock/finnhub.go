package stock

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/PresiyanaBB/nft-stock-tracker/config"
	"github.com/PresiyanaBB/nft-stock-tracker/models/stock"
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

	wsConn *websocket.Conn
)

func GetClientConns() map[*websocket.Conn]string {
	return clientConns
}

func SetClientConns(conn *websocket.Conn, symbol string) {
	mutex.Lock()
	defer mutex.Unlock()

	// Register the new client to the symbol they're subscribing to
	clientConns[conn] = symbol

	// Send a message to the client confirming the subscription
	msg := fmt.Sprintf("Subscribed to %s", symbol)
	if err := conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
		fmt.Println("Error sending subscription confirmation:", err)
	}
}

func DeleteClientConn(conn *websocket.Conn) {
	mutex.Lock()
	defer mutex.Unlock()

	// Unregister the client connection
	if _, exists := clientConns[conn]; exists {
		delete(clientConns, conn)
		fmt.Println("Client disconnected and unregistered.")
	} else {
		fmt.Println("Client connection not found.")
	}
}

func ConnectToFinnhub(env *config.EnvConfig) *websocket.Conn {
	// Use Gorilla WebSocket for client connection (Fiber does not support outer WebSocket clients)
	ws, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("wss://ws.finnhub.io?token=%s", env.APIKey), nil)
	if err != nil {
		fmt.Println("WebSocket connection failed: ", err)
	}

	fmt.Println("Connected to Finnhub WebSocket")

	// Subscribe to symbols
	for _, s := range symbols {
		msg, _ := json.Marshal(map[string]interface{}{"type": "subscribe", "symbol": s})
		if err := ws.WriteMessage(websocket.TextMessage, msg); err != nil {
			fmt.Println("Error subscribing:", err)
		}
	}

	wsConn = ws
	return ws
}

func GetWSConn() *websocket.Conn {
	if wsConn == nil {
		fmt.Println("WebSocket connection is not established.")
		return nil
	}
	return wsConn
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

		// Clean up old trades older than 3 hours
		cutoffTime := time.Now().Add(-3 * time.Hour)
		db.Where("timestamp < ?", cutoffTime).Delete(&stock.Candle{})
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
