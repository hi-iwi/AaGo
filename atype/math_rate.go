package atype

import "strconv"

// rate 不同单位间比，比如汇率、速度；  ratio 是相同单位间比，80km/h : 20km/h
type Rate16 int16 // 需要转换为 Rate 使用；-327.68% ~ 327.67%  即 -3.2768 ~ 3.2767
type Rate24 Int24 // 需要转换为 Rate 使用； -83886.08% ~ 83886.07%   即 -838.8608 ~ -838.8607
type Rate int     // 范围： -21474836.48% - 21474836.47%
const (
	Thousandth     Rate = 10              // 千分比
	Percent             = 10 * Thousandth // 百分比
	PercentFloat64      = 100.0
	DecimalAug          = 10000.0 // 小数转百分比扩大100倍

)

// @param n 本身就是转换后的值，如10000，即表示为 100*PercentFloat64，即 100%
func NewRate(n int) Rate { return Rate(n) }

// ToPercent(80.01) 表示 80.01%
// 若是整数，则直接使用  80*Percent 即可
func ToPercent(n float64) Rate { return NewRate(int(n * PercentFloat64)) }

// 范围： -327.68% ~ 32767%  即 -3.2768 ~ +3.2767
func (p Rate) Int() int         { return int(p) }
func (p Rate) Percent() float64 { return float64(p) / PercentFloat64 }
func (p Rate) Decimal() float64 { return float64(p) / DecimalAug }
func (p Rate) Mul(d int) int    { return d * int(p) }
func (p Rate) Fmt() string      { return strconv.FormatFloat(p.Percent(), 'f', -1, 32) }
func (p Rate) FmtAbs() string {
	c := p.Percent()
	if c < 0 {
		c = -c
	}
	return strconv.FormatFloat(c, 'f', -1, 32)
}

func NewRate16(n int16) Rate16 { return Rate16(n) }
func (p Rate16) Rate() Rate    { return Rate(p) }
func (p Rate16) Int16() int16  { return int16(p) }
func NewRate24(n int32) Rate24 { return Rate24(n) }
func (p Rate24) Rate() Rate    { return Rate(p) }
func (p Rate24) Int32() int32  { return int32(p) }
