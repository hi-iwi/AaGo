package atype

import (
	"strconv"
	"strings"
)

type SepPercents string
type SepMoneys string

func ToSepPercents(elems []Decimal) SepPercents {
	// strings.Concat 类同
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return SepPercents(strconv.FormatInt(int64(elems[0]), 10))
	}
	deli := ","
	n := (len(elems) - 1) + (len(elems) * MaxInt16Len)
	var b strings.Builder
	b.Grow(n)
	b.WriteString(strconv.FormatInt(int64(elems[0]), 10))
	for _, s := range elems[1:] {
		b.WriteString(deli)
		b.WriteString(strconv.FormatInt(int64(s), 10))
	}

	return SepPercents(b.String())
}

func (t SepPercents) Percents() []Decimal {
	if t == "" {
		return nil
	}
	arr := strings.Split(string(t), ",")
	v := make([]Decimal, len(arr))
	for i, a := range arr {
		p, err := strconv.ParseInt(a, 10, 32)
		if err == nil {
			v[i] = Decimal(p)
		}
	}
	return v
}

func ToSepMoneys(elems []Money) SepMoneys {
	// strings.Concat 类同
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return SepMoneys(strconv.FormatInt(int64(elems[0]), 10))
	}
	deli := ","
	n := (len(elems) - 1) + (len(elems) * MaxInt64Len)
	var b strings.Builder
	b.Grow(n)
	b.WriteString(strconv.FormatInt(int64(elems[0]), 10))
	for _, s := range elems[1:] {
		b.WriteString(deli)
		b.WriteString(strconv.FormatInt(int64(s), 10))
	}

	return SepMoneys(b.String())
}
func (t SepMoneys) Moneys() []Money {
	if t == "" {
		return nil
	}
	arr := strings.Split(string(t), ",")
	v := make([]Money, len(arr))
	for i, a := range arr {
		p, err := strconv.ParseInt(a, 10, 64)
		if err == nil {
			v[i] = Money(p)
		}
	}
	return v
}
