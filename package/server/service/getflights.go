package service

import (
	"database/sql"
	"fmt"
	"github.com/YangXingHao96/DistributedSystem/package/common"
	"github.com/YangXingHao96/DistributedSystem/package/common/constant"
	"github.com/YangXingHao96/DistributedSystem/package/server/database"
	"time"
)

func GetFlightNumbers(req map[string]interface{}, db *sql.DB, reservationMap map[string]map[int]int, addressToFlightMap map[string]map[int]time.Time, flightToAddressMap map[int]map[string]time.Time) (map[string][]byte, []byte, error) {
	source := fmt.Sprintf("%v", req[constant.Source])
	destination := fmt.Sprintf("%v", req[constant.Destination])
	userAddr := fmt.Sprintf("%v", req[constant.Address])
	flightIds, err := database.GetFlights(db, source, destination)
	if err != nil {
		return nil, nil, err
	}
	//storeMsg := "previous response: flight ids: "
	for _, flightId := range flightIds {
		fmt.Printf("flight id: %v\n", flightId)
		//	storeMsg += fmt.Sprintf("%d ", flightId)
	}
	//storeMsg += "added"
	resp := common.NewSerializeGetFlightIdBySourceDestResp(flightIds)
	responses := map[string][]byte{
		userAddr: resp,
	}
	return responses, resp, nil
}
