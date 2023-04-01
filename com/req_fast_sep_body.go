package com

import (
	"github.com/hi-iwi/AaGo/ae"
	"github.com/hi-iwi/AaGo/atype"
	"strconv"
)

func (r *Req) BodySepStrings(p string, sep string, required, allowEmptyString bool) ([]string, *ae.Error) {
	return reqSepStrings(r.Body, p, sep, required, allowEmptyString)
}

func (r *Req) BodySepInt8s(p string, sep string, required, allowZero bool) ([]int8, *ae.Error) {
	arr, e := r.BodySepStrings(p, sep, required, false)
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
func (r *Req) BodySepUint8s(p string, sep string, required, allowZero bool) ([]uint8, *ae.Error) {
	arr, e := r.BodySepStrings(p, sep, required, false)
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
func (r *Req) BodySepUint16s(p string, sep string, required, allowZero bool) ([]uint16, *ae.Error) {
	arr, e := r.BodySepStrings(p, sep, required, false)
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
func (r *Req) BodySepUint24s(p string, sep string, required, allowZero bool) ([]atype.Uint24, *ae.Error) {
	arr, e := r.BodySepStrings(p, sep, required, false)
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
func (r *Req) BodySepUint32s(p string, sep string, required, allowZero bool) ([]uint32, *ae.Error) {
	arr, e := r.BodySepStrings(p, sep, required, false)
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

func (r *Req) BodySepUints(p string, sep string, required, allowZero bool) ([]uint, *ae.Error) {
	arr, e := r.BodySepStrings(p, sep, required, false)
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

// 逗号隔开的 uint64
func (r *Req) BodySepUint64s(p string, sep string, required, allowZero bool) ([]uint64, *ae.Error) {
	arr, e := r.BodySepStrings(p, sep, required, false)
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
func (r *Req) BodySepFloat32s(p string, sep string, required, allowZero bool) ([]float32, *ae.Error) {
	arr, e := r.BodySepStrings(p, sep, required, false)
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
func (r *Req) BodySepFloat64s(p string, sep string, required, allowZero bool) ([]float64, *ae.Error) {
	arr, e := r.BodySepStrings(p, sep, required, false)
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
