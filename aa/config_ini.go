package aa

import (
	"github.com/hi-iwi/AaGo/dtype"
	"gopkg.in/ini.v1"
	"strings"
)

type Ini struct {
	data        *ini.File
	otherConfig map[string]string
}

func (app *Aa) LoadIni(filename string) error {
	c, err := ini.Load(filename)
	if err != nil {
		return err
	}
	var conf = &Ini{data: c}
	cfgMtx.Lock()
	app.Config = conf
	app.ParseToConfiguration()
	cfgMtx.Unlock()
	return nil
}

func (c *Ini) Set(k, v string) {
	c.otherConfig[k] = v
}
func (c *Ini) getIni(key string, defaultValue ...interface{}) *dtype.Dtype {
	keys := splitDots(key)

	cfgMtx.RLock()
	defer cfgMtx.RUnlock()

	var s *ini.Section
	if len(keys) == 1 {
		if s = c.data.Section(""); s.HasKey(key) {
			return dtype.New(s.Key(key).String())
		}
		return defaultDtype(defaultValue...)
	}

	k := strings.Join(keys[1:], "_")
	if s = c.data.Section(keys[0]); s.HasKey(k) {
		return dtype.New(s.Key(k).String())
	}
	return defaultDtype(defaultValue...)
}

// Get(key) or Get(key, defaultValue)
// 先从 ini 文件读取，找不到再去从其他 provider （如数据库拉下来的配置）里面找
func (c *Ini) Get(key string, defaultValue ...interface{}) *dtype.Dtype {
	v := c.getIni(key, defaultValue)
	if v.String() != "" {
		return v
	}
	if d, ok := c.otherConfig[key]; ok {
		v = dtype.New(d)
	}
	return v
}
