package service

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/YangXingHao96/DistributedSystem/package/common/constant"
	"github.com/YangXingHao96/DistributedSystem/package/server/utils"
)

var serviceMap = map[int]func(map[string]interface{}, *sql.DB) ([]byte, error){
	constant.QueryFlightsReq:      GetFlightNumbers,
	constant.QueryFlightDetailReq: GetFlightDetails,
	constant.AddFlightReq:         AddFlight,
	constant.MakeReservationReq:   MakeReservation,
	constant.CancelReservationReq: CancelReservation,
}

func HandleDuplicateRequest(req map[string]interface{}, msgMap map[string]struct{}) error {
	fmt.Printf("Processing incoming request: %v\n", req)
	err := utils.ValidateMessageId(req)
	if err != nil {
		return err
	}
	msgId := fmt.Sprintf("%v", req[constant.MessageId])
	if _, ok := msgMap[msgId]; ok {
		return errors.New("duplicated message id, request will not be process")
	}
	msgMap[msgId] = struct{}{}
	return nil
}

func HandleIncomingRequest(req map[string]interface{}, db *sql.DB) ([]byte, error) {
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
	resp, err := serviceMap[requestType](req, db)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
