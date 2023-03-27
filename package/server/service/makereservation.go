package service

import (
	"database/sql"
	"fmt"
	"github.com/YangXingHao96/DistributedSystem/package/common"
	"github.com/YangXingHao96/DistributedSystem/package/common/constant"
	"github.com/YangXingHao96/DistributedSystem/package/server/database"
)

func MakeReservation(req map[string]interface{}, db *sql.DB, reservationMap map[string]map[int]int) ([]byte, error) {
	flightNo, _ := req[constant.FlightNo].(int)
	seatCnt, _ := req[constant.SeatCnt].(int)
	userAddr := fmt.Sprintf("%v", req[constant.Address])
	remainingSeats, err := database.MakeReservation(db, flightNo, seatCnt)
	if err != nil {
		return nil, err
	}
	ack := fmt.Sprintf("reservation of %v seats made for flight number %v, remaining %v seats\n", seatCnt, flightNo, remainingSeats)
	fmt.Println(ack)
	AddReservationMap(flightNo, userAddr, seatCnt, reservationMap)
	resp := common.NewSerializeMakeReservationResp(ack)
	return resp, nil
}
