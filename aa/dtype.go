package aa

import (
	"github.com/luexu/AaGo/dtype"
)

type Dtype struct {
	Value interface{}
}

func NewDtype(data interface{}) *Dtype {
	return &Dtype{
		Value: data,
	}
}

func (p *Dtype) NotEmpty() bool {
	return dtype.NotEmpty(p.Value)
}
func (p *Dtype) Interface() interface{} {
	return p.Value
}

func (p *Dtype) Bool() (bool, error) {
	return dtype.Bool(p.Value)
}

func (p *Dtype) Slice() ([]interface{}, error) {
	return dtype.Slice(p.Value)
}

func (p *Dtype) String() string {
	return dtype.String(p.Value)
}

func (p *Dtype) Bytes() []byte {
	return dtype.Bytes(p.Value)
}

func (p *Dtype) Int() (int, error) {
	return dtype.Int(p.Value)
}

func (p *Dtype) Int64() (int64, error) {
	return dtype.Int64(p.Value)
}
func (p *Dtype) Uint64() (uint64, error) {
	return dtype.Uint64(p.Value)
}
func (p *Dtype) Float64() (float64, error) {
	return dtype.Float64(p.Value)
}

// Get get key from a map[string]interface{}
// p.Get("users.1.name") is short for p.Get("user", "1", "name")
// @warn p.Get("user", "1", "name") is diffirent with p.Get("user", 1, "name")

func (p *Dtype) Get(keys ...interface{}) (*Dtype, error) {
	v, err := dtype.NewMap(p.Value).Get(keys[0], keys[1:]...)
	return NewDtype(v), err
}
