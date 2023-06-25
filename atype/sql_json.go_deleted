package atype

import (
	"bytes"
	"encoding/json"
	"errors"
	"strconv"
)

type JsonBytes json.RawMessage
type JsonUint8s json.RawMessage
type JsonUint16s json.RawMessage
type JsonUint24s json.RawMessage
type JsonUint32s json.RawMessage
type JsonUints json.RawMessage
type JsonUint64s json.RawMessage

type JsonInt8s json.RawMessage
type JsonInt16s json.RawMessage
type JsonInt24s json.RawMessage
type JsonInt32s json.RawMessage
type JsonInts json.RawMessage
type JsonInt64s json.RawMessage

type JsonFloat64s json.RawMessage
type JsonFloat32s json.RawMessage

type JsonStrings json.RawMessage
type JsonStringMap json.RawMessage     // map[string]string
type Json2dStringMap json.RawMessage   // map[string]map[string]string
type JsonStringMaps json.RawMessage    // []map[string]string
type JsonStringMapsMap json.RawMessage // map[string][]map[string]string

/*
"Array and slice values encode as JSON arrays, except that []byte encodes as a base64-encoded string, and a nil slice encodes as the null JSON object.
json.Marshal() 不能正常转换 []byte 及 []uint8
*/
func MarshalBytes(x []byte) ([]byte, error) {
	if x == nil {
		return nil, nil
	}
	var b bytes.Buffer
	b.Grow(2 + len(x)*2 - 1)
	b.WriteByte('[')
	for i, c := range x {
		if i > 0 {
			b.WriteByte(',')
		}
		b.WriteByte('\'')
		if c == '\'' || c == '\\' {
			b.WriteByte('\\')
		}
		b.WriteByte(c)
		b.WriteByte('\'')
	}
	b.WriteByte(']')
	return b.Bytes(), nil
}

// x =>  `['a','\'', ',']
func UnmarshalBytes(x []byte) ([]byte, error) {
	if x == nil || len(x) < 2 {
		return nil, nil
	}
	n := len(x) - 1 // remove last ']'
	v := make([]byte, 0)
	for i := 1; i < n; {
		for x[i] == ' ' || x[i] == ',' {
			i++
		}
		// start with ',  and can not be ''
		if x[i] != '\'' || i >= n-1 || x[i+1] == '\'' {
			return nil, errors.New("invalid bytes json: " + string(x))
		}
		i++
		if x[i] == '\\' {
			i++
		}
		if i >= n-1 || x[i+1] != '\'' {
			return nil, errors.New("invalid bytes json: " + string(x))
		}
		v = append(v, x[i])
		i += 2
	}
	return v, nil
}
func MarshalUint8s(x []uint8) ([]byte, error) {
	if x == nil {
		return nil, nil
	}
	var b bytes.Buffer
	b.WriteByte('[')
	for i, c := range x {
		if i > 0 {
			b.WriteByte(',')
		}
		b.WriteString(strconv.FormatUint(uint64(c), 10))
	}
	b.WriteByte(']')
	return b.Bytes(), nil
}
func UnmarshalUint8s(x []byte) ([]uint8, error) {
	if x == nil || len(x) < 2 {
		return nil, nil
	}

	n := len(x) - 1 // remove last ']'
	v := make([]uint8, 0)
	for i := 1; i < n; {
		for x[i] == ' ' || x[i] == ',' {
			i++
		}
		if x[i] < '0' || x[i] > '9' {
			return nil, errors.New("invalid uint8 json: " + string(x))
		}
		var c string
		for x[i] >= '0' && x[i] <= '9' {
			c += string(x[i])
			i++
		}
		u, _ := strconv.ParseUint(c, 10, 8)
		v = append(v, uint8(u))
	}
	return v, nil
}
func ToJsonUint8s(x []uint8) JsonUint8s {
	b, _ := MarshalUint8s(x)
	return b
}
func (t JsonUint8s) Uint8s() []uint8 {
	v, _ := UnmarshalUint8s(t)
	return v
}
func ToJsonUint16s(x []uint16) JsonUint16s {
	b, _ := json.Marshal(x)
	return b
}
func (t JsonUint16s) Uint16s() []uint16 {
	var v []uint16
	json.Unmarshal(t, &v)
	return v
}
func ToJsonUint24s(x []Uint24) JsonUint24s {
	b, _ := json.Marshal(x)
	return b
}
func (t JsonUint24s) Uint24s() []Uint24 {
	var v []Uint24
	json.Unmarshal(t, &v)
	return v
}
func ToJsonUint32s(x []uint32) JsonUint32s {
	b, _ := json.Marshal(x)
	return b
}
func (t JsonUint32s) Uint32s() []uint32 {
	var v []uint32
	json.Unmarshal(t, &v)
	return v
}
func ToJsonUints(x []uint) JsonUints {
	b, _ := json.Marshal(x)
	return b
}
func (t JsonUints) Uints() []uint {
	var v []uint
	json.Unmarshal(t, &v)
	return v
}
func ToJsonUint64s(x []uint64) JsonUint64s {
	b, _ := json.Marshal(x)
	return b
}
func (t JsonUint64s) Uint64s() []uint64 {
	var v []uint64
	json.Unmarshal(t, &v)
	return v
}
func ToJsonInts(x []int) JsonInts {
	b, _ := json.Marshal(x)
	return b
}
func (t JsonInts) Ints() []int {
	var v []int
	json.Unmarshal(t, &v)
	return v
}
func ToJsonStrings(x []string) JsonStrings {
	b, _ := json.Marshal(x)
	return b
}
func (t JsonStrings) Strings() []string {
	var v []string
	json.Unmarshal(t, &v)
	return v
}
func ToJsonStringMap(x map[string]string) JsonStringMap {
	b, _ := json.Marshal(x)
	return b
}
func (t JsonStringMap) StringMap() map[string]string {
	var v map[string]string
	json.Unmarshal(t, &v)
	return v
}
func ToJson2dStringMap(x map[string]map[string]string) Json2dStringMap {
	b, _ := json.Marshal(x)
	return b
}
func (t Json2dStringMap) TStringMap() map[string]map[string]string {
	var v map[string]map[string]string
	json.Unmarshal(t, &v)
	return v
}
func ToJsonStringMaps(x []map[string]string) JsonStringMaps {
	b, _ := json.Marshal(x)
	return b
}
func (t JsonStringMaps) StringMaps() []map[string]string {
	var v []map[string]string
	json.Unmarshal(t, &v)
	return v
}
func ToJsonStringMapsMap(x map[string][]map[string]string) JsonStringMapsMap {
	b, _ := json.Marshal(x)
	return b
}
func (t JsonStringMapsMap) StringMapsMap() map[string][]map[string]string {
	var v map[string][]map[string]string
	json.Unmarshal(t, &v)
	return v
}
