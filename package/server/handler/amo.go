package handler

import (
	"database/sql"
	"fmt"
	"github.com/YangXingHao96/DistributedSystem/package/common"
	"github.com/YangXingHao96/DistributedSystem/package/common/constant"
	"github.com/YangXingHao96/DistributedSystem/package/server/service"
	"net"
	"time"
)

func HandleUDPRequestAtMostOnce(db *sql.DB) {
	// Listen for incoming packets on port 8080
	var msgIdSet = map[string][]byte{}
	udpServer, err := net.ListenPacket("udp", "localhost:2222")
	if err != nil {
		panic(err)
	}
	defer udpServer.Close()
	reservationMap := map[string]map[int]int{}
	addressToFlightMap := map[string]map[int]time.Time{}
	flightToAddressMap := map[int]map[string]time.Time{}
	stringAddressMap := map[string]net.Addr{}
	for {
		buf := make([]byte, 1024)
		n, addr, err := udpServer.ReadFrom(buf)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Received %d bytes from %s: %v\n", n, addr.String(), buf[:n])

		request := common.Deserialize(buf[:n])
		request[constant.Address] = addr.String()
		stringAddressMap[addr.String()] = addr
		resp, err := service.HandleDuplicateRequest(request, msgIdSet)
		if resp != nil {
			fmt.Println("handing repeated request")
			if _, err := udpServer.WriteTo(resp, addr); err != nil {
				fmt.Printf("An error has occured: %v\n", err)
			}
			continue
		}
		if err != nil {
			fmt.Printf("An error has occured: %v\n", err)
		}
		responses, err := service.HandleIncomingRequest(request, db, reservationMap, addressToFlightMap, flightToAddressMap)
		if err != nil {
			fmt.Printf("An error has occured: %v\n", err)
		} else {
			for key, value := range responses {
				if _, err := udpServer.WriteTo(value, stringAddressMap[key]); err != nil {
					fmt.Printf("An error has occured: %v\n", err)
				}
			}
		}
	}
}
