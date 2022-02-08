package dtype

import "encoding/json"

type JsonUint8s json.RawMessage
type JsonUint16s json.RawMessage
type JsonUint32s json.RawMessage
type JsonInts json.RawMessage
type JsonUints json.RawMessage
type JsonUint64s json.RawMessage
type JsonStrings json.RawMessage
type JsonStringMap json.RawMessage
type Json2dStringMap json.RawMessage

func (t JsonUint8s) Uint8s() []uint8 {
	var v []uint8
	json.Unmarshal(t, &v)
	return v
}
func (t JsonUint16s) Uint16s() []uint16 {
	var v []uint16
	json.Unmarshal(t, &v)
	return v
}
func (t JsonUint32s) Uint32s() []uint32 {
	var v []uint32
	json.Unmarshal(t, &v)
	return v
}
func (t JsonInts) Ints() []int {
	var v []int
	json.Unmarshal(t, &v)
	return v
}
func (t JsonUints) Uints() []uint {
	var v []uint
	json.Unmarshal(t, &v)
	return v
}
func (t JsonUint64s) Uint64s() []uint64 {
	var v []uint64
	json.Unmarshal(t, &v)
	return v
}
func (t JsonStrings) Strings() []string {
	var v []string
	json.Unmarshal(t, &v)
	return v
}

func (t JsonStringMap) StringMap() map[string]string {
	var v map[string]string
	json.Unmarshal(t, &v)
	return v
}

func (t Json2dStringMap) TStringMap() map[string]map[string]string {
	var v map[string]map[string]string
	json.Unmarshal(t, &v)
	return v
}
