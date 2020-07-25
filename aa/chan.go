package aa

import "github.com/luexu/aenum"

type StringChanEle struct {
	Qos     aenum.Qos
	TraceId string
	Data    string
}

//type IntChanEle struct {
//	Qos     aenum.Qos
//	TraceId string
//	Data    int
//}
//
//type Uint64ChanEle struct {
//	Qos     aenum.Qos
//	TraceId string
//	Data    uint64
//}

func MakeStringChanEle(qos aenum.Qos, traceId string, data string) StringChanEle {
	return StringChanEle{Qos: qos, TraceId: traceId, Data: data}
}
