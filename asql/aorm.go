package asql

import (
	"strconv"
	"strings"
)

const (
	StrUint8  = "255"
	StrUint16 = "65535"
	StrUint24 = "16777215"
	StrUint   = "4294967295"
	StrUint64 = "18446744073709551615"
)

// 用作 ON DUPLICATE KEY UPDATE v=CASE .. v+VALUES(v) .. END
type SafeDupIncrN struct {
	Field string
	Min   string
	Max   string
}

func DupUint8N(field string) SafeDupIncrN {
	return SafeDupIncrN{field, "0", StrUint8}
}
func DupUint16N(field string) SafeDupIncrN {
	return SafeDupIncrN{field, "0", StrUint16}
}
func DupUint24N(field string) SafeDupIncrN {
	return SafeDupIncrN{field, "0", StrUint24}
}
func DupUintN(field string) SafeDupIncrN {
	return SafeDupIncrN{field, "0", StrUint}
}
func DupUint64N(field string) SafeDupIncrN {
	return SafeDupIncrN{field, "0", StrUint64}
}

// @warn 如果类型是 unsigned ，一定要使用 INSERT IGNORE ，否则若值为负数，则直接报错，走不到 ON DUPLICATE KEY UPDATE 部分
func SafeDupIncrs(tb string, ns []SafeDupIncrN) string {
	var s strings.Builder
	s.Grow(len(ns) * 168)
	for i, n := range ns {
		if i > 0 {
			s.WriteByte(',')
		}
		old := tb + "." + n.Field
		ntb := "new_" + tb + "." + n.Field
		// recmds=CASE WHEN tb.recmds-$MIN<=-new_tb.recmds THEN $MIN WHEN new_tb.recmds>=$MAX-tb.recmds THEN $MAX ELSE tb.recmds+new_tb.recmds END
		s.WriteString(n.Field)
		s.WriteString("=CASE WHEN ")
		s.WriteString(old)
		s.WriteString("-")
		s.WriteString(n.Min)
		s.WriteString("<=")
		s.WriteString("-" + ntb)
		s.WriteString(" THEN ")
		s.WriteString(n.Min)
		s.WriteString(" WHEN ")
		s.WriteString(ntb)
		s.WriteString(">=")
		s.WriteString(n.Max)
		s.WriteString("-")
		s.WriteString(old)
		s.WriteString(" THEN ")
		s.WriteString(n.Max)
		s.WriteString(" ELSE ")
		s.WriteString(old)
		s.WriteString("+")
		s.WriteString(ntb)
		s.WriteString(" END")
	}
	return s.String()
}

func SafeIncr(field string, n int, max string) string {
	if n == 0 {
		return field + "=" + field
	}
	s := field + "=CASE"
	var ns string
	if n > 0 {
		ns = field + "+" + strconv.FormatUint(uint64(n), 10)
		s += " WHEN " + ns + ">=" + max + " THEN " + max
	} else {
		ns = field + "-" + strconv.FormatUint(uint64(-n), 10)
		s += " WHEN " + ns + "<=0 THEN 0"
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
	return SafeIncr(field, n, StrUint8)
}

func SafeIncrUint16(field string, n int) string {
	return SafeIncr(field, n, StrUint16)
}
func SafeIncrUint24(field string, n int) string {
	return SafeIncr(field, n, StrUint24)
}

func SafeIncrUint(field string, n int) string {
	return SafeIncr(field, n, StrUint)
}
func SafeIncrUint64(field string, n int) string {
	return SafeIncr(field, n, StrUint64)
}

func toMySqlFieldName(k string) string {
	fields := strings.Split(k, ".")
	for i, field := range fields {
		fields[i] = "`" + field + "`"
	}
	return strings.Join(fields, ".")
}
