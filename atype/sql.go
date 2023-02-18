package atype

import (
	"bytes"
	"database/sql"
	"encoding/binary"
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

/*
 一般point 需要建 spatial 索引，就需要单独到一个表里，不应该放在一起
*/
type Position struct{ sql.NullString } // []byte // postion, coordinate or point
type Ip struct{ sql.NullString }       //  net.IP               // IP Address

// https://en.wikipedia.org/wiki/Bit_numbering
type BitPos uint8       // bit position (in big endian)
type BitPosition uint16 // bit position (in big endian)
type Bitwise struct {
	BitName  string // 该位名称
	BitPos   BitPos // big endian 下，位所在位置
	BitValue bool   // 该位的值
	MaxBits  uint8
}
type Bitwiser struct {
	BitName  string      // 该位名称
	BitPos   BitPosition // big endian 下，位所在位置
	BitValue bool        // 该位的值
	MaxBits  uint8
}

type Booln uint8
type Int24 int32
type Uint24 uint32
type Year uint16      // uint16 date: yyyy
type YearMonth Uint24 // uint24 date: yyyymm  不要用 Date，主要是不需要显示dd。
type Date string      // yyyy-mm-dd
type Datetime string  // yyyy-mm-dd hh:ii:ss
type UnixTime int64   // int 形式 datetime，可与 datetime, date 互转
type Text string

// type Html template.HTML   HTML 直接使用 template.HTML
type Distri Uint24 // 6 位地址简码
type AddrId uint64 // 12 位地址码

type NullText struct{ sql.NullString }

type SepStrings string // a,b,c,d,e
type SepUint8s string  // 1,2,3,4
type SepUint16s string // 1,2,3,4
type SepUint24s string // 1,2,3,4
type SepUint32s string // 1,2,3,4
type SepInts string    // 1,2,3,4
type SepUints string   // 1,2,3,4
type SepUint64s string // 1,2,3,4
type SepPercents string
type SepMoneys string // money 表示金额的时候，是可数名词。可以用 moneys 或  monies。法律属于大额常用 moneys
type SepUmoneys string
type SepAmounts string
type SepUamounts string

const (
	MinDate     Date     = "0000-00-00"
	MaxDate     Date     = "9999-12-31"
	MinDatetime Datetime = "0000-00-00 00:00:00"
	MaxDatetime Datetime = "9999-12-31 23:59:59"
)

func delimiter(delimiters ...string) string {
	if len(delimiters) == 0 || delimiters[0] == "" {
		return ","
	}
	return delimiters[0]
}

// https://dev.mysql.com/doc/refman/8.0/en/gis-data-formats.html
//	The value length is 25 bytes, made up of these components (as can be seen from the hexadecimal value):
//	4 bytes for integer SRID (0)       4326 是GPS   WGS84，表示按 lat-lng 保存
//	1 byte for integer byte order (1 = little-endian)
//	4 bytes for integer type information (1 = Point)
//	8 bytes for double-precision X coordinate (1)
//	8 bytes for double-precision Y coordinate (−1)

func ToPositionBase(srid uint32, order byte, typ uint32, x, y float64) Position {
	var pos Position
	buf := new(bytes.Buffer)
	buf.Grow(25)
	// uint32就是4个字节
	binary.Write(buf, binary.LittleEndian, srid)
	binary.Write(buf, binary.LittleEndian, order)
	binary.Write(buf, binary.LittleEndian, typ)
	binary.Write(buf, binary.LittleEndian, x)
	binary.Write(buf, binary.LittleEndian, y)
	pos.Scan(buf.Bytes())
	return pos
}
func ToPosition(coord *Coordinate) Position {
	if coord == nil {
		return Position{}
	}
	//  4326 是GPS   WGS84，表示按 lat-lng 保存
	return ToPositionBase(4326, 1, 1, coord.Latitude, coord.Longitude)
}
func (p Position) Bytes() []byte {
	if !p.Ok() {
		return nil
	}
	return []byte(p.String)
}
func (p Position) Ok() bool {
	return p.Valid && len(p.String) == 25
}
func binaryRead(r io.Reader, littleEndian bool, data interface{}) error {
	if littleEndian {
		return binary.Read(r, binary.LittleEndian, data)
	}
	return binary.Read(r, binary.BigEndian, data)
}

func (p Position) Parse() (srid uint32, order byte, typ uint32, x float64, y float64, ok bool) {
	b := p.Bytes()
	if b == nil {
		return
	}
	buf := bytes.NewReader(b[4:5])
	binary.Read(buf, binary.LittleEndian, &order) // 只有1字节，无论bigEndian，还是littleEndian，结果都一样
	littleEndian := order == 1
	buf = bytes.NewReader(b[0:4])
	binaryRead(buf, littleEndian, &srid)
	buf = bytes.NewReader(b[5:9])
	binaryRead(buf, littleEndian, &typ)
	buf = bytes.NewReader(b[9:17])
	binaryRead(buf, littleEndian, &x)
	buf = bytes.NewReader(b[17:25])
	binaryRead(buf, littleEndian, &y)
	return
}
func (p Position) Coordinate() *Coordinate {
	_, _, _, x, y, ok := p.Parse()
	if !ok {
		return nil
	}
	return &Coordinate{
		Latitude:  x,
		Longitude: y,
	}
}
func (p Position) Point() *Point {
	_, _, _, x, y, ok := p.Parse()
	if !ok {
		return nil
	}
	return &Point{
		X: x,
		Y: y,
	}
}
func ToIp(addr string) Ip {
	var ip Ip
	if addr == "" {
		return ip
	}
	nip := net.ParseIP(addr)
	if nip != nil {
		return ip
	}

	ip.Scan(nip.String())
	return ip
}
func (ip Ip) Bytes() []byte {
	if !ip.Ok() {
		return nil
	}
	return []byte(ip.String)
}

func (ip Ip) Ok() bool {
	return ip.Valid && len(ip.String) == net.IPv4len || len(ip.String) == net.IPv6len
}

func (ip Ip) To16() string {
	b := ip.Bytes()
	if b == nil {
		return ""
	}
	ip2 := net.IP(b)
	// 包括ipv4 / ipv16
	if nip := ip2.To16(); len(nip) > 0 {
		return nip.String()
	}
	return ""
}

func (n Uint24) Uint32() uint32 { return uint32(n) }

func (b BitPos) Uint8() uint8        { return uint8(b) }
func (b BitPosition) Uint16() uint16 { return uint16(b) }

//  SET x=x|v
func (b Bitwise) SetStmt(fieldName string) string {
	if b.BitValue {
		bv := 1 << b.BitPos
		bs := strconv.FormatUint(uint64(bv), 10)
		return fieldName + "=" + fieldName + "|" + bs
	}
	return b.unsetStmt(fieldName)
}

func (b Bitwise) unsetStmt(fieldName string) string {
	max := (1 << b.MaxBits) - 1
	bv := max - (1 << b.BitPos)
	bs := strconv.FormatUint(uint64(bv), 10)
	return fieldName + "=" + fieldName + "&" + bs
}

func ToBooln(b bool) Booln {
	if b {
		return 1
	}
	return 0
}
func NewBooln(b uint8) Booln {
	if b > 0 {
		return 1
	}
	return 0
}
func (b Booln) Uint8() uint8 {
	if b.Bool() {
		return 1
	}
	return 0
}
func (b Booln) Bool() bool { return b > 0 }
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
func NewDate(d string, loc *time.Location) Date {
	if d == "" || d == MinDate.String() {
		return MinDate
	}
	_, err := time.ParseInLocation("2006-01-02", d, loc)
	if err != nil {
		return MinDate
	}
	return Date(d)
}
func ToDate(t time.Time) Date { return Date(t.Format("2006-01-02")) }
func (d Date) Valid() bool {
	return d != "" && d != MinDate && d != MaxDate && d != "1970-01-01"
}
func (d Date) String() string { return string(d) }
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
func NewDatetime(d string, loc *time.Location) Datetime {
	if d == "" || d == MinDatetime.String() {
		return MinDatetime
	}
	_, err := time.ParseInLocation("2006-01-02 15:04:05", d, loc)
	if err != nil {
		return MinDatetime
	}
	return Datetime(d)
}
func ToDatetime(t time.Time) Datetime { return Datetime(t.Format("2006-01-02 15:04:05")) }

func (d Datetime) Valid() bool {
	s := d.String()
	ok := s != "" && d != MinDatetime && s != MinDate.String()
	ok = ok && s != "1970-01-01 00:00:00" && s != "1970-01-01"
	ok = ok && d != MaxDatetime && s != MaxDate.String()
	return ok
}
func (d Datetime) String() string { return string(d) }
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

func NewUnixTime(u int64) UnixTime {
	return UnixTime(u)
}
func (u UnixTime) Int64() int64 { return int64(u) }
func (u UnixTime) Date(loc *time.Location) Date {
	if u == 0 {
		return MinDate
	}
	return ToDate(time.Unix(u.Int64(), 0).In(loc))
}
func (u UnixTime) Datetime(loc *time.Location) Datetime {
	if u == 0 {
		return MinDatetime
	}
	return ToDatetime(time.Unix(u.Int64(), 0).In(loc))
}

func ToSepStrings(elems []string, delimiters ...string) SepStrings {
	return SepStrings(strings.Join(elems, delimiter(delimiters...)))
}
func (t SepStrings) Strings(delimiters ...string) []string {
	if t == "" {
		return nil
	}
	return strings.Split(string(t), delimiter(delimiters...))
}

func ToSepUint8s(elems []uint8, delimiters ...string) SepUint8s {
	// strings.Join 类同
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return SepUint8s(strconv.FormatUint(uint64(elems[0]), 10))
	}
	deli := delimiter(delimiters...)
	n := len(deli)*(len(elems)-1) + (len(elems) * MaxUint8Len)

	var b strings.Builder
	b.Grow(n)
	b.WriteByte(elems[0])
	for _, s := range elems[1:] {
		b.WriteString(deli)
		b.WriteByte(s)
	}

	return SepUint8s(b.String())
}
func (t SepUint8s) Uint8s(delimiters ...string) []uint8 {
	if t == "" {
		return nil
	}
	arr := strings.Split(string(t), delimiter(delimiters...))
	v := make([]uint8, len(arr))
	for i, a := range arr {
		b, _ := strconv.ParseUint(a, 10, 8)
		v[i] = uint8(b)
	}
	return v
}

func ToSepUint16s(elems []uint16, delimiters ...string) SepUint16s {
	// strings.Join 类同
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return SepUint16s(strconv.FormatUint(uint64(elems[0]), 10))
	}
	deli := delimiter(delimiters...)
	n := len(deli)*(len(elems)-1) + (len(elems) * MaxUint16Len)
	var b strings.Builder
	b.Grow(n)
	b.WriteRune(rune(elems[0]))
	for _, s := range elems[1:] {
		b.WriteString(deli)
		b.WriteRune(rune(s))
	}

	return SepUint16s(b.String())
}
func (t SepUint16s) Uint16s(delimiters ...string) []uint16 {
	if t == "" {
		return nil
	}
	arr := strings.Split(string(t), delimiter(delimiters...))
	v := make([]uint16, len(arr))
	for i, a := range arr {
		b, _ := strconv.ParseUint(a, 10, 16)
		v[i] = uint16(b)
	}
	return v
}

func ToSepUint24s(elems []Uint24, delimiters ...string) SepUint24s {
	// strings.Join 类同
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return SepUint24s(strconv.FormatUint(uint64(elems[0]), 10))
	}
	deli := delimiter(delimiters...)
	n := len(deli)*(len(elems)-1) + (len(elems) * MaxUint24Len)
	var b strings.Builder
	b.Grow(n)
	b.WriteString(strconv.FormatUint(uint64(elems[0]), 10))
	for _, s := range elems[1:] {
		b.WriteString(deli)
		b.WriteString(strconv.FormatUint(uint64(s), 10))
	}

	return SepUint24s(b.String())
}

func (t SepUint24s) Uint32s(delimiters ...string) []Uint24 {
	if t == "" {
		return nil
	}
	arr := strings.Split(string(t), delimiter(delimiters...))
	v := make([]Uint24, len(arr))
	for i, a := range arr {
		b, _ := strconv.ParseUint(a, 10, 24)
		v[i] = Uint24(b)
	}
	return v
}

func ToSepUint32s(elems []uint32, delimiters ...string) SepUint32s {
	// strings.Join 类同
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return SepUint32s(strconv.FormatUint(uint64(elems[0]), 10))
	}
	deli := delimiter(delimiters...)
	n := len(deli)*(len(elems)-1) + (len(elems) * MaxUintLen)
	var b strings.Builder
	b.Grow(n)
	b.WriteString(strconv.FormatUint(uint64(elems[0]), 10))
	for _, s := range elems[1:] {
		b.WriteString(deli)
		b.WriteString(strconv.FormatUint(uint64(s), 10))
	}

	return SepUint32s(b.String())
}

func (t SepUint32s) Uint32s(delimiters ...string) []uint32 {
	if t == "" {
		return nil
	}
	arr := strings.Split(string(t), delimiter(delimiters...))
	v := make([]uint32, len(arr))
	for i, a := range arr {
		b, _ := strconv.ParseUint(a, 10, 32)
		v[i] = uint32(b)
	}
	return v
}

func ToSepInts(elems []int, delimiters ...string) SepInts {
	// strings.Join 类同
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return SepInts(strconv.FormatInt(int64(elems[0]), 10))
	}
	deli := delimiter(delimiters...)
	n := len(deli)*(len(elems)-1) + (len(elems) * MaxIntLen)
	var b strings.Builder
	b.Grow(n)
	b.WriteString(strconv.FormatInt(int64(elems[0]), 10))
	for _, s := range elems[1:] {
		b.WriteString(deli)
		b.WriteString(strconv.FormatInt(int64(s), 10))
	}

	return SepInts(b.String())
}
func (t SepInts) Ints(delimiters ...string) []int {
	if t == "" {
		return nil
	}
	arr := strings.Split(string(t), delimiter(delimiters...))
	v := make([]int, len(arr))
	for i, a := range arr {
		b, _ := strconv.ParseInt(a, 10, 64)
		v[i] = int(b)
	}
	return v
}

func ToSepUints(elems []uint, delimiters ...string) SepUints {
	// strings.Join 类同
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return SepUints(strconv.FormatUint(uint64(elems[0]), 10))
	}
	deli := delimiter(delimiters...)
	n := len(deli)*(len(elems)-1) + (len(elems) * MaxUintLen)
	var b strings.Builder
	b.Grow(n)
	b.WriteString(strconv.FormatUint(uint64(elems[0]), 10))
	for _, s := range elems[1:] {
		b.WriteString(deli)
		b.WriteString(strconv.FormatUint(uint64(s), 10))
	}

	return SepUints(b.String())
}
func (t SepUints) Uints(delimiters ...string) []uint {
	if t == "" {
		return nil
	}
	arr := strings.Split(string(t), delimiter(delimiters...))
	v := make([]uint, len(arr))
	for i, a := range arr {
		x, _ := strconv.ParseUint(a, 10, 32)
		v[i] = uint(x)
	}
	return v
}

func ToSepUint64s(elems []uint64, delimiters ...string) SepUint64s {
	// strings.Join 类同
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return SepUint64s(strconv.FormatUint(elems[0], 10))
	}
	deli := delimiter(delimiters...)
	n := len(deli)*(len(elems)-1) + (len(elems) * MaxUint64Len)
	var b strings.Builder
	b.Grow(n)
	b.WriteString(strconv.FormatUint(elems[0], 10))
	for _, s := range elems[1:] {
		b.WriteString(deli)
		b.WriteString(strconv.FormatUint(s, 10))
	}

	return SepUint64s(b.String())
}
func (t SepUint64s) Uint64s(delimiters ...string) []uint64 {
	if t == "" {
		return nil
	}
	arr := strings.Split(string(t), delimiter(delimiters...))
	v := make([]uint64, len(arr))
	for i, a := range arr {
		v[i], _ = strconv.ParseUint(a, 10, 64)
	}
	return v
}

func ToSepPercents(elems []Percent, delimiters ...string) SepPercents {

	// strings.Join 类同
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return SepPercents(strconv.FormatInt(int64(elems[0]), 10))
	}
	deli := delimiter(delimiters...)
	n := len(deli)*(len(elems)-1) + (len(elems) * MaxInt16Len)
	var b strings.Builder
	b.Grow(n)
	b.WriteString(strconv.FormatInt(int64(elems[0]), 10))
	for _, s := range elems[1:] {
		b.WriteString(deli)
		b.WriteString(strconv.FormatInt(int64(s), 10))
	}

	return SepPercents(b.String())
}

func (t SepPercents) Percents(delimiters ...string) []Percent {
	if t == "" {
		return nil
	}
	arr := strings.Split(string(t), delimiter(delimiters...))
	v := make([]Percent, len(arr))
	for i, a := range arr {
		p, err := strconv.ParseInt(a, 10, 16)
		if err == nil {
			v[i] = NewPercent(int16(p))
		}
	}
	return v
}

func ToSepMoneys(elems []Money, delimiters ...string) SepMoneys {
	// strings.Join 类同
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return SepMoneys(strconv.FormatInt(int64(elems[0]), 10))
	}
	deli := delimiter(delimiters...)
	n := len(deli)*(len(elems)-1) + (len(elems) * MaxIntLen)
	var b strings.Builder
	b.Grow(n)
	b.WriteString(strconv.FormatInt(int64(elems[0]), 10))
	for _, s := range elems[1:] {
		b.WriteString(deli)
		b.WriteString(strconv.FormatInt(int64(s), 10))
	}

	return SepMoneys(b.String())
}

func (t SepMoneys) Moneys(delimiters ...string) []Money {
	if t == "" {
		return nil
	}
	arr := strings.Split(string(t), delimiter(delimiters...))
	v := make([]Money, len(arr))
	for i, a := range arr {
		p, err := strconv.ParseInt(a, 10, 32)
		if err == nil {
			v[i] = NewMoney(int(p))
		}
	}
	return v
}
func ToSepUmoneys(elems []Umoney, delimiters ...string) SepUmoneys {
	// strings.Join 类同
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return SepUmoneys(strconv.FormatInt(int64(elems[0]), 10))
	}
	deli := delimiter(delimiters...)
	n := len(deli)*(len(elems)-1) + (len(elems) * MaxUintLen)
	var b strings.Builder
	b.Grow(n)
	b.WriteString(strconv.FormatInt(int64(elems[0]), 10))
	for _, s := range elems[1:] {
		b.WriteString(deli)
		b.WriteString(strconv.FormatInt(int64(s), 10))
	}

	return SepUmoneys(b.String())
}

func (t SepUmoneys) Umoneys(delimiters ...string) []Umoney {
	if t == "" {
		return nil
	}
	arr := strings.Split(string(t), delimiter(delimiters...))
	v := make([]Umoney, len(arr))
	for i, a := range arr {
		p, err := strconv.ParseUint(a, 10, 32)
		if err == nil {
			v[i] = NewUmoney(uint(p))
		}
	}
	return v
}

func ToSepAmounts(elems []Amount, delimiters ...string) SepAmounts {
	// strings.Join 类同
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return SepAmounts(strconv.FormatInt(int64(elems[0]), 10))
	}
	deli := delimiter(delimiters...)
	n := len(deli)*(len(elems)-1) + (len(elems) * MaxInt64Len)
	var b strings.Builder
	b.Grow(n)
	b.WriteString(strconv.FormatInt(int64(elems[0]), 10))
	for _, s := range elems[1:] {
		b.WriteString(deli)
		b.WriteString(strconv.FormatInt(int64(s), 10))
	}

	return SepAmounts(b.String())
}
func (t SepAmounts) Amounts(delimiters ...string) []Amount {
	if t == "" {
		return nil
	}
	arr := strings.Split(string(t), delimiter(delimiters...))
	v := make([]Amount, len(arr))
	for i, a := range arr {
		p, err := strconv.ParseInt(a, 10, 64)
		if err == nil {
			v[i] = NewAmount(p)
		}
	}
	return v
}
func ToSepUamounts(elems []Uamount, delimiters ...string) SepUamounts {
	// strings.Join 类同
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return SepUamounts(strconv.FormatInt(int64(elems[0]), 10))
	}
	deli := delimiter(delimiters...)
	n := len(deli)*(len(elems)-1) + (len(elems) * MaxUint64Len)

	var b strings.Builder
	b.Grow(n)
	b.WriteString(strconv.FormatInt(int64(elems[0]), 10))
	for _, s := range elems[1:] {
		b.WriteString(deli)
		b.WriteString(strconv.FormatInt(int64(s), 10))
	}

	return SepUamounts(b.String())
}

func (t SepUamounts) Amounts(delimiters ...string) []Uamount {
	if t == "" {
		return nil
	}
	arr := strings.Split(string(t), delimiter(delimiters...))
	v := make([]Uamount, len(arr))
	for i, a := range arr {
		p, err := strconv.ParseUint(a, 10, 64)
		if err == nil {
			v[i] = NewUamount(p)
		}
	}
	return v
}

func NewDistri(d Uint24) Distri {
	return Distri(d)
}
func ToDistri(d uint32) Distri {
	return Distri(d)
}
func (d Distri) Uint24() Uint24 {
	return Uint24(d)
}
func (d Distri) Uint32() uint32 {
	return uint32(d)
}
func NewAddrId(a uint64) AddrId {
	return AddrId(a)
}
func (a AddrId) Uint64() uint64 {
	return uint64(a)
}
