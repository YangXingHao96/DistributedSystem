package handler

import (
	"fmt"
	"net"
)

func HandleUDPRequest() {
	// Listen for incoming packets on port 8080
	conn, err := net.ListenPacket("udp", ":8080")
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
