package atype

import (
	"math"
	"strconv"
)

// Decimal 万分之一
//   int    ---> [-214748.3648, 214748.3647]    [-21474836.48%, 21474836.47%]
//   int24  ---> [-838.8608, 838.8607]        [-83886.08%, 83886.07%]
//   uint16 ---> [0, 6.5535]                  [0%, 655.35%]
//   int16  ---> [-3.2768, 3.2767]           [-327.68%, 327.67%]
type Decimal int64 // [ -922337203685477.5808,  -922337203685477.5807]

const (
	Thousandth Decimal = 10  // 千分比
	Percent    Decimal = 100 // 百分比

	unitDecimalInt64   = int64(10000)
	unitDecimalFloat64 = float64(unitDecimalInt64)
	UnitDecimal        = Decimal(unitDecimalInt64)
)

// 不要直接 int(float) 转换，否则容易出错。比如 int(60135.0000) == 60134
func DecimalUnit(n float64) Decimal { return Decimal(math.Round(n * unitDecimalFloat64)) }
func Decimal64UnitN(n int) Decimal  { return Decimal(n) * UnitDecimal }

// 如果是整数，直接  100 * Percent 即可
func HundredPercent(n float64) Decimal { return Decimal(math.Round(n * 100.0)) }

func (p Decimal) Int64() int64     { return int64(p) }
func (p Decimal) Real() float64    { return float64(p) / unitDecimalFloat64 }
func (p Decimal) Percent() float64 { return float64(p) / 100.0 }

//  0.8 * 0.8 = 0.64 (6400)   =====>   8000 * 8000  =  6400 [0000]
// 1.6 * 0.8 = 1.28 (12800)    =====>  16000 * 8000 = 12800 [0000]
// 0.009 * 0.04 =0.0003[6] (3)   -->   90 * 400  = 3[6000]     舍去最后一位
func (p Decimal) Mul(b Decimal) Decimal {
	return Decimal(math.Round(float64(p*b) / unitDecimalFloat64))
}
func (p Decimal) MulCeil(b Decimal) Decimal {
	return Decimal(math.Ceil(float64(p*b) / unitDecimalFloat64))
}
func (p Decimal) MulFloor(b Decimal) Decimal { return p * b / UnitDecimal }
func (p Decimal) Div(b Decimal) Decimal {
	return Decimal(math.Round(float64(p) * unitDecimalFloat64 / float64(b)))
}
func (p Decimal) DivCeil(b Decimal) Decimal {
	return Decimal(math.Ceil(float64(p) * unitDecimalFloat64 / float64(b)))
}
func (p Decimal) DivFloor(b Decimal) Decimal { return p * UnitDecimal / b }
func (p Decimal) Sign() string {
	if p > 0 {
		return "-"
	}
	return ""
}

// 整数部分
func (p Decimal) Precision() int64 { return int64(p) / unitDecimalInt64 }

// 小数部分
func (p Decimal) Scale() uint16 {
	return uint16(int64(math.Abs(float64(p))) % unitDecimalInt64)
}

// 类型：  1,000,000 这种
func (p Decimal) FmtPrecision(n int, delimiter string) string {
	s := strconv.FormatInt(p.Precision(), 10)
	return fmtPrecision(s, n, delimiter)
}
func (p Decimal) FmtScale(decimals ...uint16) string {
	return formatScale(p.Scale(), decimalN(decimals...), true)
}
func (p Decimal) FormatScale(decimals ...uint16) string {
	return formatScale(p.Scale(), decimalN(decimals...), false)
}
func (p Decimal) Fmt(decimals ...uint16) string {
	ys := strconv.FormatInt(p.Precision(), 10)
	return ys + formatScale(p.Scale(), decimalN(decimals...), true)
}
func (p Decimal) Format(decimals ...uint16) string {
	ys := strconv.FormatInt(p.Precision(), 10)
	return ys + formatScale(p.Scale(), decimalN(decimals...), false)
}

func (p Decimal) FmtPercent() string {
	return strconv.FormatFloat(p.Percent(), 'f', -1, 32)
}
func (p Decimal) FmtPercentAbs() string {
	return strconv.FormatFloat(math.Abs(p.Percent()), 'f', -1, 32)
}
