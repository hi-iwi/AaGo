package atype

import (
	"github.com/shopspring/decimal"
	"math"
)

type Percent int16 // 范围：-10000~10000 => -100.00%~100.00%

var PercentMultiplier = decimal.NewFromInt32(100) // 扩大100 * 100倍 --> 这里按百分比算，而不是小数  3* Percent 为 3% = 0.03

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

// @param n 本身就是转换后的值，如10000，即表示为 100*PercentMultiplier，即 100%
func NewPercent(n int16) Percent { return Percent(n) }

// 80.0 表示 80%
func ToPercent(n float64) Percent {
	return NewPercent(int16(PercentMultiplier.Mul(decimal.NewFromFloat(n)).IntPart()))
}
func (p Percent) Int16() int16 { return int16(p) }
func (p Percent) Int32() int32 { return int32(p) }
func (p Percent) Percent() decimal.Decimal {
	q := decimal.NewFromInt32(p.Int32())
	return q.Div(PercentMultiplier)
}
func (p Percent) Value() decimal.Decimal { return p.Percent().Div(decimal.NewFromInt32(100)) }
func (p Percent) Mul(d decimal.Decimal) decimal.Decimal {
	return p.Value().Mul(d)
}
func (p Percent) Fmt() string                { return p.Percent().String() }
func (a Money) Dec() decimal.Decimal         { return decimal.NewFromInt(int64(a)) }
func (a Money) MulPercent(p Percent) Money   { return Money(p.Mul(a.Dec()).IntPart()) }
func (a Umoney) Dec() decimal.Decimal        { return decimal.NewFromInt(int64(a)) }
func (a Umoney) MulPercent(p Percent) Umoney { return Umoney(p.Mul(a.Dec()).IntPart()) }
