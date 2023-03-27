package service

import (
	"database/sql"
	"fmt"
	"github.com/YangXingHao96/DistributedSystem/package/common/constant"
	"time"
)

func RegisterForMonitor(req map[string]interface{}, db *sql.DB, reservationMap map[string]map[int]int, addressToFlightMap map[string]map[int]time.Time, flightToAddressMap map[int]map[string]time.Time) (map[string][]byte, error) {
	userAddr := fmt.Sprintf("%v", req[constant.Address])
	flightNo, _ := req[constant.FlightNo].(int)
	monitorInterval, _ := req[constant.MonitorIntervalSec].(int)
	monitorEndDuration := time.Now().UTC().Add(time.Duration(monitorInterval) * time.Second)
	if _, ok := addressToFlightMap[userAddr]; !ok {
		addressToFlightMap[userAddr] = map[int]time.Time{}
	}
	addressToFlightMap[userAddr][flightNo] = monitorEndDuration
	if _, ok := flightToAddressMap[flightNo]; !ok {
		flightToAddressMap[flightNo] = map[string]time.Time{}
	}
	flightToAddressMap[flightNo][userAddr] = monitorEndDuration
	fmt.Printf("flight %v has been registered for monitoring for user %v, will be droped at %v\n", flightNo, userAddr, monitorEndDuration)
	resp := map[string][]byte{}
	return resp, nil
}
