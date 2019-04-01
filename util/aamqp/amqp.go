package aamqp

import "github.com/streadway/amqp"

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
