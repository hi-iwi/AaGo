package com

import (
	"github.com/hi-iwi/AaGo/ae"
	"github.com/hi-iwi/AaGo/atype"
	"regexp"
	"strconv"
	"strings"
)

func reqSepStrings(method func(string, ...interface{}) (*ReqProp, *ae.Error), p string, sep string, required ...bool) ([]string, *ae.Error) {
	require := len(required) == 0 || required[0] // 没有传，默认 require=true
	s, e := method(p, require)

	if e != nil {
		return nil, e
	}
	// 将换行符当作切割符号
	re := regexp.MustCompile(`\s*[\r\n]+\s*`)
	x := re.ReplaceAllString(s.String(), sep)
	arr := strings.Split(x, sep)
	b := make([]string, 0)
	for _, a := range arr {
		a = strings.Trim(a, " ")
		if a != "" {
			b = append(b, a)
		}
	}
	if len(b) == 0 && require {
		return nil, ae.BadParam(p)
	}
	return b, nil
}

func (r *Req) QuerySepStrings(p string, sep string, required ...bool) ([]string, *ae.Error) {
	return reqSepStrings(r.Query, p, sep, required...)
}
func (r *Req) BodySepStrings(p string, sep string, required ...bool) ([]string, *ae.Error) {
	return reqSepStrings(r.Body, p, sep, required...)
}

// 逗号隔开的 uint
func (r *Req) QuerySepUint8s(p string, sep string, required ...bool) ([]uint8, *ae.Error) {
	arr, e := r.QuerySepStrings(p, sep, required...)
	if e != nil {
		return nil, e
	}
	ids := make([]uint8, len(arr))
	var id uint64
	var err error
	for i, a := range arr {
		id, err = strconv.ParseUint(a, 10, 8)
		if err != nil {
			return nil, ae.BadParam(p)
		}
		ids[i] = uint8(id)
	}
	return ids, nil
}
func (r *Req) QuerySepUint16s(p string, sep string, required ...bool) ([]uint16, *ae.Error) {
	arr, e := r.QuerySepStrings(p, sep, required...)
	if e != nil {
		return nil, e
	}
	ids := make([]uint16, len(arr))
	var id uint64
	var err error
	for i, a := range arr {
		id, err = strconv.ParseUint(a, 10, 16)
		if err != nil {
			return nil, ae.BadParam(p)
		}
		ids[i] = uint16(id)
	}
	return ids, nil
}
func (r *Req) QuerySepUint24s(p string, sep string, required ...bool) ([]atype.Uint24, *ae.Error) {
	arr, e := r.QuerySepStrings(p, sep, required...)
	if e != nil {
		return nil, e
	}
	ids := make([]atype.Uint24, len(arr))
	var id uint64
	var err error
	for i, a := range arr {
		id, err = strconv.ParseUint(a, 10, 24)
		if err != nil {
			return nil, ae.BadParam(p)
		}
		ids[i] = atype.Uint24(id)
	}
	return ids, nil
}
func (r *Req) QuerySepUints(p string, sep string, required ...bool) ([]uint, *ae.Error) {
	arr, e := r.QuerySepStrings(p, sep, required...)
	if e != nil {
		return nil, e
	}
	ids := make([]uint, len(arr))
	var id uint64
	var err error
	for i, a := range arr {
		id, err = strconv.ParseUint(a, 10, 32)
		if err != nil {
			return nil, ae.BadParam(p)
		}
		ids[i] = uint(id)
	}
	return ids, nil
}
func (r *Req) QuerySepUint64s(p string, sep string, required ...bool) ([]uint64, *ae.Error) {
	arr, e := r.QuerySepStrings(p, sep, required...)
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
func (r *Req) BodySepUint8s(p string, sep string, required ...bool) ([]uint8, *ae.Error) {
	arr, e := r.QuerySepStrings(p, sep, required...)
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
	arr, e := r.QuerySepStrings(p, sep, required...)
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
	arr, e := r.QuerySepStrings(p, sep, required...)
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

func (r *Req) BodySepUints(p string, sep string, required ...bool) ([]uint, *ae.Error) {
	arr, e := r.QuerySepStrings(p, sep, required...)
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
func (r *Req) BodyUint64s(p string, sep string, required ...bool) ([]uint64, *ae.Error) {
	arr, e := r.QuerySepStrings(p, sep, required...)
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
