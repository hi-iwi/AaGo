package aa

import "time"

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
}

func (app *Aa) MysqlConfig(section string) (MysqlConfig, error) {
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
	host, err := app.Config.MustGetString(section + ".host")
	if err != nil {
		return MysqlConfig{}, err
	}
	tls := app.Config.GetString(section+".tls", "false")
	timeout := app.Config.GetString(section+".timeout", "10s,10s,10s")
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
	return cf, nil
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
	cf := RedisConfig{
		TLS:          tls,
		Host:         host,
		Auth:         auth,
		Db:           db,
		ConnTimeout:  ct,
		ReadTimeout:  rt,
		WriteTimeout: wt,
	}
	return cf, nil
}
