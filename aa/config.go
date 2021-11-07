package aa

import (
	"strconv"
	"strings"
	"time"

	"github.com/hi-iwi/AaGo/dtype"
)

type Config interface {
	Get(key string, defaultValue ...interface{}) *dtype.Dtype
	Set(k, v string)
}

func parseToDuration(d string) time.Duration {
	if len(d) < 2 {
		return 0
	}
	var t int
	if d[len(d)-2:] == "ms" {
		t, _ = strconv.Atoi(d[0 : len(d)-2])
		return time.Duration(t) * time.Millisecond
	}

	if d[len(d)-1:] == "s" {
		t, _ = strconv.Atoi(d[0 : len(d)-1])
	} else {
		t, _ = strconv.Atoi(d)
	}
	return time.Duration(t) * time.Second
}

func splitDots(keys ...string) []string {
	n := make([]string, 0)
	for _, key := range keys {
		n = append(n, strings.Split(key, ".")...)
	}
	return n
}

func defaultDtype(defaultValue ...interface{}) *dtype.Dtype {
	if len(defaultValue) > 0 {
		return dtype.New(defaultValue[0])
	}
	return dtype.New("")
}
