package aa

type Configuration struct {
	Service    string `yaml:"service"`
	ServerID   string `yaml:"server_id"`
	Env        string `yaml:"env"`         // dev test preprod product
	TimezoneID string `yaml:"timezone_id"` // e.g. Asia/Shanghai
	Mock       bool   // using mock
}
