package healthcheck

import (
	"database/sql"
	"errors"
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"github.com/hi-iwi/AaGo/aa"
	"github.com/hi-iwi/AaGo/dtype"
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
				Name:           app.Config.Get("name").String(),
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

func (s *health) CheckRedis(name string) (h RedisConnHealth, err error) {
	var conn redis.Conn
	f := s.app.RedisConfig(name)
	h.Config = f
	conn, err = redis.Dial("tcp", f.Host, redis.DialConnectTimeout(f.ConnTimeout), redis.DialReadTimeout(f.ReadTimeout), redis.DialWriteTimeout(f.WriteTimeout))
	if err != nil {
		h.ErrMsg = "redis dial error: " + err.Error()
		return
	}
	defer conn.Close()

	if f.Auth != "" {
		conn.Do("auth", f.Auth)
	}
	conn.Do("SELECT", f.Db)
	if _, err := redis.String(conn.Do("PING")); err != nil {
		h.ErrMsg = "redis ping error: " + err.Error()
		return
	}

	return
}

func (s *health) CheckMysql(xschema string) (h MysqlConnHealth, err error) {
	var conn *sql.DB
	f := s.app.MysqlConfig(xschema)
	h.Config = f
	ct := f.ConnTimeout.Milliseconds()
	rt := f.ReadTimeout.Milliseconds()
	wt := f.WriteTimeout.Milliseconds()
	src := fmt.Sprintf("%s:%s@tcp(%s)/%s?timeout=%dms&readTimeout=%dms&writeTimeout=%dms", f.User, f.Password, f.Host, f.Schema, ct, rt, wt)

	conn, err = sql.Open("mysql", src)
	if err != nil {
		err = errors.New("mysql connection(" + src + ") open error: " + err.Error())
		return
	}
	defer conn.Close()

	// Open doesn't open a connection. Validate DSN data:
	if err = conn.Ping(); err != nil {
		err = errors.New("mysql connection(" + src + ") ping error: " + err.Error())
		return
	}

	return
}
