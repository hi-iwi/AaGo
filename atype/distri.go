package atype

// type Html template.HTML   HTML 直接使用 template.HTML
type Province uint8 // 2 位省份地址码
type Dist uint16    // 4 位地址简码
type Distri Uint24  // 6 位地址简码
type AddrId uint64  // 12 位地址码

var (
	// 直辖市：北京、上海、天津、重庆
	Municipalities = []Distri{110000, 310000, 120000, 500000}
	// 经济特区：海南、深圳、厦门、珠海、汕头
	SEZs = []Distri{460000, 440300, 350200, 440400, 440500}
	// 自治区：新疆、西藏、宁夏、内蒙古、广西
	AutonomousRegions = []Distri{650000, 540000, 640000, 150000, 450000}
	// 特区：香港、澳门
	SARs = []Distri{810000, 820000}
)

func NewDistri(d Uint24) Distri {
	// 保持6位数字
	if d == 0 {
		return 0
	}
	if d < 100 {
		return Distri(d * 10000)
	}
	if d < 10000 {
		return Distri(d * 100)
	}
	return Distri(d)
}
func ToDistri(d uint32) Distri {
	return NewDistri(Uint24(d))
}
func (d Distri) Uint24() Uint24 {
	return Uint24(d)
}
func (d Distri) Uint32() uint32 {
	return uint32(d)
}
func (d Distri) Province() Province {
	return Province(d / 10000)
}
func (d Distri) Dist() Dist {
	return Dist(d / 100)
}
func (d Distri) AddrId() AddrId {
	return NewAddrId(uint64(d) * 1000000)
}

// 某个地区是否在另外一个地区内部
func (d Distri) Inside(p Distri) bool {
	if d == p {
		return true
	}
	b := d % 100
	if b != 0 && (d-b) == p {
		return true
	}
	b = d % 10000
	if b != 0 && (d-b) == p {
		return true
	}
	return false
}
func (d Distri) In(distris []Distri) bool {
	for _, distri := range distris {
		if d == distri {
			return true
		}
	}
	return false
}
func (d Distri) IsMunicipality() bool { return d.In(Municipalities) }
func (d Distri) IsSEZ() bool          { return d.In(SEZs) }
func (d Distri) IsAutonomous() bool   { return d.In(AutonomousRegions) }
func (d Distri) IsSAR() bool          { return d.In(SARs) }
func (d Distri) IsProvLevel() bool    { return d%10000 == 0 }
func (d Distri) IsProvince() bool     { return d.IsProvLevel() && !d.In(Municipalities) }

func (p Province) Distri() Distri {
	return Distri(uint32(p) * 10000)
}
func (p Dist) Distri() Distri {
	return Distri(uint32(p) * 100)
}

func NewAddrId(a uint64) AddrId {
	return AddrId(a)
}
func (a AddrId) Uint64() uint64 {
	return uint64(a)
}
