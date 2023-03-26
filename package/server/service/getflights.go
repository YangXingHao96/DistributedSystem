package service

import (
	"database/sql"
	"fmt"
	"github.com/YangXingHao96/DistributedSystem/package/common"
	"github.com/YangXingHao96/DistributedSystem/package/common/constant"
	"github.com/YangXingHao96/DistributedSystem/package/server/database"
)

func GetFlightNumbers(req map[string]interface{}, db *sql.DB) ([]byte, error) {
	source := fmt.Sprintf("%v", req[constant.Source])
	destination := fmt.Sprintf("%v", req[constant.Destination])

	flightIds, err := database.GetFlights(db, source, destination)
	if err != nil {
		return nil, err
	}
	for _, flightId := range flightIds {
		fmt.Printf("flight id: %v\n", flightId)
	}
	resp := common.NewSerializeGetFlightIdBySourceDestResp(flightIds)
	return resp, nil
}
