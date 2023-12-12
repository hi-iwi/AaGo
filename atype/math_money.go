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
type SmallMoney uint // 统一转换为 Money 后使用
type Money int64     // 有效范围：正负100亿元；  ±100 0000亿

type Coin Money // 1 coin = 1 money    如 chatgpt 等消耗，单次消耗低于0.1分，因此需要更大的 coin比例

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

	MinMoney = -100 * YiiYuan // -100亿
	MaxMoney = 100 * YiiYuan  // 100亿

	UnitSmallMoney = SmallMoney(UnitMoney)

	CentCoin        Coin = 100  // 1 律分
	DimeCoin        Coin = 1000 // 1 律币(角)
	UnitCoin        Coin = 10000
	unitCoinFloat64      = 10000.0
)

func MoneyUnit(n float64) Money { return Money(math.Round(n * unitMoneyFloat64)) }
func MoneyUnitN(n int) Money    { return Money(n) * UnitMoney }
func Yuans(n float64) Money     { return MoneyUnit(n) }
func YuanN(n int) Money         { return MoneyUnitN(n) }
func Dollars(n float64) Money   { return MoneyUnit(n) }
func DollarN(n int) Money       { return MoneyUnitN(n) }

func (a Money) Int64() int64      { return int64(a) }
func (a Money) Uint() uint        { return uint(a) }
func (a Money) Small() SmallMoney { return SmallMoney(a) }

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

func (a Money) Mul(n int64) Money { return a * Money(n) }

// 采用四舍五入
func (a Money) MulDecimal(p Decimal) Money {
	return Money(math.Round(float64(a*Money(p)) / unitDecimalFloat64))
}
func (a Money) MulDecimalCeil(p Decimal) Money {
	return Money(math.Ceil(float64(a*Money(p)) / unitDecimalFloat64))
}
func (a Money) MulDecimalFloor(p Decimal) Money { return a * Money(p) / Money(unitDecimalInt64) }

func (a Money) MulPercent(p float64) Money {
	return Money(math.Round(float64(a*Money(HundredPercent(p))) / unitDecimalFloat64))
}
func (a Money) MulPercentCeil(p float64) Money {
	return Money(math.Ceil(float64(a*Money(HundredPercent(p))) / unitDecimalFloat64))
}
func (a Money) MulPercentFloor(p float64) Money {
	return a * Money(HundredPercent(p)) / Money(unitDecimalInt64)
}
func (a Money) Div(n float64) Money      { return Money(math.Round(float64(a) / n)) }
func (a Money) DivFloor(n float64) Money { return Money(math.Floor(float64(a) / n)) }
func (a Money) DivCeil(n float64) Money  { return Money(math.Ceil(float64(a) / n)) }
func (a Money) Of(b Money) Decimal       { return Decimal(a.Mul(unitDecimalInt64).Div(float64(b))) }
func (a Money) OfFloor(b Money) Decimal  { return Decimal(a.Mul(unitDecimalInt64).DivFloor(float64(b))) }
func (a Money) OfCeil(b Money) Decimal   { return Decimal(a.Mul(unitDecimalInt64).DivCeil(float64(b))) }

func (a SmallMoney) P() Money   { return Money(a) } // prototype
func (a SmallMoney) Uint() uint { return uint(a) }

func CoinUnit(n float64) Coin { return Coin(math.Round(n * unitCoinFloat64)) }
func CoinUnitN(n int) Coin    { return Coin(n) * UnitCoin }

// 金币抵扣商品
//  .Off()  ==  .Offset(100*Percent)
func (c Coin) Offset(rate Decimal) Money { return Money(c).MulDecimalFloor(rate) }
func (c Coin) Off() Money                { return Money(c) }

// 用于使用计算
func (c Coin) StartCalc() Money { return Money(c) }
func (a Money) EndCalc() Coin   { return Coin(a) }

// @param ratio 汇率
func (a Money) ExchangeCoin(rate Decimal) Coin { return Coin(a.MulDecimalFloor(rate)) }
