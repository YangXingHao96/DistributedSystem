package cmd

import (
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
	return fmt.Sprintf("Flight IDs: %s\n", strings.Join(res, ","))
}
