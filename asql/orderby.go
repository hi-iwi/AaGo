package asql

import (
	"strings"

	"github.com/hi-iwi/AaGo/atype"
)

func Fields(u interface{}, fields ...string) string {
	s := ""
	if byAlias(fields...) {
		s = atype.JoinByNames(u, atype.JoinKeys, ",", fields...)
	} else {
		s = atype.JoinNamesByElements(u, atype.JoinKeys, ",", fields...)

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
