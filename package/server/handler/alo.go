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

func HandleUDPRequestAtLeastOnce(db *sql.DB) {
	// Listen for incoming packets on port 8080

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
		fmt.Println(addressToFlightMap)
		fmt.Println(flightToAddressMap)
		responses := service.HandleMonitorBackoff(addressToFlightMap, flightToAddressMap)
		for key, value := range responses {
			if _, err := udpServer.WriteTo(value, stringAddressMap[key]); err != nil {
				fmt.Printf("An error has occured: %v\n", err)
			}
		}
		udpServer.SetReadDeadline(time.Now().Add(1 * time.Second))
		buf := make([]byte, 1024)
		n, addr, err := udpServer.ReadFrom(buf)
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				continue
			}
			fmt.Println(err)
			continue
		}
		fmt.Printf("Received %d bytes from %s: %v\n", n, addr.String(), buf[:n])

		request := common.Deserialize(buf[:n])
		request[constant.Address] = addr.String()
		stringAddressMap[addr.String()] = addr
		responses, err = service.HandleIncomingRequest(request, db, reservationMap, addressToFlightMap, flightToAddressMap)
		if err != nil {
			fmt.Printf("An error has occured: %v\n", err)
			errResp := service.HandleError(err)
			if _, err := udpServer.WriteTo(errResp, addr); err != nil {
				fmt.Printf("An error has occured: %v\n", err)
			}

		} else {
			for key, value := range responses {
				if _, err := udpServer.WriteTo(value, stringAddressMap[key]); err != nil {
					fmt.Printf("An error has occured: %v\n", err)
				}
			}
		}
	}
}
