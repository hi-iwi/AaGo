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

	DecimalScale        uint8   = 4
	decimalUnitsInt64   int64   = 10000
	decimalUnitsFloat64 float64 = 10000.0
	DecimalUnits        Decimal = 10000

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
func DecimalUnitsX(n float64) Decimal { return Decimal(math.Round(n * decimalUnitsFloat64)) }

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
	return Decimal(intHandler(float64(p*b) / decimalUnitsFloat64))
}
func (p Decimal) MulRound(b Decimal) Decimal { return p.Mul(b, math.Round) }
func (p Decimal) MulCeil(b Decimal) Decimal  { return p.Mul(b, math.Ceil) }
func (p Decimal) MulFloor(b Decimal) Decimal { return p * b / DecimalUnits }

func (p Decimal) DivN(n int64, intHandler func(float64) float64) Decimal {
	return Decimal(intHandler(float64(p) / float64(n)))
}
func (p Decimal) DivRoundN(n int64) Decimal { return p.DivN(n, math.Round) }
func (p Decimal) DivCeilN(n int64) Decimal  { return p.DivN(n, math.Ceil) }
func (p Decimal) DivFloorN(n int64) Decimal { return p / Decimal(n) }

func (p Decimal) Div(b Decimal, intHandler func(float64) float64) Decimal {
	return Decimal(intHandler(float64(p) * decimalUnitsFloat64 / float64(b)))
}
func (p Decimal) DivRound(b Decimal) Decimal { return p.Div(b, math.Round) }
func (p Decimal) DivCeil(b Decimal) Decimal  { return p.Div(b, math.Ceil) }
func (p Decimal) DivFloor(b Decimal) Decimal { return p * DecimalUnits / b }
func (p Decimal) Sign() string {
	if p > 0 {
		return "-"
	}
	return ""
}

// 精度
func (p Decimal) Precision() int {
	n := len(strconv.FormatInt(p.Int64(), 10))
	if p < 0 {
		n--
	}
	return n
}

func (p Decimal) ToReal() float64 { return float64(p) / decimalUnitsFloat64 }

// 整数部分
func (p Decimal) Whole() int64 { return p.Int64() / decimalUnitsInt64 }

// 小数部分   -->  使用小数表示，不能用整数，因为  123.0001   ---> 0001
func (p Decimal) Mantissa(withSign bool) float64 {
	mantissa := float64(p%DecimalUnits) / decimalUnitsFloat64
	if !withSign && mantissa < 0 {
		mantissa = -mantissa
	}
	return mantissa
}

// FormatWhole 格式化整数部分
// 类型：  1,000,000 这种
func (p Decimal) FormatWhole(style *DecimalFormat) string {
	s := strconv.FormatInt(p.Whole(), 10)
	if style == nil {
		return s
	}
	return formatWhole(s, style.SegmentSize, style.Separator)
}

func mantissaOk(s string, scale int, trimScale bool) (string, bool) {
	if trimScale || scale == 0 {
		s = strings.TrimRight(s, "0")
	} else if len(s) < scale {
		s = padRight(s, "0", scale)
	}

	if s == "" {
		return "", false
	} else if scale == 0 || len(s) <= scale {
		return "." + s, false
	}
	return s, true
}
func (p Decimal) FormatMantissa(style *DecimalFormat) string {
	style = NewDecimalFormat(style)
	scale := int(style.Scale)

	s := strconv.FormatInt(p.Int64(), 10)
	g := len(s) - int(MoneyScale)
	if g > 0 {
		s = s[g:] // 取小数部分
	}

	var ok bool
	if s, ok = mantissaOk(s, scale, style.TrimScale); !ok {
		return s
	}

	if len(s) > scale {
		b := s[len(s)-1]
		if scale > 0 {
			b = s[scale-1]
		}
		k := style.ScaleRound.IsCeil() || (style.ScaleRound.IsRound() && strings.IndexByte("0123456789", b) > 4)
		s = s[:scale]
		if k {
			// 0.999... 就不用进位了
			repeat9 := true
			for _, ss := range s {
				if ss != '9' {
					repeat9 = false
					break
				}
			}
			if !repeat9 {
				x, _ := strconv.ParseUint("1"+s, 10, 16)
				s = strconv.FormatUint(x+1, 10)
			}
		}
		// s 发生变化
		if s, ok = mantissaOk(s, scale, style.TrimScale); !ok {
			return s
		}
	}
	return "." + s
}

func (p Decimal) Format(style *DecimalFormat) string {
	style = NewDecimalFormat(style)
	return p.FormatWhole(style) + p.FormatMantissa(style)
}

// 无参数，提供给 go template 使用
func (p Decimal) FmtWhole() string {
	return p.FormatWhole(nil)
}
func (p Decimal) FmtMantissa() string {
	return p.FormatMantissa(&DecimalFormat{Scale: 2})
}
func (p Decimal) Fmt() string {
	return p.Format(&DecimalFormat{Scale: 2})
}

func (p Decimal) Percent() float64 { return float64(p) / 100.0 }

func (p Decimal) FormatPercent() string {
	return strconv.FormatFloat(p.Percent(), 'f', -1, 32)
}
func (p Decimal) FormatPercentAbs() string {
	return strconv.FormatFloat(math.Abs(p.Percent()), 'f', -1, 32)
}
