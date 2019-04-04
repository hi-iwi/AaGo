package healthcheck

import (
	"fmt"
	"sync"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/luexu/AaGo/aa"
)

type health struct {
	app       *aa.Aa
	ConfigFmt string
}

var (
	newHealthOnce sync.Once
	healthSvc     *health
)

func NewHealth(app *aa.Aa) *health {
	newHealthOnce.Do(func() {
		healthSvc = &health{
			app:       app,
			ConfigFmt: "conn.%s",
		}
	})
	return healthSvc
}

func (s *health) Check(connections ...interface{}) Health {
	now := time.Now()
	_, offset := now.Zone()
	return Health{
		Time:        now.Format("2006-01-02 15:04:05"),
		Timezone:    offset / 3600,
		Service:     s.app.Config.Get("service").String(),
		ServerID:    s.app.Config.Get("server_id").String(),
		Connections: connections,
	}
}
func (s *health) CheckRedis(name string) (RedisConnectionHealth, error) {

	cf := s.app.Config

	connTimeout, _ := cf.Get(fmt.Sprintf(s.ConfigFmt+".conn_timeout", name), time.Second).Int()
	readTimeout, _ := cf.Get(fmt.Sprintf(s.ConfigFmt+".read_timeout", name), time.Second).Int()
	writeTimeout, _ := cf.Get(fmt.Sprintf(s.ConfigFmt+".write_timeout", name), time.Second).Int()

	h := RedisConnectionHealth{
		Name:           name,
		Scheme:         cf.Get(fmt.Sprintf(s.ConfigFmt+".scheme", name), "tcp").String(),
		Host:           cf.Get(fmt.Sprintf(s.ConfigFmt+".host", name)).String(),
		Port:           cf.Get(fmt.Sprintf(s.ConfigFmt+".port", name), "6379").String(),
		Db:             cf.Get(fmt.Sprintf(s.ConfigFmt+".db", name)).String(),
		ConnTimeoutMs:  connTimeout,
		ReadTimeoutMs:  readTimeout,
		WriteTimeoutMs: writeTimeout,
	}

	auth := cf.Get(fmt.Sprintf(s.ConfigFmt+".auth", name)).String()

	c, err := redis.DialTimeout(h.Scheme, h.Host+":"+h.Port, time.Duration(writeTimeout)*time.Millisecond, time.Duration(writeTimeout)*time.Millisecond, time.Duration(writeTimeout)*time.Millisecond)

	if err != nil {
		h.ErrMsg = fmt.Sprintf("error: %s", err)
		return h, err
	}
	defer c.Close()

	if auth != "" {
		c.Do("auth", auth)
	}

	if _, err := redis.String(c.Do("PING")); err != nil {
		h.ErrMsg = fmt.Sprintf("error: %s", err)
		return h, err
	}

	return h, err
}

func (s *health) checkMySQL() MysqlConnectionHealth {
	h := MysqlConnectionHealth{}
	return h
}
