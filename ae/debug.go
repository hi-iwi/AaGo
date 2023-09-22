package ae

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func Print(v interface{}) {
	s, err := json.Marshal(v)
	if err != nil {
		fmt.Println(v, err)
		return
	}
	fmt.Println(string(s))
}
func Caller(skip int) string {
	for {
		skip++ // 跳出Caller当前函数
		pc, file, line, ok := runtime.Caller(skip)
		if !ok {
			return ""
		}
		pcs := runtime.FuncForPC(pc).Name() // 函数名
		pi := strings.LastIndexByte(pcs, '.') + 1
		fn := pcs[pi:]
		if fn == "func1" {
			fn = ""
		} else {
			fn = " " + fn
		}
		var f string
		seps := strings.Split(file, "/")
		l := len(seps)
		if l == 1 {
			f = seps[0]
		} else if l > 1 {
			f = seps[l-2] + "/" + seps[l-1]
		}

		if f == "aa/aa.go" || f == "aa/log.go" || f == "aa/log_default.go" {
			continue
		}
		return "[" + f + ":" + strconv.Itoa(line) + fn + "]"
	}
}
