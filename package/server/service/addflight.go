package service

import (
	"database/sql"
	"fmt"
	"github.com/YangXingHao96/DistributedSystem/package/common"
	"github.com/YangXingHao96/DistributedSystem/package/common/constant"
	"github.com/YangXingHao96/DistributedSystem/package/server/database"
)

func AddFlight(req map[string]interface{}, db *sql.DB, reservationMap map[string]map[int]int) ([]byte, error) {
	flightNo, _ := req[constant.FlightNo].(int)
	source := fmt.Sprintf("%v", req[constant.Source])
	destination := fmt.Sprintf("%v", req[constant.Destination])
	departureHour, _ := req[constant.DepartureHour].(int)
	departureMin, _ := req[constant.DepartureMin].(int)
	totalSeatCnt, _ := req[constant.TotalSeats].(int)
	currentSeatCnt, _ := req[constant.CurrentSeats].(int)
	airFare, _ := req[constant.AirFare].(float32)
	err := database.AddFlight(db, flightNo, departureHour, departureMin, totalSeatCnt, currentSeatCnt, source, destination, airFare)
	if err != nil {
		return nil, err
	}
	ack := fmt.Sprintf("flight number %v added", flightNo)
	resp := common.NewSerializeAddFlightResp(ack)
	return resp, nil
}
