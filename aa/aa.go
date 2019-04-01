package aa

import (
	"sync"
)

type Aa struct {
	mu   sync.Mutex
	once sync.Once
	// self imported configurations, e.g. parsed from xml
	Config Config
	// system configuration
	Configuration Configuration
}

func New() *Aa {
	aa := &Aa{
		Config: &Yaml{},
	}

	return aa
}
