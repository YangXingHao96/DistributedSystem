package cmd

import (
	"fmt"
	"github.com/YangXingHao96/DistributedSystem/package/common"
	"github.com/YangXingHao96/DistributedSystem/package/common/constant"
	"github.com/lithammer/shortuuid/v4"
	"github.com/manifoldco/promptui"
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

	return common.NewSerializeGetFlightIdBySourceDest(shortuuid.New(), source, dest), nil
}

func fmtGetFlightIdBySourceDest(resp map[string]interface{}) {
	flightNo := resp[constant.FlightNos]
	fmt.Printf("Flight IDs: %s\n", flightNo)
}
