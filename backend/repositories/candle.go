package repositories

import (
	"context"
	"fmt"

	hs "github.com/PresiyanaBB/nft-stock-tracker/handlers/stock"
	ms "github.com/PresiyanaBB/nft-stock-tracker/models/stock"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
)

type CandleRepository struct {
	db *gorm.DB
}

func (r *CandleRepository) StockCandles(ctx context.Context, symbol string) ([]*ms.Candle, error) {
	candles := []*ms.Candle{}

	res := r.db.Model(&ms.Candle{}).Where("symbol = ?", symbol).Order("timestamp asc").Find(&candles)

	if res.Error != nil {
		return nil, res.Error
	}

	return candles, nil
}

func (r *CandleRepository) StocksHistory(ctx context.Context) (map[string][]*ms.Candle, error) {
	candles := []*ms.Candle{}

	res := r.db.Model(&ms.Candle{}).Order("timestamp asc").Find(&candles)

	if res.Error != nil {
		return nil, res.Error
	}

	groupedData := make(map[string][]*ms.Candle)

	for _, candle := range candles {
		groupedData[candle.Symbol] = append(groupedData[candle.Symbol], candle)
	}

	return groupedData, nil
}

func (r *CandleRepository) WSHandler(c *websocket.Conn) {
	// Close WebSocket connection when the client disconnects
	defer func() {
		hs.DeleteClientConn(c) // Handle client removal
		fmt.Println("Client disconnected!")
		c.Close()
	}()

	// Read messages from the WebSocket
	for {
		messageType, message, err := c.ReadMessage()
		if err != nil {
			fmt.Println("Error reading from the client:", err)
			break
		}

		// Store client subscription data
		hs.SetClientConns(c, string(message))
		fmt.Println("New Client Connected, subscribed to:", string(message))

		// Echo back the message or process further if needed
		if err := c.WriteMessage(messageType, []byte("Subscription confirmed: "+string(message))); err != nil {
			fmt.Println("Error writing WebSocket response:", err)
			break
		}
	}
}

func NewCandleRepository(db *gorm.DB) ms.CandleRepository {
	return &CandleRepository{
		db: db,
	}
}
