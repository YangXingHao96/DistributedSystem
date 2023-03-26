package handler

import (
	"database/sql"
	"fmt"
	"github.com/YangXingHao96/DistributedSystem/package/common"
	"github.com/YangXingHao96/DistributedSystem/package/server/service"
	"net"
)

func HandleUDPRequestAtLeastOnce(db *sql.DB) {
	// Listen for incoming packets on port 8080
	udpServer, err := net.ListenPacket("udp", "localhost:2222")
	if err != nil {
		panic(err)
	}
	defer udpServer.Close()
	for {
		buf := make([]byte, 1024)
		n, addr, err := udpServer.ReadFrom(buf)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Received %d bytes from %s: %s\n", n, addr.String(), string(buf[:n]))

		request := common.Deserialize(buf)
		fmt.Println(request)
		resp, err := service.HandleIncomingRequest(request, db)
		if err != nil {
			fmt.Printf("An error has occured: %v\n", err)
		} else {
			if _, err := udpServer.WriteTo(resp, addr); err != nil {
				fmt.Printf("An error has occured: %v\n", err)
			}
		}
	}
}
