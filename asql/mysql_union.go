package asql

import (
	"fmt"
	"strings"
)

func UnionAllQs( format string, ptbs []string, xargs func(ptb string) []interface{}) string {
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
