package stock

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func WSHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade incoming GET request into a Websocket connection
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrate connection:", err)
		http.Error(w, "WebSocket upgrade failed", http.StatusInternalServerError)
		return
	}
	log.Println("WebSocket connection established with client:", r.RemoteAddr)

	// Close ws connection & unregister the client when they disconnect
	defer conn.Close()
	defer func() {
		delete(clientConns, conn)
		log.Println("Client disconnected!")
	}()

	// Register the new client to the symbol they're subscribing to
	for {
		_, symbol, err := conn.ReadMessage()
		clientConns[conn] = string(symbol)
		log.Println("New Client Connected!")

		if err != nil {
			log.Println("Error reading from the client:", err)
			break
		}
	}
}
