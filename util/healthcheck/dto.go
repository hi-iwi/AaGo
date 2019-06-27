package healthcheck

type Health struct {
	Time           string            `json:"time"`
	TimezoneID     string            `json:"timezone_id"`     // e.g. Asia/Shanghai
	TimezoneOffset int               `json:"timezone_offset"` // seconds, e.g. +8 == 28800
	Service        string            `json:"service"`
	ServerID       string            `json:"server_id"`
	Connections    []interface{}     `json:"connections"`
	Runners        map[string]string `json:"runners"` // {co1: "2019-05-01 00:00:00", co2: ""} 每个死循环运行的协程，都必须经常更新时间到这里
}

type AmqpConnHealth struct {
	Name    string `json:"name"`
	Scheme  string `json:"scheme"`
	Host    string `json:"host"`
	VHost   string `json:"vhost"`
	TLS     bool   `json:"tls"`
	Timeout string `json:"timeout"`
	ErrMsg  string `json:"errmsg"`
}

type MysqlConnHealth struct {
	Name    string `json:"name"`
	Scheme  string `json:"scheme"`
	Host    string `json:"host"`
	Db      string `json:"db"`
	TLS     bool   `json:"tls"`
	Timeout string `json:"timeout"`
	ErrMsg  string `json:"errmsg"`
}

type RedisConnHealth struct {
	Name    string `json:"name"`
	Scheme  string `json:"scheme"`
	Host    string `json:"host"`
	Db      string `json:"db"`
	TLS     bool   `json:"tls"`
	Timeout string `json:"timeout"`
	ErrMsg  string `json:"errmsg"`
}
