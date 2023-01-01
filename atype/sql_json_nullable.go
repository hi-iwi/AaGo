package atype

import (
	"database/sql"
	"encoding/json"
)

type NullJson struct{ sql.NullString }
type NullUint8s struct{ sql.NullString }        // uint8 json array
type NullUint16s struct{ sql.NullString }       // uint16 json array
type NullUint24s struct{ sql.NullString }       // Uint24 json array
type NullUint32s struct{ sql.NullString }       // uint32 json array
type NullInts struct{ sql.NullString }          // int json array
type NullUints struct{ sql.NullString }         // uint json array
type NullUint64s struct{ sql.NullString }       // uint64 json array
type NullStrings struct{ sql.NullString }       // string json array
type NullStringMap struct{ sql.NullString }     // map[string]string
type Null2dStringMap struct{ sql.NullString }   // map[string]map[string]string
type NullStringMaps struct{ sql.NullString }    // []map[string]string
type NullStringMapsMap struct{ sql.NullString } // map[string][]map[string]string
type NullStringsMap struct{ sql.NullString }    // map[string][]string

func NewNullUint8s(s string) NullUint8s {
	var x NullUint8s
	if s != "" {
		x.Scan(s)
	}
	return x
}
func ToNullUint8s(v []uint8) NullUint8s {
	if len(v) == 0 {
		return NullUint8s{}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return NullUint8s{}
	}

	return NewNullUint8s(string(s))
}

func (t NullUint8s) Uint8s() []uint8 {
	if t.String == "" {
		return nil
	}
	var v []interface{}
	json.Unmarshal([]byte(t.String), &v)
	if len(v) == 0 {
		return nil
	}
	w := make([]uint8, len(v))
	for i, x := range v {
		w[i] = New(x).DefaultUint8(0)
	}
	return w
}
func NewNullUint16s(s string) NullUint16s {
	var x NullUint16s
	if s != "" {
		x.Scan(s)
	}
	return x
}
func ToNullUint16s(v []uint16) NullUint16s {
	if len(v) == 0 {
		return NullUint16s{}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return NullUint16s{}
	}

	return NewNullUint16s(string(s))
}

func (t NullUint16s) Uint16s() []uint16 {
	if t.String == "" {
		return nil
	}
	var v []interface{}
	json.Unmarshal([]byte(t.String), &v)
	if len(v) == 0 {
		return nil
	}
	w := make([]uint16, len(v))
	for i, x := range v {
		w[i] = New(x).DefaultUint16(0)
	}
	return w
}
func NewNullUint24s(s string) NullUint24s {
	var x NullUint24s
	if s != "" {
		x.Scan(s)
	}
	return x
}
func ToNullUint24s(v []uint32) NullUint24s {
	if len(v) == 0 {
		return NullUint24s{}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return NullUint24s{}
	}
	return NewNullUint24s(string(s))
}

func (t NullUint24s) Uint24s() []Uint24 {
	if t.String == "" {
		return nil
	}
	var v []interface{}
	json.Unmarshal([]byte(t.String), &v)
	if len(v) == 0 {
		return nil
	}
	w := make([]Uint24, len(v))
	for i, x := range v {
		w[i] = New(x).DefaultUint24(0)
	}
	return w
}
func NewNullUint32s(s string) NullUint32s {
	var x NullUint32s
	if s != "" {
		x.Scan(s)
	}
	return x
}
func ToNullUint32s(v []uint32) NullUint32s {
	if len(v) == 0 {
		return NullUint32s{}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return NullUint32s{}
	}
	return NewNullUint32s(string(s))
}

func (t NullUint32s) Uint32s() []uint32 {
	if t.String == "" {
		return nil
	}
	var v []interface{}
	json.Unmarshal([]byte(t.String), &v)
	if len(v) == 0 {
		return nil
	}
	w := make([]uint32, len(v))
	for i, x := range v {
		w[i] = New(x).DefaultUint32(0)
	}
	return w
}
func NewNullInts(s string) NullInts {
	var x NullInts
	if s != "" {
		x.Scan(s)
	}
	return x
}
func ToNullInts(v []int) NullInts {
	if len(v) == 0 {
		return NullInts{}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return NullInts{}
	}
	return NewNullInts(string(s))
}

func (t NullInts) Ints() []int {
	if t.String == "" {
		return nil
	}
	var v []interface{}
	json.Unmarshal([]byte(t.String), &v)
	if len(v) == 0 {
		return nil
	}
	w := make([]int, len(v))
	for i, x := range v {
		w[i] = New(x).DefaultInt(0)
	}
	return w
}
func NewNullUints(s string) NullUints {
	var x NullUints
	if s != "" {
		x.Scan(s)
	}
	return x
}
func ToNullUints(v []uint) NullUints {
	if len(v) == 0 {
		return NullUints{}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return NullUints{}
	}

	return NewNullUints(string(s))
}
func (t NullUints) Uints() []uint {
	if t.String == "" {
		return nil
	}
	var v []interface{}
	json.Unmarshal([]byte(t.String), &v)
	if len(v) == 0 {
		return nil
	}
	w := make([]uint, len(v))
	for i, x := range v {
		w[i] = New(x).DefaultUint(0)
	}
	return w
}

func NewNullUint64s(s string) NullUint64s {
	var x NullUint64s
	if s != "" {
		x.Scan(s)
	}
	return x
}
func ToNullUint64s(v []uint64) NullUint64s {
	if len(v) == 0 {
		return NullUint64s{}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return NullUint64s{}
	}

	return NewNullUint64s(string(s))
}

func (t NullUint64s) Uint64s() []uint64 {
	if t.String == "" {
		return nil
	}
	var v []interface{}
	json.Unmarshal([]byte(t.String), &v)
	if len(v) == 0 {
		return nil
	}
	w := make([]uint64, len(v))
	for i, x := range v {
		w[i] = New(x).DefaultUint64(0)
	}
	return w
}
func NewNullStrings(s string) NullStrings {
	var x NullStrings
	if s != "" {
		x.Scan(s)
	}
	return x
}
func ToNullStrings(v []string) NullStrings {
	if len(v) == 0 {
		return NullStrings{}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return NullStrings{}
	}

	return NewNullStrings(string(s))
}
func (t NullStrings) Strings() []string {
	if t.String == "" {
		return nil
	}
	var v []string
	json.Unmarshal([]byte(t.String), &v)
	return v
}
func NewNullStringMap(s string) NullStringMap {
	var x NullStringMap
	if s != "" {
		x.Scan(s)
	}
	return x
}
func ToNullStringMap(v map[string]string) NullStringMap {
	if len(v) == 0 {
		return NullStringMap{}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return NullStringMap{}
	}

	return NewNullStringMap(string(s))
}

func (t NullStringMap) StringMap() map[string]string {
	if t.String == "" {
		return nil
	}
	var v map[string]string
	json.Unmarshal([]byte(t.String), &v)
	return v
}
func NewNull2dStringMap(s string) Null2dStringMap {
	var x Null2dStringMap
	if s != "" {
		x.Scan(s)
	}
	return x
}
func ToNull2dStringMap(v map[string]map[string]string) Null2dStringMap {
	if len(v) == 0 {
		return Null2dStringMap{}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return Null2dStringMap{}
	}

	return NewNull2dStringMap(string(s))
}
func (t Null2dStringMap) TStringMap() map[string]map[string]string {
	if t.String == "" {
		return nil
	}
	var v map[string]map[string]string
	json.Unmarshal([]byte(t.String), &v)
	return v
}
func NewNullStringMaps(s string) NullStringMaps {
	var x NullStringMaps
	if s != "" {
		x.Scan(s)
	}
	return x
}
func ToNullStringMaps(v []map[string]string) NullStringMaps {
	if len(v) == 0 {
		return NullStringMaps{}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return NullStringMaps{}
	}

	return NewNullStringMaps(string(s))
}

func (t NullStringMaps) StringMaps() []map[string]string {
	if t.String == "" {
		return nil
	}
	var v []map[string]string
	json.Unmarshal([]byte(t.String), &v)
	return v
}
func NewNullStringMapsMap(s string) NullStringMapsMap {
	var x NullStringMapsMap
	if s != "" {
		x.Scan(s)
	}
	return x
}
func ToNullStringMapsMap(v map[string][]map[string]string) NullStringMapsMap {
	if len(v) == 0 {
		return NullStringMapsMap{}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return NullStringMapsMap{}
	}

	return NewNullStringMapsMap(string(s))
}
func (t NullStringMapsMap) StringMapsMap() map[string][]map[string]string {
	if t.String == "" {
		return nil
	}
	var v map[string][]map[string]string
	json.Unmarshal([]byte(t.String), &v)
	return v
}
func NewNullStringsMap(s string) NullStringsMap {
	var x NullStringsMap
	if s != "" {
		x.Scan(s)
	}
	return x
}
func ToNullStringsMap(v map[string][]string) NullStringsMap {
	if len(v) == 0 {
		return NullStringsMap{}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return NullStringsMap{}
	}

	return NewNullStringsMap(string(s))
}

func (t NullStringsMap) StringsMap() map[string][]string {
	if t.String == "" {
		return nil
	}
	var v map[string][]string
	json.Unmarshal([]byte(t.String), &v)
	return v
}
