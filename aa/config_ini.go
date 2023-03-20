package aa

import (
	"fmt"
	"github.com/hi-iwi/AaGo/atype"
	"gopkg.in/ini.v1"
	"strings"
	"sync"
)

type Ini struct {
	path        string
	data        *ini.File
	rsa         map[string][]byte // rsa 是配对出现的，不要使用 sync.Map， 直接对整个map加锁设置
	otherConfig map[string]string // 不要使用 sync.Map， 直接对整个map加锁设置
}

var (
	cfgMtx sync.RWMutex
)

func MakeIni(path string, after func(Config) Configuration) (Config, Configuration, error) {
	rsa := make(map[string][]byte)
	ocfg := make(map[string]string)
	cfg := &Ini{path: path, rsa: rsa, otherConfig: ocfg}
	conf, err := cfg.Reload(after)
	return cfg, conf, err
}

func (c *Ini) Reload(after func(Config) Configuration) (Configuration, error) {
	data, err := ini.Load(c.path)
	if err != nil {
		return Configuration{}, err
	}
	cfgMtx.Lock()
	c.data = data
	cfgMtx.Unlock()
	if err = c.loadRsa(); err != nil {
		return Configuration{}, err
	}
	return after(c), nil
}

// 这里有锁，所以要批量设置
func (c *Ini) AddConfigs(otherConfigs map[string]string) {
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

// 不要获取太细分，否则容易导致错误不容易被排查
func (c *Ini) getOtherConfig(key string) string {
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
	if rsa, _ := c.getRsa(key); len(rsa) > 0 {
		return string(rsa), nil
	}
	// 从其他配置（如数据库下载来的）读取
	if v = c.getOtherConfig(key); v != "" {
		return v, nil
	}
	return "", fmt.Errorf("must set config `%s`", key)
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

func (c *Ini) MustGet(key string) (*atype.Atype, error) {
	v, err := c.MustGetString(key)
	if err != nil {
		return nil, err
	}
	return atype.New(v), nil
}

// Get(key) or Get(key, defaultValue)
// 先从 ini 文件读取，找不到再去从其他 provider （如数据库拉下来的配置）里面找
func (c *Ini) Get(key string, defaultValue ...interface{}) *atype.Atype {
	v, _ := c.MustGetString(key)
	if v != "" {
		return atype.New(v)
	}
	if len(defaultValue) > 0 {
		return atype.New(defaultValue[0])
	}
	return atype.New("")
}
