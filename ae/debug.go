package ae

import (
	"runtime"
	"strconv"
	"strings"
)

func Caller(skip int) string {
	skip++ // 跳出Caller当前函数
	pc, file, line, _ := runtime.Caller(skip)
	var f string
	seps := strings.Split(file, "/")
	l := len(seps)
	if l == 1 {
		f = seps[0]
	} else if l == 2 {
		f = seps[l-2] + "/" + seps[l-1]
	} else if l > 2 {
		f = seps[l-3] + "/" + seps[l-2] + "/" + seps[l-1]
	}
	pcs := runtime.FuncForPC(pc).Name() // 函数名
	pi := strings.LastIndexByte(pcs, '.') + 1
	return "[" + f + ":" + strconv.Itoa(line) + " " + pcs[pi:] + "]"
}
