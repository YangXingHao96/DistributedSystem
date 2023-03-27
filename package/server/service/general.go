package service

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/YangXingHao96/DistributedSystem/package/common"
	"github.com/YangXingHao96/DistributedSystem/package/common/constant"
	"github.com/YangXingHao96/DistributedSystem/package/server/utils"
)

var serviceMap = map[int]func(map[string]interface{}, *sql.DB, map[string]map[int]int) ([]byte, error){
	constant.QueryFlightsReq:            GetFlightNumbers,
	constant.QueryFlightDetailReq:       GetFlightDetails,
	constant.AddFlightReq:               AddFlight,
	constant.MakeReservationReq:         MakeReservation,
	constant.CancelReservationReq:       CancelReservation,
	constant.GetReservationForFlightReq: GetReservation,
}

func HandleDuplicateRequest(req map[string]interface{}, msgMap map[string][]byte) ([]byte, error) {
	fmt.Printf("Processing incoming request: %v\n", req)
	err := utils.ValidateMessageId(req)
	if err != nil {
		return nil, err
	}
	msgId := fmt.Sprintf("%v", req[constant.MessageId])
	if _, ok := msgMap[msgId]; !ok {
		return nil, nil
	}

	return msgMap[msgId], nil
}

func HandleIncomingRequest(req map[string]interface{}, db *sql.DB, reservationMap map[string]map[int]int) ([]byte, error) {
	fmt.Printf("Processing incoming request: %v\n", req)
	if _, ok := req[constant.MsgType]; !ok {
		return nil, errors.New("request does not have requestType field")
	}
	requestType, ok := req[constant.MsgType].(int)
	if !ok {
		return nil, errors.New("request field requestType not of type int")
	}

	if _, exist := serviceMap[requestType]; !exist {
		return nil, errors.New(fmt.Sprintf("request type %v not supported", requestType))
	}
	err := utils.Validators[requestType](req)
	if err != nil {
		return nil, err
	}
	resp, err := serviceMap[requestType](req, db, reservationMap)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func HandleError(err error) []byte {
	return common.NewSerializeGeneralErrResp(err.Error())
}

func AddReservationMap(flightNo int, userAddr string, seatCnt int, reservationMap map[string]map[int]int) {
	innerMap, ok := reservationMap[userAddr]
	if !ok {
		innerMap = make(map[int]int)
		reservationMap[userAddr] = innerMap
	}
	currentCnt, ok := innerMap[flightNo]
	if !ok {
		innerMap[flightNo] = seatCnt
	} else {
		innerMap[flightNo] = currentCnt + seatCnt
	}
}

func RemoveReservationMap(flightNo int, userAddr string, seatCnt int, reservationMap map[string]map[int]int) error {
	innerMap, ok := reservationMap[userAddr]
	if !ok {
		return errors.New(fmt.Sprintf("user %v has not made any reservation, cannot perform cancellation", userAddr))
	}
	currentCnt, ok := innerMap[flightNo]
	if !ok {
		return errors.New(fmt.Sprintf("user %v has not made any reservation under flight %v, cannot perform cancellation", userAddr, flightNo))
	}
	if seatCnt > currentCnt {
		return errors.New(fmt.Sprintf("user %v has %v reservations for flight %v, cannot perform cancellation for %v seats", userAddr, currentCnt, flightNo, seatCnt))
	}
	innerMap[flightNo] = currentCnt - seatCnt
	if innerMap[flightNo] == 0 {
		delete(innerMap, flightNo)
	}
	return nil
}
