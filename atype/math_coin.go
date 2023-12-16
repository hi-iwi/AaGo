package atype

import "math"

type Coin Money // 1 coin = 1 money    如 chatgpt 等消耗，单次消耗低于0.1分，因此需要更大的 coin比例
const (
	CentCoin        Coin    = 100  // 1 律分
	DimeCoin        Coin    = 1000 // 1 律币(角)
	UnitCoin        Coin    = 10000
	unitCoinFloat64 float64 = 10000.0
	unitCoinInt64   int64   = 10000

	MaxCoinU32 Coin = 1<<32 - 1
	MinCoin    Coin = -1 << 63
	MaxCoin    Coin = 1<<63 - 1
)

func CoinUnit(n float64) Coin { return Coin(math.Round(n * unitCoinFloat64)) }
func CoinUnitN(n uint) Coin   { return Coin(n) * UnitCoin }

func (c Coin) Int64() int64 { return int64(c) }

// 金币抵扣商品
//  .Off()  ==  .Offset(100*Percent)
func (c Coin) Offset(rate Decimal) Money { return Money(c).MulFloor(rate) }
func (c Coin) Off() Money                { return Money(c) }

// 用于使用计算
func (c Coin) P() Money { return Money(c) }

// 该数量金币，对应商品购买数量
// 金币 sku 中，exchange_rate<Decimal> 表示单件商品兑换金币数量比例
// 单件商品可兑换金币 unit_coin<Coin> = exchange_rate.Real * UnitCoin = exchange_rate * UnitCoin / UnitDecimal
// 总兑换数量 total_coin = unit_coin * qty
// qty = total_coin / unit_coin = (total_coin * UnitDecimal) / (exchange_rate * UnitCoin)
func (c Coin) ExchangeQtyCeil(exchangeRate Decimal) uint {
	return uint(math.Ceil(float64(c.Int64()*unitDecimalInt64) / float64(exchangeRate.Int64()*unitCoinInt64)))
}
func (c Coin) ExchangeQtyFloor(exchangeRate Decimal) uint {
	return uint(math.Floor(float64(c.Int64()*unitDecimalInt64) / float64(exchangeRate.Int64()*unitCoinInt64)))
}
func (a Money) Coin() Coin { return Coin(a) }
