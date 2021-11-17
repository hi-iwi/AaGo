package aa

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/hi-iwi/AaGo/util"
)

type Configuration struct {
	Env          string // dev test preprod product
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

func (app *Aa) ParseToConfiguration() {
	app.Configuration.Env = app.Config.GetString(CkEnv)
	app.Configuration.TimeFormat = app.Config.GetString(CkTimeFormat)
	//serverID := app.Config.Get("server_id").Name()
	//app.Configuration.ServerID = serverID
	//app.Configuration.VID = svc + ":" + serverID

	if tz := app.Config.GetString(CkTimezoneID); tz != "" {
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
	n := time.Now()
	now := n.Format(c.TimeFormat) + "." + strconv.FormatInt(n.UnixMicro(), 10)
	msg := fmt.Sprintf("lauching...\nenv: %s\ntimezone_id: %s\nmock: %v\ngit_ver: %s", c.Env, c.TimezoneID, c.Mock, util.GitVersion())
	log.Println(msg)
	fmt.Println(now + " " + msg)
}

// ParseTimeout connection timeout, r timeout, w timeout, heartbeat interval
// 10s, 1000ms
func (app *Aa) ParseTimeout(t string, defaultTimeouts ...time.Duration) (conn time.Duration, read time.Duration, write time.Duration) {
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
