package asql

import (
	"regexp"
	"strings"

	"github.com/hi-iwi/AaGo/dtype"
)

func byAlias(fields ...string) bool {
	if len(fields) == 0 {
		return false
	}
	char := fields[0][0]

	return !(char >= 'A' && char <= 'Z')
}

func Comma(u interface{}, fields ...string) (s string) {
	if byAlias(fields...) {
		s = dtype.JoinByNames(u, dtype.JoinMySQL, ", ", fields...)
	} else {
		s = dtype.JoinNamesByElements(u, dtype.JoinMySQL, ", ", fields...)
	}
	return strings.Trim(defenseInjection(s), " ")
}

func CommaWithEnd(u interface{}, fields ...string) string {
	s := Comma(u, fields...)
	if len(s) == 0 {
		return s
	}
	return s + ", "
}

func CommaWithHead(u interface{}, fields ...string) string {
	s := Comma(u, fields...)
	if len(s) == 0 {
		return s
	}
	return "," + s
}

func And(u interface{}, fields ...string) (s string) {
	if byAlias(fields...) {
		s = dtype.JoinByNames(u, dtype.JoinMySQL, " AND ", fields...)
	} else {
		s = dtype.JoinNamesByElements(u, dtype.JoinMySQL, " AND ", fields...)

	}
	return strings.Trim(defenseInjection(s), " ")
}

func Or(u interface{}, fields ...string) (s string) {
	if byAlias(fields...) {
		s = dtype.JoinByNames(u, dtype.JoinMySQL, " OR ", fields...)
	} else {
		s = dtype.JoinNamesByElements(u, dtype.JoinMySQL, " OR ", fields...)
	}
	return strings.Trim(s, " ")
}

func Like(u interface{}, fields ...string) (s string) {
	if byAlias(fields...) {
		s = dtype.JoinByNames(u, dtype.JoinMySqlFullLike, " OR ", fields...)
	} else {
		s = dtype.JoinNamesByElements(u, dtype.JoinMySqlFullLike, " OR ", fields...)
	}
	return strings.Trim(defenseInjection(s), " ")
}

func AndWithWhere(u interface{}, fields ...string) string {
	s := And(u, fields...)
	if len(s) > 0 {
		s = " WHERE " + s + " "
	}
	return s
}

func OrWithWhere(u interface{}, fields ...string) string {
	s := Or(u, fields...)
	if len(s) > 0 {
		s = " WHERE " + s + " "
	}
	return s
}

func Where(conds ...string) string {
	w := ""

	for i := 0; i < len(conds); i++ {
		w += " " + conds[i]
	}

	re, _ := regexp.Compile(`\(\s*(OR|AND)*\s*\)`)
	w = re.ReplaceAllString(w, " ")

	re2, _ := regexp.Compile(`((OR|AND)\s+)+(OR|AND)`)
	w = re2.ReplaceAllString(w, "$3")

	re3, _ := regexp.Compile(`(^(\s*(OR|AND))+)|(((OR|AND)+\s*)$)`)
	w = re3.ReplaceAllString(w, "")

	re4, _ := regexp.Compile(`\((\s*(OR|AND))+`)
	w = re4.ReplaceAllString(w, "(")

	re5, _ := regexp.Compile(`(\s*(OR|AND))+\)`)
	w = re5.ReplaceAllString(w, ")")

	w = strings.Trim(w, " ")

	if len(w) > 0 {
		w = " WHERE " + w + " "
	}
	return w
}
