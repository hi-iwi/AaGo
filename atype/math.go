package atype

import (
	"math"
	"strconv"
)

type Percent int16 // 范围：-10000~10000 => -100.00%~100.00%
const MaxInt24 = 1<<23 - 1
const MinInt24 = -1 << 23
const MaxUint24 = 1<<24 - 1
const PercentMultiplier = float64(10000.0) // 扩大一万倍

func FloatToInt8(n float64) int8 {
	if n < float64(math.MinInt8) || n > float64(math.MaxInt8) {
		return 0
	}
	return int8(n) // 上面进行范围判断了，所以之类可以强转
}
func FloatToInt16(n float64) int16 {
	if n < float64(math.MinInt16) || n > float64(math.MaxInt16) {
		return 0
	}
	return int16(n) // 上面进行范围判断了，所以之类可以强转
}
func FloatToInt24(n float64) Int24 {
	if n < float64(MinInt24) || n > float64(MaxInt24) {
		return 0
	}
	return Int24(n) // 上面进行范围判断了，所以之类可以强转
}
func FloatToInt(n float64) int {
	if n < float64(math.MinInt) || n > float64(math.MaxInt) {
		return 0
	}
	return int(n) // 上面进行范围判断了，所以之类可以强转
}
func Float2Int64(n float64) int64 {
	if n < float64(math.MinInt64) || n > float64(math.MaxInt64) {
		return 0
	}
	return int64(n) // 上面进行范围判断了，所以之类可以强转
}
func Float2Uint8(n float64) uint8 {
	if n < 0.0 || n > float64(math.MaxUint8) {
		return 0
	}
	return uint8(n) // 上面进行范围判断了，所以之类可以强转
}
func Float2Uint16(n float64) uint16 {
	if n < 0.0 || n > float64(math.MaxUint16) {
		return 0
	}
	return uint16(n) // 上面进行范围判断了，所以之类可以强转
}
func Float2Uint24(n float64) Uint24 {
	if n < 0.0 || n > float64(MaxUint24) {
		return 0
	}
	return Uint24(n) // 上面进行范围判断了，所以之类可以强转
}
func Float2Uint(n float64) uint {
	if n < 0.0 || n > float64(math.MaxUint) {
		return 0
	}
	return uint(n) // 上面进行范围判断了，所以之类可以强转
}

func Float2Uint64(n float64) uint64 {
	if n < 0 || n > float64(math.MaxUint64) {
		return 0
	}
	return uint64(n) // 上面进行范围判断了，所以之类可以强转
}
func NewPercent(n int16) (Percent, bool) {
	if n < -int16(PercentMultiplier) || n > int16(PercentMultiplier) {
		return 0, false
	}
	return Percent(n), true
}
func ToPercent(v float64) (Percent, bool) {
	v *= PercentMultiplier
	return NewPercent(int16(v))
}

// 真实的值
func (p Percent) Value() float64           { return float64(p) / PercentMultiplier }
func (p Percent) Percent() float64         { return float64(p) / 100.0 }
func (p Percent) Int16() int16             { return int16(p) }
func (p Percent) By(a float64) float64     { return p.Value() * a }
func (p Percent) ByInt(a int) int          { return FloatToInt(p.By(float64(a))) }
func (p Percent) ByInt64(a int64) int64    { return Float2Int64(float64(a) * p.Value()) }
func (p Percent) ByUint(a uint) uint       { return Float2Uint(p.By(float64(a))) }
func (p Percent) ByUint64(a uint64) uint64 { return Float2Uint64(p.By(float64(a))) }
func (p Percent) Fmt() string              { return strconv.FormatFloat(p.Percent(), 'f', -1, 64) }

func (a Money) ByPercent(p Percent) Money     { return Money(p.ByInt(a.Int())) }
func (a Umoney) ByPercent(p Percent) Umoney   { return Umoney(p.ByUint(a.Uint())) }
func (a Amount) ByPercent(p Percent) Amount   { return Amount(p.ByInt64(a.Int64())) }
func (a Uamount) ByPercent(p Percent) Uamount { return Uamount(p.ByUint64(a.Uint64())) }
