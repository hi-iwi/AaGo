package aa

import (
	"strings"
	"sync"

	"github.com/luexu/dtype"
	"gopkg.in/ini.v1"
)

type Ini struct {
	mu   sync.RWMutex
	data *ini.File
}

func (app *Aa) ParseIni(filename string) error {
	c, err := ini.Load(filename)
	if err != nil {
		return err
	}
	var conf = &Ini{data: c}

	app.mu.Lock()
	app.Config = conf
	app.mu.Unlock()
	return nil
}

func (c *Ini) Get(key string, defaultValue ...interface{}) *dtype.Dtype {
	keys := splitDots(key)

	c.mu.RLock()
	defer c.mu.RUnlock()

	var s *ini.Section
	if len(keys) == 1 {
		if s = c.data.Section(""); s.HasKey(key) {
			return dtype.New(s.Key(key).String())
		}
		return defaultDtype(key, defaultValue...)
	}

	k := strings.Join(keys[1:], "_")
	if s = c.data.Section(keys[0]); s.HasKey(k) {
		return dtype.New(s.Key(k).String())
	}
	return defaultDtype(key, defaultValue...)
}
