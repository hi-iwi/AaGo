package com

import (
	"github.com/hi-iwi/AaGo/ae"
	"github.com/hi-iwi/AaGo/atype"
	"reflect"
)

func (r *Req) BodyStrings(p string, required, allowEmptyString bool) ([]string, *ae.Error) {
	q, e := r.Body(p, required)
	if e != nil {
		return nil, e
	}
	var v []string
	d := q.Raw()
	if d != nil {
		switch reflect.TypeOf(d).Kind() {
		case reflect.String: // "1,2,3"
			return r.BodySepStrings(p, ",", required, allowEmptyString)
		case reflect.Slice: // 有可能是 [1,"2",3] 这种混合的数组
			s := reflect.ValueOf(d)
			v = make([]string, 0, s.Len())
			for i := 0; i < s.Len(); i++ {
				ts := atype.String(s.Index(i).Interface()) // 不能用 s.Index(i).String()，否则返回：<interface {} Value>
				if allowEmptyString || ts != "" {
					v = append(v, ts)
				}
			}
		}
	}
	if len(v) == 0 && required {
		return nil, ae.BadParam(p)
	}
	return v, nil
}
func (r *Req) BodyUint64s(p string, required, allowZero bool) ([]uint64, *ae.Error) {
	q, e := r.Body(p, required)
	if e != nil {
		return nil, e
	}
	var v []uint64
	d := q.Raw()
	switch reflect.TypeOf(d).Kind() {
	case reflect.String:
		return r.BodySepUint64s(p, ",", required, allowZero)
	case reflect.Slice: // 有可能是 [1,"2",3] 这种混合的数组
		s := reflect.ValueOf(d)
		v = make([]uint64, 0, s.Len())
		var n uint64
		var err error
		for i := 0; i < s.Len(); i++ {
			n, err = atype.Uint64(s.Index(i).Interface())
			if err != nil {
				return nil, ae.BadParam(p)
			}
			if allowZero || n > 0 {
				v = append(v, n)
			}
		}
	}
	if len(v) == 0 && required {
		return nil, ae.BadParam(p)
	}
	return v, nil
}
func (r *Req) BodyUints(p string, required, allowZero bool) ([]uint, *ae.Error) {
	q, e := r.Body(p, required)
	if e != nil {
		return nil, e
	}
	var v []uint
	d := q.Raw()
	switch reflect.TypeOf(d).Kind() {
	case reflect.String:
		return r.BodySepUints(p, ",", required, allowZero)
	case reflect.Slice: // 有可能是 [1,"2",3] 这种混合的数组
		s := reflect.ValueOf(d)
		v = make([]uint, 0, s.Len())
		var n uint
		var err error
		for i := 0; i < s.Len(); i++ {
			n, err = atype.Uint(s.Index(i).Interface())
			if err != nil {
				return nil, ae.BadParam(p)
			}
			if allowZero || n > 0 {
				v = append(v, n)
			}
		}
	}
	if len(v) == 0 && required {
		return nil, ae.BadParam(p)
	}
	return v, nil
}
func (r *Req) BodyUint32s(p string, required, allowZero bool) ([]uint32, *ae.Error) {
	q, e := r.Body(p, required)
	if e != nil {
		return nil, e
	}
	var v []uint32
	d := q.Raw()
	switch reflect.TypeOf(d).Kind() {
	case reflect.String:
		return r.BodySepUint32s(p, ",", required, allowZero)
	case reflect.Slice: // 有可能是 [1,"2",3] 这种混合的数组
		s := reflect.ValueOf(d)
		v = make([]uint32, 0, s.Len())
		var n uint32
		var err error
		for i := 0; i < s.Len(); i++ {
			n, err = atype.Uint32(s.Index(i).Interface())
			if err != nil {
				return nil, ae.BadParam(p)
			}
			if allowZero || n > 0 {
				v = append(v, n)
			}
		}
	}
	if len(v) == 0 && required {
		return nil, ae.BadParam(p)
	}
	return v, nil
}
func (r *Req) BodyUint24s(p string, required, allowZero bool) ([]atype.Uint24, *ae.Error) {
	q, e := r.Body(p, required)
	if e != nil {
		return nil, e
	}
	var v []atype.Uint24
	d := q.Raw()
	switch reflect.TypeOf(d).Kind() {
	case reflect.String:
		return r.BodySepUint24s(p, ",", required, allowZero)
	case reflect.Slice: // 有可能是 [1,"2",3] 这种混合的数组
		s := reflect.ValueOf(d)
		v = make([]atype.Uint24, 0, s.Len())
		var n atype.Uint24
		var err error
		for i := 0; i < s.Len(); i++ {
			n, err = atype.Uint24b(s.Index(i).Interface())
			if err != nil {
				return nil, ae.BadParam(p)
			}
			if allowZero || n > 0 {
				v = append(v, n)
			}
		}
	}
	if len(v) == 0 && required {
		return nil, ae.BadParam(p)
	}
	return v, nil
}
func (r *Req) BodyUint16s(p string, required, allowZero bool) ([]uint16, *ae.Error) {
	q, e := r.Body(p, required)
	if e != nil {
		return nil, e
	}
	var v []uint16
	d := q.Raw()
	switch reflect.TypeOf(d).Kind() {
	case reflect.String:
		return r.BodySepUint16s(p, ",", required, allowZero)
	case reflect.Slice: // 有可能是 [1,"2",3] 这种混合的数组
		s := reflect.ValueOf(d)
		v = make([]uint16, 0, s.Len())
		var n uint16
		var err error
		for i := 0; i < s.Len(); i++ {
			n, err = atype.Uint16(s.Index(i).Interface())
			if err != nil {
				return nil, ae.BadParam(p)
			}
			if allowZero || n > 0 {
				v = append(v, n)
			}
		}
	}
	if len(v) == 0 && required {
		return nil, ae.BadParam(p)
	}
	return v, nil
}
func (r *Req) BodyUint8s(p string, required, allowZero bool) ([]uint8, *ae.Error) {
	q, e := r.Body(p, required)
	if e != nil {
		return nil, e
	}
	var v []uint8
	d := q.Raw()
	switch reflect.TypeOf(d).Kind() {
	case reflect.String:
		return r.BodySepUint8s(p, ",", required, allowZero)
	case reflect.Slice: // 有可能是 [1,"2",3] 这种混合的数组
		s := reflect.ValueOf(d)
		v = make([]uint8, 0, s.Len())
		var n uint8
		var err error
		for i := 0; i < s.Len(); i++ {
			n, err = atype.Uint8(s.Index(i).Interface())
			if err != nil {
				return nil, ae.BadParam(p)
			}
			if allowZero || n > 0 {
				v = append(v, n)
			}
		}
	}
	if len(v) == 0 && required {
		return nil, ae.BadParam(p)
	}
	return v, nil
}

func (r *Req) BodyInt8s(p string, required, allowZero bool) ([]int8, *ae.Error) {
	q, e := r.Body(p, required)
	if e != nil {
		return nil, e
	}
	var v []int8
	d := q.Raw()
	switch reflect.TypeOf(d).Kind() {
	case reflect.String:
		return r.BodySepInt8s(p, ",", required, allowZero)
	case reflect.Slice: // 有可能是 [1,"2",3] 这种混合的数组
		s := reflect.ValueOf(d)
		v = make([]int8, 0, s.Len())
		var n int8
		var err error
		for i := 0; i < s.Len(); i++ {
			n, err = atype.Int8(s.Index(i).Interface())
			if err != nil {
				return nil, ae.BadParam(p)
			}
			if allowZero || n > 0 {
				v = append(v, n)
			}
		}
	}
	if len(v) == 0 && required {
		return nil, ae.BadParam(p)
	}
	return v, nil
}

func (r *Req) BodyFloat64s(p string, required, allowZero bool) ([]float64, *ae.Error) {
	q, e := r.Body(p, required)
	if e != nil {
		return nil, e
	}
	var v []float64
	d := q.Raw()
	switch reflect.TypeOf(d).Kind() {
	case reflect.String:
		return r.BodySepFloat64s(p, ",", required, allowZero)
	case reflect.Slice: // 有可能是 [1,"2",3] 这种混合的数组
		s := reflect.ValueOf(d)
		v = make([]float64, 0, s.Len())
		var n float64
		var err error
		for i := 0; i < s.Len(); i++ {
			n, err = atype.Float64(s.Index(i).Interface(), 64)
			if err != nil {
				return nil, ae.BadParam(p)
			}
			if allowZero || n > 0 {
				v = append(v, n)
			}
		}
	}
	if len(v) == 0 && required {
		return nil, ae.BadParam(p)
	}
	return v, nil
}

func (r *Req) BodyFloat32s(p string, required, allowZero bool) ([]float32, *ae.Error) {
	q, e := r.Body(p, required)
	if e != nil {
		return nil, e
	}
	var v []float32
	d := q.Raw()
	switch reflect.TypeOf(d).Kind() {
	case reflect.String:
		return r.BodySepFloat32s(p, ",", required, allowZero)
	case reflect.Slice: // 有可能是 [1,"2",3] 这种混合的数组
		s := reflect.ValueOf(d)
		v = make([]float32, 0, s.Len())
		var n float32
		var err error
		for i := 0; i < s.Len(); i++ {
			n, err = atype.Float32(s.Index(i).Interface())
			if err != nil {
				return nil, ae.BadParam(p)
			}
			if allowZero || n > 0 {
				v = append(v, n)
			}
		}
	}
	if len(v) == 0 && required {
		return nil, ae.BadParam(p)
	}
	return v, nil
}
