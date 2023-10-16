package com

import "github.com/hi-iwi/AaGo/ae"

func (r *Req) QueryValid(p string, required bool, validator func(string) bool) (string, *ae.Error) {
	x, e := r.QueryString(p, required)
	if e != nil {
		return "", e
	}
	if !required && x == "" {
		return "", nil
	}
	if ok := validator(x); !ok {
		return "", ae.BadParam(p)
	}
	return x, nil
}
func (r *Req) QueryValid8(p string, required bool, validator func(uint8) bool) (uint8, *ae.Error) {
	x, e := r.QueryUint8(p, required)
	if e != nil {
		return 0, e
	}
	if !required && x == 0 {
		return 0, nil
	}
	if ok := validator(x); !ok {
		return 0, ae.BadParam(p)
	}
	return x, nil
}
func (r *Req) QueryEnum(p string, required bool, validators []string) (string, *ae.Error) {
	x, e := r.QueryString(p, required)
	if e != nil {
		return "", e
	}
	if !required && x == "" {
		return "", nil
	}
	for _, val := range validators {
		if x == val {
			return x, nil
		}
	}
	return x, ae.BadParam(p)
}
func (r *Req) QueryEnum8(p string, required bool, validators []uint8) (uint8, *ae.Error) {
	x, e := r.QueryUint8(p, required)
	if e != nil {
		return 0, e
	}
	if !required && x == 0 {
		return 0, nil
	}
	for _, val := range validators {
		if x == val {
			return x, nil
		}
	}
	return x, ae.BadParam(p)
}
func (r *Req) QueryEnum8i(p string, required bool, validators []int8) (int8, *ae.Error) {
	x, e := r.QueryInt8(p, required)
	if e != nil {
		return 0, e
	}
	if !required && x == 0 {
		return 0, nil
	}
	for _, val := range validators {
		if x == val {
			return x, nil
		}
	}
	return x, ae.BadParam(p)
}
func (r *Req) BodyValid(p string, required bool, validator func(string) bool) (string, *ae.Error) {
	x, e := r.BodyString(p, required)
	if e != nil {
		return "", e
	}
	if !required && x == "" {
		return "", nil
	}
	if ok := validator(x); !ok {
		return "", ae.BadParam(p)
	}
	return x, nil
}
func (r *Req) BodyValid8(p string, required bool, validator func(uint8) bool) (uint8, *ae.Error) {
	x, e := r.BodyUint8(p, required)
	if e != nil {
		return 0, e
	}
	if !required && x == 0 {
		return 0, nil
	}
	if ok := validator(x); !ok {
		return 0, ae.BadParam(p)
	}
	return x, nil
}
func (r *Req) BodyEnum(p string, required bool, validators []string) (string, *ae.Error) {
	x, e := r.BodyString(p, required)
	if e != nil {
		return "", e
	}
	if !required && x == "" {
		return "", nil
	}
	for _, val := range validators {
		if x == val {
			return x, nil
		}
	}
	return x, ae.BadParam(p)
}
func (r *Req) BodyEnum8(p string, required bool, validators []uint8) (uint8, *ae.Error) {
	x, e := r.BodyUint8(p, required)
	if e != nil {
		return 0, e
	}
	if !required && x == 0 {
		return 0, nil
	}
	for _, val := range validators {
		if x == val {
			return x, nil
		}
	}
	return x, ae.BadParam(p)
}
func (r *Req) BodyEnum8i(p string, required bool, validators []int8) (int8, *ae.Error) {
	x, e := r.BodyInt8(p, required)
	if e != nil {
		return 0, e
	}
	if !required && x == 0 {
		return 0, nil
	}
	for _, val := range validators {
		if x == val {
			return x, nil
		}
	}
	return x, ae.BadParam(p)
}
