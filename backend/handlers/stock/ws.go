package stock

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
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
		fmt.Println("Client conn: ", clientConns[conn])
		fmt.Println("Symbol: ", string(symbol))
		if err != nil {
			fmt.Println("Error reading from the client: ", err)
			break
		}
	}
}
