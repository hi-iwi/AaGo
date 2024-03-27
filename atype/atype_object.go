package atype

import (
	"encoding/json"
	"reflect"
)

func (p *Atype) IsNil() bool {
	if p.raw == nil {
		return true
	}
	switch reflect.TypeOf(p.raw).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(p.raw).IsNil()
	}
	return false
}
func (p *Atype) Strings() ([]string, bool) {
	v, ok := p.raw.([]string)
	if ok {
		return v, true
	}
	ar, ok := p.raw.([]any)
	if !ok {
		return nil, false
	}
	v = make([]string, len(ar))
	for i, a := range ar {
		if v[i], ok = a.(string); !ok {
			if bs, ok := a.([]byte); !ok {
				return nil, false
			} else {
				v[i] = string(bs)
			}
		}
	}
	return v, true
}
func (p *Atype) Ints() ([]int, bool) {
	v, ok := p.raw.([]int)
	if ok {
		return v, true
	}
	ar, ok := p.raw.([]any)
	if !ok {
		return nil, false
	}
	var err error
	v = make([]int, len(ar))
	for i, a := range ar {
		if v[i], err = New(a).Int(); err != nil {
			return nil, false
		}
	}
	return v, true
}
func (p *Atype) Uints() ([]uint, bool) {
	v, ok := p.raw.([]uint)
	if ok {
		return v, true
	}
	ar, ok := p.raw.([]any)
	if !ok {
		return nil, false
	}
	var err error
	v = make([]uint, len(ar))
	for i, a := range ar {
		if v[i], err = New(a).Uint(); err != nil {
			return nil, false
		}
	}
	return v, true
}
func (p *Atype) Int64s() ([]int64, bool) {
	v, ok := p.raw.([]int64)
	if ok {
		return v, true
	}
	ar, ok := p.raw.([]any)
	if !ok {
		return nil, false
	}
	var err error
	v = make([]int64, len(ar))
	for i, a := range ar {
		if v[i], err = New(a).Int64(); err != nil {
			return nil, false
		}
	}
	return v, true
}
func (p *Atype) Uint64s() ([]uint64, bool) {
	v, ok := p.raw.([]uint64)
	if ok {
		return v, true
	}
	ar, ok := p.raw.([]any)
	if !ok {
		return nil, false
	}
	var err error
	v = make([]uint64, len(ar))
	for i, a := range ar {
		if v[i], err = New(a).Uint64(); err != nil {
			return nil, false
		}
	}
	return v, true
}
func (p *Atype) Float32s() ([]float32, bool) {
	v, ok := p.raw.([]float32)
	if ok {
		return v, true
	}
	ar, ok := p.raw.([]any)
	if !ok {
		return nil, false
	}
	var err error
	v = make([]float32, len(ar))
	for i, a := range ar {
		if v[i], err = New(a).Float32(); err != nil {
			return nil, false
		}
	}
	return v, true
}
func (p *Atype) Float64s() ([]float64, bool) {
	v, ok := p.raw.([]float64)
	if ok {
		return v, true
	}
	ar, ok := p.raw.([]any)
	if !ok {
		return nil, false
	}
	var err error
	v = make([]float64, len(ar))
	for i, a := range ar {
		if v[i], err = New(a).Float64(); err != nil {
			return nil, false
		}
	}
	return v, true
}
func (p *Atype) ArrayJson(allowNil bool) (json.RawMessage, bool) {
	// 也可能客户端传的是 string ，也可能是对象原数
	bytes, ok := p.raw.([]byte)
	if ok {
		if bytes[0] == '[' {
			return bytes, true
		} else {
			return bytes, false
		}
	}

	uint8s, ok := p.raw.([]uint8)
	if ok {
		v, _ := MarshalUint8s(uint8s)
		return v, true
	}

	arr, ok := p.raw.([]any)
	if ok {
		v, _ := json.Marshal(arr)

		return v, true
	}
	if allowNil {
		if p.IsNil() {
			return nil, true
		} else if s, _ := p.raw.(string); s == "" {
			return nil, true
		}
	}

	return nil, false
}
func (p *Atype) MapJson(allowNil bool) (json.RawMessage, bool) {
	arr, ok := p.raw.(map[string]any)
	if ok {
		v, _ := json.Marshal(arr)
		return v, true
	}
	if allowNil {
		if p.IsNil() {
			return nil, true
		} else if s, _ := p.raw.(string); s == "" {
			return nil, true
		}
	}

	// 也可能客户端传的是 string ，也可能使对象原数
	bytes, ok := p.raw.([]byte)
	if ok {
		if bytes[0] == '{' {
			return bytes, true
		} else {
			return bytes, false
		}
	}

	return nil, false
}
