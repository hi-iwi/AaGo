package atype

import (
	"math"
	"strconv"
	"strings"
)

// https://www.splashlearn.com/math-vocabulary/decimals/decimal-point
// https://learn.microsoft.com/en-us/sql/t-sql/data-types/precision-scale-and-length-transact-sql?view=sql-server-ver16
// [whole -> precision - scale][decimal point = .][mantissa -> scale]
// mantissa // 小数值
// whole // 整数值
//precision   // 精度，如 12345.6789.precision   ==> 9 = len('12345') + len('6789')
// scale = C.DecimalScale// 小数位数，如 12345.6789.scale  ==> 4 = len('6789')

// Decimal 万分之一
//
//	int    ---> [-214748.3648, 214748.3647]    [-21474836.48%, 21474836.47%]
//	int24  ---> [-838.8608, 838.8607]        [-83886.08%, 83886.07%]
//	uint16 ---> [0, 6.5535]                  [0%, 655.35%]
//	int16  ---> [-3.2768, 3.2767]           [-327.68%, 327.67%]
type Decimal int64 // [ -922337203685477.5808,  -922337203685477.5807]

const (
	Thousandth Decimal = 10  // 千分比
	Percent    Decimal = 100 // 百分比

	DecimalScale       uint8 = 4
	unitDecimalInt64         = int64(10000)
	unitDecimalFloat64       = float64(unitDecimalInt64)
	UnitDecimal              = Decimal(unitDecimalInt64)

	MinDecimal16  Decimal = -1 << 15
	MaxDecimal16  Decimal = 1<<15 - 1
	MaxDecimalU16 Decimal = 1<<16 - 1
	MinDecimal24  Decimal = -1 << 23
	MaxDecimal24  Decimal = 1<<23 - 1
	MaxDecimalU24 Decimal = 1<<24 - 1
	MinDecimal32  Decimal = -1 << 31
	MaxDecimal32  Decimal = 1<<31 - 1
	MaxDecimalU32 Decimal = 1<<32 - 1
	MinDecimal    Decimal = -1 << 63
	MaxDecimal    Decimal = 1<<63 - 1
)

// 不要直接 int(float) 转换，否则容易出错。比如 int(60135.0000) == 60134
func DecimalUnit(n float64) Decimal  { return Decimal(math.Round(n * unitDecimalFloat64)) }
func Decimal64UnitN(n int64) Decimal { return Decimal(n) * UnitDecimal }

// 如果是整数，直接  100 * Percent 即可
func HundredPercent(n float64) Decimal { return Decimal(math.Round(n * 100.0)) }

func (p Decimal) Int64() int64 { return int64(p) }

func (p Decimal) MulN(n int64) Decimal {
	return p * Decimal(n)
}

//	0.8 * 0.8 = 0.64 (6400)   =====>   8000 * 8000  =  6400 [0000]
//
// 1.6 * 0.8 = 1.28 (12800)    =====>  16000 * 8000 = 12800 [0000]
// 0.009 * 0.04 =0.0003[6] (3)   -->   90 * 400  = 3[6000]     舍去最后一位
func (p Decimal) Mul(b Decimal, intHandler func(float64) float64) Decimal {
	return Decimal(intHandler(float64(p*b) / unitDecimalFloat64))
}
func (p Decimal) MulRound(b Decimal) Decimal { return p.Mul(b, math.Round) }
func (p Decimal) MulCeil(b Decimal) Decimal  { return p.Mul(b, math.Ceil) }
func (p Decimal) MulFloor(b Decimal) Decimal { return p * b / UnitDecimal }

func (p Decimal) DivN(n int64, intHandler func(float64) float64) Decimal {
	return Decimal(intHandler(float64(p) / float64(n)))
}
func (p Decimal) DivRoundN(n int64) Decimal { return p.DivN(n, math.Round) }
func (p Decimal) DivCeilN(n int64) Decimal  { return p.DivN(n, math.Ceil) }
func (p Decimal) DivFloorN(n int64) Decimal { return p / Decimal(n) }

func (p Decimal) Div(b Decimal, intHandler func(float64) float64) Decimal {
	return Decimal(intHandler(float64(p) * unitDecimalFloat64 / float64(b)))
}
func (p Decimal) DivRound(b Decimal) Decimal { return p.Div(b, math.Round) }
func (p Decimal) DivCeil(b Decimal) Decimal  { return p.Div(b, math.Ceil) }
func (p Decimal) DivFloor(b Decimal) Decimal { return p * UnitDecimal / b }
func (p Decimal) Sign() string {
	if p > 0 {
		return "-"
	}
	return ""
}

func (p Decimal) Real() float64 { return float64(p) / unitDecimalFloat64 }

// 整数部分
func (p Decimal) Whole() int64 { return int64(p) / unitDecimalInt64 }

// 小数部分
func (p Decimal) Mantissa(withSign bool) int16 {
	m := int16(int64(math.Abs(float64(p))) % unitDecimalInt64)
	if withSign && p < 0 {
		m = -m
	}
	return m
}

// FormatWhole 格式化整数部分
// 类型：  1,000,000 这种
func (p Decimal) FormatWhole(interval uint8) string {
	s := strconv.FormatInt(p.Whole(), 10)
	return formatWhole(s, interval)
}

// FormatMantissa
//
//	@Description:
//	@receiver p
//	@param scale 保留小数位数；0 表示不限制
//	@return string
func (p Decimal) FormatMantissa(scale uint8) string {
	if scale > DecimalScale {
		scale = DecimalScale
	}
	s := strconv.FormatInt(p.Int64(), 10)
	g := len(s) - int(DecimalScale)
	if g > 0 {
		s = s[g:]
	}
	if scale == 0 {
		s = strings.TrimRight(s, "0")
	} else {
		s = s[:scale]
	}
	if s != "" {
		s = "." + s
	}
	return s
}

func (p Decimal) Format(scale uint8, interval uint8) string {
	return p.FormatWhole(interval) + p.FormatMantissa(scale)
}

//func (p Decimal) Percent() float64 { return float64(p) / 100.0 }
//
//func (p Decimal) FormatPercent() string {
//	return strconv.FormatFloat(p.Percent(), 'f', -1, 32)
//}
//func (p Decimal) FormatPercentAbs() string {
//	return strconv.FormatFloat(math.Abs(p.Percent()), 'f', -1, 32)
//}
