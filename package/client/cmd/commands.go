package cmd

import (
	"fmt"
	"github.com/YangXingHao96/DistributedSystem/package/client"
)

// TODO
func getFlightIdBySourceDest(conn client.UdpConn) error {
	fmt.Println("called getFlightIdBySourceDest")
	// send some data
	if _, err := conn.Write([]byte("Hello Server! Greetings.")); err != nil {
		panic(err)
	}
	buffer := make([]byte, 1024)
	mLen, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println("Received: ", string(buffer[:mLen]))
	defer conn.Close()

	return nil
}
