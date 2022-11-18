package com

import (
	"github.com/hi-iwi/AaGo/ae"
	"github.com/hi-iwi/AaGo/atype"
	"strconv"
	"strings"
)

func (r *Req) Xhost() string {
	scheme := r.r.URL.Scheme
	if scheme == "" {
		if r.r.TLS != nil {
			scheme = "https:"
		} else {
			scheme = "http:"
		}
	}
	// 由于是通过接口做跳转，所以不可行，只会对这个接口结果跳转， 不会对页面跳转！！还需要客户端自行处理
	host := r.r.Host
	h := scheme + "//" + host
	return h
}

// 跟踪客户端数据，优先级：url --> header `X-***`  --> cookie
// 标准：Referer, VUser-Agent,
// 自定义：X-Csrf-Token, X-Request-Vuid, X-From, X-Inviter
// @warn 尽量不要通过自定义header传参，因为可能某个web server会基于安全禁止某些无法识别的header
func (r *Req) XHeader(name string, patterns ...interface{}) (v *ReqProp, e *ae.Error) {
	key := strings.Title(strings.ReplaceAll(name, "_", "-")) // 首字母大写
	if v, e = r.Header(key, patterns...); e == nil && v.NotEmpty() {
		return
	}
	key = "X-" + key
	if v, e = r.Header(key, patterns...); e == nil && v.NotEmpty() {
		return
	}
	v = NewReqProp(name, "")
	return v, v.Filter(patterns...)
}
func (r *Req) FastXheader(name string, patterns ...interface{}) *ReqProp {
	v, _ := r.XHeader(name, patterns...)
	return v
}
func (r *Req) Qparam(name string, patterns ...interface{}) (v *ReqProp, e *ae.Error) {
	key := strings.ToLower(strings.ReplaceAll(name, "-", "_"))
	if v, e = r.Query(name, patterns...); e == nil && v.NotEmpty() {
		return
	}
	if v, e = r.Query(key, patterns...); e == nil && v.NotEmpty() {
		return
	}
	if v, e = r.XHeader(name, patterns...); e == nil && v.NotEmpty() {
		return
	}
	v = NewReqProp(name, "")
	return v, v.Filter(patterns...)
}

// in url params -> in header? -> in cookie?
// e.g.  csrf_token: in url params? -> Csrf-Token: in header?  X-Csrf-Token: in header-> csrf_token: in cookie
func (r *Req) Xparam(name string, patterns ...interface{}) (v *ReqProp, e *ae.Error) {
	if v, e = r.Qparam(name, patterns...); e == nil && v.NotEmpty() {
		return
	}

	key := strings.ToLower(strings.ReplaceAll(name, "-", "_"))
	if cookie, err := r.Cookie(name); err == nil {
		v = NewReqProp(cookie.Name, cookie.Value)
		return v, v.Filter(patterns...)
	}
	if cookie, err := r.Cookie(key); err == nil {
		v = NewReqProp(cookie.Name, cookie.Value)
		return v, v.Filter(patterns...)
	}
	v = NewReqProp(name, "")
	return v, v.Filter(patterns...)
}
func (r *Req) FastXparam(name string) *ReqProp {
	v, _ := r.Xparam(name, false)
	return v
}
func reqDigit(method func(string, ...interface{}) (*ReqProp, *ae.Error), p string, positive bool, xargs ...bool) (*ReqProp, *ae.Error) {
	required := len(xargs) == 0 || xargs[0]
	reg := `^[-\d]\d*$`
	if positive {
		reg = `^\d+$`
	}
	return method(p, reg, required)
}
func (r *Req) QueryDigit(p string, positive bool, xargs ...bool) (*ReqProp, *ae.Error) {
	return reqDigit(r.Query, p, positive, xargs...)
}
func (r *Req) BodyDigit(p string, positive bool, xargs ...bool) (*ReqProp, *ae.Error) {
	return reqDigit(r.Body, p, positive, xargs...)
}

// 允许0
func (r *Req) QueryInt8(p string, required ...bool) (int8, *ae.Error) {
	_x, e := r.QueryDigit(p, false, required...)
	return _x.DefaultInt8(0), e
}
func (r *Req) QueryInt16(p string, required ...bool) (int16, *ae.Error) {
	_x, e := r.QueryDigit(p, false, required...)
	return _x.DefaultInt16(0), e
}
func (r *Req) QueryInt32(p string, required ...bool) (int32, *ae.Error) {
	_x, e := r.QueryDigit(p, false, required...)
	return _x.DefaultInt32(0), e
}
func (r *Req) QueryInt(p string, required ...bool) (int, *ae.Error) {
	_x, e := r.QueryDigit(p, false, required...)
	return _x.DefaultInt(0), e
}
func (r *Req) QueryInt64(p string, required ...bool) (int64, *ae.Error) {
	_x, e := r.QueryDigit(p, false, required...)
	return _x.DefaultInt64(0), e
}
func (r *Req) QueryUint8(p string, required ...bool) (uint8, *ae.Error) {
	_x, e := r.QueryDigit(p, true, required...)
	return _x.DefaultUint8(0), e
}
func (r *Req) QueryUint16(p string, required ...bool) (uint16, *ae.Error) {
	_x, e := r.QueryDigit(p, true, required...)
	return _x.DefaultUint16(0), e
}
func (r *Req) QueryUint32(p string, required ...bool) (uint32, *ae.Error) {
	_x, e := r.QueryDigit(p, true, required...)
	return _x.DefaultUint32(0), e
}
func (r *Req) QueryUint(p string, required ...bool) (uint, *ae.Error) {
	_x, e := r.QueryDigit(p, true, required...)
	return _x.DefaultUint(0), e
}
func (r *Req) QueryUint64(p string, required ...bool) (uint64, *ae.Error) {
	_x, e := r.QueryDigit(p, true, required...)
	return _x.DefaultUint64(0), e
}

func (r *Req) BodyInt8(p string, required ...bool) (int8, *ae.Error) {
	_x, e := r.BodyDigit(p, false, required...)
	return _x.DefaultInt8(0), e
}
func (r *Req) BodyInt16(p string, required ...bool) (int16, *ae.Error) {
	_x, e := r.BodyDigit(p, false, required...)
	return _x.DefaultInt16(0), e
}
func (r *Req) BodyInt32(p string, required ...bool) (int32, *ae.Error) {
	_x, e := r.BodyDigit(p, false, required...)
	return _x.DefaultInt32(0), e
}
func (r *Req) BodyInt(p string, required ...bool) (int, *ae.Error) {
	_x, e := r.BodyDigit(p, false, required...)
	return _x.DefaultInt(0), e
}
func (r *Req) BodyInt64(p string, required ...bool) (int64, *ae.Error) {
	_x, e := r.BodyDigit(p, false, required...)
	return _x.DefaultInt64(0), e
}
func (r *Req) BodyUint8(p string, required ...bool) (uint8, *ae.Error) {
	_x, e := r.BodyDigit(p, true, required...)
	return _x.DefaultUint8(0), e
}
func (r *Req) BodyUint16(p string, required ...bool) (uint16, *ae.Error) {
	_x, e := r.BodyDigit(p, true, required...)
	return _x.DefaultUint16(0), e
}
func (r *Req) BodyUint32(p string, required ...bool) (uint32, *ae.Error) {
	_x, e := r.BodyDigit(p, true, required...)
	return _x.DefaultUint32(0), e
}
func (r *Req) BodyUint(p string, required ...bool) (uint, *ae.Error) {
	_x, e := r.BodyDigit(p, true, required...)
	return _x.DefaultUint(0), e
}
func (r *Req) BodyUint64(p string, required ...bool) (uint64, *ae.Error) {
	_x, e := r.BodyDigit(p, true, required...)
	return _x.DefaultUint64(0), e
}
func reqStrings(method func(string, ...interface{}) (*ReqProp, *ae.Error), p string, re string, required ...bool) ([]string, *ae.Error) {
	rq := len(required) == 0 || required[0]
	var s *ReqProp
	var e *ae.Error
	if re == "" {
		s, e = method(p, rq)
	} else {
		s, e = method(p, re, rq)
	}

	if e != nil {
		return nil, e
	}
	arr := strings.Split(s.String(), ",")
	b := make([]string, 0)
	for _, a := range arr {
		if a != "" {
			b = append(b, a)
		}
	}
	if len(b) == 0 && rq {
		return nil, ae.BadParam(p)
	}
	return b, nil
}
func (r *Req) BodyStrings(p string, required ...bool) ([]string, *ae.Error) {
	return reqStrings(r.Body, p, "", required...)
}

func (r *Req) QueryStrings(p string, required ...bool) ([]string, *ae.Error) {
	return reqStrings(r.Query, p, "", required...)
}

// 逗号隔开的 string digits
func reqDigits(method func(string, ...interface{}) (*ReqProp, *ae.Error), p string, required ...bool) ([]string, *ae.Error) {
	return reqStrings(method, p, `^[\d,]$`, required...)
}

func (r *Req) BodyDigits(p string, required ...bool) ([]string, *ae.Error) {
	return reqDigits(r.Body, p, required...)
}
func (r *Req) QueryDigits(p string, required ...bool) ([]string, *ae.Error) {
	return reqDigits(r.Query, p, required...)
}

// 逗号隔开的 uint
func (r *Req) QueryUints(p string, required ...bool) ([]uint, *ae.Error) {
	arr, e := r.QueryDigits(p, required...)
	if e != nil {
		return nil, e
	}
	ids := make([]uint, len(arr))
	for i, a := range arr {
		id, _ := strconv.ParseUint(a, 10, 32)
		ids[i] = uint(id)
	}
	return ids, nil
}
func (r *Req) QueryUint64s(p string, required ...bool) ([]uint64, *ae.Error) {
	arr, e := r.QueryDigits(p, required...)
	if e != nil {
		return nil, e
	}
	ids := make([]uint64, len(arr))
	for i, a := range arr {
		ids[i], _ = strconv.ParseUint(a, 10, 32)
	}
	return ids, nil
}
func (r *Req) BodyUints(p string, required ...bool) ([]uint, *ae.Error) {
	arr, e := r.BodyDigits(p, required...)
	if e != nil {
		return nil, e
	}
	ids := make([]uint, len(arr))
	for i, a := range arr {
		id, _ := strconv.ParseUint(a, 10, 32)
		ids[i] = uint(id)
	}
	return ids, nil
}

// 逗号隔开的 uint64
func (r *Req) BodyUint64s(p string, required ...bool) ([]uint64, *ae.Error) {
	arr, e := r.BodyDigits(p, required...)
	if e != nil {
		return nil, e
	}
	ids := make([]uint64, len(arr))
	for i, a := range arr {
		ids[i], _ = strconv.ParseUint(a, 10, 32)
	}
	return ids, nil
}

// ID:uint64, required 情况ID必须>0；optional 情况，可以为0
func (r *Req) QueryId(p string, required ...bool) (uint64, *ae.Error) {
	id, e := r.QueryUint64(p, required...)
	if len(required) == 1 && !required[0] {
		return id, e // optional
	}
	if id == 0 {
		return 0, ae.NewError(400, "bad "+p)
	}
	return id, nil
}

// ID:uint64, required 情况ID必须>0；optional 情况，可以为0
func (r *Req) BodyId(p string, required ...bool) (uint64, *ae.Error) {
	id, e := r.BodyUint64(p, required...)
	if len(required) == 1 && !required[0] {
		return id, e // optional
	}
	if id == 0 {
		return 0, ae.NewError(400, "bad "+p)
	}
	return id, nil
}

func (r *Req) QueryPaging(args ...int) atype.Paging {
	page, _ := r.QueryInt(ParamPage, false)
	offset, _ := r.QueryInt(ParamOffset, false)
	limit, _ := r.QueryInt(ParamLimit, false)

	if limit < 1 {
		if len(args) > 0 {
			limit = args[0]
		} else {
			limit = 20
		}
	} else if limit > 100 {
		limit = 100
	}

	if offset > 0 {
		page = (offset / limit) + 1
	} else {
		if page < 1 {
			page = 1
		}
	}
	// change ?limit=3&offset=10 to ?limit=0&offset=10
	offset = (page - 1) * limit

	return atype.Paging{
		Page:   page,
		Offset: offset,
		Limit:  limit,
	}
}
