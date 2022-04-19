package com

import (
	"github.com/hi-iwi/AaGo/ae"
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
// 标准：Referer, User-Agent,
// 自定义：X-Csrf-Token, X-Request-Wid, X-From, X-Inviter
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
