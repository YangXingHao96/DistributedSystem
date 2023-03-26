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
	validate := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("invalid number")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Flight ID",
		Validate: validate,
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
