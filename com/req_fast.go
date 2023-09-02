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
	reg := `^[-\d]\d*$`
	if positive {
		reg = `^\d+$`
	}
	return method(p, reg, len(xargs) == 0 || xargs[0])
}
func (r *Req) QueryDigit(p string, positive bool, xargs ...bool) (*ReqProp, *ae.Error) {
	return reqDigit(r.Query, p, positive, xargs...)
}
func (r *Req) BodyDigit(p string, positive bool, xargs ...bool) (*ReqProp, *ae.Error) {
	return reqDigit(r.Body, p, positive, xargs...)
}
func (r *Req) QueryString(p string, params ...interface{}) (string, *ae.Error) {
	x, e := r.Query(p, params...)
	return x.String(), e
}
func (r *Req) QueryBool(p string, required ...interface{}) (bool, *ae.Error) {
	x, e := r.Query(p, required...)
	if e != nil {
		return false, e
	}
	return x.DefaultBool(false), nil
}
func (r *Req) QueryBooln(p string, required ...interface{}) (atype.Booln, *ae.Error) {
	b, e := r.QueryBool(p, required...)
	if e != nil {
		return 0, e
	}
	return atype.ToBooln(b), nil
}

// 允许0  --> 直接用  x.Int8() 就可以了
func (r *Req) QueryInt8(p string, required ...bool) (int8, *ae.Error) {
	x, e := r.QueryDigit(p, false, required...)
	return x.DefaultInt8(0), e
}
func (r *Req) QueryInt16(p string, required ...bool) (int16, *ae.Error) {
	x, e := r.QueryDigit(p, false, required...)
	return x.DefaultInt16(0), e
}
func (r *Req) QueryInt32(p string, required ...bool) (int32, *ae.Error) {
	x, e := r.QueryDigit(p, false, required...)
	return x.DefaultInt32(0), e
}
func (r *Req) QueryInt(p string, required ...bool) (int, *ae.Error) {
	x, e := r.QueryDigit(p, false, required...)
	return x.DefaultInt(0), e
}
func (r *Req) QueryInt64(p string, required ...bool) (int64, *ae.Error) {
	x, e := r.QueryDigit(p, false, required...)
	return x.DefaultInt64(0), e
}
func (r *Req) QueryUint8(p string, required ...bool) (uint8, *ae.Error) {
	x, e := r.QueryDigit(p, true, required...)
	return x.DefaultUint8(0), e
}
func (r *Req) QueryUint16(p string, required ...bool) (uint16, *ae.Error) {
	x, e := r.QueryDigit(p, true, required...)
	return x.DefaultUint16(0), e
}
func (r *Req) QueryUint24(p string, required ...bool) (atype.Uint24, *ae.Error) {
	x, e := r.QueryDigit(p, true, required...)
	return x.DefaultUint24(0), e
}
func (r *Req) QueryUint32(p string, required ...bool) (uint32, *ae.Error) {
	x, e := r.QueryDigit(p, true, required...)
	return x.DefaultUint32(0), e
}
func (r *Req) QueryUint(p string, required ...bool) (uint, *ae.Error) {
	x, e := r.QueryDigit(p, true, required...)
	return x.DefaultUint(0), e
}
func (r *Req) QueryUint64(p string, required ...bool) (uint64, *ae.Error) {
	x, e := r.QueryDigit(p, true, required...)
	return x.DefaultUint64(0), e
}
func (r *Req) QueryProvince(p string, required ...bool) (atype.Province, *ae.Error) {
	distri, e := r.QueryDistri(p, required...)
	if e != nil {
		return 0, e
	}
	return distri.Province(), nil
}
func (r *Req) QueryDist(p string, required ...bool) (atype.Dist, *ae.Error) {
	distri, e := r.QueryDistri(p, required...)
	if e != nil {
		return 0, e
	}
	return distri.Dist(), nil
}
func (r *Req) QueryDistri(p string, required ...bool) (atype.Distri, *ae.Error) {
	x, e := r.QueryDigit(p, false, required...)
	return atype.NewDistri(x.DefaultUint24(0)), e
}

func (r *Req) QuerySmallMoney(p string, required ...bool) (atype.SmallMoney, *ae.Error) {
	x, e := r.QueryDigit(p, false, required...)
	return atype.SmallMoney(x.DefaultUint(0)), e
}
func (r *Req) QueryMoney(p string, required ...bool) (atype.Money, *ae.Error) {
	x, e := r.QueryDigit(p, false, required...)
	return atype.Money(x.DefaultInt64(0)), e
}
func (r *Req) QueryPercent16(p string, required ...bool) (atype.Percent16, *ae.Error) {
	x, e := r.QueryDigit(p, false, required...)
	return atype.NewPercent16(x.DefaultInt16(0)), e
}
func (r *Req) QueryPercent24(p string, required ...bool) (atype.Percent24, *ae.Error) {
	x, e := r.QueryDigit(p, false, required...)
	return atype.NewPercent24(x.DefaultInt32(0)), e
}

func (r *Req) QueryPercent(p string, required ...bool) (atype.Percent, *ae.Error) {
	x, e := r.QueryDigit(p, false, required...)
	return atype.NewPercent(x.DefaultInt(0)), e
}

func (r *Req) QueryDate(p string, loc *time.Location, required ...bool) (atype.Date, *ae.Error) {
	x, e := r.Query(p, `^`+aenum.DateRegExp+`$`, len(required) == 0 || required[0])
	if e != nil {
		return "", ae.NewError(400, "invalid date ("+p+"): "+x.String())
	}
	return atype.NewDate(x.String(), loc), nil
}
func (r *Req) QueryDatetime(p string, loc *time.Location, required ...bool) (atype.Datetime, *ae.Error) {
	x, e := r.Query(p, `^`+aenum.DatetimeRegExp+`$`, len(required) == 0 || required[0])
	if e != nil {
		return "", ae.NewError(400, "invalid datetime ("+p+"): "+x.String())
	}
	return atype.NewDatetime(x.String(), loc), nil
}

func (r *Req) BodyString(p string, required ...interface{}) (string, *ae.Error) {
	x, e := r.Body(p, required...)
	return x.String(), e
}

//func (r *Req) BodyText(p string, required ...interface{}) (atype.Text, *ae.Error) {
//	x, e := r.Body(p, required...)
//	return atype.Text(x.String()), e
//}
func (r *Req) BodyHtml(p string, required ...interface{}) (template.HTML, *ae.Error) {
	x, e := r.Body(p, required...)
	return template.HTML(x.String()), e
}
func (r *Req) BodyBool(p string, required ...interface{}) (bool, *ae.Error) {
	x, e := r.Body(p, required...)
	if e != nil {
		return false, e
	}
	return x.DefaultBool(false), e
}
func (r *Req) BodyBooln(p string, required ...interface{}) (atype.Booln, *ae.Error) {
	b, e := r.BodyBool(p, required...)
	if e != nil {
		return 0, e
	}
	return atype.ToBooln(b), nil
}
func (r *Req) BodyInt8(p string, required ...bool) (int8, *ae.Error) {
	x, e := r.BodyDigit(p, false, required...)
	return x.DefaultInt8(0), e
}
func (r *Req) BodyInt16(p string, required ...bool) (int16, *ae.Error) {
	x, e := r.BodyDigit(p, false, required...)
	return x.DefaultInt16(0), e
}
func (r *Req) BodyInt24(p string, required ...bool) (atype.Int24, *ae.Error) {
	x, e := r.BodyDigit(p, false, required...)
	return atype.Int24(x.DefaultInt32(0)), e
}
func (r *Req) BodyInt32(p string, required ...bool) (int32, *ae.Error) {
	x, e := r.BodyDigit(p, false, required...)
	return x.DefaultInt32(0), e
}
func (r *Req) BodyInt(p string, required ...bool) (int, *ae.Error) {
	x, e := r.BodyDigit(p, false, required...)
	return x.DefaultInt(0), e
}
func (r *Req) BodyInt64(p string, required ...bool) (int64, *ae.Error) {
	x, e := r.BodyDigit(p, false, required...)
	return x.DefaultInt64(0), e
}
func (r *Req) BodyUint8(p string, required ...bool) (uint8, *ae.Error) {
	x, e := r.BodyDigit(p, true, required...)
	return x.DefaultUint8(0), e
}
func (r *Req) BodyUint16(p string, required ...bool) (uint16, *ae.Error) {
	x, e := r.BodyDigit(p, true, required...)
	return x.DefaultUint16(0), e
}
func (r *Req) BodyUint24(p string, required ...bool) (atype.Uint24, *ae.Error) {
	x, e := r.BodyDigit(p, true, required...)
	return atype.Uint24(x.DefaultUint32(0)), e
}
func (r *Req) BodyUint32(p string, required ...bool) (uint32, *ae.Error) {
	x, e := r.BodyDigit(p, true, required...)
	return x.DefaultUint32(0), e
}
func (r *Req) BodyUint(p string, required ...bool) (uint, *ae.Error) {
	x, e := r.BodyDigit(p, true, required...)
	return x.DefaultUint(0), e
}
func (r *Req) BodyUint64(p string, required ...bool) (uint64, *ae.Error) {
	x, e := r.BodyDigit(p, true, required...)
	return x.DefaultUint64(0), e
}
func (r *Req) BodyProvince(p string, required ...bool) (atype.Province, *ae.Error) {
	distri, e := r.BodyDistri(p, required...)
	if e != nil {
		return 0, e
	}
	return distri.Province(), nil
}
func (r *Req) BodyDist(p string, required ...bool) (atype.Dist, *ae.Error) {
	distri, e := r.BodyDistri(p, required...)
	if e != nil {
		return 0, e
	}
	return distri.Dist(), nil
}
func (r *Req) BodyDistri(p string, required ...bool) (atype.Distri, *ae.Error) {
	x, e := r.BodyDigit(p, false, required...)
	return atype.NewDistri(x.DefaultUint24(0)), e
}

func (r *Req) BodySmallMoney(p string, required ...bool) (atype.SmallMoney, *ae.Error) {
	x, e := r.BodyDigit(p, false, required...)
	return atype.SmallMoney(x.DefaultUint(0)), e
}
func (r *Req) BodyMoney(p string, required ...bool) (atype.Money, *ae.Error) {
	x, e := r.BodyDigit(p, false, required...)
	return atype.Money(x.DefaultInt64(0)), e
}
func (r *Req) BodyPercent16(p string, required ...bool) (atype.Percent16, *ae.Error) {
	x, e := r.BodyDigit(p, false, required...)
	return atype.NewPercent16(x.DefaultInt16(0)), e
}
func (r *Req) BodyPercent24(p string, required ...bool) (atype.Percent24, *ae.Error) {
	x, e := r.BodyDigit(p, false, required...)
	return atype.NewPercent24(x.DefaultInt32(0)), e
}
func (r *Req) BodyPercent(p string, required ...bool) (atype.Percent, *ae.Error) {
	x, e := r.BodyDigit(p, false, required...)
	return atype.NewPercent(x.DefaultInt(0)), e
}

func (r *Req) BodyDate(p string, loc *time.Location, required ...bool) (atype.Date, *ae.Error) {
	x, e := r.Body(p, `^`+aenum.DateRegExp+`$`, len(required) == 0 || required[0])
	if e != nil {
		return "", ae.NewError(400, "invalid date ("+p+"): "+x.String())
	}
	return atype.NewDate(x.String(), loc), nil
}
func (r *Req) BodyDatetime(p string, loc *time.Location, required ...bool) (atype.Datetime, *ae.Error) {
	x, e := r.Body(p, `^`+aenum.DatetimeRegExp+`$`, len(required) == 0 || required[0])
	if e != nil {
		return "", ae.NewError(400, "invalid datetime ("+p+"): "+x.String())
	}
	return atype.NewDatetime(x.String(), loc), nil
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

// 不可再指定offset/limit了，单一原则，通过page分页
// @param firstPageLimit 首页行数
// @param limitMax 其他页行数
func (r *Req) QueryPaging(args ...uint) atype.Paging {
	page, _ := r.QueryUint(ParamPage, false)
	return atype.NewPaging(page, args...)
}
func (r *Req) BodyImage(p string, required ...bool) (atype.Image, *ae.Error) {
	x, e := r.BodyString(p, len(required) == 0 || required[0])
	return atype.NewImage(x, true), e
}
func (r *Req) BodyAudio(p string, required ...bool) (atype.Audio, *ae.Error) {
	x, e := r.BodyString(p, len(required) == 0 || required[0])
	return atype.NewAudio(x, true), e
}
func (r *Req) BodyVideo(p string, required ...bool) (atype.Video, *ae.Error) {
	x, e := r.BodyString(p, len(required) == 0 || required[0])
	return atype.NewVideo(x, true), e
}
func (r *Req) BodyImages(p string, required ...bool) ([]atype.Image, *ae.Error) {
	x, e := r.BodyStrings(p, len(required) == 0 || required[0], false)
	if e != nil || len(x) == 0 {
		return nil, e
	}
	imgs := make([]atype.Image, len(x))
	for i, s := range x {
		imgs[i] = atype.NewImage(s, true)
	}
	return imgs, e
}

// 下面很少使用，就不要封装了。以后使用的时候，业务层直接组装就行
//func (r *Req) BodyAudios(p string, required, allowEmptyString, filenameOnly bool) (atype.Audios, *ae.Error) {
//	x, e := r.BodyStrings(p, required, allowEmptyString)
//	return atype.ToAudios(x, filenameOnly), e
//}
//func (r *Req) BodyVideos(p string, required, allowEmptyString, filenameOnly bool) (atype.Videos, *ae.Error) {
//	x, e := r.BodyStrings(p, required, allowEmptyString)
//	return atype.ToVideos(x, filenameOnly), e
//}
func (r *Req) BodyCoordinate(p string, required ...bool) (*atype.Coordinate, *ae.Error) {
	x, e := r.BodyFloat64Map(p, required...)
	if e != nil || x == nil {
		return nil, e
	}
	lat, ok := x["lat"]
	if !ok {
		return nil, ae.BadParam(p)
	}
	lng, ok := x["lng"]
	if !ok {
		return nil, ae.BadParam(p)
	}
	height, _ := x["height"]
	coord := atype.Coordinate{
		Latitude:  lat,
		Longitude: lng,
		Height:    height,
	}
	return &coord, nil
}
func (r *Req) BodyPosition(required bool) (coord *atype.Coordinate, distri atype.Distri, addr string, e *ae.Error) {
	coord, e = r.BodyCoordinate("position", required)
	if e != nil {
		return
	}
	distri, e = r.BodyDistri("distri", required)
	if e != nil {
		return
	}
	addr, e = r.BodyString("addr", required)
	return
}
