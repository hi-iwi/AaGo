package aa

import (
	"path"
	"strconv"
	"strings"
	"time"
)

type Config interface {
	Get(key string, defaultValue ...interface{}) *Dtype
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

// ParseTimeout connection timeout, r timeout, w timeout, heartbeat interval
// 10s, 1000ms
func (a *Aa) ParseTimeout(key string, defaultTimeouts ...time.Duration) (conn time.Duration, read time.Duration, write time.Duration, heartbeat time.Duration) {
	for i, t := range defaultTimeouts {
		switch i {
		case 0:
			conn = t
		case 1:
			read = t
		case 2:
			write = t
		case 3:
			heartbeat = t
		}
	}

	ts := strings.Split(strings.Replace(a.Config.Get(key).String(), " ", "", -1), ",")
	for i, t := range ts {
		switch i {
		case 0:
			conn = parseToDuration(t)
		case 1:
			read = parseToDuration(t)
		case 2:
			write = parseToDuration(t)
		case 3:
			heartbeat = parseToDuration(t)
		}
	}

	return
}

func splitDots(keys ...string) []string {
	n := make([]string, 0)
	for _, key := range keys {
		n = append(n, strings.Split(key, ".")...)
	}
	return n
}

func parseDefaultValue(vs ...interface{}) interface{} {
	if len(vs) > 0 {
		return vs[0]
	}
	return ""
}

func (a *Aa) ParseConfig(filename string) error {
	switch path.Ext(filename) {
	case ".ini":
		a.ParseIni(filename)
		// case ".yml", ".yaml":
		// 	a.ParseYml(filename)
	}

	// a.mu.Lock()
	// for k, v := range c {
	// 	a.Config[k] = NewDtype(v)
	// }
	// a.mu.Unlock()

	a.ParseToConfiguration()
	return nil
}
