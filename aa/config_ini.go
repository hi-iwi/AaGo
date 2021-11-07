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
	otherConfig map[string]string // 不要使用 sync.Map， 直接对整个map加锁设置
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
func (c *Ini) LoadOtherConfig(otherConfigs map[string]string) {
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
	v, _ := c.MustGetString(key)
	if v != "" {
		return dtype.New(v)
	}
	if len(defaultValue) > 0 {
		return dtype.New(defaultValue[0])
	}
	return dtype.New("")
}

func (c *Ini) MustGet(key string) (*dtype.Dtype, error) {
	v, err := c.MustGetString(key)
	if err != nil {
		return nil, err
	}
	return dtype.New(v), nil
}

func (c *Ini) GetString(key string, defaultValue ...string) string {
	v, _ := c.MustGetString(key)
	if v != "" {
		return v
	}
	if len(defaultValue) > 0 {
		v = defaultValue[0]
	}
	return v
}
func (c *Ini) MustGetString(key string) (string, error) {
	cfgMtx.RLock()
	v := c.getIni(key)
	cfgMtx.Unlock()

	if v == "" {
		if d, ok := c.otherConfig[key]; ok {
			v = d
		}
	}
	if v == "" {
		return "", errors.New("must set config `%s`")
	}
	return v, nil
}
