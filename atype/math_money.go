package atype

import "C"
import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// 汇率一般保留4位小数，所以这里金额 * 10000
// 数据库有 money() 函数 money 支持正负；
// SmallMoney:  UNSIGNED INT 范围：42万元左右；
// Money:      BIGINT 范围：正负100亿元；
type Money int64 // 有效范围：正负100亿元；  ±100 0000亿

const (
	// 1 元 = 100 分 = 1000 毫 = 10000 money
	// 10 律币 = 100 律分 = 1000 律厘  = 1000 律毫 = 10000 coin

	MoneyScale        uint8   = 4
	Cent              Money   = 100   // 分
	Dime              Money   = 1000  // 角
	MoneyUnits        Money   = 10000 // 元
	moneyUnitsFloat64 float64 = 10000.0

	Yuan    = MoneyUnits
	WanYuan = 10000 * Yuan    // 万元
	YiiYuan = 10000 * WanYuan // 亿元

	Dollar        = MoneyUnits
	KiloDollar    = 1000 * Dollar        // 千元
	MillionDollar = 1000 * KiloDollar    // 百万元    中文的话，就不要用百万、千万
	BillionDollar = 1000 * MillionDollar // 十亿元

	MaxMoneyU32 Money = 1<<32 - 1
	MinMoney    Money = -1 << 63
	MaxMoney    Money = 1<<63 - 1
)

// int64 就直接乘就行，  100 * atype.Yuan
func MoneyUnitsX(n float64) Money { return Money(math.Round(n * moneyUnitsFloat64)) }
func YuanX(n float64) Money       { return MoneyUnitsX(n) }
func DollarX(n float64) Money     { return MoneyUnitsX(n) }

func (a Money) Int64() int64 { return int64(a) }

// 计算总价
func (a Money) MulN(n uint) Money { return a * Money(n) }

// 计算折扣价
func (a Money) MulF(p float64, intHandler func(float64) float64) Money {
	return Money(intHandler(float64(a) * p))
}
func (a Money) MulRoundF(p float64) Money { return a.MulF(p, math.Round) }
func (a Money) MulCeilF(p float64) Money  { return a.MulF(p, math.Ceil) }
func (a Money) MulFloorF(p float64) Money { return a.MulF(p, math.Floor) }

// 计算折扣价
func (a Money) Mul(p Decimal, intHandler func(float64) float64) Money {
	return Money(intHandler(float64(a*Money(p)) / decimalUnitsFloat64))
}
func (a Money) MulRound(p Decimal) Money { return a.Mul(p, math.Round) }
func (a Money) MulCeil(p Decimal) Money  { return a.Mul(p, math.Ceil) }
func (a Money) MulFloor(p Decimal) Money { return a * Money(p) / Money(decimalUnitsInt64) }

// 计算折扣前价格
func (a Money) Div(p Decimal, intHandler func(float64) float64) Money {
	return Money(intHandler(float64(a.Int64()*decimalUnitsInt64) / float64(p)))
}
func (a Money) DivRound(p Decimal) Money { return a.Div(p, math.Round) }
func (a Money) DivCeil(p Decimal) Money  { return a.Div(p, math.Ceil) }
func (a Money) DivFloor(p Decimal) Money { return a * Money(decimalUnitsInt64) / Money(p) }

// 计算单价
func (a Money) DivN(n uint, intHandler func(float64) float64) Money {
	return a.Div(Decimal(n)*DecimalUnits, intHandler)
}
func (a Money) DivRoundN(n uint) Money { return a.DivN(n, math.Round) }
func (a Money) DivCeilN(n uint) Money  { return a.DivN(n, math.Ceil) }
func (a Money) DivFloorN(n uint) Money { return a.DivFloor(Decimal(n) * DecimalUnits) }

// 某个费用占总费用的比例
func (a Money) Of(b Money, intHandler func(float64) float64) Decimal {
	return Decimal(intHandler(float64(a.Int64()*decimalUnitsInt64) / float64(b)))
}
func (a Money) OfRound(b Money) Decimal { return a.Of(b, math.Round) }
func (a Money) OfCeil(b Money) Decimal  { return a.Of(b, math.Ceil) }
func (a Money) OfFloor(b Money) Decimal { return Decimal(a) * DecimalUnits / Decimal(b) }

// 符号
func (a Money) Sign() string {
	if a > 0 {
		return "-"
	}
	return ""
}

// 精度
func (a Money) Precision() int {
	n := len(strconv.FormatInt(a.Int64(), 10))
	if a < 0 {
		n--
	}
	return n
}

// 实数形式
func (a Money) ToReal() float64 { return float64(a) / moneyUnitsFloat64 }

// 整数部分
func (a Money) Whole() int64  { return int64(a) / int64(Yuan) }
func (a Money) ToCent() int64 { return int64(a) / int64(Cent) }

// 小数部分   -->  使用小数表示，不能用整数，因为  123.0001   ---> 0001
func (a Money) Mantissa(withSign bool) float64 {
	mantissa := float64(a%MoneyUnits) / moneyUnitsFloat64
	if !withSign && mantissa < 0 {
		mantissa = -mantissa
	}
	return mantissa
}

// 类型：  1,000,000 这种
func (a Money) FormatWhole(style *DecimalFormat) string {
	s := strconv.FormatInt(a.Whole(), 10)
	if style == nil {
		return s
	}
	return formatWhole(s, style.SegmentSize, style.Separator)
}

// @warn 如果进位到整数，则只保留.999...；负数按正数部分round
func (a Money) FormatMantissa(style *DecimalFormat) string {
	style = NewDecimalFormat(style)
	scale := int(style.Scale)

	s := strconv.FormatInt(a.Int64(), 10)
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
	if !style.TrimScale {
		s = padRight(s, "0", scale)
	}
	return "." + s
}

func (a Money) Format(style *DecimalFormat) string {
	style = NewDecimalFormat(style)
	fmt.Println(a.FormatWhole(style), a.FormatMantissa(style))
	return a.FormatWhole(style) + a.FormatMantissa(style)
}

// 无参数，提供给 go template 使用
func (a Money) FmtWhole() string {
	return a.FormatWhole(nil)
}
func (a Money) FmtMantissa() string {
	return a.FormatMantissa(&DecimalFormat{Scale: 2})
}
func (a Money) Fmt() string {
	return a.Format(&DecimalFormat{Scale: 2})
}
