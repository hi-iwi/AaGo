package aa

import (
	"fmt"
	"log"
	"time"

	"github.com/hi-iwi/AaGo/util"
)

type Configuration struct {
	//VID          string // service + ':' + service id, e.g. user01:12
	Service string `yaml:"service"`
	//ServerID     string `yaml:"server_id"`
	Env          string `yaml:"env"`         // dev test preprod product
	TimezoneID   string `yaml:"timezone_id"` // e.g. "Asia/Shanghai"
	TimeLocation *time.Location
	TimeFormat   string `yaml:"time_format"` // e.g. "2006-02-01 15:04:05"
	Mock         bool   `yaml:"mock"`        // using mock

}

const (
	CkEnv        = "env"
	CkService    = "service"
	CkTimezoneID = "timezone_id"
	CkTimeFormat = "time_format"
	CkMock       = "mock"
)

func (app *Aa) ParseToConfiguration() {
	app.mu.Lock()
	defer app.mu.Unlock()
	app.Configuration.Env = app.Config.Get(CkEnv).String()
	app.Configuration.Service = app.Config.Get(CkService).String()
	app.Configuration.TimeFormat = app.Config.Get(CkTimeFormat).String()
	//serverID := app.Config.Get("server_id").Name()
	//app.Configuration.ServerID = serverID
	//app.Configuration.VID = svc + ":" + serverID

	if tz := app.Config.Get(CkTimezoneID).String(); tz != "" {
		loc, err := time.LoadLocation(tz)
		if err != nil {
			panic("invalid timezone: " + tz + ", error: " + err.Error())
		} else {
			app.Configuration.TimezoneID = tz
			app.Configuration.TimeLocation = loc
		}
	}

	mock, _ := app.Config.Get(CkMock).Bool()
	app.Configuration.Mock = mock
}

func (c Configuration) Log() {
	msg := fmt.Sprintf("starting service %s\nenv: %s\ntimezone_id: %s\nmock: %v\ngit_ver: %s", c.Service, c.Env, c.TimezoneID, c.Mock, util.GitVersion())
	log.Println(msg)
	fmt.Println(msg)
}
