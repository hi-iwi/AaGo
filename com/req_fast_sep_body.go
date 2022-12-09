package com

import (
	"github.com/hi-iwi/AaGo/ae"
	"github.com/hi-iwi/AaGo/atype"
	"strconv"
)

func (r *Req) BodySepStrings(p string, sep string, required ...bool) ([]string, *ae.Error) {
	return reqSepStrings(r.Body, p, sep, required...)
}

func (r *Req) BodySepInt8s(p string, sep string, required ...bool) ([]int8, *ae.Error) {
	arr, e := r.BodySepStrings(p, sep, required...)
	if e != nil {
		return nil, e
	}
	ids := make([]int8, len(arr))
	var err error
	var id int64
	for i, a := range arr {
		id, err = strconv.ParseInt(a, 10, 8)
		if err != nil {
			return nil, ae.BadParam(p)
		}
		ids[i] = int8(id)
	}
	return ids, nil
}
func (r *Req) BodySepUint8s(p string, sep string, required ...bool) ([]uint8, *ae.Error) {
	arr, e := r.BodySepStrings(p, sep, required...)
	if e != nil {
		return nil, e
	}
	ids := make([]uint8, len(arr))
	var err error
	var id uint64
	for i, a := range arr {
		id, err = strconv.ParseUint(a, 10, 8)
		if err != nil {
			return nil, ae.BadParam(p)
		}
		ids[i] = uint8(id)
	}
	return ids, nil
}
func (r *Req) BodySepUint16s(p string, sep string, required ...bool) ([]uint16, *ae.Error) {
	arr, e := r.BodySepStrings(p, sep, required...)
	if e != nil {
		return nil, e
	}
	ids := make([]uint16, len(arr))
	var err error
	var id uint64
	for i, a := range arr {
		id, err = strconv.ParseUint(a, 10, 16)
		if err != nil {
			return nil, ae.BadParam(p)
		}
		ids[i] = uint16(id)
	}
	return ids, nil
}
func (r *Req) BodySepUint24s(p string, sep string, required ...bool) ([]atype.Uint24, *ae.Error) {
	arr, e := r.BodySepStrings(p, sep, required...)
	if e != nil {
		return nil, e
	}
	ids := make([]atype.Uint24, len(arr))
	var err error
	var id uint64
	for i, a := range arr {
		id, err = strconv.ParseUint(a, 10, 24)
		if err != nil {
			return nil, ae.BadParam(p)
		}
		ids[i] = atype.Uint24(id)
	}
	return ids, nil
}
func (r *Req) BodySepUint32s(p string, sep string, required ...bool) ([]uint32, *ae.Error) {
	arr, e := r.BodySepStrings(p, sep, required...)
	if e != nil {
		return nil, e
	}
	ids := make([]uint32, len(arr))
	var err error
	var id uint64
	for i, a := range arr {
		id, err = strconv.ParseUint(a, 10, 32)
		if err != nil {
			return nil, ae.BadParam(p)
		}
		ids[i] = uint32(id)
	}
	return ids, nil
}

func (r *Req) BodySepUints(p string, sep string, required ...bool) ([]uint, *ae.Error) {
	arr, e := r.BodySepStrings(p, sep, required...)
	if e != nil {
		return nil, e
	}
	ids := make([]uint, len(arr))
	var err error
	var id uint64
	for i, a := range arr {
		id, err = strconv.ParseUint(a, 10, 32)
		if err != nil {
			return nil, ae.BadParam(p)
		}
		ids[i] = uint(id)
	}
	return ids, nil
}

// 逗号隔开的 uint64
func (r *Req) BodySepUint64s(p string, sep string, required ...bool) ([]uint64, *ae.Error) {
	arr, e := r.BodySepStrings(p, sep, required...)
	if e != nil {
		return nil, e
	}
	ids := make([]uint64, len(arr))
	var err error
	for i, a := range arr {
		ids[i], err = strconv.ParseUint(a, 10, 64)
		if err != nil {
			return nil, ae.BadParam(p)
		}
	}
	return ids, nil
}
func (r *Req) BodySepFloat32s(p string, sep string, required ...bool) ([]float32, *ae.Error) {
	arr, e := r.BodySepStrings(p, sep, required...)
	if e != nil {
		return nil, e
	}
	ids := make([]float32, len(arr))
	for i, a := range arr {
		f, err := strconv.ParseFloat(a, 32)
		if err != nil {
			return nil, ae.BadParam(p)
		}
		ids[i] = float32(f)
	}
	return ids, nil
}
func (r *Req) BodySepFloat64s(p string, sep string, required ...bool) ([]float64, *ae.Error) {
	arr, e := r.BodySepStrings(p, sep, required...)
	if e != nil {
		return nil, e
	}
	ids := make([]float64, len(arr))
	var err error
	for i, a := range arr {
		ids[i], err = strconv.ParseFloat(a, 64)
		if err != nil {
			return nil, ae.BadParam(p)
		}
	}
	return ids, nil
}
