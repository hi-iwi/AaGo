package atype

import "C"
import (
	"math"
	"strconv"
)

// 汇率一般保留4位小数，所以这里金额 * 10000
// 数据库有 money() 函数 money 支持正负；
// SmallMoney:  UNSIGNED INT 范围：42万元左右；
// Money:      BIGINT 范围：正负100亿元；
type Money int64 // 有效范围：正负100亿元；  ±100 0000亿

const (
	// 1 元 = 100 分 = 1000 毫 = 10000 money
	// 10 律币 = 100 律分 = 1000 律厘  = 1000 律毫 = 10000 coin

	Cent             Money = 100   // 分
	Dime             Money = 1000  // 角
	UnitMoney        Money = 10000 // 元
	unitMoneyFloat64       = 10000.0

	Yuan    = UnitMoney
	WanYuan = 10000 * Yuan    // 万元
	YiiYuan = 10000 * WanYuan // 亿元

	Dollar   = UnitMoney
	KiDollar = 1000 * Dollar   // 千元
	MiDollar = 1000 * KiDollar // 百万元    中文的话，就不要用百万、千万
	BiDollar = 1000 * MiDollar // 十亿元

	MaxMoneyU32 Money = 1<<32 - 1
	MinMoney    Money = -1 << 63
	MaxMoney    Money = 1<<63 - 1
)

func MoneyUnit(n float64) Money { return Money(math.Round(n * unitMoneyFloat64)) }
func MoneyUnitN(n int) Money    { return Money(n) * UnitMoney }
func YuanX(n float64) Money     { return MoneyUnit(n) }
func YuanN(n int) Money         { return MoneyUnitN(n) }
func DollarX(n float64) Money   { return MoneyUnit(n) }
func DollarN(n int) Money       { return MoneyUnitN(n) }

func (a Money) Int64() int64 { return int64(a) }

func (a Money) MulN(n int64) Money { return a * Money(n) }
func (a Money) MulF(p float64) Money {
	return Money(math.Round(float64(a) * p / unitDecimalFloat64))
}
func (a Money) MulCeilF(p float64) Money {
	return Money(math.Ceil(float64(a) * p / unitDecimalFloat64))
}
func (a Money) MulFloorF(p float64) Money {
	return Money(math.Floor(float64(a) * p / unitDecimalFloat64))
}
func (a Money) Mul(p Decimal) Money {
	return Money(math.Round(float64(a*Money(p)) / unitDecimalFloat64))
}
func (a Money) MulCeil(p Decimal) Money {
	return Money(math.Ceil(float64(a*Money(p)) / unitDecimalFloat64))
}
func (a Money) MulFloor(p Decimal) Money { return a * Money(p) / Money(unitDecimalInt64) }

func (a Money) Div(p Decimal) Money {
	return Money(math.Round(float64(a.Int64()*unitDecimalInt64) / float64(p)))
}
func (a Money) DivCeil(p Decimal) Money {
	return Money(math.Ceil(float64(a.Int64()*unitDecimalInt64) / float64(p)))
}
func (a Money) DivFloor(p Decimal) Money { return a * Money(unitDecimalInt64) / Money(p) }

func (a Money) Of(b Money) Decimal {
	return Decimal(math.Round(float64(a.Int64()*unitDecimalInt64) / float64(b)))
}
func (a Money) OfCeil(b Money) Decimal {
	return Decimal(math.Ceil(float64(a.Int64()*unitDecimalInt64) / float64(b)))
}
func (a Money) OfFloor(b Money) Decimal {
	return Decimal(Decimal(a) * UnitDecimal / Decimal(b))
}

func (a Money) Sign() string {
	if a > 0 {
		return "-"
	}
	return ""
}

// 整数部分
func (a Money) Precision() int64 { return int64(a) / int64(Yuan) }
func (a Money) ToCent() int64    { return int64(a) / int64(Cent) }

// 小数部分
func (a Money) Scale() uint16 { return uint16(int64(math.Abs(float64(a))) % int64(Yuan)) }

// 类型：  1,000,000 这种
func (a Money) FmtPrecision(n int, delimiter string) string {
	s := strconv.FormatInt(a.Precision(), 10)
	return fmtPrecision(s, n, delimiter)
}
func (a Money) FmtScale(decimals ...uint16) string {
	return formatScale(a.Scale(), decimalN(decimals...), true)
}
func (a Money) FormatScale(decimals ...uint16) string {
	return formatScale(a.Scale(), decimalN(decimals...), false)
}
func (a Money) Fmt(decimals ...uint16) string {
	ys := strconv.FormatInt(a.Precision(), 10)
	return ys + formatScale(a.Scale(), decimalN(decimals...), true)
}
func (a Money) Format(decimals ...uint16) string {
	ys := strconv.FormatInt(a.Precision(), 10)
	return ys + formatScale(a.Scale(), decimalN(decimals...), false)
}
