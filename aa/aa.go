package aa

import (
	"sync"
	"time"

	"github.com/hi-iwi/alog"
)

type Aa struct {
	mu sync.Mutex
	//once sync.Once
	// self imported configurations, e.g. parsed from ini
	Config Config
	// system configuration
	Configuration Configuration
	Log           alog.Log
}

func New() *Aa {
	zone, _ := time.Now().Zone()
	aa := &Aa{
		Log: alog.NewXlog(),
		Configuration: Configuration{
			TimezoneID:   zone,
			TimeLocation: time.Local,
		},
	}

	return aa
}
