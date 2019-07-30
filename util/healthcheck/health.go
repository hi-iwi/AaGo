package healthcheck

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"github.com/luexu/AaGo/aa"
	"github.com/luexu/dtype"
	"github.com/streadway/amqp"
)

type health struct {
	app *aa.Aa
	mtx sync.RWMutex
	h   Health
}

var (
	newHealthOnce sync.Once
	healthSvc     *health
)

func NewHealth(app *aa.Aa) *health {
	_, offset := time.Now().In(app.Configuration.TimeLocation).Zone()

	newHealthOnce.Do(func() {
		healthSvc = &health{
			app: app,
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

func (s *health) getConf(name string, suffix string, defaultValue ...interface{}) *dtype.Dtype {
	k := "driver." + name + "_" + suffix
	return s.app.Config.Get(k, defaultValue...)
}

func (s *health) CheckRedis(name string) (RedisConnHealth, error) {
	tls, _ := s.getConf(name, "tls", false).Bool()
	scheme := s.getConf(name, "scheme", "tcp").String()
	host := s.getConf(name, "host").String()
	db := s.getConf(name, "db", "0").String()
	auth := s.getConf(name, "auth").String()

	ct, rt, wt, _ := s.app.ParseTimeout("driver."+name+"_timeout", 3*time.Second, 3*time.Second, 3*time.Second)

	h := RedisConnHealth{
		Name:    name,
		Scheme:  scheme,
		Host:    host,
		Db:      db,
		TLS:     tls,
		Timeout: s.getConf(name, "timeout").String(),
	}

	c, err := redis.DialTimeout(h.Scheme, host, ct, rt, wt)

	if err != nil {
		h.ErrMsg = "redis dial error: " + err.Error()
		return h, err
	}
	defer c.Close()

	if auth != "" {
		c.Do("auth", auth)
	}

	if _, err := redis.String(c.Do("PING")); err != nil {
		h.ErrMsg = "redis ping error: " + err.Error()
		return h, err
	}

	return h, err
}

func (s *health) CheckMysql(name string) (MysqlConnHealth, error) {

	tls, _ := s.getConf(name, "tls", false).Bool()
	scheme := s.getConf(name, "scheme", "tcp").String()
	host := s.getConf(name, "host").String()
	db := s.getConf(name, "db").String()
	user := s.getConf(name, "user").String()
	password := s.getConf(name, "password").String()
	//loc := url.QueryEscape(s.app.Config.Get("timezone_id", "UTC").String())
	charset := s.getConf(name, "charset", "utf8mb4").String()

	ct, rt, wt, _ := s.app.ParseTimeout("driver."+name+"_timeout", 3*time.Second, 3*time.Second, 3*time.Second)

	h := MysqlConnHealth{
		Name:    name,
		Scheme:  scheme,
		Host:    host,
		Db:      db,
		TLS:     tls,
		Timeout: s.getConf(name, "timeout").String(),
	}

	ctn := strconv.Itoa(int(ct / time.Millisecond))
	rtn := strconv.Itoa(int(rt / time.Millisecond))
	wtn := strconv.Itoa(int(wt / time.Millisecond))
	src := user + ":" + password + "@" + scheme + "(" + host + ")/" + db + "?charset=" + charset + "&timeout=" + ctn + "ms&readTimeout=" + rtn + "ms&writeTimeout=" + wtn + "ms"

	conn, err := sql.Open("mysql", src)
	if err != nil {
		return h, errors.New("mysql connection(" + src + ") open error: " + err.Error())
	}
	defer conn.Close()

	// Open doesn't open a connection. Validate DSN data:
	if err = conn.Ping(); err != nil {
		return h, errors.New("mysql connection(" + src + ") ping error: " + err.Error())
	}

	return h, nil
}

func (s *health) CheckAmqp(name string) (AmqpConnHealth, error) {
	tls, _ := s.getConf(name, "tls", false).Bool()
	scheme := s.getConf(name, "scheme", "tcp").String()
	host := s.getConf(name, "host").String()
	vhost := s.getConf(name, "vhost").String()
	user := s.getConf(name, "user").String()
	password := s.getConf(name, "password").String()

	if vhost[0] == byte('/') {
		vhost = vhost[1:]
	}

	h := AmqpConnHealth{
		Name:    name,
		Scheme:  scheme,
		Host:    host,
		VHost:   vhost,
		TLS:     tls,
		Timeout: s.getConf(name, "timeout").String(),
	}

	url := "amqp://" + user + ":" + password + "@" + host + "/" + vhost

	conn, err := amqp.Dial(url)
	if err != nil {
		return h, fmt.Errorf("failed to connect to AMQP broker %s: %s", url, err)
	}
	defer conn.Close()

	return h, nil
}
