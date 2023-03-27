package service

import (
	"database/sql"
	"fmt"
	"github.com/YangXingHao96/DistributedSystem/package/common"
	"github.com/YangXingHao96/DistributedSystem/package/common/constant"
	"github.com/YangXingHao96/DistributedSystem/package/server/database"
	"time"
)

func GetFlightDetails(req map[string]interface{}, db *sql.DB, reservationMap map[string]map[int]int, addressToFlightMap map[string]map[int]time.Time, flightToAddressMap map[int]map[string]time.Time) (map[string][]byte, error) {
	flightNo, _ := req[constant.FlightNo].(int)
	flightDetails, err := database.GetFlightDetail(db, flightNo)
	userAddr := fmt.Sprintf("%v", req[constant.Address])
	if err != nil {
		return nil, err
	}
	if flightDetails.FlightNo == 0 {
		flightDetails.SeatAvailability = 0
		flightDetails.Source = ""
		flightDetails.Destination = ""
	}
	fmt.Printf("flight detail retrieved, flight number: %v, flight source: %v, flight destination: %v, flight seat availability: %v\n", flightDetails.FlightNo, flightDetails.Source, flightDetails.Destination, flightDetails.SeatAvailability)
	resp := common.NewSerializeQueryFlightDetailResp(flightDetails.FlightNo, flightDetails.Source, flightDetails.Destination, flightDetails.SeatAvailability)
	responses := map[string][]byte{
		userAddr: resp,
	}
	return responses, nil
}
