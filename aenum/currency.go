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
func (c Currency) Valid() bool {
	return true
}
func (currency Currency) String() string {
	return strconv.FormatUint(uint64(currency), 10)
}
func (currency Currency) Name() string {
	switch currency {
	case USD:
		return "USD"
	case CNY:
		return "CNY"
	case HKD:
		return "HKD"
	}
	return "unknown currency"
}
