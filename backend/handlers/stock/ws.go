package stock

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

func WSHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to upgrade connection: ", err)
	}

	defer conn.Close()
	defer func() {
		delete(clientConns, conn)
		fmt.Println("Client disconnected")
	}()

	for {
		_, symbol, err := conn.ReadMessage()
		clientConns[conn] = string(symbol)
		fmt.Println("New client connected")

		if err != nil {
			fmt.Println("Error reading from the client: ", err)
			break
		}
	}
}
