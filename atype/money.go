package atype

import (
	"fmt"
	"math"
	"strconv"
)

// 汇率一般保留4位小数，所以这里金额 * 10000
// 数据库有 money() 函数 money 支持正负；
type Money int    // 范围：±21万元
type Umoney uint  // 范围：42万元左右； price 用 UMoney
type Amount int64 // 范围：±900万亿元。
type Uamount uint64

func (a Amount) Int64() int64 {
	return int64(a)
}

// 整数部分
func (a Amount) Precision() int64 {
	return int64(a) / 10000
}

// 小数部分
func (a Amount) Scale() uint16 {
	return uint16(int64(math.Abs(float64(a))) % 10000)
}
func decimal(decimals ...uint16) uint16 {
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
func moneySeperator(seperator ...string) string {
	sep := ","
	if len(seperator) > 0 {
		sep = seperator[0]
	}
	return sep
}
func fmtScale(scale, decimal uint16, trim bool) string {
	const d uint16 = 4 //  4位小数
	if trim && (decimal == 0 || scale == 0) {
		return ""
	}

	x := math.Pow10(int(d - decimal))
	y := int(math.Floor(float64(scale) / x)) // 四舍五入是违法的，只能舍弃
	if trim && (y == 0) {
		return ""
	}
	return "." + fmt.Sprintf("%0*d", decimal, y)
}
func fmtPrecision(s string, n int, seperator string) string {
	if n == 0 || len(s) < n {
		return s
	}
	var s2 string
	j := 0
	for i := len(s) - 1; i > -1; i-- {
		if j > 0 && j%n == 0 {
			s2 = seperator + s2
		}
		s2 = string(s[i]) + s2
		j++
	}
	return s2
}

// 类型：  1,000,000 这种
func (a Amount) FmtPrecision(n int, seperators ...string) string {
	s := strconv.FormatInt(int64(a), 10)
	sep := moneySeperator(seperators...)
	return fmtPrecision(s, n, sep)
}
func (a Amount) FmtScale(decimals ...uint16) string {
	return fmtScale(a.Scale(), decimal(decimals...), true)
}
func (a Amount) FormatScale(decimals ...uint16) string {
	return fmtScale(a.Scale(), decimal(decimals...), false)
}
func (a Amount) Fmt(decimals ...uint16) string {
	ys := strconv.FormatInt(a.Precision(), 10)
	return ys + fmtScale(a.Scale(), decimal(decimals...), true)
}
func (a Amount) Format(decimals ...uint16) string {
	ys := strconv.FormatInt(a.Precision(), 10)
	return ys + fmtScale(a.Scale(), decimal(decimals...), false)
}

func (a Uamount) Uint64() uint64 {
	return uint64(a)
}

// 整数部分
func (a Uamount) Precision() uint64 {
	return uint64(a) / 10000
}

// 小数部分
func (a Uamount) Scale() uint16 {
	return uint16(a % 10000)
}

// 类型：  1,000,000 这种
func (a Uamount) FmtPrecision(n int, seperators ...string) string {
	s := strconv.FormatUint(uint64(a), 10)
	sep := moneySeperator(seperators...)
	return fmtPrecision(s, n, sep)
}
func (a Uamount) Fmt(decimals ...uint16) string {
	decimal := uint16(2) // 保留2位小数
	if len(decimals) > 0 {
		decimal = decimals[0]
	}
	ys := strconv.FormatUint(a.Precision(), 10)
	return ys + fmtScale(a.Scale(), decimal, true)
}
func (a Uamount) Format(decimals ...uint16) string {
	decimal := uint16(2) // 保留2位小数
	if len(decimals) > 0 {
		decimal = decimals[0]
	}
	ys := strconv.FormatUint(a.Precision(), 10)
	return ys + fmtScale(a.Scale(), decimal, false)
}

func (a Money) Int() int {
	return int(a)
}

// 整数部分
func (a Money) Precision() int {
	return int(a) / 10000
}

// 小数部分
func (a Money) Scale() uint16 {
	return uint16(int64(math.Abs(float64(a))) % 10000)
}

// 类型：  1,000,000 这种
func (a Money) FmtPrecision(n int, seperators ...string) string {
	s := strconv.FormatInt(int64(a), 10)
	sep := moneySeperator(seperators...)
	return fmtPrecision(s, n, sep)
}
func (a Money) Fmt(decimals ...uint16) string {
	decimal := uint16(2) // 保留2位小数
	if len(decimals) > 0 {
		decimal = decimals[0]
	}
	ys := strconv.Itoa(a.Precision())
	return ys + fmtScale(a.Scale(), decimal, true)
}
func (a Money) Format(decimals ...uint16) string {
	decimal := uint16(2) // 保留2位小数
	if len(decimals) > 0 {
		decimal = decimals[0]
	}
	ys := strconv.Itoa(a.Precision())
	return ys + fmtScale(a.Scale(), decimal, false)
}

func (a Umoney) Uint() uint {
	return uint(a)
}

// 整数部分
func (a Umoney) Precision() uint {
	return uint(a) / 10000
}

// 小数部分
func (a Umoney) Scale() uint16 {
	return uint16(a % 10000)
}

// 类型：  1,000,000 这种
func (a Umoney) FmtPrecision(n int, seperators ...string) string {
	s := strconv.FormatUint(uint64(a), 10)
	sep := moneySeperator(seperators...)
	return fmtPrecision(s, n, sep)
}
func (a Umoney) Fmt(decimals ...uint16) string {
	decimal := uint16(2) // 保留2位小数
	if len(decimals) > 0 {
		decimal = decimals[0]
	}
	ys := strconv.FormatUint(uint64(a.Precision()), 10)
	return ys + fmtScale(a.Scale(), decimal, true)
}
func (a Umoney) Format(decimals ...uint16) string {
	decimal := uint16(2) // 保留2位小数
	if len(decimals) > 0 {
		decimal = decimals[0]
	}
	ys := strconv.FormatUint(uint64(a.Precision()), 10)
	return ys + fmtScale(a.Scale(), decimal, false)
}
