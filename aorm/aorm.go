package aorm

import (
	"strconv"
	"strings"
)

func SafeIncr(field string, n int, max string) string {
	if n == 0 {
		return field + "=field"
	}
	s := field + "=CASE"
	var ns string
	if n > 0 {
		ns = field + "+" + strconv.FormatUint(uint64(n), 10)
		s += " WHEN " + ns + ">" + max + " THEN " + max
	} else {
		ns = field + "-" + strconv.FormatUint(uint64(-n), 10)
		s += " WHEN " + ns + "<0 THEN 0"
	}
	s += " ELSE " + ns + " END"
	return s
}

// INSERT INTO tb (v) VALUES ()
func SafeUintString(n int) string {
	if n <= 0 {
		return "0"
	}
	return strconv.FormatUint(uint64(n), 10)
}

func SafeIncrUint8(field string, n int) string {
	return SafeIncr(field, n, "255")
}

func SafeIncrUint16(field string, n int) string {
	return SafeIncr(field, n, "65535")
}
func SafeIncrUint24(field string, n int) string {
	return SafeIncr(field, n, "16777215")
}

func SafeIncrUint(field string, n int) string {
	return SafeIncr(field, n, "4294967295")
}
func SafeIncrUint64(field string, n int) string {
	return SafeIncr(field, n, "18446744073709551615")
}

func toMySqlFieldName(k string) string {
	fields := strings.Split(k, ".")
	for i, field := range fields {
		fields[i] = "`" + field + "`"
	}
	return strings.Join(fields, ".")
}
