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
	rsa         map[string][]byte // rsa 是配对出现的，不要使用 sync.Map， 直接对整个map加锁设置
	otherConfig map[string]string // 不要使用 sync.Map， 直接对整个map加锁设置
}

func (app *Aa) LoadIni(path string) error {
	app.Config = &Ini{path: path}
	err := app.Config.Reload(app)
	if err != nil {
		return err
	}
	return nil
}

func (c *Ini) Reload(app *Aa) error {
	data, err := ini.Load(c.path)
	if err != nil {
		return err
	}
	cfgMtx.Lock()
	c.data = data
	app.ParseToConfiguration()
	cfgMtx.Unlock()
	err = c.loadRsa()
	return err
}

// 这里有锁，所以要批量设置
func (c *Ini) AddOtherConfigs(otherConfigs map[string]string) {
	cfgMtx.Lock()
	defer cfgMtx.Unlock()
	for k, v := range otherConfigs {
		c.otherConfig[k] = v
	}
}

func (c *Ini) getIni(key string) string {
	cfgMtx.RLock()
	defer cfgMtx.RUnlock()

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
func (c *Ini) GetOtherConfig(key string) string {
	cfgMtx.RLock()
	defer cfgMtx.RUnlock()
	d, _ := c.otherConfig[key]
	return d
}

func (c *Ini) MustGetString(key string) (string, error) {
	v := c.getIni(key)
	if v != "" {
		return v, nil
	}
	// 从RSA读取
	if rsa, _ := c.GetRsa(key); len(rsa) > 0 {
		return string(rsa), nil
	}
	// 从其他配置（如数据库下载来的）读取
	if v = c.GetOtherConfig(key); v != "" {
		return v, nil
	}
	return "", errors.New("must set config `%s`")
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

func (c *Ini) MustGet(key string) (*dtype.Dtype, error) {
	v, err := c.MustGetString(key)
	if err != nil {
		return nil, err
	}
	return dtype.New(v), nil
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
