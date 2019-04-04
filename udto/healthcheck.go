package udto

type Health struct {
	Time        string        `json:"time"`
	Timezone    string        `json:"timezone"`
	Service     string        `json:"service"`
	ServerID    string        `json:"server_id"`
	Connections []interface{} `json:"connections"`
}

type AmqpConnectionHealth struct {
	Scheme       string `json:"scheme"`
	Host         string `json:"host"`
	Port         string `json:"port"`
	VHost        string `json:"vhost"`
	TLS          bool   `json:"tls"`
	ReadTimeout  int    `json:"read_timeout"`
	WriteTimeout int    `json:"write_timeout"`
	Status       string `json:"status"`
	Msg          string `json:"msg"`
}
type MysqlConnectionHealth struct {
	Scheme       string `json:"scheme"`
	Host         string `json:"host"`
	Port         string `json:"port"`
	Db           int    `json:"db"`
	TLS          bool   `json:"tls"`
	ReadTimeout  int    `json:"read_timeout"`
	WriteTimeout int    `json:"write_timeout"`
	Status       string `json:"status"`
	Msg          string `json:"msg"`
}
type RedisConnectionHealth struct {
	Scheme       string `json:"scheme"`
	Host         string `json:"host"`
	Port         string `json:"port"`
	Db           int    `json:"db"`
	TLS          bool   `json:"tls"`
	ReadTimeout  int    `json:"read_timeout"`
	WriteTimeout int    `json:"write_timeout"`
	Status       string `json:"status"`
	Msg          string `json:"msg"`
}
