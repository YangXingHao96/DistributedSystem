package cmd

import (
	"errors"
	"fmt"
	"github.com/YangXingHao96/DistributedSystem/package/common"
	"github.com/YangXingHao96/DistributedSystem/package/common/constant"
	"github.com/lithammer/shortuuid/v4"
	"github.com/manifoldco/promptui"
	"math"
	"strconv"
	"strings"
	"time"
)

func promptGetFlightIdBySourceDest() ([]byte, error) {
	prompt := promptui.Prompt{
		Label: "Flight source",
	}
	source, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	prompt = promptui.Prompt{
		Label: "Flight destination",
	}
	dest, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	prompt = promptui.Prompt{
		Label: "Message Request ID (Optional, leave blank if unsure)",
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
		Label: "Flight ID",
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
		Label: "Message Request ID (Optional, leave blank if unsure)",
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
	dest := resp[constant.Destination]
	availSeats := resp[constant.AvailableSeats].(int)
	flightTime := resp[constant.FlightTime].(int)
	layout := "01/02/2006 3:04:05 PM"
	flightTimeStr := time.Unix(int64(flightTime), 0).UTC().Format(layout)
	airFare := resp[constant.AirFare].(float32)
	return fmt.Sprintf("Flight ID: %d\nSource: %s\nDestination: %s\nAvailable Seats left: %d\nFlight Time: %s\nAir Fare: %.2f", flightNo, source, dest, availSeats, flightTimeStr, airFare)
}

func promptAddFlight() ([]byte, error) {
	prompt := promptui.Prompt{
		Label: "Flight ID",
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
		Label: "Flight source",
	}
	source, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	prompt = promptui.Prompt{
		Label: "Flight destination",
	}
	dest, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	layout := "02/01/2006 03:04:05 PM"
	prompt = promptui.Prompt{
		Label: "Flight Departure Time in the format \"DD/MM/YYYY hr:min:second AM/PM\", Eg: 01/02/2006 03:04:05 PM",
		Validate: func(input string) error {
			t, err := time.Parse(layout, input)
			if err != nil {
				return err
			}
			if t.Unix() > math.MaxInt32 {
				return errors.New("cannot store this date, try something before 2038")
			}
			return nil
		},
	}
	depTimeStr, err := prompt.Run()
	if err != nil {
		return nil, err
	}
	t, err := time.Parse(layout, depTimeStr)
	depTime := t.Unix()

	prompt = promptui.Prompt{
		Label: "Flight Airfare",
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
		Label: "Total seat count",
		Validate: func(input string) error {
			cnt, err := strconv.ParseInt(input, 10, 32)
			if err != nil {
				return errors.New("invalid number")
			}
			if cnt <= 0 {
				return errors.New("total seat count cannot be less than or equal 0")
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
		Label: "Current number of seats booked",
		Validate: func(input string) error {
			cnt, err := strconv.ParseInt(input, 10, 32)
			if err != nil {
				return errors.New("invalid number")
			}
			if cnt < 0 {
				errors.New("current number of seats booked cannot be less than 0")
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
		Label: "Message Request ID (Optional, leave blank if unsure)",
	}
	msgId, err := prompt.Run()
	if err != nil {
		return nil, err
	}
	if msgId == "" {
		msgId = shortuuid.New()
	}

	data := common.NewSerializeAddFlightReq(msgId, x, source, dest, int(depTime), float32(airFare), ttlSeatCnt, curSeatCnt)
	return data, nil
}

func fmtSimpleAck(resp map[string]interface{}) string {
	return resp[constant.Ack].(string)
}

func promptMakeReservation() ([]byte, error) {
	prompt := promptui.Prompt{
		Label: "Flight ID",
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
		Label: "Number of seats to reserve",
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
		Label: "Message Request ID (Optional, leave blank if unsure)",
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
		Label: "Flight ID",
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
		Label: "Number of seats to cancel",
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
		Label: "Message Request ID (Optional, leave blank if unsure)",
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
		Label: "Flight ID",
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
		Label: "Message Request ID (Optional, leave blank if unsure)",
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
	return fmt.Sprintf("Flight ID: %d\nSeats reserved: %d\n", flightNo, seatCnt)
}

func promptRegisterMonitorReq() ([]byte, error) {
	prompt := promptui.Prompt{
		Label: "Flight ID",
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
		Label: "Monitor time interval (in seconds)",
		Validate: func(input string) error {
			_, err := strconv.ParseInt(input, 10, 32)
			if err != nil {
				return errors.New("invalid number")
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
		Label: "Message Request ID (Optional, leave blank if unsure)",
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
		return fmt.Sprintf("\n❗ Flight %d seat availability alert\nNew seat availability: %d\n", flightNo, seatsAvailable)
	}

	return fmt.Sprintf("Stopped monitoring. Reason from server: %s", resp[constant.Ack].(string))
}

func fmtGeneralErrResp(resp map[string]interface{}) string {
	return fmt.Sprintf("Something went wrong processing your request. Reason from server: %s\n", resp[constant.ErrMessage].(string))
}
