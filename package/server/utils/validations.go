package utils

import (
	"errors"
	"github.com/YangXingHao96/DistributedSystem/package/common/constant"
	"reflect"
)

var Validators = map[int]func(map[string]interface{}) error{
	constant.QueryFlightsReq:      validateQueryFlightsRequest,
	constant.QueryFlightDetailReq: validateQueryFlightDetailsRequest,
	constant.AddFlightReq:         validateAddFlightRequest,
	constant.MakeReservationReq:   validateMakeReservationRequest,
	constant.CancelReservationReq: validateCancelReservationRequest,
}

func validateMessageId(req map[string]interface{}) error {
	if _, ok := req[constant.MessageId]; !ok {
		return errors.New("request message id cannot be empty")
	}
	if reflect.TypeOf(req[constant.MessageId]).Kind() == reflect.String {
		return errors.New("request message id not of type string")
	}
	return nil
}

func validateQueryFlightsRequest(req map[string]interface{}) error {
	err := validateMessageId(req)
	if err != nil {
		return err
	}
	if _, ok := req[constant.Source]; !ok {
		return errors.New("request query flights source cannot be empty")
	}
	if reflect.TypeOf(req[constant.Source]).Kind() == reflect.String {
		return errors.New("request query flights source not of type string")
	}
	if _, ok := req[constant.Destination]; !ok {
		return errors.New("request query flights destination cannot be empty")
	}
	if reflect.TypeOf(req[constant.Destination]).Kind() == reflect.String {
		return errors.New("request query flights destination not of type string")
	}
	return nil
}

func validateQueryFlightDetailsRequest(req map[string]interface{}) error {
	return nil
}

func validateAddFlightRequest(req map[string]interface{}) error {
	return nil
}

func validateMakeReservationRequest(req map[string]interface{}) error {
	return nil
}

func validateCancelReservationRequest(req map[string]interface{}) error {
	return nil
}
