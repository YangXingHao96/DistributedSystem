package handler

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWebSocketAtLeastOnce(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to upgrade connection:", err)
		return
	}
	defer conn.Close()

	for {
		// Read message from client
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Failed to read message:", err)
			break
		}

		// Print message received
		fmt.Println("Received message:", string(message))

		// Send response back to client
		err = conn.WriteMessage(websocket.TextMessage, []byte("Hello, client!"))
		if err != nil {
			fmt.Println("Failed to write message:", err)
			break
		}
	}
}

func HandleWebSocketAtMostOnce(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to upgrade connection:", err)
		return
	}
	defer conn.Close()

	for {
		// Read message from client
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Failed to read message:", err)
			break
		}

		// Print message received
		fmt.Println("Received message:", string(message))

		// Send response back to client
		err = conn.WriteMessage(websocket.TextMessage, []byte("Hello, client!"))
		if err != nil {
			fmt.Println("Failed to write message:", err)
			break
		}
	}
}
