package healthcheck

type Health struct {
	Time           string        `json:"time"`
	TimezoneID     string        `json:"timezone_id"`     // e.g. Asia/Shanghai
	TimezoneOffset int           `json:"timezone_offset"` // seconds, e.g. +8 == 28800
	Service        string        `json:"service"`
	ServerID       string        `json:"server_id"`
	Connections    []interface{} `json:"connections"`
}

type AmqpConnHealth struct {
	Name           string `json:"name"`
	Scheme         string `json:"scheme"`
	Host           string `json:"host"`
	Port           string `json:"port"`
	VHost          string `json:"vhost"`
	TLS            bool   `json:"tls"`
	TimeoutMs      int    `json:"timeout_ms"`
	ReadTimeoutMs  int    `json:"read_timeout_ms"`
	WriteTimeoutMs int    `json:"write_timeout_ms"`
	ErrMsg         string `json:"errmsg"`
}

type MysqlConnHealth struct {
	Name           string `json:"name"`
	Scheme         string `json:"scheme"`
	Host           string `json:"host"`
	Port           string `json:"port"`
	Db             string `json:"db"`
	TLS            bool   `json:"tls"`
	TimeoutMs      int    `json:"timeout_ms"`
	ReadTimeoutMs  int    `json:"read_timeout_ms"`
	WriteTimeoutMs int    `json:"write_timeout_ms"`
	ErrMsg         string `json:"errmsg"`
}

type RedisConnHealth struct {
	Name           string `json:"name"`
	Scheme         string `json:"scheme"`
	Host           string `json:"host"`
	Port           string `json:"port"`
	Db             string `json:"db"`
	TLS            bool   `json:"tls"`
	TimeoutMs      int    `json:"timeout_ms"`
	ReadTimeoutMs  int    `json:"read_timeout_ms"`
	WriteTimeoutMs int    `json:"write_timeout_ms"`
	ErrMsg         string `json:"errmsg"`
}
