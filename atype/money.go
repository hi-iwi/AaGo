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

func fmtDecimal(ps string, dec, decimal uint16, trim bool) string {
	const d uint16 = 4 //  4位小数
	if decimal == 0 || dec == 0 {
		return ps
	}
	if decimal >= d {
		return ps + "." + fmt.Sprintf("%04d", dec)
	}

	m := dec / uint16(math.Pow10(int(d-decimal)))
	// 四舍五入是违法的，只能舍弃
	return ps + "." + fmt.Sprintf("%0*d", decimal, m)
}
func (a Amount) Fmt(decimals ...uint16) string {
	decimal := uint16(2) // 保留2位小数
	if len(decimals) > 0 {
		decimal = decimals[0]
	}
	ys := strconv.FormatInt(a.Precision(), 10)
	return fmtDecimal(ys, a.Scale(), decimal, true)
}
func (a Amount) Format(decimals ...uint16) string {
	decimal := uint16(2) // 保留2位小数
	if len(decimals) > 0 {
		decimal = decimals[0]
	}
	ys := strconv.FormatInt(a.Precision(), 10)
	return fmtDecimal(ys, a.Scale(), decimal, false)
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

func (a Uamount) Fmt(decimals ...uint16) string {
	decimal := uint16(2) // 保留2位小数
	if len(decimals) > 0 {
		decimal = decimals[0]
	}
	ys := strconv.FormatUint(a.Precision(), 10)
	return fmtDecimal(ys, a.Scale(), decimal, true)
}
func (a Uamount) Format(decimals ...uint16) string {
	decimal := uint16(2) // 保留2位小数
	if len(decimals) > 0 {
		decimal = decimals[0]
	}
	ys := strconv.FormatUint(a.Precision(), 10)
	return fmtDecimal(ys, a.Scale(), decimal, false)
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
func (a Money) Fmt(decimals ...uint16) string {
	decimal := uint16(2) // 保留2位小数
	if len(decimals) > 0 {
		decimal = decimals[0]
	}
	ys := strconv.Itoa(a.Precision())
	return fmtDecimal(ys, a.Scale(), decimal, true)
}
func (a Money) Format(decimals ...uint16) string {
	decimal := uint16(2) // 保留2位小数
	if len(decimals) > 0 {
		decimal = decimals[0]
	}
	ys := strconv.Itoa(a.Precision())
	return fmtDecimal(ys, a.Scale(), decimal, false)
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

func (a Umoney) Fmt(decimals ...uint16) string {
	decimal := uint16(2) // 保留2位小数
	if len(decimals) > 0 {
		decimal = decimals[0]
	}
	ys := strconv.FormatUint(uint64(a.Precision()), 10)
	return fmtDecimal(ys, a.Scale(), decimal, true)
}
func (a Umoney) Format(decimals ...uint16) string {
	decimal := uint16(2) // 保留2位小数
	if len(decimals) > 0 {
		decimal = decimals[0]
	}
	ys := strconv.FormatUint(uint64(a.Precision()), 10)
	return fmtDecimal(ys, a.Scale(), decimal, false)
}
