package aa

import (
	"github.com/hi-iwi/AaGo/dtype"
	"strconv"
	"strings"
	"time"
)

type Config interface {
	Reload() error
	Add(otherConfigs map[string]string) // 这里有锁，所以要批量设置
	Get(key string, defaultValue ...interface{}) *dtype.Dtype
	MustGet(key string) (*dtype.Dtype, error)
	GetString(key string, defaultValue ...string) string
	MustGetString(key string) (string, error)
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
