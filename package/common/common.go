package common

import (
	"encoding/binary"
	"fmt"
	"github.com/YangXingHao96/DistributedSystem/package/common/constant"
	"math"
	"strings"
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

func NewSerializeQueryFlightDetailResp(flightNo int, source string, dest string, availSeats int) []byte {
	serSource := serializeStr(source)
	serSourceSize := serializeInt32(len(serSource))
	serDest := serializeStr(dest)
	serDestSize := serializeInt32(len(serDest))

	buf := make([]byte, 0)
	buf = append(buf, serializeInt32(constant.QueryFlightDetailReq)[0])
	buf = append(buf, serializeInt32(flightNo)...)
	buf = append(buf, serSourceSize...)
	buf = append(buf, serSource...)
	buf = append(buf, serDestSize...)
	buf = append(buf, serDest...)
	buf = append(buf, serializeInt32(availSeats)...)

	return append(serializeInt32(len(buf)+4), buf...)
}

var deserFuncMapping = map[int]func(b []byte) map[string]interface{} {
	constant.QueryFlightsReq: deserQueryFlightReq,
	constant.QueryFlightsResp: deserQueryFlightResp,
	constant.QueryFlightDetailReq: deserQueryFlightDetailReq,
	constant.QueryFlightDetailResp: deserQueryFlightDetailResp,
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
	buf := []byte(s)
	for len(buf) % 4 != 0 {
		buf = append(buf, byte('_'))
	}

	return buf
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

func deserQueryFlightReq(b []byte) map[string]interface{} {
	x := 5
	_, x, msgIdB := extract(b, x)
	_, x, sourceB := extract(b, x)
	_, _, destB := extract(b, x)
	return map[string]interface{}{
		constant.MsgType: constant.QueryFlightsReq,
		constant.MessageId: strings.TrimRight(string(msgIdB), "_"),
		constant.Source: strings.TrimRight(string(sourceB), "_"),
		constant.Destination: strings.TrimRight(string(destB), "_"),
	}
}

func deserQueryFlightResp(b []byte) map[string]interface{} {
	flightNos := deserializeInt32(b[5:9])
	flightIds := extractInt32Arr(b, 9, flightNos)
	return map[string]interface{}{
		constant.MsgType: constant.QueryFlightsResp,
		constant.FlightNos: flightIds,
	}
}

func deserQueryFlightDetailReq(b []byte) map[string]interface{} {
	x := 5
	_, x, msgIdB := extract(b, x)
	return map[string]interface{}{
		constant.MsgType: constant.QueryFlightDetailReq,
		constant.MessageId: strings.TrimRight(string(msgIdB), "_"),
		constant.FlightNo: deserializeInt32(b[x:x+4]),
	}
}

func deserQueryFlightDetailResp(b []byte) map[string]interface{} {
	x := 9
	_, x, sourceB := extract(b, x)
	_, x, destB := extract(b, x)
	return map[string]interface{}{
		constant.MsgType: constant.QueryFlightDetailResp,
		constant.FlightNo: deserializeInt32(b[5:9]),
		constant.Source: strings.TrimRight(string(sourceB), "_"),
		constant.Destination: strings.TrimRight(string(destB), "_"),
		constant.AvailableSeats: deserializeInt32(b[x:x+4]),
	}
}

func extract(b []byte, ptr int) (int, int, []byte) {
	size := deserializeInt32(b[ptr:ptr+4])
	ptr += 4
	data := b[ptr:ptr+size]
	ptr += size

	return size, ptr, data
}

func extractInt32Arr(b []byte, ptr int, n int) []int {
	res := make([]int, 0)
	for i := 0; i<n; i++ {
		res = append(res, deserializeInt32(b[ptr+i*4:ptr+(i+1)*4]))
	}

	return res
}
