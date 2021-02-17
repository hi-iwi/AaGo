package aorm

import (
	"strings"

	"github.com/hi-iwi/AaGo/dtype"
)

func Fields(u interface{}, fields ...string) string {
	s := ""
	if byAlias(fields...) {
		s = dtype.JoinByAlias(u, dtype.JoinKeys, ",", fields...)
	} else {
		s = dtype.JoinAliasByElements(u, dtype.JoinKeys, ",", fields...)

	}
	return strings.Trim(s, " ")
}

func OrderBy(u interface{}, orm ORM) string {
	if len(orm.OrderByDesc) > 0 {
		fs := strings.Split(orm.OrderByDesc, ",")
		return " ORDER BY " + Fields(u, fs...) + " DESC "
	} else if len(orm.OrderBy) > 0 {
		fs := strings.Split(orm.OrderByDesc, ",")
		return " ORDER BY " + Fields(u, fs...) + " "
	}
	return ""
}
