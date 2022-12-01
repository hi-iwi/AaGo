package aenum

import "strconv"

type Currency uint16

const (
	UnknownCurrency = Currency(0)
	USD             = Currency(America)
	CNY             = Currency(China)
	HKD             = Currency(HongKong)
	MOP             = Currency(Macau)
	JPY             = Currency(Japan)
	KHR             = Currency(Cambodia)
	SGD             = Currency(Singapore)
)

func NewCurrency(x uint16) (Currency, bool) {
	c := Currency(x)
	return c, true
}
func (c Currency) Uint16() uint16   { return uint16(c) }
func (c Currency) String() string   { return strconv.FormatUint(uint64(c), 10) }
func (c Currency) Is(x uint16) bool { return c.Uint16() == x }
func (c Currency) In(args ...Currency) bool {
	for _, a := range args {
		if a == c {
			return true
		}
	}
	return false
}

func (c Currency) Code() string {
	switch c {
	case USD:
		return "USD"
	case CNY:
		return "CNY"
	case HKD:
		return "HKD"
	}
	return "unknown currency"
}
