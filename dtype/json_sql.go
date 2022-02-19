package dtype

import (
	"database/sql"
	"encoding/json"
	"github.com/hi-iwi/AaGo/adto"
	"strconv"
	"strings"
)

type ObjScan interface {
	Scan(value interface{}) error
}
type NullJson struct{ sql.NullString }
type NullUint8s struct{ sql.NullString }        // uint8 json array
type NullUint16s struct{ sql.NullString }       // uint16 json array
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

type NullSepStrings struct{ sql.NullString } // a,b,c,d,e
type NullSepUint8s struct{ sql.NullString }  // 1,2,3,4
type NullSepUint16s struct{ sql.NullString } // 1,2,3,4
type NullSepUint32s struct{ sql.NullString } // 1,2,3,4
type NullSepInts struct{ sql.NullString }    // 1,2,3,4
type NullSepUints struct{ sql.NullString }   // 1,2,3,4
type NullSepUint64s struct{ sql.NullString } // 1,2,3,4

type NullImgSrc struct{ sql.NullString }     // adto.ImgSrc  @warn 为了节省空间，这里使用数组方式存储
type NullImgSrcs struct{ sql.NullString }    // []adto.ImgSrc  @warn 为了节省空间，这里使用数组方式存储
type NullVideoSrc struct{ sql.NullString }   // adto.VideoSrc  @warn 为了节省空间，这里使用数组方式存储
type NullVideosSrcs struct{ sql.NullString } // []adto.VideoSrc  @warn 为了节省空间，这里使用数组方式存储

// 保证空字符串不能正常的对象
func ScanNullObject(obj ObjScan, data string) {
	if data == "" {
		obj.Scan(nil)
	}
	obj.Scan(data)
}
func ToNullUint8s(v []uint8) NullUint8s {
	if len(v) == 0 {
		return NullUint8s{}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return NullUint8s{}
	}
	var x NullUint8s
	ScanNullObject(&x, string(s))
	return x
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

func ToNullUint16s(v []uint16) NullUint16s {
	if len(v) == 0 {
		return NullUint16s{}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return NullUint16s{}
	}
	var x NullUint16s
	ScanNullObject(&x, string(s))
	return x
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

func ToNullUint32s(v []uint32) NullUint32s {
	if len(v) == 0 {
		return NullUint32s{}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return NullUint32s{}
	}
	var x NullUint32s
	ScanNullObject(&x, string(s))
	return x
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

func ToNullInts(v []int) NullInts {
	if len(v) == 0 {
		return NullInts{}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return NullInts{}
	}
	var x NullInts
	ScanNullObject(&x, string(s))
	return x
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

func ToNullUints(v []uint) NullUints {
	if len(v) == 0 {
		return NullUints{}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return NullUints{}
	}
	var x NullUints
	ScanNullObject(&x, string(s))
	return x
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

func ToNullUint64s(v []uint64) NullUint64s {
	if len(v) == 0 {
		return NullUint64s{}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return NullUint64s{}
	}
	var x NullUint64s
	ScanNullObject(&x, string(s))
	return x
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

func ToNullStrings(v []string) NullStrings {
	if len(v) == 0 {
		return NullStrings{}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return NullStrings{}
	}
	var x NullStrings
	ScanNullObject(&x, string(s))
	return x
}
func (t NullStrings) Strings() []string {
	if t.String == "" {
		return nil
	}
	var v []string
	json.Unmarshal([]byte(t.String), &v)
	return v
}

func ToNullNullStringMap(v map[string]string) NullStringMap {
	if len(v) == 0 {
		return NullStringMap{}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return NullStringMap{}
	}
	var x NullStringMap
	ScanNullObject(&x, string(s))
	return x
}

func (t NullStringMap) StringMap() map[string]string {
	if t.String == "" {
		return nil
	}
	var v map[string]string
	json.Unmarshal([]byte(t.String), &v)
	return v
}

func ToNull2dStringMap(v map[string]map[string]string) Null2dStringMap {
	if len(v) == 0 {
		return Null2dStringMap{}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return Null2dStringMap{}
	}
	var x Null2dStringMap
	ScanNullObject(&x, string(s))
	return x
}
func (t Null2dStringMap) TStringMap() map[string]map[string]string {
	if t.String == "" {
		return nil
	}
	var v map[string]map[string]string
	json.Unmarshal([]byte(t.String), &v)
	return v
}

func ToNullStringMaps(v []map[string]string) NullStringMaps {
	if len(v) == 0 {
		return NullStringMaps{}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return NullStringMaps{}
	}
	var x NullStringMaps
	ScanNullObject(&x, string(s))
	return x
}

func (t NullStringMaps) StringMaps() []map[string]string {
	if t.String == "" {
		return nil
	}
	var v []map[string]string
	json.Unmarshal([]byte(t.String), &v)
	return v
}

func ToNullStringMapsMap(v map[string][]map[string]string) NullStringMapsMap {
	if len(v) == 0 {
		return NullStringMapsMap{}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return NullStringMapsMap{}
	}
	var x NullStringMapsMap
	ScanNullObject(&x, string(s))
	return x
}
func (t NullStringMapsMap) StringMapsMap() map[string][]map[string]string {
	if t.String == "" {
		return nil
	}
	var v map[string][]map[string]string
	json.Unmarshal([]byte(t.String), &v)
	return v
}

func ToNullStringsMap(v map[string][]string) NullStringsMap {
	if len(v) == 0 {
		return NullStringsMap{}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return NullStringsMap{}
	}
	var x NullStringsMap
	ScanNullObject(&x, string(s))
	return x
}

func (t NullStringsMap) StringsMap() map[string][]string {
	if t.String == "" {
		return nil
	}
	var v map[string][]string
	json.Unmarshal([]byte(t.String), &v)
	return v
}

func ToNullSepStrings(v []string) NullSepStrings {
	if len(v) == 0 {
		return NullSepStrings{}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return NullSepStrings{}
	}
	var x NullSepStrings
	ScanNullObject(&x, string(s))
	return x
}
func (t NullSepStrings) Strings(sep string) []string {
	if t.String == "" {
		return nil
	}
	return strings.Split(t.String, sep)
}

func ToNullSepUint8s(v []uint8) NullSepUint8s {
	if len(v) == 0 {
		return NullSepUint8s{}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return NullSepUint8s{}
	}
	var x NullSepUint8s
	ScanNullObject(&x, string(s))
	return x
}
func (t NullSepUint8s) Uint8s(sep string) []uint8 {
	if t.String == "" {
		return nil
	}
	arr := strings.Split(t.String, sep)
	v := make([]uint8, len(arr))
	for i, a := range arr {
		b, _ := strconv.ParseUint(a, 10, 8)
		v[i] = uint8(b)
	}
	return v
}

func ToNullSepUint16s(v []uint16) NullSepUint16s {
	if len(v) == 0 {
		return NullSepUint16s{}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return NullSepUint16s{}
	}
	var x NullSepUint16s
	ScanNullObject(&x, string(s))
	return x
}
func (t NullSepUint16s) Uint16s(sep string) []uint16 {
	if t.String == "" {
		return nil
	}
	arr := strings.Split(t.String, sep)
	v := make([]uint16, len(arr))
	for i, a := range arr {
		b, _ := strconv.ParseUint(a, 10, 16)
		v[i] = uint16(b)
	}
	return v
}

func ToNullSepUint32s(v []uint32) NullSepUint32s {
	if len(v) == 0 {
		return NullSepUint32s{}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return NullSepUint32s{}
	}
	var x NullSepUint32s
	ScanNullObject(&x, string(s))
	return x
}
func (t NullSepUint32s) Uint32s(sep string) []uint32 {
	if t.String == "" {
		return nil
	}
	arr := strings.Split(t.String, sep)
	v := make([]uint32, len(arr))
	for i, a := range arr {
		b, _ := strconv.ParseUint(a, 10, 32)
		v[i] = uint32(b)
	}
	return v
}

func ToNullSepInts(v []int) NullSepInts {
	if len(v) == 0 {
		return NullSepInts{}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return NullSepInts{}
	}
	var x NullSepInts
	ScanNullObject(&x, string(s))
	return x
}
func (t NullSepInts) Ints(sep string) []int {
	if t.String == "" {
		return nil
	}
	arr := strings.Split(t.String, sep)
	v := make([]int, len(arr))
	for i, a := range arr {
		b, _ := strconv.ParseInt(a, 10, 64)
		v[i] = int(b)
	}
	return v
}

func ToNullSepUint64s(v []uint64) NullSepUint64s {
	if len(v) == 0 {
		return NullSepUint64s{}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return NullSepUint64s{}
	}
	var x NullSepUint64s
	ScanNullObject(&x, string(s))
	return x
}
func (t NullSepUint64s) Uint16s(sep string) []uint64 {
	if t.String == "" {
		return nil
	}
	arr := strings.Split(t.String, sep)
	v := make([]uint64, len(arr))
	for i, a := range arr {
		v[i], _ = strconv.ParseUint(a, 10, 64)
	}
	return v
}

// @warn 这里不是json，为了节省存储空间，这里使用 [path,size,width,height] 数组方式存储
func ToNullImgSrc(v *adto.ImgSrc) NullImgSrc {
	if v == nil {
		return NullImgSrc{}
	}
	m := [4]interface{}{v.Path, v.Size, v.Width, v.Height}
	s, _ := json.Marshal(m)
	if len(s) == 0 {
		return NullImgSrc{}
	}
	var x NullImgSrc
	ScanNullObject(&x, string(s))
	return x
}

func (t NullImgSrc) ImgSrc() *adto.ImgSrc {
	if t.String == "" {
		return nil
	}

	// 为了节省存储空间，这里使用 [path,size,width,height] 数组方式存储
	//  If you sent the JSON value through browser then any number you sent that will be the type float64
	var m [4]interface{}
	err := json.Unmarshal([]byte(t.String), &m)
	if err != nil {
		return nil
	}
	return &adto.ImgSrc{
		Path:   New(m[0]).String(),
		Size:   New(m[1]).DefaultUint32(0),
		Width:  New(m[2]).DefaultUint16(0),
		Height: New(m[3]).DefaultUint16(0),
	}
}

func ToNullImgSrcs(v []adto.ImgSrc) NullImgSrcs {
	if len(v) == 0 {
		return NullImgSrcs{}
	}
	// 为了节省存储空间，这里使用 [[path,size,width,height],[path,size,width,height]...] 数组方式存储
	m := make([][4]interface{}, len(v))
	for i, w := range v {
		m[i] = [4]interface{}{w.Path, w.Size, w.Width, w.Height}
	}
	s, _ := json.Marshal(m)
	if len(s) == 0 {
		return NullImgSrcs{}
	}
	var x NullImgSrcs
	ScanNullObject(&x, string(s))
	return x
}
func (t NullImgSrcs) ImgSrcs() []adto.ImgSrc {
	if t.String == "" {
		return nil
	}

	// 为了节省存储空间，这里使用 [[path,size,width,height],[path,size,width,height]...] 数组方式存储
	//  If you sent the JSON value through browser then any number you sent that will be the type float64
	var ms [][4]interface{}
	err := json.Unmarshal([]byte(t.String), &ms)
	if err != nil {
		return nil
	}
	v := make([]adto.ImgSrc, len(ms))
	for i, m := range ms {
		v[i] = adto.ImgSrc{
			Path:   New(m[0]).String(),
			Size:   New(m[1]).DefaultUint32(0),
			Width:  New(m[2]).DefaultUint16(0),
			Height: New(m[3]).DefaultUint16(0),
		}
	}
	return v
}
