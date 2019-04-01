package queue

import "time"

type Qos uint8

const (
	Qos0 = Qos(0)
	Qos1 = Qos(1)
	Qos2 = Qos(2)
)

type Queue interface {
	Pub(channel string, message string, qos Qos) error
	Sub(channels interface{}, hanldler func(payload string), timeout time.Duration) error
}
