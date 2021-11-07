package aa

import (
	"errors"
	"github.com/hi-iwi/AaGo/dtype"
	"gopkg.in/ini.v1"
	"strings"
)

type Ini struct {
	path        string
	data        *ini.File
	otherConfig map[string]string
}

func (app *Aa) LoadIni(path string) error {
	cfgMtx.Lock()
	defer cfgMtx.Unlock()
	app.Config = &Ini{path: path}
	app.Config.Reload()
	app.ParseToConfiguration()

	return nil
}
func (c *Ini) Reload() error {
	data, err := ini.Load(c.path)
	if err != nil {
		return err
	}
	c.data = data
	return nil
}

// 这里有锁，所以要批量设置
func (c *Ini) Add(otherConfigs map[string]string) {
	cfgMtx.Lock()
	defer cfgMtx.Unlock()
	for k, v := range otherConfigs {
		c.otherConfig[k] = v
	}
}
func (c *Ini) getIni(key string) string {
	keys := splitDots(key)
	var s *ini.Section
	if len(keys) == 1 {
		if s = c.data.Section(""); s.HasKey(key) {
			return s.Key(key).String()
		}
		return ""
	}

	k := strings.Join(keys[1:], "_")
	if s = c.data.Section(keys[0]); s.HasKey(k) {
		return s.Key(k).String()
	}
	return ""
}

// Get(key) or Get(key, defaultValue)
// 先从 ini 文件读取，找不到再去从其他 provider （如数据库拉下来的配置）里面找
func (c *Ini) Get(key string, defaultValue ...interface{}) *dtype.Dtype {
	v, err := c.MustGet(key)
	if err != nil {
		return defaultDtype(defaultValue...)
	}
	return v
}

func (c *Ini) MustGet(key string) (*dtype.Dtype, error) {
	cfgMtx.RLock()
	defer cfgMtx.Unlock()

	v := c.getIni(key)
	if v == "" {
		v, _ = c.otherConfig[key]
	}
	if v == "" {
		return nil, errors.New("must set config `%s`")
	}
	return dtype.New(v), nil
}
