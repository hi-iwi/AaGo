package util

import "math"

type Price struct {
	value float64
}

func NewPrice(v float64) *Price {
	return &Price{value: v}
}
func (p *Price) Val() float64 {
	return math.Round(p.value*100.0) / 100.0
}
func (p *Price) Add(b float64) *Price {
	v := (p.value + b) * 100.0
	p.value = math.Round(v) / 100.0
	return p
}

func (p *Price) Sub(b float64) *Price {
	v := (p.value - b) * 100.0
	p.value = math.Round(v) / 100.0
	return p
}

func (p *Price) Mul(b float64) *Price {
	v := (p.value * b) * 100.0
	p.value = math.Round(v) / 100.0
	return p
}

// Percent  参数是百分比，如 80.5 表示 80.5%
func (p *Price) XPercent(percent float64) *Price {
	v := (p.value * percent) * 100.0
	p.value = math.Round(v) / 10000.0
	return p
}

func (p *Price) Percent(base float64) float64 {
	x := math.Round(p.value * 10000.0 / float64(base))
	return x / 100.0
}

func (p *Price) Div(b float64) *Price {
	v := (p.value / b) * 100.0
	p.value = math.Round(v) / 100.0
	return p
}
