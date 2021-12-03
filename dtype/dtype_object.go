package dtype

import "encoding/json"

func (p *Dtype) Strings() ([]string, bool) {
	v, ok := p.raw.([]string)
	return v, ok
}
func (p *Dtype) Ints() ([]int, bool) {
	v, ok := p.raw.([]int)
	return v, ok
}
func (p *Dtype) Uints() ([]uint, bool) {
	v, ok := p.raw.([]uint)
	return v, ok
}
func (p *Dtype) Int64s() ([]int64, bool) {
	v, ok := p.raw.([]int64)
	return v, ok
}
func (p *Dtype) Uint64s() ([]uint64, bool) {
	v, ok := p.raw.([]uint64)
	return v, ok
}
func (p *Dtype) Float32s() ([]float32, bool) {
	v, ok := p.raw.([]float32)
	return v, ok
}
func (p *Dtype) Float64s() ([]float64, bool) {
	v, ok := p.raw.([]float64)
	return v, ok
}
func (p *Dtype) ArrayJson() (json.RawMessage, bool) {
	arr, ok := p.raw.([]interface{})
	if ok {
		v, _ := json.Marshal(arr)
		return v, true
	}
	// 也可能客户端传的是 string ，也可能使对象原数
	bytes, ok := p.raw.([]byte)
	if ok {
		if bytes[0] == '[' {
			return bytes, true
		} else {
			return bytes, false
		}
	}

	return nil, false
}
func (p *Dtype) MapJson() (json.RawMessage, bool) {
	arr, ok := p.raw.(map[string]interface{})
	if ok {
		v, _ := json.Marshal(arr)
		return v, true
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