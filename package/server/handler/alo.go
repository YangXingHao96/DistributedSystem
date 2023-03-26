package handler

import (
	"database/sql"
	"fmt"
	"net"
)

func HandleUDPRequestAtLeastOnce(db *sql.DB) {
	// Listen for incoming packets on port 8080
	conn, err := net.ListenPacket("udp", "localhost:2222")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Server listening on", conn.LocalAddr())

	// Loop to continuously receive incoming packets
	for {
		buffer := make([]byte, 1024)
		n, addr, err := conn.ReadFrom(buffer)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		fmt.Printf("Received %d bytes from %s: %s\n", n, addr.String(), string(buffer[:n]))
	}
}
