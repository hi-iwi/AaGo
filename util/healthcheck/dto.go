package healthcheck

type Health struct {
	Time        string        `json:"time"`
	Timezone    int           `json:"timezone"` // 8  -8
	Service     string        `json:"service"`
	ServerID    string        `json:"server_id"`
	Connections []interface{} `json:"connections"`
}

type AmqpConnectionHealth struct {
	Name           string `json:"name"`
	Scheme         string `json:"scheme"`
	Host           string `json:"host"`
	Port           string `json:"port"`
	VHost          string `json:"vhost"`
	TLS            bool   `json:"tls"`
	ConnTimeoutMs  int    `json:"conn_timeout_ms"`
	ReadTimeoutMs  int    `json:"read_timeout_ms"`
	WriteTimeoutMs int    `json:"write_timeout_ms"`
	ErrMsg         string `json:"errmsg"`
}

type MysqlConnectionHealth struct {
	Name           string `json:"name"`
	Scheme         string `json:"scheme"`
	Host           string `json:"host"`
	Port           string `json:"port"`
	Db             int    `json:"db"`
	TLS            bool   `json:"tls"`
	ConnTimeoutMs  int    `json:"conn_timeout_ms"`
	ReadTimeoutMs  int    `json:"read_timeout_ms"`
	WriteTimeoutMs int    `json:"write_timeout_ms"`
	ErrMsg         string `json:"errmsg"`
}

type RedisConnectionHealth struct {
	Name           string `json:"name"`
	Scheme         string `json:"scheme"`
	Host           string `json:"host"`
	Port           string `json:"port"`
	Db             string `json:"db"`
	TLS            bool   `json:"tls"`
	ConnTimeoutMs  int    `json:"conn_timeout_ms"`
	ReadTimeoutMs  int    `json:"read_timeout_ms"`
	WriteTimeoutMs int    `json:"write_timeout_ms"`
	ErrMsg         string `json:"errmsg"`
}
