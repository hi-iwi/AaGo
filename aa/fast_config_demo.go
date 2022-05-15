package aa

import (
	"github.com/go-redis/redis/v8"
	"github.com/hi-iwi/AaGo/dtype"
	"time"
)

type MysqlPoolConfig struct {
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
}

// https://github.com/go-sql-driver/mysql/
type MysqlConfig struct {
	Schema   string // dbname
	User     string
	Password string
	// Scheme   string // tcp|unix，只支持tcp，unix仅本地可用
	TLS  string // 默认 false，Valid Values:   true, false, skip-verify, preferred, <name>
	Host string
	// Charset  string  不建议用，应该服务器默认设置

	// mysql客户端在尝试与mysql服务器建立连接时，mysql服务器返回错误握手协议前等待客户端数据包的最大时限。默认10秒。
	ConnTimeout  time.Duration // 使用时，需要设置单位，s, ms等。Timeout for establishing connections, aka dial timeout
	ReadTimeout  time.Duration // 使用时，需要设置单位，s, ms等。I/O read timeout.
	WriteTimeout time.Duration // 使用时，需要设置单位，s, ms等。I/O write timeout.

	Pool MysqlPoolConfig
}

func (app *Aa) tryGetMysqlCfg(section string, key string) (string, error) {
	k := section + "." + key
	v, err := app.Config.MustGetString(k)
	if err == nil {
		return v, nil
	}

	return app.Config.MustGetString("mysql." + key)
}
func (app *Aa) MysqlConfig(section string) (MysqlConfig, error) {
	host, err := app.tryGetMysqlCfg(section, "host")
	if err != nil {
		return MysqlConfig{}, err
	}
	schema, err := app.tryGetMysqlCfg(section, "schema")
	if err != nil {
		return MysqlConfig{}, err
	}
	user, err := app.tryGetMysqlCfg(section, "user")
	if err != nil {
		return MysqlConfig{}, err
	}
	password, err := app.tryGetMysqlCfg(section, "password")
	if err != nil {
		return MysqlConfig{}, err
	}

	tls, _ := app.tryGetMysqlCfg(section, "tls")
	timeout, _ := app.tryGetMysqlCfg(section, "timeout")
	ct, rt, wt := app.ParseTimeout(timeout)
	poolMaxIdleConns, _ := app.tryGetMysqlCfg(section, "pool_max_idle_conns")
	poolMaxOpenConns, _ := app.tryGetMysqlCfg(section, "pool_max_open_conns")
	poolConnMaxLifetime, _ := app.tryGetMysqlCfg(section, "pool_conn_max_life_time")
	poolConnMaxIdleTime, _ := app.tryGetMysqlCfg(section, "pool_conn_max_idle_time")
	cf := MysqlConfig{
		Schema:       schema,
		User:         user,
		Password:     password,
		TLS:          tls,
		Host:         host,
		ConnTimeout:  ct,
		ReadTimeout:  rt,
		WriteTimeout: wt,
		Pool: MysqlPoolConfig{
			MaxIdleConns:    dtype.New(poolMaxIdleConns).DefaultInt(0),
			MaxOpenConns:    dtype.New(poolMaxOpenConns).DefaultInt(0),
			ConnMaxLifetime: time.Duration(dtype.New(poolConnMaxLifetime).DefaultInt64(0)) * time.Second,
			ConnMaxIdleTime: time.Duration(dtype.New(poolConnMaxIdleTime).DefaultInt64(0)) * time.Second,
		},
	}
	return cf, nil
}

func (app *Aa) tryGeRedisCfg(section string, key string) (string, error) {
	k := section + "." + key
	v, err := app.Config.MustGetString(k)
	if err == nil {
		return v, nil
	}
	return app.Config.MustGetString("redis." + key)
}

func (app *Aa) RedisConfig(section string) (*redis.Options, error) {
	addr, err := app.tryGeRedisCfg(section, "addr")
	if err != nil {
		return nil, err
	}
	network, _ := app.tryGeRedisCfg(section, "network")
	username, _ := app.tryGeRedisCfg(section, "username") // username 可以为空
	password, _ := app.tryGeRedisCfg(section, "password") // password 可以为空
	db, _ := app.tryGeRedisCfg(section, "db")
	maxRetries, _ := app.tryGeRedisCfg(section, "max_retries")
	minRetryBackoff, _ := app.tryGeRedisCfg(section, "min_retry_backoff")
	maxRetryBackoff, _ := app.tryGeRedisCfg(section, "max_retry_backoff")
	dialTimeout, _ := app.tryGeRedisCfg(section, "dial_timeout")
	readTimeout, _ := app.tryGeRedisCfg(section, "read_timeout")
	writeTimeout, _ := app.tryGeRedisCfg(section, "write_timeout")
	poolFIFO, _ := app.tryGeRedisCfg(section, "pool_fifo")
	poolSize, _ := app.tryGeRedisCfg(section, "pool_size")
	minIdleConns, _ := app.tryGeRedisCfg(section, "min_idle_conns")
	maxConnAge, _ := app.tryGeRedisCfg(section, "max_conn_age")
	poolTimeout, _ := app.tryGeRedisCfg(section, "pool_timeout")
	idleTimeout, _ := app.tryGeRedisCfg(section, "idle_timeout")
	idleCheckFrequency, _ := app.tryGeRedisCfg(section, "idle_check_frequency")

	opt := redis.Options{
		Network:            network,
		Addr:               addr, //  127.0.0.1:6379
		Dialer:             nil,
		OnConnect:          nil,
		Username:           username,
		Password:           password,
		DB:                 dtype.New(db).DefaultInt(0),
		MaxRetries:         dtype.New(maxRetries).DefaultInt(0),
		MinRetryBackoff:    time.Duration(dtype.New(minRetryBackoff).DefaultInt64(0)),
		MaxRetryBackoff:    time.Duration(dtype.New(maxRetryBackoff).DefaultInt64(0)),
		DialTimeout:        time.Duration(dtype.New(dialTimeout).DefaultInt64(0)),
		ReadTimeout:        time.Duration(dtype.New(readTimeout).DefaultInt64(0)),
		WriteTimeout:       time.Duration(dtype.New(writeTimeout).DefaultInt64(0)),
		PoolFIFO:           dtype.New(poolFIFO).DefaultBool(false),
		PoolSize:           dtype.New(poolSize).DefaultInt(0),
		MinIdleConns:       dtype.New(minIdleConns).DefaultInt(0),
		MaxConnAge:         time.Duration(dtype.New(maxConnAge).DefaultInt64(0)),
		PoolTimeout:        time.Duration(dtype.New(poolTimeout).DefaultInt64(0)),
		IdleTimeout:        time.Duration(dtype.New(idleTimeout).DefaultInt64(0)),
		IdleCheckFrequency: time.Duration(dtype.New(idleCheckFrequency).DefaultInt64(0)),
		TLSConfig:          nil,
		Limiter:            nil,
	}
	return &opt, nil
}
