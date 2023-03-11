package atype

import "C"
import (
	"fmt"
	"github.com/shopspring/decimal"
	"math"
	"strconv"
	"strings"
)

// 汇率一般保留4位小数，所以这里金额 * 10000
// 数据库有 money() 函数 money 支持正负；
type Money int64 // 有效范围：正负100亿元；  ±100 0000亿

type Price Money // uint 范围：42万元左右； price 用 UMoney

type Percent int16                                // 范围：-10000~10000 => -100.00%~100.00%
var PercentMultiplier = decimal.NewFromInt32(100) // 扩大100 * 100倍 --> 这里按百分比算，而不是小数  3* Percent 为 3% = 0.03

const (
	Cent        Money = 100           // 分
	Dime              = 10 * Cent     // 角
	Yuan              = 10 * Dime     // 元
	KiloYuan          = 1000 * Yuan   // 千元
	WanYuan           = 10000 * Yuan  // 万元
	MillionYuan       = 100 * WanYuan // 百万元    中文的话，就不要用百万、千万
	//QianWanYuan  Money = 100000000000   // 千万元
	YiYuan      = 100 * MillionYuan // 亿元
	BillionYuan = 10 * YiYuan       // 十亿元

	MinMoney = -100 * YiYuan // -100亿
	MaxMoney = 100 * YiYuan  // 100亿

)

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
func (p Percent) Fmt() string              { return p.Percent().String() }
func (a Money) Dec() decimal.Decimal       { return decimal.NewFromInt(int64(a)) }
func (a Money) MulPercent(p Percent) Money { return Money(p.Mul(a.Dec()).IntPart()) }

func NewMoney(m int64) Money  { return Money(m) }
func ToMoney(y float64) Money { return Money(y * float64(Yuan)) }
func (a Money) Int64() int64  { return int64(a) }
func (a Money) Price() Price  { return Price(a) }
func (p Price) Money() Money  { return Money(p) }

// 整数部分
func (a Money) Precision() int64 { return int64(a) / int64(Yuan) }

// 小数部分
func (a Money) Scale() uint16 { return uint16(int64(math.Abs(float64(a))) % int64(Yuan)) }
func decimalN(decimals ...uint16) uint16 {
	const d uint16 = 4   //  4位小数
	decimal := uint16(2) // 保留2位小数
	if len(decimals) > 0 {
		decimal = decimals[0]
	}
	if decimal > d {
		return d
	}
	return decimal
}
func moneydelimiter(delimiter ...string) string {
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
	sep := moneydelimiter(delimiters...)
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

func (a Money) minus(b Money) Money {
	// b 必须≥0， a可大于、等于、小于0
	if a < 0 || b < 0 {
		panic(fmt.Sprintf("overflow money %d.minus(%d)", a, b))
		return 0
	}
	c := a - b
	if c < MinMoney || c > MaxMoney {
		panic(fmt.Sprintf("overflow money %d.minus(%d)", a, b))
		return 0
	}
	return c
}

func (a Money) Add(b Money) Money {
	if b < 0 {
		if a < 0 {
			return -(-a).add(-b) // a<0&&b<0 ==> -((-a)+(-b))
		}
		return a.minus(-b) // a>=0 && b<0  ==>  a-(-b)
	} else if a < 0 {
		return b.minus(-a) // a<0 && b>=0 ==>  b-(-a)
	}
	return a.add(b) // a>=0 && b >=0

}
func (a Money) Minus(b Money) Money {
	if b < 0 {
		if a < 0 {
			return (-b).minus(-a) // a<0&&b<0 ==> (-b)-(-a)
		}
		return a.add(-b) // a>=0&&b<0  ==> a+(-b)
	} else if a < 0 {
		return b.add(-a) // a<0 && b>=0 ==> -((-a)+b)
	}
	return a.minus(b)
}

// 必须大于0
func (a Money) AddN(b Money) Money {
	if a < -b {
		panic(fmt.Sprintf("overflow money %d.AddN(%d)", a, b))
		return 0
	}
	return a.Add(b)
}
func (a Money) MinusN(b Money) Money {
	if a < b {
		panic(fmt.Sprintf("overflow money %d.MinusN(%d)", a, b))
		return 0
	}
	return a.Minus(b)
}
