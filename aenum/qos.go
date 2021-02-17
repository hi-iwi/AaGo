package aenum

type Qos uint8

const (
	Qos0 Qos = 0
	Qos1 Qos = 1
	Qos2 Qos = 2
)

func (qos Qos) String() string {
	return string(qos)
}
func (qos Qos) Stringify() string {
	switch qos {
	case Qos0:
		return "Qos0"
	case Qos1:
		return "Qos1"
	case Qos2:
		return "Qos2"
	}
	return "UnknownQOS"
}
