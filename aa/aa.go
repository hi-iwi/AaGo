package aa

import (
	"github.com/RussellLuo/timingwheel"
	"sync"
)

var (
	smtx sync.RWMutex
)

type App struct {
	// self imported configurations, e.g. parsed from ini
	Config Config
	// system configuration
	configuration Configuration
	Log           Log
}

type Aa struct {
	App
	wheelTimer *timingwheel.TimingWheel
}

func NewApp(ini string) (*App, error) {
	cfg, conf, err := LoadIni(ini, AfterConfigLoaded)
	if err != nil {
		return nil, err
	}
	app := &App{
		Config:        cfg,
		configuration: conf,
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
		App:        *app,
		wheelTimer: timer,
	}
	return a, nil
}

func (app *App) ReloadConfig() error {
	smtx.Lock()
	defer smtx.Unlock()
	conf, err := app.Config.Reload(AfterConfigLoaded)
	if err != nil {
		return err
	}
	app.configuration = conf
	return nil
}

func (app *App) Cfg() *Configuration {
	smtx.RLock()
	defer smtx.RUnlock()
	return &app.configuration
}

