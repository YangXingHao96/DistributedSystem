package service

import (
	"database/sql"
	"fmt"
	"github.com/YangXingHao96/DistributedSystem/package/common"
	"github.com/YangXingHao96/DistributedSystem/package/common/constant"
	"github.com/YangXingHao96/DistributedSystem/package/server/database"
	"time"
)

func CancelReservation(req map[string]interface{}, db *sql.DB, reservationMap map[string]map[int]int, addressToFlightMap map[string]map[int]time.Time, flightToAddressMap map[int]map[string]time.Time) (map[string][]byte, []byte, error) {
	flightNo, _ := req[constant.FlightNo].(int)
	seatCnt, _ := req[constant.SeatCnt].(int)
	userAddr := fmt.Sprintf("%v", req[constant.Address])
	err := RemoveReservationMap(flightNo, userAddr, seatCnt, reservationMap)
	if err != nil {
		return nil, nil, err
	}
	remainingSeats, err := database.CancelReservation(db, flightNo, seatCnt)
	if err != nil {
		return nil, nil, err
	}

	ack := fmt.Sprintf("cancel reservation of %v seats made for flight number %v, remaining %v seats\n", seatCnt, flightNo, remainingSeats)
	storedMsg := "previous response: " + ack
	fmt.Println(ack)
	resp := common.NewSerializeCancelReservationResp(ack)
	storeResp := common.NewSerializeCancelReservationResp(storedMsg)
	responses := map[string][]byte{
		userAddr: resp,
	}
	if addrTimeMap, ok := flightToAddressMap[flightNo]; ok {
		broadCastResp := common.NewSerializeMonitorUpdateResp(flightNo, remainingSeats)
		for key, _ := range addrTimeMap {
			responses[key] = broadCastResp
		}
	}
	return responses, storeResp, nil
}
