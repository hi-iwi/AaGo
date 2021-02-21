package aa

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hi-iwi/AaGo/util"
)

type Configuration struct {
	Name         string `yaml:"name"`
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
	app.Configuration.Name = app.Config.Get(CkService).String()
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
	msg := fmt.Sprintf("starting service %s\nenv: %s\ntimezone_id: %s\nmock: %v\ngit_ver: %s", c.Name, c.Env, c.TimezoneID, c.Mock, util.GitVersion())
	log.Println(msg)
	fmt.Println(msg)
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

// https://github.com/go-sql-driver/mysql/
type MysqlConfig struct {
	Schema   string // dbname
	User     string
	Password string
	// Scheme   string // tcp|unix，只支持tcp，unix仅本地可用
	TLS  string // 默认 flase，Valid Values:   true, false, skip-verify, preferred, <name>
	Host string
	// Charset  string  不建议用，应该服务器默认设置

	// mysql客户端在尝试与mysql服务器建立连接时，mysql服务器返回错误握手协议前等待客户端数据包的最大时限。默认10秒。
	ConnTimeout  time.Duration // 使用时，需要设置单位，s, ms等。Timeout for establishing connections, aka dial timeout
	ReadTimeout  time.Duration // 使用时，需要设置单位，s, ms等。I/O read timeout.
	WriteTimeout time.Duration // 使用时，需要设置单位，s, ms等。I/O write timeout.
}

func (app *Aa) MysqlConfig(xschema string) MysqlConfig {
	tls := app.Config.Get("driver.mysql_tls_"+xschema, "false").String()
	schema := app.Config.Get("driver.mysql_schema_"+xschema, xschema).String()
	user := app.Config.Get("driver.mysql_user_" + xschema).String()
	password := app.Config.Get("driver.mysql_password_" + xschema).String()
	host := app.Config.Get("driver.mysql_host_" + xschema).String()
	timeout := app.Config.Get("driver.mysql_timeout_"+xschema, "10s,10s,10s").String()
	ct, rt, wt := app.ParseTimeout(timeout)
	cf := MysqlConfig{
		Schema:       schema,
		User:         user,
		Password:     password,
		TLS:          tls,
		Host:         host,
		ConnTimeout:  ct,
		ReadTimeout:  rt,
		WriteTimeout: wt,
	}
	return cf
}

type RedisConfig struct {
	TLS  bool
	Host string
	Auth string
	Db   uint8 // 默认0，系统配置为16个，方便flush all，但是不常用

	ConnTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func (app *Aa) RedisConfig(n string) RedisConfig {
	tls, _ := app.Config.Get("driver.redis_tls_"+n, false).Bool()
	host := app.Config.Get("driver.redis_host_" + n).String()
	auth := app.Config.Get("driver.redis_auth_" + n).String()
	db, _ := app.Config.Get("driver.redis_db_"+n, 0).Uint8()
	timeout := app.Config.Get("driver.redis_timeout_"+n, " 3s, 3s, 3s").String()
	ct, rt, wt := app.ParseTimeout(timeout)
	cf := RedisConfig{
		TLS:          tls,
		Host:         host,
		Auth:         auth,
		Db:           db,
		ConnTimeout:  ct,
		ReadTimeout:  rt,
		WriteTimeout: wt,
	}
	return cf
}
