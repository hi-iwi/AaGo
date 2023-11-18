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
type SmallMoney uint // 统一转换为 Money 后使用
type Money int64     // 有效范围：正负100亿元；  ±100 0000亿

type Coin Money // 1 coin = 1 money    如 chatgpt 等消耗，单次消耗低于0.1分，因此需要更大的 coin比例

const (
	Cent    Money = 100           // 分
	Dime          = 10 * Cent     // 角
	Yuan          = 10 * Dime     // 元
	KiYuan        = 1000 * Yuan   // 千元
	WanYuan       = 10000 * Yuan  // 万元
	MiYuan        = 100 * WanYuan // 百万元    中文的话，就不要用百万、千万
	//QianWanYuan  Money = 100000000000   // 千万元
	YiYuan = 100 * MiYuan // 亿元
	BiYuan = 10 * YiYuan  // 十亿元

	MinMoney = -100 * YiYuan // -100亿
	MaxMoney = 100 * YiYuan  // 100亿

	UnitCoin = Coin(Yuan)
)

// 金币抵扣商品
//  .Off()  ==  .Offset(100*Percent)
func (c Coin) Offset(rate Rate) Money { return Money(c).MulRateFloor(rate) }
func (c Coin) Off() Money             { return Money(c) }

// 用于使用计算
func (c Coin) Calc() Money     { return Money(c) }
func (a Money) CalcCoin() Coin { return Coin(a) }

// @param ratio 汇率
func (a Money) ExchangeCoin(rate Rate) Coin { return Coin(a.MulRateFloor(rate)) }

// 采用四舍五入
func (a Money) MulRate(p Rate) Money            { return a.Mul(int64(p)).Div(int64(DecimalAug)) }
func (a Money) MulRateCeil(p Rate) Money        { return a.Mul(int64(p)).DivCeil(int64(DecimalAug)) }
func (a Money) MulRateFloor(p Rate) Money       { return a.Mul(int64(p)).DivFloor(int64(DecimalAug)) }
func (a Money) MulPercent(p float64) Money      { return a.MulRate(ToPercent(p)) }
func (a Money) MulPercentCeil(p float64) Money  { return a.MulRateCeil(ToPercent(p)) }
func (a Money) MulPercentFloor(p float64) Money { return a.MulRateFloor(ToPercent(p)) }

func NewSmallMoney(n uint) SmallMoney  { return SmallMoney(n) }
func (a SmallMoney) Money() Money      { return Money(a) }
func (a SmallMoney) Uint() uint        { return uint(a) }
func NewMoney(m int64) Money           { return Money(m) }
func ToMoney(y float64) Money          { return Money(y * float64(Yuan)) }
func (a Money) Int64() int64           { return int64(a) }
func (a Money) Uint() uint             { return uint(a) }
func (a Money) SmallMoney() SmallMoney { return SmallMoney(a) }

// 整数部分
func (a Money) Precision() int64 { return int64(a) / int64(Yuan) }
func (a Money) ToCent() int64    { return int64(a) / int64(Cent) }

// 小数部分
func (a Money) Scale() uint16 { return uint16(int64(math.Abs(float64(a))) % int64(Yuan)) }
func decimalN(decimals ...uint16) uint16 {
	const d uint16 = 4 //  4位小数
	dec := uint16(2)   // 保留2位小数
	if len(decimals) > 0 {
		dec = decimals[0]
	}
	if dec > d {
		return d
	}
	return dec
}
func moneyDelimiter(delimiter ...string) string {
	sep := ","
	if len(delimiter) > 0 {
		sep = delimiter[0]
	}
	return sep
}

func formatScale(scale, decimal uint16, trim bool) string {
	const n uint16 = 4 //  4位小数
	if decimal == 0 || (trim && scale == 0) {
		return ""
	}

	x := math.Pow10(int(n - decimal))
	y := int(math.Floor(float64(scale) / x)) // 四舍五入是违法的，只能舍弃
	d := fmt.Sprintf("%0*d", decimal, y)
	if trim {
		d = strings.TrimRight(d, "0")
		if d == "" {
			return ""
		}
	}
	return "." + d
}

func fmtPrecision(s string, n int, delimiter string) string {
	if n == 0 || len(s) < n {
		return s
	}
	var s2 string
	j := 0
	for i := len(s) - 1; i > -1; i-- {
		if j > 0 && j%n == 0 {
			s2 = delimiter + s2
		}
		s2 = string(s[i]) + s2
		j++
	}
	return s2
}

// 类型：  1,000,000 这种
func (a Money) FmtPrecision(n int, delimiters ...string) string {
	s := strconv.FormatInt(int64(a), 10)
	sep := moneyDelimiter(delimiters...)
	return fmtPrecision(s, n, sep)
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

func (a Money) add(b Money) Money {
	// b 必须≥0， a可大于、等于、小于0
	if a < 0 || b < 0 || a > MaxMoney-b {
		panic(fmt.Sprintf("overflow money %d.add(%d)", a, b))
		return 0
	}
	return a + b
}

func (a Money) sub(b Money) Money {
	// b 必须≥0， a可大于、等于、小于0
	if a < 0 || b < 0 {
		panic(fmt.Sprintf("overflow money %d.sub(%d)", a, b))
		return 0
	}
	c := a - b
	if c < MinMoney || c > MaxMoney {
		panic(fmt.Sprintf("overflow money %d.sub(%d)", a, b))
		return 0
	}
	return c
}

func (a Money) Add(b Money) Money {
	if b < 0 {
		if a < 0 {
			return -(-a).add(-b) // a<0&&b<0 ==> -((-a)+(-b))
		}
		return a.sub(-b) // a>=0 && b<0  ==>  a-(-b)
	} else if a < 0 {
		return b.sub(-a) // a<0 && b>=0 ==>  b-(-a)
	}
	return a.add(b) // a>=0 && b >=0

}
func (a Money) Sub(b Money) Money {
	if b < 0 {
		if a < 0 {
			return (-b).sub(-a) // a<0&&b<0 ==> (-b)-(-a)
		}
		return a.add(-b) // a>=0&&b<0  ==> a+(-b)
	} else if a < 0 {
		return b.add(-a) // a<0 && b>=0 ==> -((-a)+b)
	}
	return a.sub(b)
}

// 必须大于0
func (a Money) AddN(b Money) Money {
	if a < -b {
		panic(fmt.Sprintf("overflow money %d.AddN(%d)", a, b))
		return 0
	}
	return a.Add(b)
}
func (a Money) SubN(b Money) Money {
	if a < b {
		panic(fmt.Sprintf("overflow money %d.SubN(%d)", a, b))
		return 0
	}
	return a.Sub(b)
}
func (a Money) Mul(n int64) Money {
	b := NewMoney(a.Int64() * n)
	if b < MinMoney || b > MaxMoney {
		panic(fmt.Sprintf("overflow money %d.Mul(%d)", a, n))
		return 0
	}
	return b
}

// 四舍五入
func (a Money) Div(n int64) Money {
	x := a.Int64()
	i := x / n
	m := a.Int64() % n
	if m > 0 && math.Round(float64(m)/float64(n)) == 1 {
		i++
	}
	b := NewMoney(i)
	if b < MinMoney || b > MaxMoney {
		panic(fmt.Sprintf("overflow money %d.Div(%d)", a, n))
		return 0
	}
	return b
}
func (a Money) DivFloor(n int64) Money {
	b := NewMoney(a.Int64() / n)
	if b < MinMoney || b > MaxMoney {
		panic(fmt.Sprintf("overflow money %d.DivFloor(%d)", a, n))
		return 0
	}
	return b
}
func (a Money) DivCeil(n int64) Money {
	x := a.Int64()
	i := x / n
	m := a.Int64() % n
	if m > 0 {
		i++
	}
	b := NewMoney(i)
	if b < MinMoney || b > MaxMoney {
		panic(fmt.Sprintf("overflow money %d.DivCeil(%d)", a, n))
		return 0
	}
	return b
}

// 四舍五入
func (a Money) Of(b Money) Rate {
	return NewRate(int(a.Mul(int64(DecimalAug)).Div(b.Int64())))
}
func (a Money) OfFloor(b Money) Rate {
	return NewRate(int(a.Mul(int64(DecimalAug)).DivFloor(b.Int64())))
}
func (a Money) OfCeil(b Money) Rate {
	return NewRate(int(a.Mul(int64(DecimalAug)).DivCeil(b.Int64())))
}
