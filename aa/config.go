package aa

import (
	"github.com/hi-iwi/AaGo/atype"
	"strconv"
	"strings"
	"time"
)

type Config interface {
	Reload(after func(Config) Configuration) (Configuration, error)

	AddConfigs(map[string]string) // 这里有锁

	//getOtherConfig(key string) string    // 不要获取太细分，否则容易导致错误不容易被排查
	AddRsaConfigs(rsaConfigs map[string][]byte)
	//GetRsa(name string) ([]byte, bool) // 不要获取太细分，否则容易导致错误不容易被排查
	MustGetString(key string) (string, error)
	GetString(key string, defaultValue ...string) string
	MustGet(key string) (*atype.Atype, error)
	Get(key string, defaultValue ...interface{}) *atype.Atype
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
