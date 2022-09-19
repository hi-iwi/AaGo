package com

import (
	"github.com/hi-iwi/AaGo/ae"
	"github.com/hi-iwi/AaGo/atype"
	"reflect"
)

func (r *Req) BodyJsonStrings(p string, required bool) ([]string, *ae.Error) {
	q, e := r.Body(p, required)
	if e != nil {
		return nil, e
	}
	var v []string
	d := q.Raw()
	switch reflect.TypeOf(d).Kind() {
	case reflect.Slice: // 有可能是 [1,"2",3] 这种混合的数组
		s := reflect.ValueOf(d)
		v = make([]string, s.Len())
		for i := 0; i < s.Len(); i++ {
			v[i] = atype.String(s.Index(i).Interface())
		}
	}
	if len(v) == 0 && required {
		return nil, ae.BadParam(p)
	}
	return v, nil
}
func (r *Req) BodyJsonUint64s(p string, required bool) ([]uint64, *ae.Error) {
	q, e := r.Body(p, required)
	if e != nil {
		return nil, e
	}
	var v []uint64
	d := q.Raw()
	switch reflect.TypeOf(d).Kind() {
	case reflect.Slice: // 有可能是 [1,"2",3] 这种混合的数组
		s := reflect.ValueOf(d)
		v = make([]uint64, s.Len())
		var err error
		for i := 0; i < s.Len(); i++ {
			v[i], err = atype.Uint64(s.Index(i).Interface())
			if err != nil {
				return nil, ae.BadParam(p)
			}
		}
	}
	if len(v) == 0 && required {
		return nil, ae.BadParam(p)
	}
	return v, nil
}
func (r *Req) BodyJsonUints(p string, required bool) ([]uint, *ae.Error) {
	q, e := r.Body(p, required)
	if e != nil {
		return nil, e
	}
	var v []uint
	d := q.Raw()
	switch reflect.TypeOf(d).Kind() {
	case reflect.Slice: // 有可能是 [1,"2",3] 这种混合的数组
		s := reflect.ValueOf(d)
		v = make([]uint, s.Len())
		var err error
		for i := 0; i < s.Len(); i++ {
			v[i], err = atype.Uint(s.Index(i).Interface())
			if err != nil {
				return nil, ae.BadParam(p)
			}
		}
	}
	if len(v) == 0 && required {
		return nil, ae.BadParam(p)
	}
	return v, nil
}
func (r *Req) BodyJsonUint32s(p string, required bool) ([]uint32, *ae.Error) {
	q, e := r.Body(p, required)
	if e != nil {
		return nil, e
	}
	var v []uint32
	d := q.Raw()
	switch reflect.TypeOf(d).Kind() {
	case reflect.Slice: // 有可能是 [1,"2",3] 这种混合的数组
		s := reflect.ValueOf(d)
		v = make([]uint32, s.Len())
		var err error
		for i := 0; i < s.Len(); i++ {
			v[i], err = atype.Uint32(s.Index(i).Interface())
			if err != nil {
				return nil, ae.BadParam(p)
			}
		}
	}
	if len(v) == 0 && required {
		return nil, ae.BadParam(p)
	}
	return v, nil
}
func (r *Req) BodyJsonUint24s(p string, required bool) ([]atype.Uint24, *ae.Error) {
	q, e := r.Body(p, required)
	if e != nil {
		return nil, e
	}
	var v []atype.Uint24
	d := q.Raw()
	switch reflect.TypeOf(d).Kind() {
	case reflect.Slice: // 有可能是 [1,"2",3] 这种混合的数组
		s := reflect.ValueOf(d)
		v = make([]atype.Uint24, s.Len())
		var err error
		for i := 0; i < s.Len(); i++ {
			v[i], err = atype.Uint24b(s.Index(i).Interface())
			if err != nil {
				return nil, ae.BadParam(p)
			}
		}
	}
	if len(v) == 0 && required {
		return nil, ae.BadParam(p)
	}
	return v, nil
}
func (r *Req) BodyJsonUint16s(p string, required bool) ([]uint16, *ae.Error) {
	q, e := r.Body(p, required)
	if e != nil {
		return nil, e
	}
	var v []uint16
	d := q.Raw()
	switch reflect.TypeOf(d).Kind() {
	case reflect.Slice: // 有可能是 [1,"2",3] 这种混合的数组
		s := reflect.ValueOf(d)
		v = make([]uint16, s.Len())
		var err error
		for i := 0; i < s.Len(); i++ {
			v[i], err = atype.Uint16(s.Index(i).Interface())
			if err != nil {
				return nil, ae.BadParam(p)
			}
		}
	}
	if len(v) == 0 && required {
		return nil, ae.BadParam(p)
	}
	return v, nil
}
func (r *Req) BodyJsonUint8s(p string, required bool) ([]uint8, *ae.Error) {
	q, e := r.Body(p, required)
	if e != nil {
		return nil, e
	}
	var v []uint8
	d := q.Raw()
	switch reflect.TypeOf(d).Kind() {
	case reflect.Slice: // 有可能是 [1,"2",3] 这种混合的数组
		s := reflect.ValueOf(d)
		v = make([]uint8, s.Len())
		var err error
		for i := 0; i < s.Len(); i++ {
			v[i], err = atype.Uint8(s.Index(i).Interface())
			if err != nil {
				return nil, ae.BadParam(p)
			}
		}
	}
	if len(v) == 0 && required {
		return nil, ae.BadParam(p)
	}
	return v, nil
}

func (r *Req) BodyJsonInt8s(p string, required bool) ([]int8, *ae.Error) {
	q, e := r.Body(p, required)
	if e != nil {
		return nil, e
	}
	var v []int8
	d := q.Raw()
	switch reflect.TypeOf(d).Kind() {
	case reflect.Slice: // 有可能是 [1,"2",3] 这种混合的数组
		s := reflect.ValueOf(d)
		v = make([]int8, s.Len())
		var err error
		for i := 0; i < s.Len(); i++ {
			v[i], err = atype.Int8(s.Index(i).Interface())
			if err != nil {
				return nil, ae.BadParam(p)
			}
		}
	}
	if len(v) == 0 && required {
		return nil, ae.BadParam(p)
	}
	return v, nil
}

func (r *Req) BodyJsonFloat64s(p string, required bool) ([]float64, *ae.Error) {
	q, e := r.Body(p, required)
	if e != nil {
		return nil, e
	}
	var v []float64
	d := q.Raw()
	switch reflect.TypeOf(d).Kind() {
	case reflect.Slice: // 有可能是 [1,"2",3] 这种混合的数组
		s := reflect.ValueOf(d)
		v = make([]float64, s.Len())
		var err error
		for i := 0; i < s.Len(); i++ {
			v[i], err = atype.Float64(s.Index(i).Interface())
			if err != nil {
				return nil, ae.BadParam(p)
			}
		}
	}
	if len(v) == 0 && required {
		return nil, ae.BadParam(p)
	}
	return v, nil
}

func (r *Req) BodyJsonFloat32s(p string, required bool) ([]float32, *ae.Error) {
	q, e := r.Body(p, required)
	if e != nil {
		return nil, e
	}
	var v []float32
	d := q.Raw()
	switch reflect.TypeOf(d).Kind() {
	case reflect.Slice: // 有可能是 [1,"2",3] 这种混合的数组
		s := reflect.ValueOf(d)
		v = make([]float32, s.Len())
		var err error
		for i := 0; i < s.Len(); i++ {
			v[i], err = atype.Float32(s.Index(i).Interface())
			if err != nil {
				return nil, ae.BadParam(p)
			}
		}
	}
	if len(v) == 0 && required {
		return nil, ae.BadParam(p)
	}
	return v, nil
}
