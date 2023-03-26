package utils

import (
	"github.com/YangXingHao96/DistributedSystem/package/common/constant"
)

var Validators = map[int]func(map[string]interface{}) *error{
	constant.QueryFlights:      validateQueryFlightsRequest,
	constant.QueryFlightDetail: validateQueryFlightDetailsRequest,
	constant.AddFlight:         validateAddFlightRequest,
	constant.MakeReservation:   validateMakeReservationRequest,
	constant.CancelReservation: validateCancelReservationRequest,
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func validateQueryFlightsRequest(req map[string]interface{}) *error {
	return nil
}

func validateQueryFlightDetailsRequest(req map[string]interface{}) *error {
	return nil
}

func validateAddFlightRequest(req map[string]interface{}) *error {
	return nil
}

func validateMakeReservationRequest(req map[string]interface{}) *error {
	return nil
}

func validateCancelReservationRequest(req map[string]interface{}) *error {
	return nil
}
