package aenum

import (
	"strconv"
	"strings"
)

type Sex uint8

const (
	NilSex   Sex = 0
	Male     Sex = 1
	Female   Sex = 2
	OtherSex Sex = 255
)

func NewSex(s interface{}) (Sex, bool) {
	var ss string
	switch v := s.(type) {
	case string:
		ss = strings.ToUpper(v)
	case uint8:
		ss = strconv.FormatUint(uint64(v), 10)
	}

	switch ss {
	case "0", "U", "UNKNOWN":
		return NilSex, true
	case "1", "M", "MALE":
		return Male, true
	case "2", "F", "FEMALE":
		return Female, true
	case "255":
		return OtherSex, true
	}
	return NilSex, false
}

func (x Sex) Uint8() uint8   { return uint8(x) }
func (x Sex) String() string { return strconv.FormatUint(uint64(x), 10) }
func (x Sex) Is(x2 Sex) bool { return x == x2 }
func (x Sex) In(args ...Sex) bool {
	for _, a := range args {
		if a == x {
			return true
		}
	}
	return false
}
func (x Sex) Name() string {
	switch x {
	case NilSex:
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
