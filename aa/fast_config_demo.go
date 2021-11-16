package aa

import "time"

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
[mysql_helloworld]
host=localhost
schema=helloworld
user=hi
password=hello
tls=false
timeout=10s,10s,10s
pool_max_idle_conns=0
pool_max_open_conns=0
pool_conn_max_life_time=0
pool_conn_max_idle_time=0
*/
func (app *Aa) MysqlConfig(section string) (MysqlConfig, error) {
	host, err := app.Config.MustGetString(section + ".host")
	if err != nil {
		return MysqlConfig{}, err
	}
	schema, err := app.Config.MustGetString(section + ".schema")
	if err != nil {
		return MysqlConfig{}, err
	}
	user, err := app.Config.MustGetString(section + ".user")
	if err != nil {
		return MysqlConfig{}, err
	}
	password, err := app.Config.MustGetString(section + ".password")
	if err != nil {
		return MysqlConfig{}, err
	}

	tls := app.Config.GetString(section+".tls", "false")
	timeout := app.Config.GetString(section+".timeout", "10s,10s,10s")
	ct, rt, wt := app.ParseTimeout(timeout)
	poolMaxIdleConns := app.Config.Get(section + ".pool_max_idle_conns").DefaultInt(0)
	pooMaxOpenConns := app.Config.Get(section + ".pool_max_open_conns").DefaultInt(0)
	poolConnMaxLifetime := app.Config.Get(section + ".pool_conn_max_life_time").DefaultInt64(0)
	pooConnMaxIdleTime := app.Config.Get(section + ".pool_conn_max_idle_time").DefaultInt64(0)
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
			MaxIdleConns:    poolMaxIdleConns,
			MaxOpenConns:    pooMaxOpenConns,
			ConnMaxLifetime: time.Duration(poolConnMaxLifetime) * time.Second,
			ConnMaxIdleTime: time.Duration(pooConnMaxIdleTime) * time.Second,
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
[redis_helloworld]
host=localhost
auth=
tls=false
db=0
timeout=3s,3s,3s
pool_max_idle=0
pool_max_active=0
pool_wait=false
pool_conn_life_time=0
*/



func (app *Aa) RedisConfig(section string) (RedisConfig, error) {
	host, err := app.Config.MustGetString(section + ".host")
	if err != nil {
		return RedisConfig{}, err
	}
	auth := app.Config.GetString(section + ".auth")
	tls, _ := app.Config.Get(section+".tls", false).Bool()
	db, _ := app.Config.Get(section+".db", 0).Uint8()
	timeout := app.Config.Get(section+".timeout", " 3s, 3s, 3s").String()
	ct, rt, wt := app.ParseTimeout(timeout)

	poolMaxIdle := app.Config.Get(section + ".pool_max_idle").DefaultInt(0)
	poolMaxActive := app.Config.Get(section + ".pool_max_active").DefaultInt(0)
	poolIdleTimeout := app.Config.Get(section + ".pool_idle_timeout").DefaultInt64(0)
	poolWait := app.Config.Get(section + ".pool_wait").DefaultBool(false)
	poolConnLifeTime := app.Config.Get(section + ".pool_conn_life_time").DefaultInt64(0)
	cf := RedisConfig{
		TLS:          tls,
		Host:         host,
		Auth:         auth,
		Db:           db,
		ConnTimeout:  ct,
		ReadTimeout:  rt,
		WriteTimeout: wt,
		Pool: RedisPoolConfig{
			MaxIdle:         poolMaxIdle,
			MaxActive:       poolMaxActive,
			IdleTimeout:     time.Duration(poolIdleTimeout) * time.Second,
			Wait:            poolWait,
			MaxConnLifetime: time.Duration(poolConnLifeTime) * time.Second,
		},
	}
	return cf, nil
}
