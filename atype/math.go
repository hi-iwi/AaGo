package atype

import (
	"math"
)

func FloatToInt8(n float64) int8 {
	if n < float64(math.MinInt8) || n > float64(math.MaxInt8) {
		return 0
	}
	return int8(n) // 上面进行范围判断了，所以之类可以强转
}
func FloatToInt16(n float64) int16 {
	if n < float64(math.MinInt16) || n > float64(math.MaxInt16) {
		return 0
	}
	return int16(n) // 上面进行范围判断了，所以之类可以强转
}
func FloatToInt24(n float64) Int24 {
	if n < float64(MinInt24) || n > float64(MaxInt24) {
		return 0
	}
	return Int24(n) // 上面进行范围判断了，所以之类可以强转
}
func FloatToInt(n float64) int {
	if n < float64(math.MinInt) || n > float64(math.MaxInt) {
		return 0
	}
	return int(n) // 上面进行范围判断了，所以之类可以强转
}
func Float2Int64(n float64) int64 {
	if n < float64(math.MinInt64) || n > float64(math.MaxInt64) {
		return 0
	}
	return int64(n) // 上面进行范围判断了，所以之类可以强转
}
func Float2Uint8(n float64) uint8 {
	if n < 0.0 || n > float64(math.MaxUint8) {
		return 0
	}
	return uint8(n) // 上面进行范围判断了，所以之类可以强转
}
func Float2Uint16(n float64) uint16 {
	if n < 0.0 || n > float64(math.MaxUint16) {
		return 0
	}
	return uint16(n) // 上面进行范围判断了，所以之类可以强转
}
func Float2Uint24(n float64) Uint24 {
	if n < 0.0 || n > float64(MaxUint24) {
		return 0
	}
	return Uint24(n) // 上面进行范围判断了，所以之类可以强转
}
func Float2Uint(n float64) uint {
	if n < 0.0 || n > float64(math.MaxUint) {
		return 0
	}
	return uint(n) // 上面进行范围判断了，所以之类可以强转
}

func Float2Uint64(n float64) uint64 {
	if n < 0 || n > float64(math.MaxUint64) {
		return 0
	}
	return uint64(n) // 上面进行范围判断了，所以之类可以强转
}
