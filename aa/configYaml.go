package aa

import (
	"io/ioutil"
	"sync"

	"github.com/luexu/dtype"
	"gopkg.in/yaml.v2"
)

type Yaml struct {
	mu   sync.RWMutex
	data map[interface{}]interface{}
}

func (a *Aa) ParseYml(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	var conf = &Yaml{}

	if err := yaml.Unmarshal(data, &conf.data); err != nil {
		return err
	}
	a.mu.Lock()
	a.Config = conf
	a.mu.Unlock()
	return nil
}

func (c *Yaml) Get(key string, defaultValue ...interface{}) *dtype.Dtype {
	keys := splitDots(key)
	c.mu.RLock()
	defer c.mu.RUnlock()
	var v interface{}
	var d = c.data
	var ok bool
	for i, key := range keys {
		if v, ok = d[key]; ok {
			if i == len(keys)-1 {
				break
			}
			if d, ok = v.(map[interface{}]interface{}); !ok {
				return defaultDtype(key, defaultValue...)
			}
		} else {
			return defaultDtype(key, defaultValue...)
		}
	}
	return defaultDtype(key, defaultValue...)
}
