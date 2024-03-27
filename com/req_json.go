package com

import (
	"github.com/hi-iwi/AaGo/ae"
	"github.com/hi-iwi/AaGo/atype"
)

func (r *Req) BodyInterfaceMap(p string, requireds ...bool) (map[string]any, *ae.Error) {
	required := len(requireds) == 0 || requireds[0]
	x, e := r.Body(p, required)
	if e != nil {
		return nil, e
	}
	if x.IsNil() || x.String() == "" {
		if required {
			return nil, ae.BadParam(p)
		}
		return nil, nil
	}
	raw := x.Raw()
	b, ok := raw.(map[string]any)
	if !ok {
		return nil, ae.BadParam(p)
	}
	if len(b) == 0 {
		if required {
			return nil, ae.BadParam(p)
		}
		return nil, nil
	}
	return b, nil
}

func (r *Req) BodyInterfaces(p string, requireds ...bool) ([]any, *ae.Error) {
	required := len(requireds) == 0 || requireds[0]
	x, e := r.Body(p, required)
	if e != nil {
		return nil, e
	}
	if x.IsNil() || x.String() == "" {
		if required {
			return nil, ae.BadParam(p)
		}
		return nil, nil
	}
	raw := x.Raw()
	b, ok := raw.([]any)
	if !ok {
		return nil, ae.BadParam(p)
	}
	if len(b) == 0 {
		if required {
			return nil, ae.BadParam(p)
		}
		return nil, nil
	}
	return b, nil
}
func (r *Req) BodyFloat64Map(p string, requireds ...bool) (map[string]float64, *ae.Error) {
	b, e := r.BodyInterfaceMap(p, requireds...)
	if e != nil {
		return nil, e
	}
	maps, err := atype.ConvFloat64Map(b)
	if err != nil {
		return nil, ae.BadParam(p)
	}
	required := len(requireds) == 0 || requireds[0]
	if required && maps == nil {
		return nil, ae.BadParam(p)
	}
	return maps, nil
}
func (r *Req) BodyStringMap(p string, requireds ...bool) (map[string]string, *ae.Error) {
	b, e := r.BodyInterfaceMap(p, requireds...)
	if e != nil {
		return nil, e
	}
	maps := atype.ConvStringMap(b)
	required := len(requireds) == 0 || requireds[0]
	if required && maps == nil {
		return nil, ae.BadParam(p)
	}
	return maps, nil
}
func (r *Req) BodyStringsMap(p string, requireds ...bool) (map[string][]string, *ae.Error) {
	b, e := r.BodyInterfaceMap(p, requireds...)
	if e != nil {
		return nil, e
	}
	maps := atype.ConvStringsMap(b)
	required := len(requireds) == 0 || requireds[0]
	if required && maps == nil {
		return nil, ae.BadParam(p)
	}
	return maps, nil
}
func (r *Req) BodyComplexStringMap(p string, requireds ...bool) (map[string]map[string]string, *ae.Error) {
	b, e := r.BodyInterfaceMap(p, requireds...)
	if e != nil {
		return nil, e
	}
	maps := atype.ConvComplexStringMap(b)
	required := len(requireds) == 0 || requireds[0]
	if required && maps == nil {
		return nil, ae.BadParam(p)
	}
	return maps, nil
}
func (r *Req) BodyComplexStringsMap(p string, requireds ...bool) (map[string][][]string, *ae.Error) {
	b, e := r.BodyInterfaceMap(p, requireds...)
	if e != nil {
		return nil, e
	}
	maps := atype.ConvComplexStringsMap(b)
	required := len(requireds) == 0 || requireds[0]
	if required && maps == nil {
		return nil, ae.BadParam(p)
	}
	return maps, nil
}
func (r *Req) BodyConvStringMaps(p string, requireds ...bool) ([]map[string]string, *ae.Error) {
	b, e := r.BodyInterfaces(p, requireds...)
	if e != nil {
		return nil, e
	}
	maps := atype.ConvStringMaps(b)
	required := len(requireds) == 0 || requireds[0]
	if required && maps == nil {
		return nil, ae.BadParam(p)
	}
	return maps, nil
}

func (r *Req) BodyComplexMaps(p string, requireds ...bool) ([]map[string]any, *ae.Error) {
	b, e := r.BodyInterfaces(p, requireds...)
	if e != nil {
		return nil, e
	}
	maps := atype.ConvComplexMaps(b)
	required := len(requireds) == 0 || requireds[0]
	if required && maps == nil {
		return nil, ae.BadParam(p)
	}
	return maps, nil
}
