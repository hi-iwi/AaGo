package aa

import (
	"sync"
	"time"
)

var (
	cfgMtx sync.RWMutex
)

type Aa struct {

	//once sync.Once
	// self imported configurations, e.g. parsed from ini
	Config Config
	// system configuration
	Configuration Configuration
	Log           Log
}

func New() *Aa {
	zone, _ := time.Now().Zone()
	aa := &Aa{
		Log: NewDefaultLog(),
		Configuration: Configuration{
			TimezoneID:   zone,
			TimeLocation: time.Local,
		},
	}

	return aa
}
