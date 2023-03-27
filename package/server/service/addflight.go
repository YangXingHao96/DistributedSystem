package service

import (
	"database/sql"
	"fmt"
	"github.com/YangXingHao96/DistributedSystem/package/common"
	"github.com/YangXingHao96/DistributedSystem/package/common/constant"
	"github.com/YangXingHao96/DistributedSystem/package/server/database"
	"time"
)

func AddFlight(req map[string]interface{}, db *sql.DB, reservationMap map[string]map[int]int, addressToFlightMap map[string]map[int]time.Time, flightToAddressMap map[int]map[string]time.Time) (map[string][]byte, error) {
	flightNo, _ := req[constant.FlightNo].(int)
	source := fmt.Sprintf("%v", req[constant.Source])
	destination := fmt.Sprintf("%v", req[constant.Destination])
	flightTime, _ := req[constant.FlightTime].(int)
	totalSeatCnt, _ := req[constant.TotalSeats].(int)
	currentSeatCnt, _ := req[constant.CurrentSeats].(int)
	airFare, _ := req[constant.AirFare].(float32)
	userAddr := fmt.Sprintf("%v", req[constant.Address])
	err := database.AddFlight(db, flightNo, flightTime, totalSeatCnt, currentSeatCnt, source, destination, airFare)
	if err != nil {
		return nil, err
	}
	ack := fmt.Sprintf("flight number %v added", flightNo)
	resp := common.NewSerializeAddFlightResp(ack)
	responses := map[string][]byte{
		userAddr: resp,
	}
	return responses, nil
}
