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
	DecimalAug                = 10000.0 // Decimal 表示万分之一
	Thousandth     Percentage = 10      // 千分比
	Percent        Percentage = 100     // 百分比
	ThousandthRate            = float32(1000.0)
	PercentRate               = float32(100.0)
)

func NewDecimal64(n int64) Decimal64  { return Decimal64(n) }
func ToDecimal64(n float64) Decimal64 { return Decimal64(n * DecimalAug) }
func (p Decimal64) Int() int          { return int(p) }
func (p Decimal64) Decimal() float64  { return float64(p) / DecimalAug }

// 整数部分
func (p Decimal64) Precision() int64 { return int64(p) / int64(DecimalAug) }

// 小数部分
func (p Decimal64) Scale() uint16 { return uint16(int64(math.Abs(float64(p))) % int64(DecimalAug)) }

// 类型：  1,000,000 这种
func (p Decimal64) FmtPrecision(n int, delimiter string) string {
	s := strconv.FormatInt(int64(p), 10)
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

func NewDecimal(n int) Decimal     { return Decimal(n) }
func ToDecimal(n float64) Decimal  { return Decimal(n * DecimalAug) }
func (p Decimal) Int() int         { return int(p) }
func (p Decimal) Decimal() float64 { return float64(p) / DecimalAug }
func (p Decimal) P() Decimal64     { return Decimal64(p) } // prototype

func NewDecimal24(n int32) Decimal24  { return Decimal24(n) }
func ToDecimal24(n float32) Decimal24 { return Decimal24(n * DecimalAug) }
func (p Decimal24) Int32() int32      { return int32(p) }
func (p Decimal24) Decimal() float32  { return float32(p) / DecimalAug }
func (p Decimal24) P() Decimal64      { return Decimal64(p) } // prototype

func NewDecimal16(n int16) Decimal16  { return Decimal16(n) }
func ToDecimal16(n float32) Decimal16 { return Decimal16(n * DecimalAug) }
func (p Decimal16) Int16() int16      { return int16(p) }
func (p Decimal16) Decimal() float32  { return float32(p) / DecimalAug }
func (p Decimal16) P() Decimal64      { return Decimal64(p) } // prototype

func NewPercentage(n int) Percentage { return Percentage(n) }
func ConvertPercent(n float32) Percentage {
	return NewPercentage(int(n * PercentRate))
}
func (p Percentage) Decimal() Decimal { return Decimal(p) }
func (p Percentage) Int() int         { return int(p) }
func (p Percentage) Percent() float32 { return float32(p) / PercentRate }
func (p Percentage) Fmt() string {
	return strconv.FormatFloat(float64(p.Percent()), 'f', -1, 32)
}
func (p Percentage) FmtAbs() string {
	c := p.Percent()
	if c < 0 {
		c = -c
	}
	return strconv.FormatFloat(float64(c), 'f', -1, 32)
}

// e.g. NewPercentage(100*atype.Percent)
func NewPercentage16(n uint16) Percentage16   { return Percentage16(n) }
func ConvertPercent16(n float32) Percentage16 { return NewPercentage16(uint16(n * PercentRate)) }
func (p Percentage16) Decimal() Decimal24     { return Decimal24(p) }
func (p Percentage16) Uint16() uint16         { return uint16(p) }
func (p Percentage16) Percent() float32       { return float32(p) / PercentRate }
func (p Percentage16) Fmt() string            { return Percentage(p).Fmt() }
func (p Percentage16) FmtAbs() string         { return Percentage(p).FmtAbs() }
