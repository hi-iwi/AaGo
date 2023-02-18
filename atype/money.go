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
type Money int      // 范围：±21万元
type Umoney uint    // 范围：42万元左右； price 用 UMoney
type Amount int64   // 范围：±900万亿元。      不用于计算
type Uamount uint64 //     不用于计算

const (
	Cent         Money = 100         // 分
	Dime         Money = 1000        // 角
	Yuan         Money = 10000       // 元
	ThousandYuan Money = 10000000    // 千元
	WanYuan      Money = 100000000   // 万元
	MillionYuan  Money = 10000000000 // 百万元    中文的话，就不要用百万、千万
	//QianWanYuan  Money = 100000000000   // 千万元
	YiYuan      Money = 1000000000000  // 亿元
	BillionYuan Money = 10000000000000 // 十亿元，就不要用十亿，而是 10 * YiYuan

)

func ToYuan(y float64) Money {
	return Money(y * float64(Yuan))
}
func NewYuan(y float64) Umoney {
	return Umoney(y * float64(Yuan))
}

func NewAmount(m int64) Amount {
	return Amount(m)
}

func (a Amount) Int64() int64 {
	return int64(a)
}

// 整数部分
func (a Amount) Precision() int64 {
	return int64(a) / int64(Yuan)
}

// 小数部分
func (a Amount) Scale() uint16 {
	return uint16(int64(math.Abs(float64(a))) % int64(Yuan))
}
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
	if trim && (decimal == 0 || scale == 0) {
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
func (a Amount) FmtPrecision(n int, delimiters ...string) string {
	s := strconv.FormatInt(int64(a), 10)
	sep := moneydelimiter(delimiters...)
	return fmtPrecision(s, n, sep)
}
func (a Amount) FmtScale(decimals ...uint16) string {
	return formatScale(a.Scale(), decimalN(decimals...), true)
}
func (a Amount) FormatScale(decimals ...uint16) string {
	return formatScale(a.Scale(), decimalN(decimals...), false)
}
func (a Amount) Fmt(decimals ...uint16) string {
	ys := strconv.FormatInt(a.Precision(), 10)
	return ys + formatScale(a.Scale(), decimalN(decimals...), true)
}
func (a Amount) Format(decimals ...uint16) string {
	ys := strconv.FormatInt(a.Precision(), 10)
	return ys + formatScale(a.Scale(), decimalN(decimals...), false)
}

func NewUamount(m uint64) Uamount {
	return Uamount(m)
}

func (a Uamount) Uint64() uint64 {
	return uint64(a)
}

// 整数部分
func (a Uamount) Precision() uint64 {
	return a.Uint64() / uint64(Yuan)
}

// 小数部分
func (a Uamount) Scale() uint16 {
	return uint16(a.Uint64() % uint64(Yuan))
}

// 类型：  1,000,000 这种
func (a Uamount) FmtPrecision(n int, delimiters ...string) string {
	s := strconv.FormatUint(uint64(a), 10)
	sep := moneydelimiter(delimiters...)
	return fmtPrecision(s, n, sep)
}

// 不保留小数尾部的0
func (a Uamount) Fmt(decimals ...uint16) string {
	ys := strconv.FormatUint(a.Precision(), 10)
	return ys + formatScale(a.Scale(), decimalN(decimals...), true)
}

// 保留小数尾部0
func (a Uamount) Format(decimals ...uint16) string {
	ys := strconv.FormatUint(a.Precision(), 10)
	return ys + formatScale(a.Scale(), decimalN(decimals...), false)
}

func NewMoney(m int) Money {
	return Money(m)
}

func (a Money) Int() int {
	return int(a)
}

// 整数部分
func (a Money) Precision() int {
	return int(a) / int(Yuan)
}

// 小数部分
func (a Money) Scale() uint16 {
	return uint16(int64(math.Abs(float64(a))) % int64(Yuan))
}

// 类型：  1,000,000 这种
func (a Money) FmtPrecision(n int, delimiters ...string) string {
	s := strconv.FormatInt(int64(a), 10)
	sep := moneydelimiter(delimiters...)
	return fmtPrecision(s, n, sep)
}

// 不保留小数尾部的0
func (a Money) Fmt(decimals ...uint16) string {
	ys := strconv.Itoa(a.Precision())
	return ys + formatScale(a.Scale(), decimalN(decimals...), true)
}

// 保留小数尾部的0
func (a Money) Format(decimals ...uint16) string {
	ys := strconv.Itoa(a.Precision())
	return ys + formatScale(a.Scale(), decimalN(decimals...), false)
}

func NewUmoney(m uint) Umoney {
	return Umoney(m)
}

func (a Umoney) Uint() uint {
	return uint(a)
}

// 整数部分
func (a Umoney) Precision() uint {
	return a.Uint() / uint(Yuan)
}

// 小数部分
func (a Umoney) Scale() uint16 {
	return uint16(a.Uint() % uint(Yuan))
}

// 类型：  1,000,000 这种
func (a Umoney) FmtPrecision(n int, delimiter ...string) string {
	s := strconv.FormatUint(uint64(a), 10)
	sep := moneydelimiter(delimiter...)
	return fmtPrecision(s, n, sep)
}

// 不保留小数尾部的0
func (a Umoney) Fmt(decimals ...uint16) string {
	ys := strconv.FormatUint(uint64(a.Precision()), 10)
	return ys + formatScale(a.Scale(), decimalN(decimals...), true)
}

// 保留小数尾部的0
func (a Umoney) Format(decimals ...uint16) string {
	ys := strconv.FormatUint(uint64(a.Precision()), 10)
	return ys + formatScale(a.Scale(), decimalN(decimals...), false)
}
