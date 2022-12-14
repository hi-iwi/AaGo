package com

import (
	"github.com/hi-iwi/AaGo/ae"
	"github.com/hi-iwi/AaGo/aenum"
	"github.com/hi-iwi/AaGo/atype"
	"html/template"
	"strconv"
	"strings"
	"time"
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
func (r *Req) QueryString(p string, params ...interface{}) (string, *ae.Error) {
	_x, e := r.Query(p, params...)
	return _x.String(), e
}

// 允许0  --> 直接用  _x.Int8() 就可以了
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

func (r *Req) QueryMoney(p string, required ...bool) (atype.Money, *ae.Error) {
	_x, e := r.QueryDigit(p, false, required...)
	return atype.Money(_x.DefaultInt(0)), e
}

func (r *Req) QueryUmoney(p string, required ...bool) (atype.Umoney, *ae.Error) {
	_x, e := r.QueryDigit(p, false, required...)
	return atype.Umoney(_x.DefaultUint(0)), e
}
func (r *Req) QueryAmount(p string, required ...bool) (atype.Amount, *ae.Error) {
	_x, e := r.QueryDigit(p, false, required...)
	return atype.Amount(_x.DefaultInt64(0)), e
}

func (r *Req) QueryUamount(p string, required ...bool) (atype.Uamount, *ae.Error) {
	_x, e := r.QueryDigit(p, false, required...)
	return atype.Uamount(_x.DefaultUint64(0)), e
}

func (r *Req) QueryDate(p string, loc *time.Location, required ...bool) (atype.Date, *ae.Error) {
	rq := true
	if len(required) == 1 {
		rq = required[0]
	}
	_x, e := r.Query(p, `^`+aenum.DateRegExp+`$`, rq)
	if e != nil {
		return "", ae.NewError(400, "invalid date ("+p+"): "+_x.String())
	}
	return atype.NewDate(_x.String(), loc), nil
}
func (r *Req) QueryDatetime(p string, loc *time.Location, required ...bool) (atype.Datetime, *ae.Error) {
	rq := true
	if len(required) == 1 {
		rq = required[0]
	}
	_x, e := r.Query(p, `^`+aenum.DatetimeRegExp+`$`, rq)
	if e != nil {
		return "", ae.NewError(400, "invalid datetime ("+p+"): "+_x.String())
	}
	return atype.NewDatetime(_x.String(), loc), nil
}
func (r *Req) BodyString(p string, required ...interface{}) (string, *ae.Error) {
	_x, e := r.Query(p, required...)
	return _x.String(), e
}
func (r *Req) BodyText(p string, required ...interface{}) (atype.Text, *ae.Error) {
	_x, e := r.Query(p, required...)
	return atype.Text(_x.String()), e
}
func (r *Req) BodyHtml(p string, required ...interface{}) (template.HTML, *ae.Error) {
	_x, e := r.Query(p, required...)
	return template.HTML(_x.String()), e
}
func (r *Req) BodyInt8(p string, required ...bool) (int8, *ae.Error) {
	_x, e := r.BodyDigit(p, false, required...)
	return _x.DefaultInt8(0), e
}
func (r *Req) BodyInt16(p string, required ...bool) (int16, *ae.Error) {
	_x, e := r.BodyDigit(p, false, required...)
	return _x.DefaultInt16(0), e
}
func (r *Req) BodyInt24(p string, required ...bool) (atype.Int24, *ae.Error) {
	_x, e := r.BodyDigit(p, false, required...)
	return atype.Int24(_x.DefaultInt32(0)), e
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
func (r *Req) BodyUint24(p string, required ...bool) (atype.Uint24, *ae.Error) {
	_x, e := r.BodyDigit(p, true, required...)
	return atype.Uint24(_x.DefaultUint32(0)), e
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

func (r *Req) BodyMoney(p string, required ...bool) (atype.Money, *ae.Error) {
	_x, e := r.BodyDigit(p, false, required...)
	return atype.Money(_x.DefaultInt(0)), e
}

func (r *Req) BodyUmoney(p string, required ...bool) (atype.Umoney, *ae.Error) {
	_x, e := r.BodyDigit(p, false, required...)
	return atype.Umoney(_x.DefaultUint(0)), e
}
func (r *Req) BodyAmount(p string, required ...bool) (atype.Amount, *ae.Error) {
	_x, e := r.BodyDigit(p, false, required...)
	return atype.Amount(_x.DefaultInt64(0)), e
}
func (r *Req) BodyUamount(p string, required ...bool) (atype.Uamount, *ae.Error) {
	_x, e := r.BodyDigit(p, false, required...)
	return atype.Uamount(_x.DefaultUint64(0)), e
}
func (r *Req) BodyDate(p string, loc *time.Location, required ...bool) (atype.Date, *ae.Error) {
	rq := true
	if len(required) == 1 {
		rq = required[0]
	}
	_x, e := r.Body(p, `^`+aenum.DateRegExp+`$`, rq)
	if e != nil {
		return "", ae.NewError(400, "invalid date ("+p+"): "+_x.String())
	}
	return atype.NewDate(_x.String(), loc), nil
}
func (r *Req) BodyDatetime(p string, loc *time.Location, required ...bool) (atype.Datetime, *ae.Error) {
	rq := true
	if len(required) == 1 {
		rq = required[0]
	}
	_x, e := r.Body(p, `^`+aenum.DatetimeRegExp+`$`, rq)
	if e != nil {
		return "", ae.NewError(400, "invalid datetime ("+p+"): "+_x.String())
	}
	return atype.NewDatetime(_x.String(), loc), nil
}

// {id:uint64}  or {sid:string}
func (r *Req) QueryId(p string, params ...interface{}) (sid string, id uint64, e *ae.Error) {
	var x *ReqProp
	if x, e = r.Query(p, params...); e != nil {
		return
	}
	sid = x.String()
	if sid == "" || sid == "0" {
		return
	}
	for _, s := range sid {
		if s < '0' || s > '9' {
			return
		}
	}
	id, _ = strconv.ParseUint(sid, 10, 64)
	return
}

func (r *Req) QueryPaging(limitMax uint, firstPages ...uint) atype.Paging {
	page, _ := r.QueryUint(ParamPage, false)
	offset, _ := r.QueryUint(ParamOffset, false)
	limit, _ := r.QueryUint(ParamLimit, false)

	if limit < 1 || limit > limitMax {
		limit = limitMax
	}

	if offset > 0 {
		page = (offset / limit) + 1
	} else {
		if page < 1 {
			page = 1
			if len(firstPages) > 0 {
				page = firstPages[0]
				limit = page * limitMax
			}
		} else {
			// change ?limit=3&offset=10 to ?limit=0&offset=10
			offset = (page - 1) * limit
		}
	}

	return atype.Paging{
		Page:   page,
		Offset: offset,
		Limit:  limit,
	}
}
