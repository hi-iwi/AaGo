package dtype

import (
	"bytes"
	"database/sql"
	"encoding/binary"
	"encoding/json"
	"io"
	"net"
	"strconv"
	"strings"
	"time"
)

type ObjScan interface {
	Scan(value interface{}) error
}

type Coordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}
type Position []byte // postion, coordinate or point
type Ip []byte       // IP Address

type Bitwise struct {
	BitName       string // 该位名称
	BitRightIndex uint8  // 位所在位置
	BitValue      bool   // 该位的值
	MaxBits       uint8
}

type Boolean uint8
type Uint24 uint32
type Year uint16      // uint16 date: yyyy
type YearMonth uint32 // uint24 date: yyyymm  不要用 Date，主要是不需要显示dd。
type Date string      // yyyy-mm-dd
type Datetime string  // yyyy-mm-dd hh:ii:ss
type UnixTime int64   // int 形式 datetime，可与 datetime, date 互转
type Text struct{ sql.NullString }
type CityId uint32 // uint24
type AddrId uint64

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
// 保证空字符串不能正常的对象
func ScanNullObject(obj ObjScan, data string) {
	if data == "" {
		obj.Scan(nil)
	}
	obj.Scan(data)
}

// https://dev.mysql.com/doc/refman/8.0/en/gis-data-formats.html
//	The value length is 25 bytes, made up of these components (as can be seen from the hexadecimal value):
//	4 bytes for integer SRID (0)       4326 是GPS   WGS84，表示按 lat-lng 保存
//	1 byte for integer byte order (1 = little-endian)
//	4 bytes for integer type information (1 = Point)
//	8 bytes for double-precision X coordinate (1)
//	8 bytes for double-precision Y coordinate (−1)

func ToPositionBase(srid uint32, order byte, typ uint32, x, y float64) Position {
	buf := new(bytes.Buffer)
	buf.Grow(25)
	// uint32就是4个字节
	binary.Write(buf, binary.LittleEndian, srid)
	binary.Write(buf, binary.LittleEndian, order)
	binary.Write(buf, binary.LittleEndian, typ)
	binary.Write(buf, binary.LittleEndian, x)
	binary.Write(buf, binary.LittleEndian, y)
	return buf.Bytes()
}
func ToPosition(coord Coordinate) Position {
	//  4326 是GPS   WGS84，表示按 lat-lng 保存
	return ToPositionBase(4326, 1, 1, coord.Latitude, coord.Longitude)
}

func binaryRead(r io.Reader, littleEndian bool, data interface{}) error {
	if littleEndian {
		return binary.Read(r, binary.LittleEndian, data)
	}
	return binary.Read(r, binary.BigEndian, data)
}
func (p Position) Valid() bool {
	return len(p) == 25
}
func (p Position) Parse() (srid uint32, order byte, typ uint32, x float64, y float64) {
	if p.Valid() {
		buf := bytes.NewReader(p[4:5])
		binary.Read(buf, binary.LittleEndian, &order) // 只有1字节，无论bigEndian，还是littleEndian，结果都一样
		littleEndian := order == 1
		buf = bytes.NewReader(p[0:4])
		binaryRead(buf, littleEndian, &srid)
		buf = bytes.NewReader(p[5:9])
		binaryRead(buf, littleEndian, &typ)
		buf = bytes.NewReader(p[9:17])
		binaryRead(buf, littleEndian, &x)
		buf = bytes.NewReader(p[17:25])
		binaryRead(buf, littleEndian, &y)
	}
	return
}
func (p Position) Coordinate() Coordinate {
	_, _, _, x, y := p.Parse()
	return Coordinate{
		Latitude:  x,
		Longitude: y,
	}
}
func (p Position) Point() Point {
	_, _, _, x, y := p.Parse()
	return Point{
		X: x,
		Y: y,
	}
}
func (ip Ip) Valid() bool {
	return len(ip) == net.IPv4len || len(ip) == net.IPv6len
}
func (ip Ip) String() string {
	var addr string
	if len(ip) > 1 {
		if nip := net.IP(ip).To16(); len(nip) > 0 {
			addr = nip.String()
		}
	}
	return addr
}

func ToIp(addr string) Ip {
	if addr != "" {
		nip := net.ParseIP(addr)
		if nip != nil {
			return Ip(nip)
		}
	}
	return []byte{0} // varbinary(16) NULL DEFAULT 0x00 ;   ==> 等价于 DEFAULT '\0'
}

func PtrUint64(n *uint64) uint64 {
	if n == nil {
		return 0
	}
	return *n
}
func PtrUint32(n *uint32) uint32 {
	if n == nil {
		return 0
	}
	return *n
}
func PtrUint(n *uint) uint {
	if n == nil {
		return 0
	}
	return *n
}
func (n Uint24) Uint32() uint32 {return uint32(n)}
func PtrUint24(n *Uint24) Uint24 {
	if n == nil {
		return 0
	}
	return *n
}
func PtrUint16(n *uint16) uint16 {
	if n == nil {
		return 0
	}
	return *n
}
func PtrUint8(n *uint8) uint8 {
	if n == nil {
		return 0
	}
	return *n
}
func PtrInt64(n *int64) int64 {
	if n == nil {
		return 0
	}
	return *n
}
func PtrInt32(n *int32) int32 {
	if n == nil {
		return 0
	}
	return *n
}
func PtrInt(n *int) int {
	if n == nil {
		return 0
	}
	return *n
}
func PtrInt16(n *int16) int16 {
	if n == nil {
		return 0
	}
	return *n
}
func PtrInt8(n *int8) int8 {
	if n == nil {
		return 0
	}
	return *n
}
func PtrFloat64(n *float64) float64 {
	if n == nil {
		return 0.00
	}
	return *n
}
func PtrFloat32(n *float32) float32 {
	if n == nil {
		return 0.00
	}
	return *n
}

//  SET x=x|v
func (b Bitwise) SetStmt(fieldName string) string {
	if b.BitValue {
		bv := 1 << b.BitRightIndex
		bs := strconv.FormatUint(uint64(bv), 10)
		return fieldName + "=" + fieldName + "|" + bs
	}
	return b.unsetStmt(fieldName)
}

func (b Bitwise) unsetStmt(fieldName string) string {
	max := (1 << b.MaxBits) - 1
	bv := max - (1 << b.BitRightIndex)
	bs := strconv.FormatUint(uint64(bv), 10)
	return fieldName + "=" + fieldName + "&" + bs
}

func ToBoolean(b bool) Boolean {
	if b {
		return 1
	}
	return 0
}
func (b Boolean) Uint8() uint8 { return uint8(b) }
func (b Boolean) Bool() bool { return b > 0 }
func ToYearMonth(year int, month time.Month) YearMonth {
	if year < 0 {
		return 0
	}
	ym := year*100 + int(month)
	return YearMonth(ym)
}

// years/months 可为负数
func (ym YearMonth) Add(years int, months int, loc *time.Location) YearMonth {
	tm := ym.Time(loc).AddDate(years, months, 0)
	return ToYearMonth(tm.Year(), tm.Month())
}
func (ym YearMonth) Uin32() uint32 { return uint32(ym) }
func (ym YearMonth) Date() (int, time.Month) {
	year := int(ym) / 100
	month := time.Month(ym % 100)
	return year, month
}
func (ym YearMonth) Time(loc *time.Location) time.Time {
	y, m := ym.Date()
	return time.Date(y, m, 0, 00, 00, 00, 0, loc)
}

// time.Now().In()  loc 直接通过 in 传递
func ToDate(t time.Time) Date { return Date(t.Format("2006-01-02")) }

func (d Date) Time(loc *time.Location) (time.Time, error) {
	return time.ParseInLocation("2006-01-02", string(d), loc)
}
func (d Date) Int64(loc *time.Location) int64 {
	tm, err := time.ParseInLocation("2006-01-02", string(d), loc)
	if err != nil {
		return 0
	}
	return tm.Unix()
}
func (d Date) Unix(loc *time.Location) UnixTime {
	t, _ := d.Time(loc)
	return UnixTime(t.Unix())
}

// time.Now().In()  loc 直接通过 in 传递
func ToDatetime(t time.Time) Datetime { return Datetime(t.Format("2006-01-02 15:04:05")) }
func (d Datetime) Time(loc *time.Location) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:05", string(d), loc)
}
func (d Datetime) Int64(loc *time.Location) int64 {
	tm, err := time.ParseInLocation("2006-01-02 15:04:05", string(d), loc)
	if err != nil {
		return 0
	}
	return tm.Unix()
}
func (d Datetime) Unix(loc *time.Location) UnixTime {
	t, _ := d.Time(loc)
	return UnixTime(t.Unix())
}

func (u UnixTime) Int64() int64 {return int64(u)}

func (u UnixTime) Date(loc *time.Location) Date {
	return ToDate(time.Unix(u.Int64(), 0).In(loc))
}
func (u UnixTime) Datetime(loc *time.Location) Datetime {
	return ToDatetime(time.Unix(u.Int64(), 0).In(loc))
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

func ToNullStringMap(v map[string]string) NullStringMap {
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

func ToNullSepStrings(elems []string, sep string) NullSepStrings {
	var x NullSepStrings
	ScanNullObject(&x, strings.Join(elems, sep))
	return x
}
func (t NullSepStrings) Strings(sep string) []string {
	if t.String == "" {
		return nil
	}
	return strings.Split(t.String, sep)
}

func ToNullSepUint8s(elems []uint8, sep string) NullSepUint8s {
	// strings.Join 类同
	switch len(elems) {
	case 0:
		return NullSepUint8s{}
	case 1:
		var x NullSepUint8s
		x.Scan(strconv.FormatUint(uint64(elems[0]), 10))
		return x
	}

	n := len(sep)*(len(elems)-1) + (len(elems) * 3) // uint8 -> 0~256, max 3 bytes

	var b strings.Builder
	b.Grow(n)
	b.WriteByte(elems[0])
	for _, s := range elems[1:] {
		b.WriteString(sep)
		b.WriteByte(s)
	}

	var x NullSepUint8s
	ScanNullObject(&x, b.String())
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

func ToNullSepUint16s(elems []uint16, sep string) NullSepUint16s {
	// strings.Join 类同
	switch len(elems) {
	case 0:
		return NullSepUint16s{}
	case 1:
		var x NullSepUint16s
		x.Scan(strconv.FormatUint(uint64(elems[0]), 10))
		return x
	}

	n := len(sep)*(len(elems)-1) + (len(elems) * 5) // uint16 -> 0~65535, max 5 bytes

	var b strings.Builder
	b.Grow(n)
	b.WriteRune(rune(elems[0]))
	for _, s := range elems[1:] {
		b.WriteString(sep)
		b.WriteRune(rune(s))
	}

	var x NullSepUint16s
	ScanNullObject(&x, b.String())
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

func ToNullSepUint32s(elems []uint32, sep string) NullSepUint32s {
	// strings.Join 类同
	switch len(elems) {
	case 0:
		return NullSepUint32s{}
	case 1:
		var x NullSepUint32s
		x.Scan(strconv.FormatUint(uint64(elems[0]), 10))
		return x
	}

	n := len(sep)*(len(elems)-1) + (len(elems) * 10) // uint32 -> 0~4294967295, max 10 bytes

	var b strings.Builder
	b.Grow(n)
	b.WriteString(strconv.FormatUint(uint64(elems[0]), 10))
	for _, s := range elems[1:] {
		b.WriteString(sep)
		b.WriteString(strconv.FormatUint(uint64(s), 10))
	}

	var x NullSepUint32s
	ScanNullObject(&x, b.String())
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

func ToNullSepInts(elems []int, sep string) NullSepInts {
	// strings.Join 类同
	switch len(elems) {
	case 0:
		return NullSepInts{}
	case 1:
		var x NullSepInts
		x.Scan(strconv.FormatInt(int64(elems[0]), 10))
		return x
	}

	n := len(sep)*(len(elems)-1) + (len(elems) * 11) // int -> -2147483648到2147483647, max 11 bytes

	var b strings.Builder
	b.Grow(n)
	b.WriteString(strconv.FormatInt(int64(elems[0]), 10))
	for _, s := range elems[1:] {
		b.WriteString(sep)
		b.WriteString(strconv.FormatInt(int64(s), 10))
	}

	var x NullSepInts
	ScanNullObject(&x, b.String())
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

func ToNullSepUint64s(elems []uint64, sep string) NullSepUint64s {
	// strings.Join 类同
	switch len(elems) {
	case 0:
		return NullSepUint64s{}
	case 1:
		var x NullSepUint64s
		x.Scan(strconv.FormatUint(elems[0], 10))
		return x
	}

	n := len(sep)*(len(elems)-1) + (len(elems) * 11) // int -> -2147483648到2147483647, max 11 bytes

	var b strings.Builder
	b.Grow(n)
	b.WriteString(strconv.FormatUint(elems[0], 10))
	for _, s := range elems[1:] {
		b.WriteString(sep)
		b.WriteString(strconv.FormatUint(s, 10))
	}

	var x NullSepUint64s
	ScanNullObject(&x, b.String())
	return x
}
func (t NullSepUint64s) Uint64s(sep string) []uint64 {
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
