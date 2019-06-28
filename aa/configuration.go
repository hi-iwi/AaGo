package aa

import (
	"fmt"
	"log"
	"time"

	"github.com/luexu/AaGo/util"
)

type Configuration struct {
	Service      string `yaml:"service"`
	ServerID     string `yaml:"server_id"`
	Env          string `yaml:"env"`         // dev test preprod product
	TimezoneID   string `yaml:"timezone_id"` // e.g. "Asia/Shanghai"
	TimeLocation *time.Location
	TimeFormat   string `yaml:"time_format"` // e.g. "2006-02-01 15:04:05"
	Mock         bool   `yaml:"mock"`        // using mock
}

func (a *Aa) ParseToConfiguration() {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.Configuration.Service = a.Config.Get("service").String()
	a.Configuration.ServerID = a.Config.Get("server_id").String()
	a.Configuration.Env = a.Config.Get("env").String()

	if tz := a.Config.Get("timezone_id").String(); tz != "" {
		loc, err := time.LoadLocation(tz)
		if err != nil {
			log.Printf("invalid timezone: %s, error: %v", tz, err)
		} else {
			a.Configuration.TimezoneID = tz
			a.Configuration.TimeLocation = loc
		}
	}

	mock, _ := a.Config.Get("mock").Bool()

	a.Configuration.Mock = mock
}

func (c Configuration) Log() {
	msg := fmt.Sprintf("service %s has started! env: %s server_id: %s timezone_id: %s mock: %v git_ver: %s", c.Service, c.Env, c.ServerID, c.TimezoneID, c.Mock, util.GitVersion())
	log.Println(msg)
	fmt.Printf("%s %s\n", time.Now().In(c.TimeLocation).Format("2006-01-02 15:04:05"), msg)
}
