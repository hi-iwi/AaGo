package atype

import (
	"strconv"
	"strings"
)

type SepPercents string
type SepAmounts string

func ToSepPercents(elems []Percent) SepPercents {
	// strings.Join 类同
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

func (t SepPercents) Percents() []Percent {
	if t == "" {
		return nil
	}
	arr := strings.Split(string(t), ",")
	v := make([]Percent, len(arr))
	for i, a := range arr {
		p, err := strconv.ParseInt(a, 10, 16)
		if err == nil {
			v[i] = NewPercent(int16(p))
		}
	}
	return v
}

func ToSepAmounts(elems []Amount) SepAmounts {
	// strings.Join 类同
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return SepAmounts(strconv.FormatInt(int64(elems[0]), 10))
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

	return SepAmounts(b.String())
}
func (t SepAmounts) Amounts() []Amount {
	if t == "" {
		return nil
	}
	arr := strings.Split(string(t), ",")
	v := make([]Amount, len(arr))
	for i, a := range arr {
		p, err := strconv.ParseInt(a, 10, 64)
		if err == nil {
			v[i] = NewAmount(p)
		}
	}
	return v
}
