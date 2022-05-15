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
	return c, c.Valid()
}
func (c Currency) Valid() bool {return true}
func (c Currency) Raw() uint16 {return uint16(c)}
func (c Currency) String() string {return strconv.FormatUint(uint64(c), 10)}
func (c Currency) Name() string {
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
