package atype

import (
	"database/sql"
	"encoding/json"
)

type NullJson struct{ sql.NullString }          // any
type NullUint8s struct{ sql.NullString }        // uint8 json array
type NullUint16s struct{ sql.NullString }       // uint16 json array
type NullUint24s struct{ sql.NullString }       // Uint24 json array
type NullUint32s struct{ sql.NullString }       // uint32 json array
type NullInts struct{ sql.NullString }          // int json array
type NullUints struct{ sql.NullString }         // uint json array
type NullUint64s struct{ sql.NullString }       // uint64 json array
type NullStrings struct{ sql.NullString }       // string json array
type StringMap struct{ sql.NullString }         // map[string]string   // JSON 规范，key 必须为字符串
type StringMapsMap struct{ sql.NullString }     // map[string][]map[string]string
type StringsMap struct{ sql.NullString }        // map[string][]string
type ComplexStringMap struct{ sql.NullString }  // map[string]map[string]string
type ComplexStringsMap struct{ sql.NullString } // map[string][][]string
type StringMaps struct{ sql.NullString }        // []map[string]string
type ComplexStringMaps struct{ sql.NullString } // []map[string][]map[string]string

type ComplexMaps struct{ sql.NullString } // []map[string]any

func NewNullJson(s []byte) NullJson {
	var x NullJson
	if len(s) > 0 {
		x.Scan(string(s))
	}
	return x
}

func ToNullJson(v any) NullJson {
	if v == nil {
		return NullJson{}
	}
	s, _ := json.Marshal(v)
	return NewNullJson(s)
}

func (t ComplexMaps) Interface() any {
	if !t.Valid || t.String == "" {
		return nil
	}
	var v any
	json.Unmarshal([]byte(t.String), &v)
	return v
}

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
	s, _ := MarshalUint8s(v)
	if len(s) == 0 {
		return NullUint8s{}
	}

	return NewNullUint8s(string(s))
}

func (t NullUint8s) Uint8s() []uint8 {
	if !t.Valid || t.String == "" {
		return nil
	}
	w, _ := UnmarshalUint8s([]byte(t.String))
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
	if !t.Valid || t.String == "" {
		return nil
	}
	var v []any
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
	if !t.Valid || t.String == "" {
		return nil
	}
	var v []any
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
	if !t.Valid || t.String == "" {
		return nil
	}
	var v []any
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
	if !t.Valid || t.String == "" {
		return nil
	}
	var v []any
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
	if !t.Valid || t.String == "" {
		return nil
	}
	var v []any
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
	if !t.Valid || t.String == "" {
		return nil
	}
	var v []any
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
	if !t.Valid || t.String == "" {
		return nil
	}
	var v []string
	json.Unmarshal([]byte(t.String), &v)
	return v
}
func NewStringMap(s string) StringMap {
	var x StringMap
	if s != "" {
		x.Scan(s)
	}
	return x
}
func ToStringMap(v map[string]string) StringMap {
	if len(v) == 0 {
		return StringMap{}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return StringMap{}
	}

	return NewStringMap(string(s))
}

func (t StringMap) StringMap() map[string]string {
	if !t.Valid || t.String == "" {
		return nil
	}
	var v map[string]string
	json.Unmarshal([]byte(t.String), &v)
	return v
}
func NewComplexStringMap(s string) ComplexStringMap {
	var x ComplexStringMap
	if s != "" {
		x.Scan(s)
	}
	return x
}
func ToComplexStringMap(v map[string]map[string]string) ComplexStringMap {
	if len(v) == 0 {
		return ComplexStringMap{}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return ComplexStringMap{}
	}

	return NewComplexStringMap(string(s))
}
func (t ComplexStringMap) TStringMap() map[string]map[string]string {
	if !t.Valid || t.String == "" {
		return nil
	}
	var v map[string]map[string]string
	json.Unmarshal([]byte(t.String), &v)
	return v
}
func NewStringMaps(s string) StringMaps {
	var x StringMaps
	if s != "" {
		x.Scan(s)
	}
	return x
}
func ToStringMaps(v []map[string]string) StringMaps {
	if len(v) == 0 {
		return StringMaps{}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return StringMaps{}
	}

	return NewStringMaps(string(s))
}

func (t StringMaps) StringMaps() []map[string]string {
	if !t.Valid || t.String == "" {
		return nil
	}
	var v []map[string]string
	json.Unmarshal([]byte(t.String), &v)
	return v
}
func NewStringMapsMap(s string) StringMapsMap {
	var x StringMapsMap
	if s != "" {
		x.Scan(s)
	}
	return x
}
func ToStringMapsMap(v map[string][]map[string]string) StringMapsMap {
	if len(v) == 0 {
		return StringMapsMap{}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return StringMapsMap{}
	}

	return NewStringMapsMap(string(s))
}
func (t StringMapsMap) StringMapsMap() map[string][]map[string]string {
	if !t.Valid || t.String == "" {
		return nil
	}
	var v map[string][]map[string]string
	json.Unmarshal([]byte(t.String), &v)
	return v
}
func NewStringsMap(s string) StringsMap {
	var x StringsMap
	if s != "" {
		x.Scan(s)
	}
	return x
}
func ToStringsMap(v map[string][]string) StringsMap {
	if len(v) == 0 {
		return StringsMap{}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return StringsMap{}
	}

	return NewStringsMap(string(s))
}

func (t StringsMap) StringsMap() map[string][]string {
	if !t.Valid || t.String == "" {
		return nil
	}
	var v map[string][]string
	json.Unmarshal([]byte(t.String), &v)
	return v
}
func NewComplexStringsMap(s string) ComplexStringsMap {
	var x ComplexStringsMap
	if s != "" {
		x.Scan(s)
	}
	return x
}
func ToComplexStringsMap(v map[string][][]string) ComplexStringsMap {
	if len(v) == 0 {
		return ComplexStringsMap{}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return ComplexStringsMap{}
	}

	return NewComplexStringsMap(string(s))
}

func (t ComplexStringsMap) StringsMap() map[string][][]string {
	if !t.Valid || t.String == "" {
		return nil
	}
	var v map[string][][]string
	json.Unmarshal([]byte(t.String), &v)
	return v
}

func NewComplexStringMaps(s string) ComplexStringMaps {
	var x ComplexStringMaps
	if s != "" {
		x.Scan(s)
	}
	return x
}
func ToComplexStringMaps(v []map[string][]map[string]string) ComplexStringMaps {
	if len(v) == 0 {
		return ComplexStringMaps{}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return ComplexStringMaps{}
	}

	return NewComplexStringMaps(string(s))
}

func (t ComplexStringMaps) StringMaps() []map[string][]map[string]string {
	if !t.Valid || t.String == "" {
		return nil
	}
	var v []map[string][]map[string]string
	json.Unmarshal([]byte(t.String), &v)
	return v
}

func NewComplexMaps(s string) ComplexMaps {
	var x ComplexMaps
	if s != "" {
		x.Scan(s)
	}
	return x
}
func ToComplexMaps(v []map[string]any) ComplexMaps {
	if len(v) == 0 {
		return ComplexMaps{}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return ComplexMaps{}
	}

	return NewComplexMaps(string(s))
}

func (t ComplexMaps) Maps() []map[string]any {
	if !t.Valid || t.String == "" {
		return nil
	}
	var v []map[string]any
	json.Unmarshal([]byte(t.String), &v)
	return v
}
