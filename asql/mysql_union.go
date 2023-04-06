package asql

import (
	"fmt"
	"github.com/hi-iwi/AaGo/atype"
	"strconv"
	"strings"
)

func UnionAllQs(format string, ptbs []string, xargs func(ptb string) []interface{}) string {
	var qs strings.Builder
	for i, ptb := range ptbs {
		if i > 0 {
			qs.WriteString(" UNION ALL ")
		}
		args := xargs(ptb)
		qs.WriteString(fmt.Sprintf(format, args...))
	}
	return qs.String()
}

func UnionInUints(ids []uint, f func(uint) string) ([]string, string) {
	tables := make([]string, 0)
	var conds strings.Builder
	conds.Grow((atype.MaxUintLen + 1) * len(ids))
	for i, id := range ids {
		if i > 0 {
			conds.WriteByte(',')
		}
		conds.WriteString(strconv.FormatUint(uint64(id), 10))

		table := f(id)
		var exists bool
		for _, t := range tables {
			if t == table {
				exists = true
			}
		}
		if !exists {
			tables = append(tables, f(id))
		}
	}
	return tables, conds.String()
}
func UnionInUint64s(ids []uint64, f func(uint64) string) ([]string, string) {
	tables := make([]string, 0)
	var conds strings.Builder
	conds.Grow((atype.MaxUintLen + 1) * len(ids))
	for i, id := range ids {
		if i > 0 {
			conds.WriteByte(',')
		}
		conds.WriteString(strconv.FormatUint(id, 10))

		table := f(id)
		var exists bool
		for _, t := range tables {
			if t == table {
				exists = true
			}
		}
		if !exists {
			tables = append(tables, f(id))
		}
	}
	return tables, conds.String()
}
