package atype

import "math"

type Coin Money // 1 coin = 1 money    如 chatgpt 等消耗，单次消耗低于0.1分，因此需要更大的 coin比例
const (
	CentCoin        Coin = 100  // 1 律分
	DimeCoin        Coin = 1000 // 1 律币(角)
	UnitCoin        Coin = 10000
	unitCoinFloat64      = 10000.0

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

// @param ratio 汇率
func (a Money) ExchangeCoin(rate Decimal) Coin { return Coin(a.MulFloor(rate)) }
func (a Money) Coin() Coin                     { return Coin(a) }
