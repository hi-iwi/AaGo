package aa

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
)

func (c *Ini) loadRsa() error {
	root := c.GetString(CkRsaRoot)
	if root == "" {
		return nil
	}
	var (
		name string
		err  error
		dir  []fs.FileInfo
		p    string // file path
	)

	if dir, err = ioutil.ReadDir(root); err != nil {
		return err
	}

	rsas := make(map[string][]byte, len(dir))
	// 因为RSA是配对出现的，所以要整体加载
	for _, fi := range dir {
		if !fi.IsDir() {
			name = fi.Name()
			if len(name) > 0 && name[0] != '.' {
				p = root + "/" + name
				dat, _ := os.ReadFile(p)
				if len(dat) == 0 {
					return fmt.Errorf("invalid rsa file `%s`", p)
				}
				rsas[name] = dat
			}
		}
	}

	c.AddRsaConfigs(rsas)
	return nil
}

func (c *Ini) AddRsaConfigs(rsaConfigs map[string][]byte) {
	cfgMtx.Lock()
	cfgMtx.Unlock()
	for name, v := range rsaConfigs {
		if len(v) > 0 {
			c.rsa[name] = v
		}
	}
}

func (c *Ini) GetRsa(name string) ([]byte, bool) {
	cfgMtx.RLock()
	defer cfgMtx.RUnlock()
	v, ok := c.rsa[name]
	return v, ok
}
