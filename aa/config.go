package aa

import (
	"encoding/json"
	"github.com/hi-iwi/AaGo/dtype"
	"strconv"
	"strings"
	"time"
)

type Config interface {
	Reload(app *Aa) error
	LoadAIni(cfgs map[string]json.RawMessage) error      // 加载 a ini 一维json配置
	AddOtherConfigs(otherConfigs map[string]string) // 这里有锁，所以要批量设置
	//getOtherConfig(key string) string    // 不要获取太细分，否则容易导致错误不容易被排查
	AddRsaConfigs(rsaConfigs map[string][]byte)
	//GetRsa(name string) ([]byte, bool) // 不要获取太细分，否则容易导致错误不容易被排查
	MustGetString(key string) (string, error)
	GetString(key string, defaultValue ...string) string
	MustGet(key string) (*dtype.Dtype, error)
	Get(key string, defaultValue ...interface{}) *dtype.Dtype
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
