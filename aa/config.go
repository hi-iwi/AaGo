package aa

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"sync"

	"gopkg.in/yaml.v2"
)

type Config interface {
	Get(key string, defaultValue ...interface{}) *Dtype
}

type Yaml struct {
	mu   sync.RWMutex
	data map[interface{}]interface{}
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
	return nil
}

func (c *Yaml) Get(key string, defaultValue ...interface{}) *Dtype {
	keys := splitDots(key)
	dv := parseDefaultValue(defaultValue...)
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
				return NewDtype(dv)
			}
		} else {
			return NewDtype(dv)
		}
	}
	return NewDtype(v)
}

func (a *Aa) ParseConfig(filename string) (Config, error) {
	var conf = &Yaml{}
	yamlAbsPath, err := filepath.Abs(filename)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadFile(yamlAbsPath)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, &conf.data); err != nil {
		return nil, err
	}
	// a.mu.Lock()
	// for k, v := range c {
	// 	a.Config[k] = NewDtype(v)
	// }
	// a.mu.Unlock()
	a.mu.Lock()
	a.Config = conf
	a.mu.Unlock()

	a.ParseToConfiguration()
	return conf, nil
}
