package main

import (
	"DistributedSystem/package/server/handler"
	"DistributedSystem/package/server/service/database"
	"fmt"
	"net/http"
)

func main() {
	database.Init()
	http.HandleFunc("/", handler.HandleWebSocketAtMostOnce)
	fmt.Println("Starting WebSocket server on port 8080")
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Println("Failed to start WebSocket server:", err)
	}
}
