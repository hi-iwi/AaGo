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

func SexValidator(s uint8) bool {
	x := Sex(s)
	return x == Male || x == Female || x == OtherSex
}
func NewSex(s uint8) Sex {
	if ok := SexValidator(s); ok {
		return Sex(s)
	}
	return NilSex
}
func ToSex(s string) Sex {
	s = strings.ToUpper(s)
	switch s {
	case "0", "U", "UNKNOWN":
		return NilSex
	case "1", "M", "MALE", "男":
		return Male
	case "2", "F", "FEMALE", "女":
		return Female
	case "255":
		return OtherSex
	}
	return NilSex
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
		return "*"
	case Male:
		return "male"
	case Female:
		return "female"
	case OtherSex:
		return "*"
	}
	return x.String()
}
func (x Sex) NameCn() string {
	switch x {
	case NilSex:
		return "*"
	case Male:
		return "男"
	case Female:
		return "女"
	case OtherSex:
		return "*"
	}
	return x.String()
}
