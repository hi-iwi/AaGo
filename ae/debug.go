package ae

import (
	"runtime"
	"strconv"
	"strings"
)

func Caller(skip int) string {
	var s string
	for {
		skip++ // 跳出Caller当前函数
		pc, file, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}
		a := strings.LastIndexByte(file, '/') + 1 // 文件名
		pcs := runtime.FuncForPC(pc).Name()       // 函数名
		pi := strings.LastIndexByte(pcs, '.') + 1
		s += "[" + file[a:] + ":" + strconv.Itoa(line) + " " + pcs[pi:] + "]"
	}
	return s
}
