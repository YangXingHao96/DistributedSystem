package cmd

import (
	"errors"
	"fmt"
	"github.com/YangXingHao96/DistributedSystem/package/common"
	"github.com/YangXingHao96/DistributedSystem/package/common/constant"
	"github.com/lithammer/shortuuid/v4"
	"github.com/manifoldco/promptui"
	"strconv"
	"strings"
)

func promptGetFlightIdBySourceDest() ([]byte, error) {
	prompt := promptui.Prompt{
		Label:    "Flight source",
	}
	source, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	prompt = promptui.Prompt{
		Label:    "Flight destination",
	}
	dest, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	data := common.NewSerializeGetFlightIdBySourceDest(shortuuid.New(), source, dest)
	return data, nil
}

func fmtGetFlightIdBySourceDest(resp map[string]interface{}) string {
	flightNo := resp[constant.FlightNos].([]int)
	res := make([]string, 0)
	for _, f := range flightNo {
		res = append(res, strconv.Itoa(f))
	}
	if len(res) == 0 {
		res = append(res, "No matching entry.")
	}
	return fmt.Sprintf("Flight IDs: %s\n", strings.Join(res, ","))
}

func promptGetFlightDetail() ([]byte, error) {
	prompt := promptui.Prompt{
		Label:    "Flight ID",
		Validate: func(input string) error {
			_, err := strconv.ParseInt(input, 10, 32)
			if err != nil {
				return errors.New("invalid number")
			}
			return nil
		},
	}
	flightId, err := prompt.Run()
	if err != nil {
		return nil, err
	}
	x, err := strconv.Atoi(flightId)
	if err != nil {
		return nil, err
	}

	data := common.NewSerializeQueryFlightDetailReq(shortuuid.New(), x)
	return data, nil
}

func fmtGetFlightDetail(resp map[string]interface{}) string {
	flightNo := resp[constant.FlightNo].(int)
	if flightNo == 0 {
		return fmt.Sprintf("No matching entry found for flight id: %d", flightNo)
	}
	source := resp[constant.Source]
	dest := resp[constant.Source]
	availSeats := resp[constant.AvailableSeats].(int)
	return fmt.Sprintf("Flight ID: %d\nSource: %s\nDestination: %s\nAvailable Seats left: %d", flightNo, source, dest, availSeats)
}

func promptAddFlight() ([]byte, error) {
	prompt := promptui.Prompt{
		Label:    "Flight ID",
		Validate: func(input string) error {
			_, err := strconv.ParseInt(input, 10, 32)
			if err != nil {
				return errors.New("invalid number")
			}
			return nil
		},
	}
	flightId, err := prompt.Run()
	if err != nil {
		return nil, err
	}
	x, err := strconv.Atoi(flightId)
	if err != nil {
		return nil, err
	}

	prompt = promptui.Prompt{
		Label:    "Flight source",
	}
	source, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	prompt = promptui.Prompt{
		Label:    "Flight destination",
	}
	dest, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	prompt = promptui.Prompt{
		Label:    "Flight Departure Hour",
		Validate: func(input string) error {
			x, err := strconv.ParseInt(input, 10, 32)
			if err != nil {
				return errors.New("invalid number")
			}
			if x < 0 || x > 24 {
				return errors.New("hour must be in range [0,23]")
			}
			return nil
		},
	}
	depHrStr, err := prompt.Run()
	if err != nil {
		return nil, err
	}
	depHr, err := strconv.Atoi(depHrStr)
	if err != nil {
		return nil, err
	}

	prompt = promptui.Prompt{
		Label:    "Flight Departure Minute",
		Validate: func(input string) error {
			x, err := strconv.ParseInt(input, 10, 32)
			if err != nil {
				return errors.New("invalid number")
			}
			if x < 0 || x > 60 {
				return errors.New("minute must be in range [0,59]")
			}
			return nil
		},
	}
	depMinStr, err := prompt.Run()
	if err != nil {
		return nil, err
	}
	depMin, err := strconv.Atoi(depMinStr)
	if err != nil {
		return nil, err
	}

	prompt = promptui.Prompt{
		Label:    "Flight Airfare",
		Validate: func(input string) error {
			_, err := strconv.ParseFloat(input, 32)
			if err != nil {
				return errors.New("invalid number")
			}
			return nil
		},
	}
	airFareStr, err := prompt.Run()
	if err != nil {
		return nil, err
	}
	airFare, err := strconv.ParseFloat(airFareStr, 32)
	if err != nil {
		return nil, err
	}

	prompt = promptui.Prompt{
		Label:    "Total seat count",
		Validate: func(input string) error {
			_, err := strconv.ParseInt(input, 10, 32)
			if err != nil {
				return errors.New("invalid number")
			}
			return nil
		},
	}
	ttlSeatCntStr, err := prompt.Run()
	if err != nil {
		return nil, err
	}
	ttlSeatCnt, err := strconv.Atoi(ttlSeatCntStr)
	if err != nil {
		return nil, err
	}

	prompt = promptui.Prompt{
		Label:    "Current seat count",
		Validate: func(input string) error {
			_, err := strconv.ParseInt(input, 10, 32)
			if err != nil {
				return errors.New("invalid number")
			}
			return nil
		},
	}
	curSeatCntStr, err := prompt.Run()
	if err != nil {
		return nil, err
	}
	curSeatCnt, err := strconv.Atoi(curSeatCntStr)
	if err != nil {
		return nil, err
	}

	data := common.NewSerializeAddFlightReq(shortuuid.New(), x, source, dest, depHr, depMin, float32(airFare), ttlSeatCnt, curSeatCnt)
	return data, nil
}

func fmtSimpleAck(resp map[string]interface{}) string {
	return resp[constant.Ack].(string)
}

func promptMakeReservation() ([]byte, error) {
	prompt := promptui.Prompt{
		Label:    "Flight ID",
		Validate: func(input string) error {
			_, err := strconv.ParseInt(input, 10, 32)
			if err != nil {
				return errors.New("invalid number")
			}
			return nil
		},
	}
	flightId, err := prompt.Run()
	if err != nil {
		return nil, err
	}
	x, err := strconv.Atoi(flightId)
	if err != nil {
		return nil, err
	}

	prompt = promptui.Prompt{
		Label:    "Number of seats to reserve",
		Validate: func(input string) error {
			_, err := strconv.ParseInt(input, 10, 32)
			if err != nil {
				return errors.New("invalid number")
			}
			return nil
		},
	}
	seatCntStr, err := prompt.Run()
	if err != nil {
		return nil, err
	}
	seatCnt, err := strconv.Atoi(seatCntStr)
	if err != nil {
		return nil, err
	}

	data := common.NewSerializeMakeReservationReq(shortuuid.New(), x, seatCnt)
	return data, nil
}
