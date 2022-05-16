package aa

import (
	"github.com/RussellLuo/timingwheel"
	"sync"
)

var (
	cfgMtx sync.RWMutex
)

type App struct {
	// self imported configurations, e.g. parsed from ini
	Config Config
	// system configuration
	Configuration Configuration
	Log           Log
}

type Aa struct {
	App
	Timer *timingwheel.TimingWheel
}

func NewApp(ini string) (*App, error) {
	cfg, conf, err := LoadIni(ini, AfterConfigLoaded)
	if err != nil {
		return nil, err
	}
	app := &App{
		Config:        cfg,
		Configuration: conf,
		Log:           NewDefaultLog(),
	}
	return app, nil
}

func New(ini string, timer *timingwheel.TimingWheel) (*Aa, error) {
	app, err := NewApp(ini)
	if err != nil {
		return nil, err
	}
	a := &Aa{
		App:   *app,
		Timer: timer,
	}
	return a, nil
}
