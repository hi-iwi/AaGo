package aenum

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

func (currency Currency) String() string {
	return string(currency)
}
func (currency Currency) Stringify() string {
	switch currency {
	case USD:
		return "USD"
	case CNY:
		return "CNY"
	case HKD:
		return "HKD"
	}
	return "UnknownCurrency"
}
