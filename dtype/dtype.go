package dtype

import (
	"database/sql"
)

// Dbyte String('A') will returns "97". So you have to use String(Dbyte('A')) to return "A"
type Dbyte byte

type Dtype struct {
	Value interface{}
}

func New(data interface{}) *Dtype {
	return &Dtype{
		Value: data,
	}
}

// Get get key from a map[string]interface{}
// p.Get("users.1.name") is short for p.Get("user", "1", "name")
// @warn p.Get("user", "1", "name") is diffirent with p.Get("user", 1, "name")

func (p *Dtype) Get(keys ...interface{}) (*Dtype, error) {
	v, err := NewMap(p.Value).Get(keys[0], keys[1:]...)
	return New(v), err
}

func (p *Dtype) SqlNullString() sql.NullString {
	return sql.NullString{String: p.String(), Valid: p.NotEmpty()}
}

func (p *Dtype) SqlNullInt64() sql.NullInt64 {
	v, _ := p.Int64()
	return sql.NullInt64{Int64: v, Valid: p.NotEmpty()}
}

func (p *Dtype) SqlNullFloat64() sql.NullFloat64 {
	v, _ := p.Float64()
	return sql.NullFloat64{Float64: v, Valid: p.NotEmpty()}
}

func (p *Dtype) IsEmpty() bool {
	return IsEmpty(p.Value)
}
func (p *Dtype) NotEmpty() bool {
	return NotEmpty(p.Value)
}
func (p *Dtype) Interface() interface{} {
	return p.Value
}

func (p *Dtype) Bool() (bool, error) {
	return Bool(p.Value)
}

func (p *Dtype) DefaultBool(defaultValue bool) bool {
	v, err := p.Bool()
	if err != nil {
		return defaultValue
	}
	return v
}

func (p *Dtype) Slice() ([]interface{}, error) {
	return Slice(p.Value)
}

func (p *Dtype) DefaultSlice(defaultValue []interface{}) []interface{} {
	v, err := p.Slice()
	if err != nil {
		return defaultValue
	}
	return v
}

func (p *Dtype) String() string {
	return String(p.Value)
}
func (p *Dtype) DefaultString(defaultValue string) string {
	v := p.String()
	if v == "" {
		return defaultValue
	}
	return v
}

func (p *Dtype) Bytes() []byte {
	return Bytes(p.Value)
}

func (p *Dtype) DefaultBytes(defaultValue []byte) []byte {
	v := p.Bytes()
	if len(v) == 0 {
		return defaultValue
	}
	return v
}

func (p *Dtype) Int8() (int8, error) {
	return Int8(p.Value)
}

func (p *Dtype) DefaultInt8(defaultValue int8) int8 {
	v, err := p.Int8()
	if err != nil {
		return defaultValue
	}
	return v
}

func (p *Dtype) Int16() (int16, error) {
	return Int16(p.Value)
}
func (p *Dtype) DefaultInt16(defaultValue int16) int16 {
	v, err := p.Int16()
	if err != nil {
		return defaultValue
	}
	return v
}

func (p *Dtype) Int32() (int32, error) {
	return Int32(p.Value)
}

func (p *Dtype) DefaultInt32(defaultValue int32) int32 {
	v, err := p.Int32()
	if err != nil {
		return defaultValue
	}
	return v
}

func (p *Dtype) Int() (int, error) {
	return Int(p.Value)
}

func (p *Dtype) DefaultInt(defaultValue int) int {
	v, err := p.Int()
	if err != nil {
		return defaultValue
	}
	return v
}

func (p *Dtype) Int64() (int64, error) {
	return Int64(p.Value)
}

func (p *Dtype) DefaultInt64(defaultValue int64) int64 {
	v, err := p.Int64()
	if err != nil {
		return defaultValue
	}
	return v
}

func (p *Dtype) Uint8() (uint8, error) {
	return Uint8(p.Value)
}

func (p *Dtype) DefaultUint8(defaultValue uint8) uint8 {
	v, err := p.Uint8()
	if err != nil {
		return defaultValue
	}
	return v
}

func (p *Dtype) Uint16() (uint16, error) {
	return Uint16(p.Value)
}

func (p *Dtype) DefaultUint16(defaultValue uint16) uint16 {
	v, err := p.Uint16()
	if err != nil {
		return defaultValue
	}
	return v
}
func (p *Dtype) Uint32() (uint32, error) {
	return Uint32(p.Value)
}

func (p *Dtype) DefaultUint32(defaultValue uint32) uint32 {
	v, err := p.Uint32()
	if err != nil {
		return defaultValue
	}
	return v
}
func (p *Dtype) Uint() (uint, error) {
	return Uint(p.Value)
}

func (p *Dtype) DefaultUint(defaultValue uint) uint {
	v, err := p.Uint()
	if err != nil {
		return defaultValue
	}
	return v
}

func (p *Dtype) Uint64() (uint64, error) {
	return Uint64(p.Value)
}

func (p *Dtype) DefaultUint64(defaultValue uint64) uint64 {
	v, err := p.Uint64()
	if err != nil {
		return defaultValue
	}
	return v
}

func (p *Dtype) Float32() (float32, error) {
	return Float32(p.Value)
}

func (p *Dtype) DefaultFloat32(defaultValue float32) float32 {
	v, err := p.Float32()
	if err != nil {
		return defaultValue
	}
	return v
}

func (p *Dtype) Float64() (float64, error) {
	return Float64(p.Value)
}

func (p *Dtype) DefaultFloat64(defaultValue float64) float64 {
	v, err := p.Float64()
	if err != nil {
		return defaultValue
	}
	return v
}
