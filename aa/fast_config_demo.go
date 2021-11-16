package aa

import (
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

//
//// @example
/*
[mysql]
host=localhost
user=hi
password=hello
tls=false
timeout=10s,10s,10s
pool_max_idle_conns=0
pool_max_open_conns=0
pool_conn_max_life_time=0
pool_conn_max_idle_time=0
[mysql_helloworld]
schema=helloworld
[mysql_helloworld2]
host=localhost2
schema=helloworld2
user=hi2
password=hello2
tls=false
timeout=10s,10s,10s
pool_max_idle_conns=0
pool_max_open_conns=0
pool_conn_max_life_time=0
pool_conn_max_idle_time=0
*/
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

type RedisPoolConfig struct {
	// Maximum number of idle connections in the pool.
	MaxIdle int

	// Maximum number of connections allocated by the pool at a given time.
	// When zero, there is no limit on the number of connections in the pool.
	MaxActive int

	// Close connections after remaining idle for this duration. If the value
	// is zero, then idle connections are not closed. Applications should set
	// the timeout to a value less than the server's timeout.
	IdleTimeout time.Duration

	// If Wait is true and the pool is at the MaxActive limit, then Get() waits
	// for a connection to be returned to the pool before returning.
	Wait bool

	// Close connections older than this duration. If the value is zero, then
	// the pool does not close connections based on age.
	MaxConnLifetime time.Duration
}
type RedisConfig struct {
	TLS  bool
	Host string
	Auth string
	Db   uint8 // 默认0，系统配置为16个，方便flush all，但是不常用

	ConnTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Pool         RedisPoolConfig
}

// @example

/*
[redis]
host=localhost
auth=
tls=false
db=0
timeout=3s,3s,3s
pool_max_idle=0
pool_max_active=0
pool_idle_timeout=0
pool_wait=false
pool_conn_life_time=0
[redis_helloworld]
db=1
[redis_helloworld2]
host=localhost2
auth=2
tls=false
db=2
timeout=3s,3s,3s
pool_max_idle=0
pool_max_active=0
pool_idle_timeout=0
pool_wait=false
pool_conn_life_time=0
*/
func (app *Aa) tryGeRedisCfg(section string, key string) (string, error) {
	k := section + "." + key
	v, err := app.Config.MustGetString(k)
	if err == nil {
		return v, nil
	}

	return app.Config.MustGetString("redis." + key)
}

func (app *Aa) RedisConfig(section string) (RedisConfig, error) {
	host, err := app.tryGeRedisCfg(section, "host")
	if err != nil {
		return RedisConfig{}, err
	}
	auth, _ := app.tryGeRedisCfg(section, "auth") // auth 可以为空
	tls, _ := app.tryGeRedisCfg(section, "tls")
	db, _ := app.tryGeRedisCfg(section, "db")
	timeout, _ := app.tryGeRedisCfg(section, "timeout")
	ct, rt, wt := app.ParseTimeout(timeout)

	poolMaxIdle, _ := app.tryGeRedisCfg(section, "pool_max_idle")
	poolMaxActive, _ := app.tryGeRedisCfg(section, "pool_max_active")
	poolIdleTimeout, _ := app.tryGeRedisCfg(section, "pool_idle_timeout")
	poolWait, _ := app.tryGeRedisCfg(section, "pool_wait")
	poolConnLifeTime, _ := app.tryGeRedisCfg(section, "pool_conn_life_time")
	cf := RedisConfig{
		TLS:          dtype.New(tls).DefaultBool(false),
		Host:         host,
		Auth:         auth,
		Db:           dtype.New(db).DefaultUint8(0),
		ConnTimeout:  ct,
		ReadTimeout:  rt,
		WriteTimeout: wt,
		Pool: RedisPoolConfig{
			MaxIdle:         dtype.New(poolMaxIdle).DefaultInt(0),
			MaxActive:       dtype.New(poolMaxActive).DefaultInt(0),
			IdleTimeout:     time.Duration(dtype.New(poolIdleTimeout).DefaultInt64(0)) * time.Second,
			Wait:            dtype.New(poolWait).DefaultBool(false),
			MaxConnLifetime: time.Duration(dtype.New(poolConnLifeTime).DefaultInt64(0)) * time.Second,
		},
	}
	return cf, nil
}
