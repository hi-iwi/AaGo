package aa

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hi-iwi/AaGo/util"
)

type Env string

func (env Env) String() string      { return string(env) }
func (env Env) IsLocal() bool       { return env == "local" || env == "loc" }
func (env Env) IsDevelopment() bool { return env == "development" || env == "dev" }
func (env Env) IsIntegration() bool { return env == "integration" }
func (env Env) IsTesting() bool     { return env == "testing" || env == "test" || env == "qc" }
func (env Env) IsPreProduction() bool {
	return env == "pre-production" || env == "pre" || env == "demo"
}
func (env Env) IsProduction() bool { return env == "production" || env == "pro" || env == "live" }

type Configuration struct {
	/*
		https://en.wikipedia.org/wiki/Deployment_environment
		local -> development/trunk -> integration -> testing/test/qc/internal acceptnace -> staging/stage/model/pre-production/demo -> production/live
		development -> test -> pre-production -> production
	*/
	Env          Env
	TimezoneID   string // e.g. "Asia/Shanghai"
	TimeLocation *time.Location
	TimeFormat   string // e.g. "2006-02-01 15:04:05"
	Mock         bool   // using mock
}

const (
	CkRsaRoot    = "rsa_root"
	CkEnv        = "env"
	CkTimezoneID = "timezone_id"
	CkTimeFormat = "time_format"
	CkMock       = "mock"
)

func AfterConfigLoaded(cfg Config) Configuration {
	log.Println("config loaded")
	return ParseToConfiguration(cfg)
}

func ParseToConfiguration(cfg Config) Configuration {
	zone, _ := time.Now().Zone()
	c := Configuration{
		Env:          Env(cfg.GetString(CkEnv)),
		TimezoneID:   zone,
		TimeLocation: time.Local,
		TimeFormat:   cfg.GetString(CkTimeFormat),
		Mock:         false,
	}

	//serverID := app.Config.Get("server_id").Name()
	//app.Configuration.ServerID = serverID
	//app.Configuration.VID = svc + ":" + serverID

	if tz := cfg.GetString(CkTimezoneID); tz != "" {
		loc, err := time.LoadLocation(tz)
		if err != nil {
			panic("invalid timezone: " + tz + ", error: " + err.Error())
		} else {
			c.TimezoneID = tz
			c.TimeLocation = loc
		}
	}

	mock, _ := cfg.Get(CkMock).Bool()
	c.Mock = mock
	return c
}

func (c *Configuration) Log() {
	msg := fmt.Sprintf("lauching...\nenv: %s\ntimezone_id: %s\nmock: %v\ngit_ver: %s", c.Env, c.TimezoneID, c.Mock, util.GitVersion())
	Println(msg)
}

// ParseTimeout connection timeout, r timeout, w timeout, heartbeat interval
// 10s, 1000ms
func (app *App) ParseTimeout(t string, defaultTimeouts ...time.Duration) (conn time.Duration, read time.Duration, write time.Duration) {
	for i, t := range defaultTimeouts {
		switch i {
		case 0:
			conn = t
		case 1:
			read = t
		case 2:
			write = t
		}
	}

	ts := strings.Split(strings.Replace(t, " ", "", -1), ",")
	for i, t := range ts {
		switch i {
		case 0:
			conn = parseToDuration(t)
		case 1:
			read = parseToDuration(t)
		case 2:
			write = parseToDuration(t)
		}
	}

	return
}
