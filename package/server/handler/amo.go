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

func HandleUDPRequestAtMostOnce(db *sql.DB, timeout bool, host string, port string, timeoutPercentage int) {
	// Listen for incoming packets on port 8080
	storedResponses := map[string][]byte{}
	addr := host + ":" + port
	udpServer, err := net.ListenPacket("udp", addr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("server listening on address: %v\n", addr)
	defer udpServer.Close()
	reservationMap := map[string]map[int]int{}
	addressToFlightMap := map[string]map[int]time.Time{}
	flightToAddressMap := map[int]map[string]time.Time{}
	stringAddressMap := map[string]net.Addr{}
	for {
		//add these 2 lines to monitor the callback
		//fmt.Println(addressToFlightMap)
		//fmt.Println(flightToAddressMap)
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

		if service.SimulateRandomTimeOut(timeout, timeoutPercentage) {
			fmt.Println("simulate server timeout, no action will be performed")
			continue
		}

		request := common.Deserialize(buf[:n])
		request[constant.Address] = addr.String()
		stringAddressMap[addr.String()] = addr

		resp, err := service.HandleDuplicateRequest(request, storedResponses)
		if resp != nil {
			fmt.Println("handling repeated request")
			if _, err := udpServer.WriteTo(resp, addr); err != nil {
				fmt.Printf("An error has occured: %v\n", err)
			}
			continue
		}
		// We will not store the mapping of error response to specific messages send
		if err != nil {
			resp = service.HandleError(err)
			fmt.Printf("An error has occured: %v\n", err)
			if _, err := udpServer.WriteTo(resp, addr); err != nil {
				fmt.Printf("An error has occured: %v\n", err)
			}
			continue
		}

		responses, resp, err = service.HandleIncomingRequest(request, db, reservationMap, addressToFlightMap, flightToAddressMap)
		if err != nil {
			fmt.Printf("An error has occured: %v\n", err)
			resp = service.HandleError(err)
			if _, err := udpServer.WriteTo(resp, addr); err != nil {
				fmt.Printf("An error has occured: %v\n", err)
			}
		} else {
			service.RegisterResponses(request, resp, storedResponses)
			for key, value := range responses {
				if _, err := udpServer.WriteTo(value, stringAddressMap[key]); err != nil {
					fmt.Printf("An error has occured: %v\n", err)
				}
			}
		}
	}
}
