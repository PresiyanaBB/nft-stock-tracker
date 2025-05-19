package stock

import (
	"encoding/json"
	"fmt"
	"github.com/PresiyanaBB/crypto-price-tracker/models/stock"
	"github.com/gorilla/websocket"
	"time"
)

func BroadcastUpdates() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	var latestUpdate *stock.BroadcastMessage

	for {
		select {
		case update := <-broadcast:
			if update.UpdateType == stock.Closed {
				broadcastToClients(update)
			} else {
				latestUpdate = update
			}
		case <-ticker.C:
			if latestUpdate != nil {
				broadcastToClients(latestUpdate)
			}
			latestUpdate = nil
		}
	}
}

func broadcastToClients(update *stock.BroadcastMessage) {
	jsonUpdate, _ := json.Marshal(update)

	for clientConn, symbol := range clientConns {
		if update.Candle.Symbol == symbol {
			err := clientConn.WriteMessage(websocket.TextMessage, jsonUpdate)
			if err != nil {
				fmt.Println("Error sending message to client: ", err)
				clientConn.Close()
				delete(clientConns, clientConn)
			}
		}
	}
}
