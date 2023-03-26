package service

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/YangXingHao96/DistributedSystem/package/common/constant"
)

var serviceMap = map[int]func(map[string]interface{}, *sql.DB) ([]byte, error){
	constant.QueryFlights:      GetFlightNumbers,
	constant.QueryFlightDetail: GetFlightDetails,
	constant.AddFlight:         AddFlight,
	constant.MakeReservation:   MakeReservation,
	constant.CancelReservation: CancelReservation,
}

func HandleIncomingRequest(req map[string]interface{}, db *sql.DB) ([]byte, error) {
	if _, ok := req[constant.RequestType]; !ok {
		return nil, errors.New("request does not have requestType field")
	}
	requestType, ok := req[constant.RequestType].(int)
	if !ok {
		return nil, errors.New("request field requestType not of type int")
	}
	if _, exist := serviceMap[requestType]; !exist {
		return nil, errors.New(fmt.Sprintf("request type %v not supported", requestType))
	}
	resp, err := serviceMap[requestType](req, db)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
