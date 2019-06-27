package aa

import (
	"strings"
	"sync"

	"gopkg.in/ini.v1"
)

type Ini struct {
	mu   sync.RWMutex
	data *ini.File
}

func (a *Aa) ParseIni(filename string) error {
	c, err := ini.Load(filename)
	if err != nil {
		return err
	}
	var conf = &Ini{data: c}

	a.mu.Lock()
	a.Config = conf
	a.mu.Unlock()
	return nil
}

func (c *Ini) Get(key string, defaultValue ...interface{}) *Dtype {
	keys := splitDots(key)
	dv := parseDefaultValue(defaultValue...)
	c.mu.RLock()
	defer c.mu.RUnlock()

	var s *ini.Section
	if len(keys) == 1 {
		if s = c.data.Section(""); s.HasKey(key) {
			return NewDtype(s.Key(key).String())
		}
		return NewDtype(dv)
	}

	k := strings.Join(keys[1:], "_")
	if s = c.data.Section(keys[0]); s.HasKey(k) {
		return NewDtype(s.Key(k).String())
	}
	return NewDtype(dv)
}
