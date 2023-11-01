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

type Coin uint // 1 coin = 1 cent  = Money(100)  范围 4200万元左右

type Percent16 int16 // 需要转换为 Percent 使用；-327.68% ~ 327.67%  即 -3.2768 ~ 3.2767
type Percent24 Int24 // 需要转换为 Percent 使用； -83886.08% ~ 83886.07%   即 -838.8608 ~ -838.8607
type Percent int     // 范围： -21474836.48% - 21474836.47%

var PercentAug float64 = 100      // 扩大100 * 100倍 --> 这里按百分比算，而不是小数  3* Percent 为 3% = 0.03
var DecimalAug = PercentAug * 100 // 小数转百分比扩大100倍

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

)

// @param exchange 1 元可以兑换多少coin
func (a SmallMoney) ToCoin(exchange Coin) Coin { return a.ToCoin(exchange) }
func (a Money) ToCoin(exchange Coin) Coin      { return Coin(a/Yuan) * exchange }

// @param n 本身就是转换后的值，如10000，即表示为 100*PercentAug，即 100%
func NewPercent(n int) Percent { return Percent(n) }

// ToPercent(80.01) 表示 80.01%
func ToPercent(n float64) Percent { return NewPercent(int(n * PercentAug)) }

func NewPercent16(n int16) Percent16 { return Percent16(n) }
func (p Percent16) Percent() Percent { return Percent(p) }
func (p Percent16) Int16() int16     { return int16(p) }
func NewPercent24(n int32) Percent24 { return Percent24(n) }
func (p Percent24) Percent() Percent { return Percent(p) }
func (p Percent24) Int32() int32     { return int32(p) }

// 范围： -327.68% ~ 32767%  即 -3.2768 ~ +3.2767
func (p Percent) Int16() int16         { return int16(p) }
func (p Percent) Percent16() Percent16 { return Percent16(p) }
func (p Percent) Int32() int32         { return int32(p) }
func (p Percent) Percent24() Percent24 { return Percent24(p) }
func (p Percent) Int() int             { return int(p) }

func (p Percent) Percent() float64  { return float64(p.Int()) / PercentAug }
func (p Percent) Decimal() float64  { return float64(p.Int()) / DecimalAug }
func (p Percent) Mul(d int64) int64 { return d * int64(p) }
func (p Percent) Fmt() string       { return strconv.FormatFloat(p.Percent(), 'f', -1, 32) }
func (p Percent) FmtAbs() string {
	c := p.Percent()
	if c < 0 {
		c = -c
	}
	return strconv.FormatFloat(c, 'f', -1, 32)
}

// 采用四舍五入
func (a Money) MulPercent(p Percent) Money     { return a.Mul(int64(p)).Div(int64(DecimalAug)) }
func (a Money) MulPercentCeil(p Percent) Money { return a.Mul(int64(p)).DivCeil(int64(DecimalAug)) }
func (a Money) MulPercentFloor(p Percent) Money {
	return a.Mul(int64(p)).DivFloor(int64(DecimalAug))
}
func (a Money) MulPct(p float64) Money      { return a.MulPercent(ToPercent(p)) }
func (a Money) MulPctCeil(p float64) Money  { return a.MulPercentCeil(ToPercent(p)) }
func (a Money) MulPctFloor(p float64) Money { return a.MulPercentFloor(ToPercent(p)) }

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
func (a Money) Of(b Money) Percent {
	return NewPercent(int(a.Mul(int64(DecimalAug)).Div(b.Int64())))
}
func (a Money) OfFloor(b Money) Percent {
	return NewPercent(int(a.Mul(int64(DecimalAug)).DivFloor(b.Int64())))
}
func (a Money) OfCeil(b Money) Percent {
	return NewPercent(int(a.Mul(int64(DecimalAug)).DivCeil(b.Int64())))
}
