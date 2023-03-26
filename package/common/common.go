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

	return append(serializeInt32(len(buf)), buf...)
}

func NewSerializeGetFlightIdBySourceDestResp(flightNo []int) []byte {
	buf := make([]byte, 0)
	buf = append(buf, serializeInt32(len(flightNo))...)
	buf = append(buf, serializeInt32(constant.QueryFlightsResp)[0])
	buf = append(buf, serializeInt32(len(flightNo))...)
	for _, f := range flightNo {
		buf = append(buf, serializeInt32(f)...)
	}

	return append(serializeInt32(len(buf)), buf...)
}

var deserFuncMapping = map[int]func(b []byte) map[string]interface{} {
	constant.QueryFlightsReq: deserQueryFlightReq,
	constant.QueryFlightsResp: deserQueryFlightResp,
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
	for len(b) < 4 {
		b = append(b, 0)
	}
	return int(binary.LittleEndian.Uint32(b))
}

func deserQueryFlightReq(b []byte) map[string]interface{} {
	x := 5
	_, x, msgIdB := extract(b, x)
	_, x, sourceB := extract(b, x)
	_, _, destB := extract(b, x)
	return map[string]interface{}{
		constant.MsgType: constant.QueryFlightsReq,
		constant.MessageId: string(msgIdB),
		constant.Source: string(sourceB),
		constant.Destination: string(destB),
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

func extract(b []byte, ptr int) (int, int, []byte) {
	size := deserializeInt32(b[ptr:ptr+4])
	ptr += size
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
