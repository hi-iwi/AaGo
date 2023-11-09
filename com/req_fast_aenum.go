package com

import (
	"github.com/hi-iwi/AaGo/ae"
	"github.com/hi-iwi/AaGo/aenum"
)

func reqStatus(method func(string, ...bool) (int8, *ae.Error), p string, xargs ...bool) (aenum.Status, *ae.Error) {
	required := len(xargs) == 0 || xargs[0]
	sts, e := method(p, required)
	if e != nil {
		return 0, e
	}

	status, ok := aenum.NewStatus(sts)
	if !ok {
		return 0, ae.BadParam(p)
	}
	return status, nil
}
func (r *Req) QueryStatus(p string, xargs ...bool) (aenum.Status, *ae.Error) {
	return reqStatus(r.QueryInt8, p, xargs...)
}
func (r *Req) BodyStatus(p string, xargs ...bool) (aenum.Status, *ae.Error) {
	return reqStatus(r.BodyInt8, p, xargs...)
}
func reqCountry(method func(string, ...bool) (uint16, *ae.Error), p string, xargs ...bool) (aenum.Country, *ae.Error) {
	required := len(xargs) == 0 || xargs[0]
	sts, e := method(p, required)
	if e != nil {
		return 0, e
	}

	status, ok := aenum.NewCountry(sts)
	if !ok {
		return 0, ae.BadParam(p)
	}
	return status, nil
}
func (r *Req) QueryCountry(p string, xargs ...bool) (aenum.Country, *ae.Error) {
	return reqCountry(r.QueryUint16, p, xargs...)
}
func (r *Req) BodyCountry(p string, xargs ...bool) (aenum.Country, *ae.Error) {
	return reqCountry(r.BodyUint16, p, xargs...)
}
func reqSex(method func(string, ...bool) (uint8, *ae.Error), p string, xargs ...bool) (aenum.Sex, *ae.Error) {
	required := len(xargs) == 0 || xargs[0]
	sts, e := method(p, required)
	if e != nil {
		return 0, e
	}
	
	return aenum.NewSex(sts), nil
}
func (r *Req) QuerySex(p string, xargs ...bool) (aenum.Sex, *ae.Error) {
	return reqSex(r.QueryUint8, p, xargs...)
}
func (r *Req) BodySex(p string, xargs ...bool) (aenum.Sex, *ae.Error) {
	return reqSex(r.BodyUint8, p, xargs...)
}
