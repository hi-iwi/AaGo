package ae

import (
	"runtime"
	"strconv"
	"strings"
)

func Caller(skip int) string {
	skip++ // 跳出Caller当前函数
	pc, file, line, _ := runtime.Caller(skip)
	a := strings.LastIndexByte(file, '/') + 1 // 文件名
	pcs := runtime.FuncForPC(pc).Name()       // 函数名
	pi := strings.LastIndexByte(pcs, '.') + 1
	return "[" + file[a:] + ":" + strconv.Itoa(line) + " " + pcs[pi:] + "]"
}
