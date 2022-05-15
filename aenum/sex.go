package aenum

import (
	"github.com/hi-iwi/AaGo/dtype"
	"strconv"
	"strings"
)

type Sex uint8

const (
	UnknownSex Sex = 0
	Male       Sex = 1
	Female     Sex = 2
	OtherSex   Sex = 255
)

func NewSex(s interface{}) (Sex, bool) {
	ss := strings.ToUpper(dtype.String(s))
	switch ss {
	case "0", "U", "UNKNOWN":
		return UnknownSex, true
	case "1", "M", "MALE":
		return Male, true
	case "2", "F", "FEMALE":
		return Female, true
	case "255":
		return OtherSex, true
	}
	return UnknownSex, false
}

func (x Sex) Valid() bool    { return x <= Female || x == OtherSex }
func (x Sex) Uint8() uint8   { return uint8(x) }
func (x Sex) String() string { return strconv.FormatUint(uint64(x), 10) }

func (x Sex) Name() string {
	switch x {
	case UnknownSex:
		return "unknown sex"
	case Male:
		return "male"
	case Female:
		return "female"
	case OtherSex:
		return "other sex"
	}
	return x.String()
}
func (x Sex) Is(x2 Sex) bool { return x == x2 }
