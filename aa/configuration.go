package aa

import (
	"fmt"
	"log"

	"github.com/luexu/AaGo/util"
)

type Configuration struct {
	Service    string `yaml:"service"`
	ServerID   string `yaml:"server_id"`
	Env        string `yaml:"env"`         // dev test preprod product
	TimezoneID string `yaml:"timezone_id"` // e.g. Asia/Shanghai
	Mock       bool   `yaml:"mock"`        // using mock
}

func (a *Aa) ParseToConfiguration() {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.Configuration.Service = a.Config.Get("service").String()
	a.Configuration.ServerID = a.Config.Get("server_id").String()
	a.Configuration.Env = a.Config.Get("env").String()
	a.Configuration.TimezoneID = a.Config.Get("timezone_id").String()
	mock, _ := a.Config.Get("mock").Bool()
	a.Configuration.Mock = mock
}

func (c Configuration) Log() {
	msg := fmt.Sprintf("service %s has started! env: %s, server id: %s, timezone id: %s, mock: %v, git ver: %s", c.Service, c.Env, c.ServerID, c.TimezoneID, c.Mock, util.GitVersion())
	log.Println(msg)
	fmt.Println(msg)
}
