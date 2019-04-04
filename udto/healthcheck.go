package udto

type Health struct {
	Time        string        `json:"time"`
	Timezone    string        `json:"timezone"`
	Service     string        `json:"service"`
	ServerID    string        `json:"server_id"`
	Connections []interface{} `json:"connections"`
}

type AmqpConnectionHealth struct {
	Name         string `json:"name"`
	Scheme       string `json:"scheme"`
	Host         string `json:"host"`
	Port         string `json:"port"`
	VHost        string `json:"vhost"`
	TLS          bool   `json:"tls"`
	ReadTimeout  int    `json:"read_timeout"`
	WriteTimeout int    `json:"write_timeout"`
	ErrMsg       string `json:"errmsg"`
}

type MysqlConnectionHealth struct {
	Name         string `json:"name"`
	Scheme       string `json:"scheme"`
	Host         string `json:"host"`
	Port         string `json:"port"`
	Db           int    `json:"db"`
	TLS          bool   `json:"tls"`
	ReadTimeout  int    `json:"read_timeout"`
	WriteTimeout int    `json:"write_timeout"`
	ErrMsg       string `json:"errmsg"`
}

type RedisConnectionHealth struct {
	Name         string `json:"name"`
	Scheme       string `json:"scheme"`
	Host         string `json:"host"`
	Port         string `json:"port"`
	Db           string `json:"db"`
	TLS          bool   `json:"tls"`
	ReadTimeout  int    `json:"read_timeout"`
	WriteTimeout int    `json:"write_timeout"`
	ErrMsg       string `json:"errmsg"`
}
