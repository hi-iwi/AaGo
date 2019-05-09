package healthcheck

import (
	"database/sql"
	"fmt"
	"net/url"
	"sync"
	"time"

	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/luexu/AaGo/aa"
	"github.com/streadway/amqp"
)

type health struct {
	app       *aa.Aa
	ConfigFmt string
	mtx       sync.RWMutex
	h         Health
}

var (
	newHealthOnce sync.Once
	healthSvc     *health
)

func NewHealth(app *aa.Aa) *health {
	_, offset := time.Now().In(app.Configuration.TimeLocation).Zone()

	newHealthOnce.Do(func() {
		healthSvc = &health{
			app:       app,
			ConfigFmt: "conn.%s",
			h: Health{
				TimezoneID:     app.Configuration.TimezoneID,
				TimezoneOffset: offset,
				Service:        app.Config.Get("service").String(),
				ServerID:       app.Config.Get("server_id").String(),
			},
		}
	})
	return healthSvc
}

func (s *health) UpdateRunner(name string) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.h.Runners[name] = time.Now().In(s.app.Configuration.TimeLocation).Format("2006-01-02 15:04:05")
}

func (s *health) Check(connections ...interface{}) Health {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.h.Time = time.Now().In(s.app.Configuration.TimeLocation).Format("2006-01-02 15:04:05")
	s.h.Connections = connections
	return s.h
}

func (s *health) CheckRedis(name string) (RedisConnHealth, error) {
	cf := s.app.Config
	tls, _ := cf.Get(fmt.Sprintf(s.ConfigFmt+".port", name), false).Bool()
	connTimeout, _ := cf.Get(fmt.Sprintf(s.ConfigFmt+".timeout_ms", name), 0).Int()
	readTimeout, _ := cf.Get(fmt.Sprintf(s.ConfigFmt+".read_timeout_ms", name), 0).Int()
	writeTimeout, _ := cf.Get(fmt.Sprintf(s.ConfigFmt+".write_timeout_ms", name), 0).Int()

	h := RedisConnHealth{
		Name:           name,
		Scheme:         cf.Get(fmt.Sprintf(s.ConfigFmt+".scheme", name), "tcp").String(),
		Host:           cf.Get(fmt.Sprintf(s.ConfigFmt+".host", name)).String(),
		Port:           cf.Get(fmt.Sprintf(s.ConfigFmt+".port", name), "6379").String(),
		Db:             cf.Get(fmt.Sprintf(s.ConfigFmt+".db", name)).String(),
		TLS:            tls,
		TimeoutMs:      connTimeout,
		ReadTimeoutMs:  readTimeout,
		WriteTimeoutMs: writeTimeout,
	}

	auth := cf.Get(fmt.Sprintf(s.ConfigFmt+".auth", name)).String()

	c, err := redis.DialTimeout(h.Scheme, h.Host+":"+h.Port, time.Duration(writeTimeout)*time.Millisecond, time.Duration(writeTimeout)*time.Millisecond, time.Duration(writeTimeout)*time.Millisecond)

	if err != nil {
		h.ErrMsg = fmt.Sprintf("redis dial error: %s", err)
		return h, err
	}
	defer c.Close()

	if auth != "" {
		c.Do("auth", auth)
	}

	if _, err := redis.String(c.Do("PING")); err != nil {
		h.ErrMsg = fmt.Sprintf("redis ping error: %s", err)
		return h, err
	}

	return h, err
}

func (s *health) CheckMysql(name string) (MysqlConnHealth, error) {
	cf := s.app.Config
	tls, _ := cf.Get(fmt.Sprintf(s.ConfigFmt+".port", name), false).Bool()
	connTimeout, _ := cf.Get(fmt.Sprintf(s.ConfigFmt+".timeout_ms", name), 0).Int()
	readTimeout, _ := cf.Get(fmt.Sprintf(s.ConfigFmt+".read_timeout_ms", name), 0).Int()
	writeTimeout, _ := cf.Get(fmt.Sprintf(s.ConfigFmt+".write_timeout_ms", name), 0).Int()

	h := MysqlConnHealth{
		Name:           name,
		Scheme:         cf.Get(fmt.Sprintf(s.ConfigFmt+".scheme", name), "tcp").String(),
		Host:           cf.Get(fmt.Sprintf(s.ConfigFmt+".host", name)).String(),
		Port:           cf.Get(fmt.Sprintf(s.ConfigFmt+".port", name), "3306").String(),
		Db:             cf.Get(fmt.Sprintf(s.ConfigFmt+".db", name)).String(),
		TLS:            tls,
		TimeoutMs:      connTimeout,
		ReadTimeoutMs:  readTimeout,
		WriteTimeoutMs: writeTimeout,
	}

	user := cf.Get(fmt.Sprintf(s.ConfigFmt+".user", name)).String()
	password := cf.Get(fmt.Sprintf(s.ConfigFmt+".password", name)).String()
	timezoneID := cf.Get("timezone_id", "UTC").String()
	src := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?loc=%s&timeout=%dms&readTimeout=%dms&writeTimeout=%dms", user, password, h.Scheme, h.Host, h.Port, h.Db, url.QueryEscape(timezoneID), h.TimeoutMs, h.ReadTimeoutMs, h.WriteTimeoutMs)

	db, err := sql.Open("mysql", src)
	if err != nil {
		return h, fmt.Errorf("mysql connection(%s) open error: %s", src, err)
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	if err = db.Ping(); err != nil {
		return h, fmt.Errorf("mysql connection(%s) ping error: %s", src, err)
	}

	return h, nil
}

func (s *health) CheckAmqp(name string) (AmqpConnHealth, error) {
	cf := s.app.Config
	tls, _ := cf.Get(fmt.Sprintf(s.ConfigFmt+".port", name), false).Bool()
	connTimeout, _ := cf.Get(fmt.Sprintf(s.ConfigFmt+".timeout_ms", name), 0).Int()
	readTimeout, _ := cf.Get(fmt.Sprintf(s.ConfigFmt+".read_timeout_ms", name), 0).Int()
	writeTimeout, _ := cf.Get(fmt.Sprintf(s.ConfigFmt+".write_timeout_ms", name), 0).Int()
	heartbeat, _ := cf.Get(fmt.Sprintf(s.ConfigFmt+".heartbeat", name), 0).Int()

	vhost := cf.Get(fmt.Sprintf(s.ConfigFmt+".vhost", name)).String()

	if vhost[0] == byte('/') {
		vhost = vhost[1:]
	}

	h := AmqpConnHealth{
		Name:           name,
		Scheme:         cf.Get(fmt.Sprintf(s.ConfigFmt+".scheme", name), "tcp").String(),
		Host:           cf.Get(fmt.Sprintf(s.ConfigFmt+".host", name)).String(),
		Port:           cf.Get(fmt.Sprintf(s.ConfigFmt+".port", name), "5672").String(),
		VHost:          vhost,
		TLS:            tls,
		TimeoutMs:      connTimeout,
		HeartbeatMs:    heartbeat,
		ReadTimeoutMs:  readTimeout,
		WriteTimeoutMs: writeTimeout,
	}

	user := cf.Get(fmt.Sprintf(s.ConfigFmt+".user", name)).String()
	password := cf.Get(fmt.Sprintf(s.ConfigFmt+".password", name)).String()

	url := fmt.Sprintf("amqp://%s:%s@%s:%s/%s", user, password, h.Host, h.Port, h.VHost)

	conn, err := amqp.Dial(url)
	if err != nil {
		return h, fmt.Errorf("failed to connect to AMQP broker %s: %s", url, err)
	}
	defer conn.Close()

	return h, nil
}
