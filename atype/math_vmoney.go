package atype

import "math"

// virtual money
type VMoney Money // 1 coin = 1 money    如 chatgpt 等消耗，单次消耗低于0.1分，因此需要更大的 coin比例
const (
	VCent              VMoney  = 100  // 1 律分
	VDime              VMoney  = 1000 // 1 律币(角)
	VMoneyUnits        VMoney  = 10000
	vmoneyUnitsFloat64 float64 = 10000.0
	vmoneyUnitsInt64   int64   = 10000

	MaxVMoneyU32 VMoney = 1<<32 - 1
	MinVMoney    VMoney = -1 << 63
	MaxVMoney    VMoney = 1<<63 - 1
)

func VmoneyUnitsX(n float64) VMoney { return VMoney(math.Round(n * vmoneyUnitsFloat64)) }

func (c VMoney) Int64() int64 { return int64(c) }

// 实数形式
func (c VMoney) Real() float64 { return float64(c) / vmoneyUnitsFloat64 }

// 金币抵扣商品
//
//	.Off()  ==  .Offset(100*Percent)
func (c VMoney) Offset(rate Decimal) Money { return Money(c).MulFloor(rate) }
func (c VMoney) Off() Money                { return Money(c) }

// 用于使用计算
func (c VMoney) P() Money { return Money(c) }

// 该数量金币，对应商品购买数量
// 金币 sku 中，exchange_rate<Decimal> 表示单件商品兑换金币数量比例
// 单件商品可兑换金币 unit_coin<VMoney> = exchange_rate.ToReal * VMoneyUnits = exchange_rate * VMoneyUnits / DecimalUnits
// 总兑换数量 total_coin = unit_coin * qty
// qty = total_coin / unit_coin = (total_coin * DecimalUnits) / (exchange_rate * VMoneyUnits)
func ExchangeVMoney(exchangeRate Decimal, qty uint) VMoney {
	return VMoney(exchangeRate.Int64() * int64(qty) * vmoneyUnitsInt64 / decimalUnitsInt64)
}
func (c VMoney) ExchangeQty(exchangeRate Decimal, intHandler func(float64) float64) uint {
	return uint(intHandler(float64(c.Int64()*decimalUnitsInt64) / float64(exchangeRate.Int64()*vmoneyUnitsInt64)))
}
func (c VMoney) ExchangeQtyCeil(exchangeRate Decimal) uint {
	return c.ExchangeQty(exchangeRate, math.Ceil)
}
func (c VMoney) ExchangeQtyFloor(exchangeRate Decimal) uint {
	return c.ExchangeQty(exchangeRate, math.Floor)
}
func (a Money) VMoney() VMoney { return VMoney(a) }
