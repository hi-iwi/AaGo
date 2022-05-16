package aa

import (
	"github.com/RussellLuo/timingwheel"
	"sync"
)

var (
	cfgMtx sync.RWMutex
)

type Aa struct {
	// self imported configurations, e.g. parsed from ini
	Config Config
	// system configuration
	Configuration Configuration
	Log           Log
}

type AaWithTimer struct {
	Aa
	Timer *timingwheel.TimingWheel
}

func New(ini string) (*Aa, error) {
	cfg, conf, err := LoadIni(ini, AfterConfigLoaded)
	if err != nil {
		return nil, err
	}
	app := &Aa{
		Config:        cfg,
		Configuration: conf,
		Log:           NewDefaultLog(),
	}
	return app, nil
}

func NewWithTimer(ini string, timer *timingwheel.TimingWheel) (*AaWithTimer, error) {
	app, err := New(ini)
	if err != nil {
		return nil, err
	}
	a := &AaWithTimer{
		Aa:    *app,
		Timer: timer,
	}
	return a, nil
}
