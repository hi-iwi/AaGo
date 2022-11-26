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
func (a Amount) Fmt(dig uint16) string {
	const d uint16 = 4 //  4位小数
	y := a.Precision()
	dec := a.Scale()
	ys := strconv.FormatInt(y, 10)
	if dig == 0 || dec == 0 {
		return ys
	}
	if dig >= d {
		return ys + "." + fmt.Sprintf("%04d", dec)
	}

	m := dec / uint16(math.Pow10(int(d-dig)))
	// 四舍五入是违法的，只能舍弃
	return ys + "." + fmt.Sprintf("%0*d", dig, m)
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
func (a Uamount) Fmt(dig uint16) string {
	const d uint16 = 4 //  4位小数
	y := a.Precision()
	dec := a.Scale()
	ys := strconv.FormatUint(y, 10)
	if dig == 0 || dec == 0 {
		return ys
	}
	if dig >= d {
		return ys + "." + fmt.Sprintf("%04d", dec)
	}

	m := dec / uint16(math.Pow10(int(d-dig)))
	// 四舍五入是违法的，只能舍弃
	return ys + "." + fmt.Sprintf("%0*d", dig, m)
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
func (a Money) Fmt(dig uint16) string {
	const d uint16 = 4 //  4位小数
	y := a.Precision()
	dec := a.Scale()
	ys := strconv.Itoa(y)
	if dig == 0 || dec == 0 {
		return ys
	}
	if dig >= d {
		return ys + "." + fmt.Sprintf("%04d", dec)
	}

	m := dec / uint16(math.Pow10(int(d-dig)))
	// 四舍五入是违法的，只能舍弃
	return ys + "." + fmt.Sprintf("%0*d", dig, m)
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
func (a Umoney) Fmt(dig uint16) string {
	const d uint16 = 4 //  4位小数
	y := a.Precision()
	dec := a.Scale()
	ys := strconv.FormatUint(uint64(y), 10)
	if dig == 0 || dec == 0 {
		return ys
	}
	if dig >= d {
		return ys + "." + fmt.Sprintf("%04d", dec)
	}

	m := dec / uint16(math.Pow10(int(d-dig)))
	// 四舍五入是违法的，只能舍弃
	return ys + "." + fmt.Sprintf("%0*d", dig, m)
}
