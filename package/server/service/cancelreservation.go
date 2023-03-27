package service

import (
	"database/sql"
	"fmt"
	"github.com/YangXingHao96/DistributedSystem/package/common"
	"github.com/YangXingHao96/DistributedSystem/package/common/constant"
	"github.com/YangXingHao96/DistributedSystem/package/server/database"
)

func CancelReservation(req map[string]interface{}, db *sql.DB, reservationMap map[string]map[int]int) ([]byte, error) {
	flightNo, _ := req[constant.FlightNo].(int)
	seatCnt, _ := req[constant.SeatCnt].(int)
	userAddr := fmt.Sprintf("%v", req[constant.Address])
	err := RemoveReservationMap(flightNo, userAddr, seatCnt, reservationMap)
	if err != nil {
		return nil, err
	}
	remainingSeats, err := database.CancelReservation(db, flightNo, seatCnt)
	if err != nil {
		return nil, err
	}
	ack := fmt.Sprintf("cancel reservation of %v seats made for flight number %v, remaining %v seats\n", seatCnt, flightNo, remainingSeats)
	fmt.Println(ack)
	resp := common.NewSerializeCancelReservationResp(ack)
	return resp, nil
}
