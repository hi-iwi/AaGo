package aa

type Configuration struct {
	Service    string
	ServiceID  string
	Env        string // dev test preprod product
	TimezoneID string // e.g. Asia/Shanghai
}
