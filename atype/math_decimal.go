package atype

import (
	"math"
	"strconv"
)

// Decimal 万分之一
type Decimal64 int64 // [ -922337203685477.5808,  -922337203685477.5807]
type Decimal int     // [-214748.3648, 214748.3647]
type Decimal24 Int24 // [-838.8608, 838.8607]
type Decimal16 int16 // [-3.2768, 3.2767]

type Percentage int      //  [-21474836.48%, 21474836.47%]
type Percentage16 uint16 // [0%,655.35%]

const (
	Thousandth     Percentage = 10  // 千分比
	Percent        Percentage = 100 // 百分比
	ThousandthRate            = 1000.0
	PercentRate               = 100.0

	unitDecimalInt64   = int64(10000)
	unitDecimalFloat64 = float64(unitDecimalInt64) // Decimal 表示万分之一
	UnitDecimal        = Decimal64(unitDecimalInt64)
)

// 不要直接 int(float) 转换，否则容易出错。比如 int(60135.0000) == 60134
func Decimal64Unit(n float64) Decimal64 { return Decimal64(math.Round(n * unitDecimalFloat64)) }
func Decimal64UnitN(n int) Decimal64    { return Decimal64(n) * UnitDecimal }

func (p Decimal64) Int64() int64     { return int64(p) }
func (p Decimal64) Decimal() float64 { return float64(p) / unitDecimalFloat64 }

// 整数部分
func (p Decimal64) Precision() int64 { return int64(p) / unitDecimalInt64 }

// 小数部分
func (p Decimal64) Scale() uint16 {
	return uint16(int64(math.Abs(float64(p))) % unitDecimalInt64)
}

// 类型：  1,000,000 这种
func (p Decimal64) FmtPrecision(n int, delimiter string) string {
	s := strconv.FormatInt(p.Precision(), 10)
	return fmtPrecision(s, n, delimiter)
}
func (p Decimal64) FmtScale(decimals ...uint16) string {
	return formatScale(p.Scale(), decimalN(decimals...), true)
}
func (p Decimal64) FormatScale(decimals ...uint16) string {
	return formatScale(p.Scale(), decimalN(decimals...), false)
}
func (p Decimal64) Fmt(decimals ...uint16) string {
	ys := strconv.FormatInt(p.Precision(), 10)
	return ys + formatScale(p.Scale(), decimalN(decimals...), true)
}
func (p Decimal64) Format(decimals ...uint16) string {
	ys := strconv.FormatInt(p.Precision(), 10)
	return ys + formatScale(p.Scale(), decimalN(decimals...), false)
}

//  0.8 * 0.8 = 0.64 (6400)   =====>   8000 * 8000  =  6400 [0000]
// 1.6 * 0.8 = 1.28 (12800)    =====>  16000 * 8000 = 12800 [0000]
// 0.009 * 0.04 =0.0003[6] (3)   -->   90 * 400  = 3[6000]     舍去最后一位
func (p Decimal64) Mul(b Decimal64) Decimal64 {
	return Decimal64(math.Round(float64(p*b) / unitDecimalFloat64))
}
func (p Decimal64) MulCeil(b Decimal64) Decimal64 {
	return Decimal64(math.Ceil(float64(p*b) / unitDecimalFloat64))
}
func (p Decimal64) MulFloor(b Decimal64) Decimal64 { return p * b / UnitDecimal }
func (p Decimal64) Div(b Decimal64) Decimal64 {
	return Decimal64(math.Round(float64(p) * unitDecimalFloat64 / float64(b)))
}
func (p Decimal64) DivCeil(b Decimal64) Decimal64 {
	return Decimal64(math.Ceil(float64(p) * unitDecimalFloat64 / float64(b)))
}
func (p Decimal64) DivFloor(b Decimal64) Decimal64 { return p * UnitDecimal / b }

func DecimalUnit(n float64) Decimal { return Decimal(math.Round(n * unitDecimalFloat64)) }
func (p Decimal) Int() int          { return int(p) }
func (p Decimal) Decimal() float64  { return float64(p) / unitDecimalFloat64 }
func (p Decimal) P() Decimal64      { return Decimal64(p) } // prototype

func Decimal24Unit(n float64) Decimal24 { return Decimal24(math.Round(n * unitDecimalFloat64)) }
func (p Decimal24) Int32() int32        { return int32(p) }
func (p Decimal24) Decimal() float64    { return float64(p) / unitDecimalFloat64 }
func (p Decimal24) P() Decimal64        { return Decimal64(p) } // prototype

func Decimal16Unit(n float64) Decimal16 { return Decimal16(math.Round(n * unitDecimalFloat64)) }
func (p Decimal16) Int16() int16        { return int16(p) }
func (p Decimal16) Decimal() float64    { return float64(p) / unitDecimalFloat64 }
func (p Decimal16) P() Decimal64        { return Decimal64(p) } // prototype

func HundredPercent(n float64) Percentage { return Percentage(math.Round(n * PercentRate)) }
func (p Percentage) Decimal() Decimal     { return Decimal(p) }
func (p Percentage) Int() int             { return int(p) }
func (p Percentage) Percent() float64     { return float64(p) / PercentRate }
func (p Percentage) Fmt() string          { return strconv.FormatFloat(float64(p.Percent()), 'f', -1, 32) }
func (p Percentage) FmtAbs() string {
	c := p.Percent()
	if c < 0 {
		c = -c
	}
	return strconv.FormatFloat(float64(c), 'f', -1, 32)
}

// e.g. 100*atype.Percent
func HundredPercent16(n float64) Percentage16 { return Percentage16(math.Round(n * PercentRate)) }
func (p Percentage16) Decimal() Decimal24     { return Decimal24(p) }
func (p Percentage16) Uint16() uint16         { return uint16(p) }
func (p Percentage16) Percent() float64       { return float64(p) / PercentRate }
func (p Percentage16) Fmt() string            { return Percentage(p).Fmt() }
func (p Percentage16) FmtAbs() string         { return Percentage(p).FmtAbs() }
