package com

import (
	"github.com/hi-iwi/AaGo/ae"
	"github.com/hi-iwi/AaGo/atype"
	"regexp"
	"strconv"
	"strings"
)

func reqSepStrings(method func(string, ...interface{}) (*ReqProp, *ae.Error), p string, sep string, required, allowEmptyString bool) ([]string, *ae.Error) {
	s, e := method(p, required)

	if e != nil {
		return nil, e
	}
	// 将换行符当作切割符号
	re := regexp.MustCompile(`\s*[\r\n]+\s*`)
	x := re.ReplaceAllString(s.String(), sep)
	arr := strings.Split(x, sep)
	b := make([]string, 0)
	for _, a := range arr {
		if allowEmptyString {
			b = append(b, a)
		} else {
			a = strings.Trim(a, " ")
			if a != "" {
				b = append(b, a)
			}
		}

	}
	if len(b) == 0 && required {
		return nil, ae.BadParam(p)
	}
	return b, nil
}

func (r *Req) QuerySepStrings(p string, sep string, required, allowEmptyString bool) ([]string, *ae.Error) {
	return reqSepStrings(r.Query, p, sep, required, allowEmptyString)
}

func (r *Req) QuerySepInt8s(p string, sep string, required, allowZero bool) ([]int8, *ae.Error) {
	arr, e := r.QuerySepStrings(p, sep, required, false)
	if e != nil {
		return nil, e
	}
	v := make([]int8, 0, len(arr))
	var err error
	var n int64
	for _, a := range arr {
		n, err = strconv.ParseInt(a, 10, 8)
		if err != nil {
			return nil, ae.BadParam(p)
		}
		if allowZero || n > 0 {
			v = append(v, int8(n))
		}
	}
	return v, nil
}

// 逗号隔开的 uint
func (r *Req) QuerySepUint8s(p string, sep string, required, allowZero bool) ([]uint8, *ae.Error) {
	arr, e := r.QuerySepStrings(p, sep, required, false)
	if e != nil {
		return nil, e
	}
	v := make([]uint8, 0, len(arr))
	var err error
	var n uint64
	for _, a := range arr {
		n, err = strconv.ParseUint(a, 10, 8)
		if err != nil {
			return nil, ae.BadParam(p)
		}
		if allowZero || n > 0 {
			v = append(v, uint8(n))
		}
	}
	return v, nil
}
func (r *Req) QuerySepUint16s(p string, sep string, required, allowZero bool) ([]uint16, *ae.Error) {
	arr, e := r.QuerySepStrings(p, sep, required, false)
	if e != nil {
		return nil, e
	}
	v := make([]uint16, 0, len(arr))
	var err error
	var n uint64
	for _, a := range arr {
		n, err = strconv.ParseUint(a, 10, 16)
		if err != nil {
			return nil, ae.BadParam(p)
		}
		if allowZero || n > 0 {
			v = append(v, uint16(n))
		}
	}
	return v, nil
}
func (r *Req) QuerySepUint24s(p string, sep string, required, allowZero bool) ([]atype.Uint24, *ae.Error) {
	arr, e := r.QuerySepStrings(p, sep, required, false)
	if e != nil {
		return nil, e
	}
	v := make([]atype.Uint24, 0, len(arr))
	var err error
	var n uint64
	for _, a := range arr {
		n, err = strconv.ParseUint(a, 10, 24)
		if err != nil {
			return nil, ae.BadParam(p)
		}
		if allowZero || n > 0 {
			v = append(v, atype.Uint24(n))
		}
	}
	return v, nil
}
func (r *Req) QuerySepUint32s(p string, sep string, required, allowZero bool) ([]uint32, *ae.Error) {
	arr, e := r.QuerySepStrings(p, sep, required, false)
	if e != nil {
		return nil, e
	}
	v := make([]uint32, 0, len(arr))
	var err error
	var n uint64
	for _, a := range arr {
		n, err = strconv.ParseUint(a, 10, 32)
		if err != nil {
			return nil, ae.BadParam(p)
		}
		if allowZero || n > 0 {
			v = append(v, uint32(n))
		}
	}
	return v, nil
}
func (r *Req) QuerySepUints(p string, sep string, required, allowZero bool) ([]uint, *ae.Error) {
	arr, e := r.QuerySepStrings(p, sep, required, false)
	if e != nil {
		return nil, e
	}
	v := make([]uint, 0, len(arr))
	var err error
	var n uint64
	for _, a := range arr {
		n, err = strconv.ParseUint(a, 10, 32)
		if err != nil {
			return nil, ae.BadParam(p)
		}
		if allowZero || n > 0 {
			v = append(v, uint(n))
		}
	}
	return v, nil
}
func (r *Req) QuerySepUint64s(p string, sep string, required, allowZero bool) ([]uint64, *ae.Error) {
	arr, e := r.QuerySepStrings(p, sep, required, false)
	if e != nil {
		return nil, e
	}
	v := make([]uint64, 0, len(arr))
	var err error
	var n uint64
	for _, a := range arr {
		n, err = strconv.ParseUint(a, 10, 64)
		if err != nil {
			return nil, ae.BadParam(p)
		}
		if allowZero || n > 0 {
			v = append(v, n)
		}
	}
	return v, nil
}
func (r *Req) QuerySepFloat32s(p string, sep string, required, allowZero bool) ([]float32, *ae.Error) {
	arr, e := r.QuerySepStrings(p, sep, required, false)
	if e != nil {
		return nil, e
	}
	v := make([]float32, 0, len(arr))
	var err error
	var n float64
	for _, a := range arr {
		n, err = strconv.ParseFloat(a, 32)
		if err != nil {
			return nil, ae.BadParam(p)
		}
		if allowZero || n > 0 {
			v = append(v, float32(n))
		}
	}
	return v, nil
}
func (r *Req) QuerySepFloat64s(p string, sep string, required, allowZero bool) ([]float64, *ae.Error) {
	arr, e := r.QuerySepStrings(p, sep, required, false)
	if e != nil {
		return nil, e
	}
	v := make([]float64, 0, len(arr))
	var err error
	var n float64
	for _, a := range arr {
		n, err = strconv.ParseFloat(a, 64)
		if err != nil {
			return nil, ae.BadParam(p)
		}
		if allowZero || n > 0 {
			v = append(v, n)
		}
	}
	return v, nil
}
