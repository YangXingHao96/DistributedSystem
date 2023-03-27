package cmd

import (
	"errors"
	"fmt"
	"github.com/YangXingHao96/DistributedSystem/package/client"
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

	prompt = promptui.Prompt{
		Label:    "Message Request ID (Optional, leave blank if unsure)",
	}
	msgId, err := prompt.Run()
	if err != nil {
		return nil, err
	}
	if msgId == "" {
		msgId = shortuuid.New()
	}

	data := common.NewSerializeGetFlightIdBySourceDest(msgId, source, dest)
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

	prompt = promptui.Prompt{
		Label:    "Message Request ID (Optional, leave blank if unsure)",
	}
	msgId, err := prompt.Run()
	if err != nil {
		return nil, err
	}
	if msgId == "" {
		msgId = shortuuid.New()
	}


	data := common.NewSerializeQueryFlightDetailReq(msgId, x)
	return data, nil
}

func fmtGetFlightDetail(resp map[string]interface{}) string {
	flightNo := resp[constant.FlightNo].(int)
	if flightNo == 0 {
		return "No matching entry found for given flight id"
	}
	source := resp[constant.Source]
	dest := resp[constant.Source]
	availSeats := resp[constant.AvailableSeats].(int)
	return fmt.Sprintf("Flight ID: %d\nSource: %s\nDestination: %s\nAvailable Seats left: %d\n", flightNo, source, dest, availSeats)
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
		Label:    "Current number of seats booked",
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

	prompt = promptui.Prompt{
		Label:    "Message Request ID (Optional, leave blank if unsure)",
	}
	msgId, err := prompt.Run()
	if err != nil {
		return nil, err
	}
	if msgId == "" {
		msgId = shortuuid.New()
	}


	data := common.NewSerializeAddFlightReq(msgId, x, source, dest, depHr, depMin, float32(airFare), ttlSeatCnt, curSeatCnt)
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

	prompt = promptui.Prompt{
		Label:    "Message Request ID (Optional, leave blank if unsure)",
	}
	msgId, err := prompt.Run()
	if err != nil {
		return nil, err
	}
	if msgId == "" {
		msgId = shortuuid.New()
	}


	data := common.NewSerializeMakeReservationReq(msgId, x, seatCnt)
	return data, nil
}

func promptCancelReservation() ([]byte, error) {
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
		Label:    "Number of seats to cancel",
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

	prompt = promptui.Prompt{
		Label:    "Message Request ID (Optional, leave blank if unsure)",
	}
	msgId, err := prompt.Run()
	if err != nil {
		return nil, err
	}
	if msgId == "" {
		msgId = shortuuid.New()
	}


	data := common.NewSerializeCancelReservationReq(msgId, x, seatCnt)
	return data, nil
}

func promptGetReservationForFlight() ([]byte, error) {
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
		Label:    "Message Request ID (Optional, leave blank if unsure)",
	}
	msgId, err := prompt.Run()
	if err != nil {
		return nil, err
	}
	if msgId == "" {
		msgId = shortuuid.New()
	}

	data := common.NewSerializeGetReservationForFlightReq(msgId, x)
	return data, nil
}

func fmtGetReservationForFlight(resp map[string]interface{}) string {
	flightNo := resp[constant.FlightNo].(int)
	seatCnt := resp[constant.SeatCnt].(int)
	return fmt.Sprintf("Flight ID: %s\nSeats reserved: %d\n", flightNo, seatCnt)
}

func promptRegisterMonitorReq() ([]byte, error) {
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
		Label:    "Monitor time interval (in seconds)",
		Validate: func(input string) error {
			x, err := strconv.ParseInt(input, 10, 32)
			if err != nil {
				return errors.New("invalid number")
			}
			if int(x) * 1000 >= client.ReadTimeoutMs {
				return errors.New(fmt.Sprintf("monitor interval cannot be longer than configured read timeout: %d seconds", client.ReadTimeoutMs / 1000))
			}
			return nil
		},
	}
	monitorIntervalSecStr, err := prompt.Run()
	if err != nil {
		return nil, err
	}
	monitorIntervalSec, err := strconv.Atoi(monitorIntervalSecStr)
	if err != nil {
		return nil, err
	}

	prompt = promptui.Prompt{
		Label:    "Message Request ID (Optional, leave blank if unsure)",
	}
	msgId, err := prompt.Run()
	if err != nil {
		return nil, err
	}
	if msgId == "" {
		msgId = shortuuid.New()
	}

	data := common.NewSerializeRegisterForMonitorReq(msgId, x, monitorIntervalSec)
	return data, nil
}

func fmtMonitorFlightResp(resp map[string]interface{}) string {
	msgType := resp[constant.MsgType].(int)
	if msgType == constant.MonitorUpdateResp {
		flightNo := resp[constant.FlightNo].(int)
		seatsAvailable := resp[constant.AvailableSeats].(int)
		return fmt.Sprintf("❗Flight %d seat availability alert\nNew seat availability: %d\n", flightNo, seatsAvailable)
	}

	return fmt.Sprintf("Stopped monitoring. Reason from server: %s", resp[constant.Ack].(string))
}

func fmtGeneralErrResp(resp map[string]interface{}) string {
	return fmt.Sprintf("Something went wrong processing your request. Reason from server: %s", resp[constant.ErrMessage].(string))
}
