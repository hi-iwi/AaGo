package ae

import (
	"runtime"
	"strconv"
	"strings"
)

func Caller(skip int) string {
	skip++ // 跳出Caller当前函数
	pc, file, line, _ := runtime.Caller(skip)
	pcs := runtime.FuncForPC(pc).Name() // 函数名
	a := strings.Split(file, "/")       // 文件名
	return a[len(a)-1] + ":" + strconv.Itoa(line) + " " + pcs
}
