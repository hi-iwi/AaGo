package asql

import (
	"github.com/hi-iwi/AaGo/atype"
	"strconv"
	"strings"
)

func In(field string, ids map[uint64]struct{}) string {
	if len(ids) == 0 {
		return "1!=1"
	}
	if len(ids) == 1 {
		for id, _ := range ids {
			return field + "=" + strconv.FormatUint(id, 10)
		}
	}
	var s strings.Builder
	s.Grow(len(field) + len(" IN ()") + (1+atype.MaxUint64Len)*len(ids))
	s.WriteString(field + " IN (")
	isFirst := true
	for id, _ := range ids {
		if isFirst {
			isFirst = false
		} else {
			s.WriteByte(',')
		}
		s.WriteString(strconv.FormatUint(id, 10))
	}
	s.WriteByte(')')
	return s.String()
}
func InUint(field string, ids map[uint]struct{}) string {
	if len(ids) == 0 {
		return "1!=1"
	}
	if len(ids) == 1 {
		for id, _ := range ids {
			return field + "=" + strconv.FormatUint(uint64(id), 10)
		}
	}
	var s strings.Builder
	s.Grow(len(field) + len(" IN ()") + (1+atype.MaxUintLen)*len(ids))
	s.WriteString(field + " IN (")
	isFirst := true
	for id, _ := range ids {
		if isFirst {
			isFirst = false
		} else {
			s.WriteByte(',')
		}
		s.WriteString(strconv.FormatUint(uint64(id), 10))
	}
	s.WriteByte(')')
	return s.String()
}
func InValues(field string, ids map[string]struct{}) string {
	if len(ids) == 0 {
		return "1!=1"
	}
	if len(ids) == 1 {
		for id, _ := range ids {
			return field + `="` + id + `"`
		}
	}

	var s strings.Builder
	s.WriteString(field)
	s.WriteString(" IN (")
	isFirst := true
	for id, _ := range ids {
		if isFirst {
			isFirst = false
		} else {
			s.WriteByte(',')
		}
		s.WriteByte('"')
		s.WriteString(id)
		s.WriteByte('"')
	}
	s.WriteByte(')')
	return s.String()
}

type ArgStmt struct {
	Field   string
	Value   interface{}
	Valid   bool
	Ignores []string // 忽略部分不同步的字段
}

/*
  组合sql语句，用于修改符合valid条件的字段
  @return ["a=?","b=?"], [$a,$b, $condId?]
*/
func ArgPairs(condId interface{}, args []ArgStmt) (string, []interface{}, bool) {
	var n int
loop1:
	for i, arg := range args {
		if !arg.Valid {
			continue
		}
		if arg.Ignores != nil {
			for _, no := range arg.Ignores {
				if no == arg.Field {
					args[i].Valid = false // 重置为忽略
					continue loop1
				}
			}
		}
		n++
	}
	if n == 0 {
		return "", nil, false
	}
	n2 := n
	if condId != nil {
		n2++
	}
	var fs strings.Builder
	//fs.Grow()
	fas := make([]interface{}, n2)
	var i int
	for _, arg := range args {
		if !arg.Valid {
			continue
		}
		if i > 0 {
			fs.WriteByte(',')
		}
		fs.WriteString(arg.Field)
		fs.WriteString("=?")
		fas[i] = arg.Value
		i++
	}
	if condId != nil {
		fas[i] = condId
	}
	return fs.String(), fas, true
}
