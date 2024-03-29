package ae

import (
	"runtime"
	"strconv"
	"strings"
)

func CallerMsg(errmsg string, skip int) (string, string) {
	caller := Caller(skip)
	if errmsg == "context canceled" {
		skip++
		caller2 := Caller(skip)
		if caller2 != caller {
			caller = Caller(skip) + "->" + caller
		}
	}
	return errmsg, caller
}
func Caller(skip int) string {
	var msg string
loop:
	for {
		skip++ // 跳出Caller当前函数
		pc, file, line, ok := runtime.Caller(skip)
		if !ok {
			return msg
		}
		pcs := runtime.FuncForPC(pc).Name() // 函数名
		pi := strings.LastIndexByte(pcs, '.') + 1
		fn := pcs[pi:]
		var f string
		seps := strings.Split(file, "/")
		l := len(seps)
		if l == 1 {
			f = seps[0]
		} else if l == 2 {
			f = seps[0] + "/" + seps[1]
		} else {
			f = seps[l-3] + "/" + seps[l-2] + "/" + seps[l-1]
		}
		for _, sep := range seps {
			// AaGo 框架上移到业务代码
			s := strings.ToLower(sep)
			if strings.Index(s, "!aa!go@") == 0 {
				continue loop
			}
		}
		if fn == "func1" {
			fn = ""
		} else {
			fn = " " + fn
		}
		msg = "[" + f + ":" + strconv.Itoa(line) + fn + "]" + msg
		return msg
	}
}
