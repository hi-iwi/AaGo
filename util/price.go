package util

import (
	"fmt"
	"math"
)

type Price struct {
	intType uint8 // 0 ceil;  1 round  2 floor
	value   float64
}

// 0.14  会被存储为 14.000000000000002   ，所以对于这种精度问题，一定要先去掉这种，再做round
func bigCeil(f float64) float64 {
	g := math.Floor(f)
	if f-g < 0.001 {
		return g
	}
	return math.Ceil(f)
}
func bigFloor(f float64) float64 {
	g := math.Ceil(f)
	if g-f < 0.001 {
		return g
	}
	return math.Floor(f)
}

func NewPrice(v float64, intType ...uint8) *Price {
	it := uint8(0)
	if len(intType) == 1 {
		it = intType[0]
	}

	return &Price{
		intType: it,
		value:   v,
	}
}

// 数据增广
func (p *Price) aug(v float64, multiple float64) float64 {
	switch p.intType {
	case 0:
		return bigCeil(v * multiple)
	case 1:
		return math.Round(v * multiple)
	case 2:
		return bigFloor(v * multiple)
	}
	return 0
}

func (p *Price) Fixed() string {
	return fmt.Sprintf("%.2f", p.Val())
}
func (p *Price) Val() float64 {
	return p.aug(p.value, 100.0) / 100.0
}
func (p *Price) Add(b float64) *Price {
	v := (p.value * 100.0) + (b * 100.0)
	p.value = p.aug(v, 1.0) / 100.0
	return p
}

func (p *Price) Sub(b float64) *Price {
	v := (p.value * 100.0) - (b * 100.0)
	p.value = p.aug(v, 1.0) / 100.0
	return p
}

func (p *Price) Mul(b float64) *Price {
	v := (p.value * 100.0) * (b * 100.0)
	p.value = p.aug(v, 1.0) / 10000.0
	return p
}

func (p *Price) Div(b float64) *Price {
	v := ((p.value * 100.0) / (b * 100.0)) * 100.0
	p.value = p.aug(v, 1.0) / 100.0
	return p
}

// Percent  参数是百分比，如 80.5 表示 80.5%
func (p *Price) XPercent(percent float64) *Price {
	v := p.value * percent
	p.value = p.aug(v, 1.0) / 100.0
	return p
}

func (p *Price) Percent(base float64) float64 {
	x := p.aug(p.value*10000.0/float64(base), 1.0)
	return x / 100.0
}
