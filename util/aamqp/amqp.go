package aamqp

import (
	"time"

	"github.com/streadway/amqp"
)

const (
	defaultProduct           = "https://github.com/luexu/AaGo"
	defaultVersion           = "Aario"
	defaultConnectionTimeout = 5 * time.Second
	defaultHeartbeat         = 10 * time.Second
	defaultLocale            = "en_US"
)

type ConnectionConfig struct {
	Timeout   time.Duration
	Heartbeat time.Duration
	Locale    string
}

func (cc ConnectionConfig) withDefault() ConnectionConfig {
	if cc.Locale == "" {
		cc.Locale = defaultLocale
	}
	if cc.Timeout <= 0*time.Second {
		cc.Timeout = defaultConnectionTimeout
	}
	if cc.Heartbeat <= 0*time.Second {
		cc.Heartbeat = defaultHeartbeat
	}
	return cc
}

type Exchange struct {
	Name       string
	Kind       string
	Durable    bool
	AutoDelete bool // delete when usused
	Internal   bool
	NoWait     bool
	Args       amqp.Table
}

type Queue struct {
	Name       string
	Durable    bool
	AutoDelete bool // delete when usused
	Exclusive  bool
	NoWait     bool
	Args       amqp.Table
}

type QueueBinding struct {
	Key      string
	Exchange string
	NoWait   bool
	Args     amqp.Table
}

type BasicQos struct {
	PrefetchSize  int
	PrefetchCount int
	Global        bool
}

type ConsumeParams struct {
	AutoAck   bool
	Exclusive bool
	NoLocal   bool
	NoWait    bool
	Args      amqp.Table
}

type PublishParams struct {
	Exchange  string
	Key       string
	Mandatory bool
	Immediate bool
	Msg       amqp.Publishing
}
