package utils

import (
	"errors"
	"github.com/YangXingHao96/DistributedSystem/package/common/constant"
	"reflect"
)

var Validators = map[int]func(map[string]interface{}) error{
	constant.QueryFlightsReq:            validateQueryFlightsRequest,
	constant.QueryFlightDetailReq:       validateQueryFlightDetailsRequest,
	constant.AddFlightReq:               validateAddFlightRequest,
	constant.MakeReservationReq:         validateMakeReservationRequest,
	constant.CancelReservationReq:       validateCancelReservationRequest,
	constant.GetReservationForFlightReq: validateGetReservationRequest,
	constant.RegisterForMonitorReq:      validateRegisterForMonitor,
}

func ValidateMessageId(req map[string]interface{}) error {
	if _, ok := req[constant.MessageId]; !ok {
		return errors.New("request message id cannot be empty")
	}
	if reflect.TypeOf(req[constant.MessageId]).Kind() != reflect.String {
		return errors.New("request message id not of type string")
	}
	return nil
}

func validateQueryFlightsRequest(req map[string]interface{}) error {
	err := ValidateMessageId(req)
	if err != nil {
		return err
	}
	if _, ok := req[constant.Source]; !ok {
		return errors.New("request queryFlights source cannot be empty")
	}
	if reflect.TypeOf(req[constant.Source]).Kind() != reflect.String {
		return errors.New("request queryFlights source not of type string")
	}
	if _, ok := req[constant.Destination]; !ok {
		return errors.New("request queryFlights destination cannot be empty")
	}
	if reflect.TypeOf(req[constant.Destination]).Kind() != reflect.String {
		return errors.New("request queryFlights destination not of type string")
	}
	return nil
}

func validateQueryFlightDetailsRequest(req map[string]interface{}) error {
	err := ValidateMessageId(req)
	if err != nil {
		return err
	}
	if _, ok := req[constant.FlightNo]; !ok {
		return errors.New("request queryFlightDetails flight number cannot be empty")
	}
	if reflect.TypeOf(req[constant.FlightNo]).Kind() != reflect.Int {
		return errors.New("request queryFlightDetails flight number not of type int")
	}
	return nil
}

func validateAddFlightRequest(req map[string]interface{}) error {
	err := ValidateMessageId(req)

	var ok bool
	if err != nil {
		return err
	}
	if _, ok := req[constant.FlightNo]; !ok {
		return errors.New("request addFlight flight number cannot be empty")
	}
	if reflect.TypeOf(req[constant.FlightNo]).Kind() != reflect.Int {
		return errors.New("request addFlight flight number not of type int")
	}
	if _, ok := req[constant.Source]; !ok {
		return errors.New("request addFlight source cannot be empty")
	}
	if reflect.TypeOf(req[constant.Source]).Kind() != reflect.String {
		return errors.New("request addFlight source not of type string")
	}
	if _, ok := req[constant.Destination]; !ok {
		return errors.New("request addFlight destination cannot be empty")
	}
	if reflect.TypeOf(req[constant.Destination]).Kind() != reflect.String {
		return errors.New("request addFlight destination not of type string")
	}
	if _, ok := req[constant.FlightTime]; !ok {
		return errors.New("request flightTime departure hour cannot be empty")
	}
	if reflect.TypeOf(req[constant.FlightTime]).Kind() != reflect.Int {
		return errors.New("request flightTime departure minute not of type integer")
	}
	if _, ok := req[constant.AirFare]; !ok {
		return errors.New("request addFlight airfare cannot be empty")
	}
	if reflect.TypeOf(req[constant.AirFare]).Kind() != reflect.Float32 {
		return errors.New("request addFlight airfare not of type float32")
	}
	if _, ok = req[constant.TotalSeats]; !ok {
		return errors.New("request addFlight total seat counts cannot be empty")
	}
	if reflect.TypeOf(req[constant.TotalSeats]).Kind() != reflect.Int {
		return errors.New("request addFlight total seat counts not of type integer")
	}
	totalSeats, _ := req[constant.TotalSeats].(int)
	if totalSeats <= 0 {
		return errors.New("request addFlight total seat counts cannot be less than or equal 0")
	}
	if _, ok := req[constant.CurrentSeats]; !ok {
		return errors.New("request addFlight current seat counts cannot be empty")
	}
	if reflect.TypeOf(req[constant.CurrentSeats]).Kind() != reflect.Int {
		return errors.New("request addFlight current seat counts not of type integer")
	}
	currentSeats, _ := req[constant.CurrentSeats].(int)
	if currentSeats < 0 {
		return errors.New("request addFlight current seat counts cannot be less than 0")
	}
	if currentSeats > totalSeats {
		return errors.New("request addFlight current seat counts cannot be greater than total seats on flight")
	}

	return nil
}

func validateMakeReservationRequest(req map[string]interface{}) error {
	err := ValidateMessageId(req)
	if err != nil {
		return err
	}
	if _, ok := req[constant.FlightNo]; !ok {
		return errors.New("request makeReservation flight number cannot be empty")
	}
	if reflect.TypeOf(req[constant.FlightNo]).Kind() != reflect.Int {
		return errors.New("request makeReservation flight number not of type int")
	}
	if _, ok := req[constant.SeatCnt]; !ok {
		return errors.New("request makeReservation seat count cannot be empty")
	}
	if reflect.TypeOf(req[constant.SeatCnt]).Kind() != reflect.Int {
		return errors.New("request makeReservation seat count not of type int")
	}
	return nil
}

func validateCancelReservationRequest(req map[string]interface{}) error {
	err := ValidateMessageId(req)
	if err != nil {
		return err
	}
	if _, ok := req[constant.FlightNo]; !ok {
		return errors.New("request cancelReservation flight number cannot be empty")
	}
	if reflect.TypeOf(req[constant.FlightNo]).Kind() != reflect.Int {
		return errors.New("request cancelReservation flight number not of type int")
	}
	if _, ok := req[constant.SeatCnt]; !ok {
		return errors.New("request cancelReservation seat count cannot be empty")
	}
	if reflect.TypeOf(req[constant.SeatCnt]).Kind() != reflect.Int {
		return errors.New("request cancelReservation seat count not of type int")
	}
	return nil
}

func validateGetReservationRequest(req map[string]interface{}) error {
	err := ValidateMessageId(req)
	if err != nil {
		return err
	}
	if _, ok := req[constant.FlightNo]; !ok {
		return errors.New("request getReservation flight number cannot be empty")
	}
	if reflect.TypeOf(req[constant.FlightNo]).Kind() != reflect.Int {
		return errors.New("request getReservation flight number not of type int")
	}
	return nil
}

func validateRegisterForMonitor(req map[string]interface{}) error {
	err := ValidateMessageId(req)
	if err != nil {
		return err
	}
	if _, ok := req[constant.FlightNo]; !ok {
		return errors.New("request registerForMonitor flight number cannot be empty")
	}
	if reflect.TypeOf(req[constant.FlightNo]).Kind() != reflect.Int {
		return errors.New("request registerForMonitor flight number not of type int")
	}
	if _, ok := req[constant.MonitorIntervalSec]; !ok {
		return errors.New("request registerForMonitor monitor interval cannot be empty")
	}
	if reflect.TypeOf(req[constant.MonitorIntervalSec]).Kind() != reflect.Int {
		return errors.New("request registerForMonitor monitor interval not of type int")
	}
	return nil
}
