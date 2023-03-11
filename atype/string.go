package atype

import (
	"strconv"
	"time"
)

// int64 -> datetime
func S2Datetime(s string, base int, loc *time.Location) Datetime {
	n, err := strconv.ParseInt(s, base, 64)
	if err != nil {
		return MinDatetime
	}
	return UnixTime(n).Datetime(loc)
}
func S2Date(s string, base int, loc *time.Location) Date {
	n, err := strconv.ParseInt(s, base, 64)
	if err != nil {
		return MinDate
	}
	return UnixTime(n).Date(loc)
}
func S2Distri(s string, base int) Distri {
	n, _ := strconv.ParseInt(s, base, 24)
	return Distri(n)
}
func S2Percent(s string, base int) Percent {
	n, _ := strconv.ParseUint(s, base, 16)
	return Percent(n)
}

func S2Money(s string, base int) Money {
	n, _ := strconv.ParseInt(s, base, 64)
	return Money(n)
}
func S2Booln(s string, base int) Booln {
	n, _ := strconv.ParseUint(s, base, 8)
	return Booln(n)
}
func S2Uint8(s string, base int) uint8 {
	n, _ := strconv.ParseUint(s, base, 8)
	return uint8(n)
}
func S2Uint16(s string, base int) uint16 {
	n, _ := strconv.ParseUint(s, base, 16)
	return uint16(n)
}

func S2Uint24(s string, base int) Uint24 {
	n, _ := strconv.ParseUint(s, base, 24)
	return Uint24(n)
}
func S2Uint(s string, base int) uint {
	n, _ := strconv.ParseUint(s, base, 32)
	return uint(n)
}
func S2Uint64(s string, base int) uint64 {
	n, _ := strconv.ParseUint(s, base, 64)
	return n
}
func S2Int8(s string, base int) int8 {
	n, _ := strconv.ParseInt(s, base, 8)
	return int8(n)
}
func S2Int16(s string, base int) int16 {
	n, _ := strconv.ParseInt(s, base, 16)
	return int16(n)
}

func S2Int24(s string, base int) Int24 {
	n, _ := strconv.ParseInt(s, base, 24)
	return Int24(n)
}
func S2Int(s string, base int) int {
	n, _ := strconv.ParseInt(s, base, 32)
	return int(n)
}
func S2Int64(s string, base int) int64 {
	n, _ := strconv.ParseInt(s, base, 64)
	return n
}
