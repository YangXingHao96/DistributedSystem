package service

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/YangXingHao96/DistributedSystem/package/common"
	"github.com/YangXingHao96/DistributedSystem/package/common/constant"
	"github.com/YangXingHao96/DistributedSystem/package/server/database"
)

func GetReservation(req map[string]interface{}, db *sql.DB, reservationMap map[string]map[int]int) ([]byte, error) {
	userAddr := fmt.Sprintf("%v", req[constant.Address])
	flightNo, _ := req[constant.FlightNo].(int)
	flightDetail, err := database.GetFlightDetail(db, flightNo)
	var userMap map[int]int
	var reservationCnt int
	if err != nil {
		return nil, err
	}
	if flightDetail != nil && flightDetail.FlightNo == 0 {
		return nil, errors.New("flight not found in the system")
	}
	userMap, exist := reservationMap[userAddr]
	if !exist {
		return nil, errors.New(fmt.Sprintf("user %v does not have any reservation", userAddr))
	}
	reservationCnt, exist = userMap[flightNo]
	if !exist {
		return nil, errors.New(fmt.Sprintf("user %v does not have any reservation under flight %v", userAddr, flightNo))
	}
	resp := common.NewSerializeGetReservationForFlightResp(flightNo, reservationCnt)

	return resp, nil

}
