package aa

import (
	"context"
	"github.com/hi-iwi/AaGo/ae"
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

// 快捷方式，对服务器错误记录日志
func (app *Aa) Try(ctx context.Context, e *ae.Error) bool {
	if e != nil && e.IsServerError() {
		app.Log.Error(ctx, e.Error())
		return false
	}
	return true
}

// 快捷记录错误
func (app *Aa) TryLog(ctx context.Context, err error) {
	if err != nil {
		app.Log.Error(ctx, err.Error())
	}
}

// 快捷panic
func (app *Aa) TryPanic(ctx context.Context, e *ae.Error) {
	if e != nil {
		app.Log.Error(ctx, e.Error())
		panic(e.Error())
	}
}
