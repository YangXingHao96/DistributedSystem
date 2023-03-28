package common

import (
	"encoding/binary"
	"fmt"
	"github.com/YangXingHao96/DistributedSystem/package/common/constant"
	"math"
)

func NewSerializeGetFlightIdBySourceDest(msgId, source, dest string) []byte {
	serMsgId := serializeStr(msgId)
	serMsgIdSize := serializeInt32(len(serMsgId))
	serSource := serializeStr(source)
	serSourceSize := serializeInt32(len(serSource))
	serDest := serializeStr(dest)
	serDestSize := serializeInt32(len(serDest))

	buf := make([]byte, 0)
	buf = append(buf, serializeInt32(constant.QueryFlightsReq)[0])
	buf = append(buf, serMsgIdSize...)
	buf = append(buf, serMsgId...)
	buf = append(buf, serSourceSize...)
	buf = append(buf, serSource...)
	buf = append(buf, serDestSize...)
	buf = append(buf, serDest...)

	return append(serializeInt32(len(buf)+4), buf...)
}

func NewSerializeGetFlightIdBySourceDestResp(flightNo []int) []byte {
	buf := make([]byte, 0)
	buf = append(buf, serializeInt32(constant.QueryFlightsResp)[0])
	buf = append(buf, serializeInt32(len(flightNo))...)
	for _, f := range flightNo {
		buf = append(buf, serializeInt32(f)...)
	}

	return append(serializeInt32(len(buf)+4), buf...)
}

func NewSerializeQueryFlightDetailReq(msgId string, flightNo int) []byte {
	buf := make([]byte, 0)
	buf = append(buf, serializeInt32(constant.QueryFlightDetailReq)[0])
	serMsgId := serializeStr(msgId)
	serMsgIdSize := serializeInt32(len(serMsgId))
	buf = append(buf, serMsgIdSize...)
	buf = append(buf, serMsgId...)
	buf = append(buf, serializeInt32(flightNo)...)

	return append(serializeInt32(len(buf)+4), buf...)
}

func NewSerializeQueryFlightDetailResp(flightNo int, source string, dest string, availSeats int, flightTime int, airFare float32) []byte {
	serSource := serializeStr(source)
	serSourceSize := serializeInt32(len(serSource))
	serDest := serializeStr(dest)
	serDestSize := serializeInt32(len(serDest))

	buf := make([]byte, 0)
	buf = append(buf, serializeInt32(constant.QueryFlightDetailResp)[0])
	buf = append(buf, serializeInt32(flightNo)...)
	buf = append(buf, serSourceSize...)
	buf = append(buf, serSource...)
	buf = append(buf, serDestSize...)
	buf = append(buf, serDest...)
	buf = append(buf, serializeInt32(availSeats)...)
	buf = append(buf, serializeInt32(flightTime)...)
	buf = append(buf, serializeFloat32(airFare)...)

	return append(serializeInt32(len(buf)+4), buf...)
}

func NewSerializeAddFlightReq(msgId string, flightNo int, source string, dest string, flightTime int, airFare float32, totalSeatCount int, currentSeatCount int) []byte {
	buf := make([]byte, 0)
	buf = append(buf, serializeInt32(constant.AddFlightReq)[0])
	serMsgId := serializeStr(msgId)
	serMsgIdSize := serializeInt32(len(serMsgId))
	buf = append(buf, serMsgIdSize...)
	buf = append(buf, serMsgId...)
	buf = append(buf, serializeInt32(flightNo)...)

	serSource := serializeStr(source)
	serSourceSize := serializeInt32(len(serSource))
	serDest := serializeStr(dest)
	serDestSize := serializeInt32(len(serDest))
	buf = append(buf, serSourceSize...)
	buf = append(buf, serSource...)
	buf = append(buf, serDestSize...)
	buf = append(buf, serDest...)
	buf = append(buf, serializeInt32(flightTime)...)
	buf = append(buf, serializeFloat32(airFare)...)
	buf = append(buf, serializeInt32(totalSeatCount)...)
	buf = append(buf, serializeInt32(currentSeatCount)...)
	return append(serializeInt32(len(buf)+4), buf...)
}

func NewSerializeAddFlightResp(ack string) []byte {
	buf := make([]byte, 0)
	buf = append(buf, serializeInt32(constant.AddFlightResp)[0])
	serAck := serializeStr(ack)
	serAckSize := serializeInt32(len(serAck))
	buf = append(buf, serAckSize...)
	buf = append(buf, serAck...)

	return append(serializeInt32(len(buf)+4), buf...)
}

func NewSerializeMakeReservationReq(msgId string, flightNo int, seatCnt int) []byte {
	buf := make([]byte, 0)
	buf = append(buf, serializeInt32(constant.MakeReservationReq)[0])
	buf = append(buf, serializeReservationReq(msgId, flightNo, seatCnt)...)

	return append(serializeInt32(len(buf)+4), buf...)
}

func NewSerializeMakeReservationResp(ack string) []byte {
	buf := make([]byte, 0)
	buf = append(buf, serializeInt32(constant.MakeReservationResp)[0])
	buf = append(buf, serializeSimpleAckResp(ack)...)

	return append(serializeInt32(len(buf)+4), buf...)
}

func NewSerializeCancelReservationReq(msgId string, flightNo int, seatCnt int) []byte {
	buf := make([]byte, 0)
	buf = append(buf, serializeInt32(constant.CancelReservationReq)[0])
	buf = append(buf, serializeReservationReq(msgId, flightNo, seatCnt)...)

	return append(serializeInt32(len(buf)+4), buf...)
}

func NewSerializeCancelReservationResp(ack string) []byte {
	buf := make([]byte, 0)
	buf = append(buf, serializeInt32(constant.CancelReservationResp)[0])
	buf = append(buf, serializeSimpleAckResp(ack)...)

	return append(serializeInt32(len(buf)+4), buf...)
}

func NewSerializeGetReservationForFlightReq(msgId string, flightNo int) []byte {
	buf := make([]byte, 0)
	buf = append(buf, serializeInt32(constant.GetReservationForFlightReq)[0])
	serMsgId := serializeStr(msgId)
	serMsgIdSize := serializeInt32(len(serMsgId))
	buf = append(buf, serMsgIdSize...)
	buf = append(buf, serMsgId...)
	buf = append(buf, serializeInt32(flightNo)...)

	return append(serializeInt32(len(buf)+4), buf...)
}

func NewSerializeGetReservationForFlightResp(flightNo int, seatCnt int) []byte {
	buf := make([]byte, 0)
	buf = append(buf, serializeInt32(constant.GetReservationForFlightResp)[0])
	buf = append(buf, serializeInt32(flightNo)...)
	buf = append(buf, serializeInt32(seatCnt)...)

	return append(serializeInt32(len(buf)+4), buf...)
}

func NewSerializeRegisterForMonitorReq(msgId string, flightNo int, monitorIntervalSec int) []byte {
	buf := make([]byte, 0)
	buf = append(buf, serializeInt32(constant.RegisterForMonitorReq)[0])
	serMsgId := serializeStr(msgId)
	serMsgIdSize := serializeInt32(len(serMsgId))
	buf = append(buf, serMsgIdSize...)
	buf = append(buf, serMsgId...)
	buf = append(buf, serializeInt32(flightNo)...)
	buf = append(buf, serializeInt32(monitorIntervalSec)...)

	return append(serializeInt32(len(buf)+4), buf...)
}

func NewSerializeMonitorUpdateResp(flightNo int, availableSeats int) []byte {
	buf := make([]byte, 0)
	buf = append(buf, serializeInt32(constant.MonitorUpdateResp)[0])
	buf = append(buf, serializeInt32(flightNo)...)
	buf = append(buf, serializeInt32(availableSeats)...)

	return append(serializeInt32(len(buf)+4), buf...)
}

func NewSerializeMonitorBackoffResp(ack string) []byte {
	buf := make([]byte, 0)
	buf = append(buf, serializeInt32(constant.MonitorBackoffResp)[0])
	buf = append(buf, serializeSimpleAckResp(ack)...)

	return append(serializeInt32(len(buf)+4), buf...)
}

func NewSerializeGeneralErrResp(errorMsg string) []byte {
	buf := make([]byte, 0)
	buf = append(buf, serializeInt32(constant.GeneralErrResp)[0])
	buf = append(buf, serializeSimpleAckResp(errorMsg)...)

	return append(serializeInt32(len(buf)+4), buf...)
}

func serializeReservationReq(msgId string, flightNo int, seatCnt int) []byte {
	buf := make([]byte, 0)
	serMsgId := serializeStr(msgId)
	serMsgIdSize := serializeInt32(len(serMsgId))
	buf = append(buf, serMsgIdSize...)
	buf = append(buf, serMsgId...)
	buf = append(buf, serializeInt32(flightNo)...)
	buf = append(buf, serializeInt32(seatCnt)...)

	return buf
}

func serializeSimpleAckResp(ack string) []byte {
	serAck := serializeStr(ack)
	serAckSize := serializeInt32(len(serAck))
	buf := make([]byte, 0)
	buf = append(buf, serAckSize...)
	buf = append(buf, serAck...)

	return buf
}

var deserFuncMapping = map[int]func(b []byte) map[string]interface{}{
	constant.QueryFlightsReq:             deserQueryFlightReq,
	constant.QueryFlightsResp:            deserQueryFlightResp,
	constant.QueryFlightDetailReq:        deserQueryFlightDetailReq,
	constant.QueryFlightDetailResp:       deserQueryFlightDetailResp,
	constant.AddFlightReq:                deserAddFlightReq,
	constant.AddFlightResp:               deserAddFlightResp,
	constant.MakeReservationReq:          deserReservationReq,
	constant.MakeReservationResp:         deserReservationResp,
	constant.CancelReservationReq:        deserReservationReq,
	constant.CancelReservationResp:       deserReservationResp,
	constant.GetReservationForFlightReq:  deserGetReservationForFlightReq,
	constant.GetReservationForFlightResp: deserGetReservationForFlightResp,
	constant.RegisterForMonitorReq:       deserRegisterForMonitorReq,
	constant.MonitorUpdateResp:           deserMonitorUpdateResp,
	constant.MonitorBackoffResp:          deserMonitorBackOffResp,
	constant.GeneralErrResp:              deserGeneralErrorResp,
}

func Deserialize(b []byte) map[string]interface{} {
	msgType := deserializeInt32(b[4:5])
	f, exist := deserFuncMapping[msgType]
	if !exist {
		panic(fmt.Sprintf("%d err: not impl", msgType))
	}

	return f(b)
}

func serializeStr(s string) []byte {
	size := len(s)
	buf := []byte(s)
	for len(buf)%4 != 0 {
		buf = append(buf, byte('_'))
	}

	return append(serializeInt32(size), buf...)
}

func serializeInt32(x int) []byte {
	if x > math.MaxInt32 {
		panic("x is larger than max int32")
	}
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(x))

	return bs
}

func deserializeInt32(b []byte) int {
	t := make([]byte, len(b))
	copy(t, b)
	for len(t) < 4 {
		t = append(t, 0)
	}
	return int(binary.LittleEndian.Uint32(t))
}

func serializeFloat32(x float32) []byte {
	var buf [4]byte
	binary.LittleEndian.PutUint32(buf[:], math.Float32bits(x))
	return buf[:]
}

func deserializeFloat32(b []byte) float32 {
	return math.Float32frombits(binary.LittleEndian.Uint32(b))
}

func deserQueryFlightReq(b []byte) map[string]interface{} {
	x := 5
	_, x, msgIdB := extract(b, x)
	_, x, sourceB := extract(b, x)
	_, _, destB := extract(b, x)
	return map[string]interface{}{
		constant.MsgType:     constant.QueryFlightsReq,
		constant.MessageId:   string(msgIdB),
		constant.Source:      string(sourceB),
		constant.Destination: string(destB),
	}
}

func deserQueryFlightResp(b []byte) map[string]interface{} {
	flightNos := deserializeInt32(b[5:9])
	flightIds := extractInt32Arr(b, 9, flightNos)
	return map[string]interface{}{
		constant.MsgType:   constant.QueryFlightsResp,
		constant.FlightNos: flightIds,
	}
}

func deserQueryFlightDetailReq(b []byte) map[string]interface{} {
	x := 5
	_, x, msgIdB := extract(b, x)
	return map[string]interface{}{
		constant.MsgType:   constant.QueryFlightDetailReq,
		constant.MessageId: string(msgIdB),
		constant.FlightNo:  deserializeInt32(b[x : x+4]),
	}
}

func deserQueryFlightDetailResp(b []byte) map[string]interface{} {
	x := 9
	_, x, sourceB := extract(b, x)
	_, x, destB := extract(b, x)
	availSeats := deserializeInt32(b[x : x+4])
	x += 4
	flightTime := deserializeInt32(b[x : x+4])
	x += 4
	airFare := deserializeFloat32(b[x : x+4])
	return map[string]interface{}{
		constant.MsgType:        constant.QueryFlightDetailResp,
		constant.FlightNo:       deserializeInt32(b[5:9]),
		constant.Source:         string(sourceB),
		constant.Destination:    string(destB),
		constant.AvailableSeats: availSeats,
		constant.FlightTime:     flightTime,
		constant.AirFare:        airFare,
	}
}

func deserAddFlightReq(b []byte) map[string]interface{} {
	x := 5
	_, x, msgIdB := extract(b, x)
	flightNo := deserializeInt32(b[x : x+4])
	x += 4
	_, x, sourceB := extract(b, x)
	_, x, destB := extract(b, x)
	flightTime := deserializeInt32(b[x : x+4])
	x += 4
	airFare := deserializeFloat32(b[x : x+4])
	x += 4
	ttlSeatCnt := deserializeInt32(b[x : x+4])
	x += 4
	curSeatCnt := deserializeInt32(b[x : x+4])
	return map[string]interface{}{
		constant.MsgType:      constant.AddFlightReq,
		constant.MessageId:    string(msgIdB),
		constant.FlightNo:     flightNo,
		constant.Source:       string(sourceB),
		constant.Destination:  string(destB),
		constant.FlightTime:   flightTime,
		constant.AirFare:      airFare,
		constant.TotalSeats:   ttlSeatCnt,
		constant.CurrentSeats: curSeatCnt,
	}
}

func deserAddFlightResp(b []byte) map[string]interface{} {
	_, _, ackB := extract(b, 5)
	return map[string]interface{}{
		constant.MsgType: constant.AddFlightResp,
		constant.Ack:     string(ackB),
	}
}

func deserReservationReq(b []byte) map[string]interface{} {
	x := 5
	_, x, msgIdB := extract(b, x)
	flightNo := deserializeInt32(b[x : x+4])
	x += 4
	seatCnt := deserializeInt32(b[x : x+4])
	return map[string]interface{}{
		constant.MsgType:   deserializeInt32(b[4:5]),
		constant.MessageId: string(msgIdB),
		constant.FlightNo:  flightNo,
		constant.SeatCnt:   seatCnt,
	}
}

func deserReservationResp(b []byte) map[string]interface{} {
	_, _, ackB := extract(b, 5)
	return map[string]interface{}{
		constant.MsgType: deserializeInt32(b[4:5]),
		constant.Ack:     string(ackB),
	}
}

func deserGetReservationForFlightReq(b []byte) map[string]interface{} {
	x := 5
	_, x, msgIdB := extract(b, x)
	flightNo := deserializeInt32(b[x : x+4])
	return map[string]interface{}{
		constant.MsgType:   deserializeInt32(b[4:5]),
		constant.MessageId: string(msgIdB),
		constant.FlightNo:  flightNo,
	}
}

func deserGetReservationForFlightResp(b []byte) map[string]interface{} {
	return map[string]interface{}{
		constant.MsgType:  deserializeInt32(b[4:5]),
		constant.FlightNo: deserializeInt32(b[5:9]),
		constant.SeatCnt:  deserializeInt32(b[9:13]),
	}
}

func deserRegisterForMonitorReq(b []byte) map[string]interface{} {
	x := 5
	_, x, msgIdB := extract(b, x)
	flightNo := deserializeInt32(b[x : x+4])
	x += 4
	monitorIntervalSec := deserializeInt32(b[x : x+4])
	return map[string]interface{}{
		constant.MsgType:            deserializeInt32(b[4:5]),
		constant.MessageId:          string(msgIdB),
		constant.FlightNo:           flightNo,
		constant.MonitorIntervalSec: monitorIntervalSec,
	}
}

func deserMonitorUpdateResp(b []byte) map[string]interface{} {
	return map[string]interface{}{
		constant.MsgType:        deserializeInt32(b[4:5]),
		constant.FlightNo:       deserializeInt32(b[5:9]),
		constant.AvailableSeats: deserializeInt32(b[9:13]),
	}
}

func deserMonitorBackOffResp(b []byte) map[string]interface{} {
	_, _, ackB := extract(b, 5)
	return map[string]interface{}{
		constant.MsgType: deserializeInt32(b[4:5]),
		constant.Ack:     string(ackB),
	}
}

func deserGeneralErrorResp(b []byte) map[string]interface{} {
	_, _, errMsg := extract(b, 5)
	return map[string]interface{}{
		constant.MsgType:    deserializeInt32(b[4:5]),
		constant.ErrMessage: string(errMsg),
	}
}

func extract(b []byte, ptr int) (int, int, []byte) {
	size := deserializeInt32(b[ptr : ptr+4])
	ptr += 4
	finalPtr := ptr + size
	strSize := deserializeInt32(b[ptr : ptr+4])
	ptr += 4
	data := b[ptr : ptr+strSize]

	return size, finalPtr, data
}

func extractInt32Arr(b []byte, ptr int, n int) []int {
	res := make([]int, 0)
	for i := 0; i < n; i++ {
		res = append(res, deserializeInt32(b[ptr+i*4:ptr+(i+1)*4]))
	}

	return res
}
