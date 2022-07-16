package atype

import (
	"database/sql"
)

// Abyte String('A') will returns "97". So you have to use String(Abyte('A')) to return "A"
type Abyte byte

// fastjson 命名就是 autotype
type Atype struct {
	raw interface{}
}

func New(data interface{}) *Atype {
	return &Atype{
		raw: data,
	}
}

func (p *Atype) Raw() interface{} {
	return p.raw
}
func (p *Atype) Reload(v interface{}) {
	p.raw = v
}

// Get get key from a map[string]interface{}
// p.Get("users.1.name") is short for p.Get("user", "1", "name")
// @warn p.Get("user", "1", "name") is diffirent with p.Get("user", 1, "name")

func (p *Atype) Get(keys ...interface{}) (*Atype, error) {
	v, err := NewMap(p.raw).Get(keys[0], keys[1:]...)
	return New(v), err
}

func (p *Atype) SqlNullString() sql.NullString {
	return sql.NullString{String: p.String(), Valid: p.NotEmpty()}
}

func (p *Atype) SqlNullInt64() sql.NullInt64 {
	v, _ := p.Int64()
	return sql.NullInt64{Int64: v, Valid: p.NotEmpty()}
}

func (p *Atype) SqlNullFloat64() sql.NullFloat64 {
	v, _ := p.Float64()
	return sql.NullFloat64{Float64: v, Valid: p.NotEmpty()}
}

func (p *Atype) IsEmpty() bool {
	return IsEmpty(p.raw)
}
func (p *Atype) NotEmpty() bool {
	return NotEmpty(p.raw)
}

func (p *Atype) Bool() (bool, error) {
	return Bool(p.raw)
}

func (p *Atype) DefaultBool(defaultValue bool) bool {
	v, err := p.Bool()
	if err != nil {
		return defaultValue
	}
	return v
}

func (p *Atype) Slice() ([]interface{}, error) {
	return Slice(p.raw)
}

func (p *Atype) DefaultSlice(defaultValue []interface{}) []interface{} {
	v, err := p.Slice()
	if err != nil {
		return defaultValue
	}
	return v
}

func (p *Atype) String() string {
	return String(p.raw)
}
func (p *Atype) DefaultString(defaultValue string) string {
	v := p.String()
	if v == "" {
		return defaultValue
	}
	return v
}

func (p *Atype) Bytes() []byte {
	return Bytes(p.raw)
}

func (p *Atype) DefaultBytes(defaultValue []byte) []byte {
	v := p.Bytes()
	if len(v) == 0 {
		return defaultValue
	}
	return v
}

func (p *Atype) Int8() (int8, error) {
	return Int8(p.raw)
}

func (p *Atype) DefaultInt8(defaultValue int8) int8 {
	v, err := p.Int8()
	if err != nil {
		return defaultValue
	}
	return v
}

func (p *Atype) Int16() (int16, error) {
	return Int16(p.raw)
}
func (p *Atype) DefaultInt16(defaultValue int16) int16 {
	v, err := p.Int16()
	if err != nil {
		return defaultValue
	}
	return v
}

func (p *Atype) Int32() (int32, error) {
	return Int32(p.raw)
}

func (p *Atype) DefaultInt32(defaultValue int32) int32 {
	v, err := p.Int32()
	if err != nil {
		return defaultValue
	}
	return v
}

func (p *Atype) Int() (int, error) {
	return Int(p.raw)
}

func (p *Atype) DefaultInt(defaultValue int) int {
	v, err := p.Int()
	if err != nil {
		return defaultValue
	}
	return v
}

func (p *Atype) Int64() (int64, error) {
	return Int64(p.raw)
}

func (p *Atype) DefaultInt64(defaultValue int64) int64 {
	v, err := p.Int64()
	if err != nil {
		return defaultValue
	}
	return v
}

func (p *Atype) Uint8() (uint8, error) {
	return Uint8(p.raw)
}

func (p *Atype) DefaultUint8(defaultValue uint8) uint8 {
	v, err := p.Uint8()
	if err != nil {
		return defaultValue
	}
	return v
}

func (p *Atype) Uint16() (uint16, error) {
	return Uint16(p.raw)
}

func (p *Atype) DefaultUint16(defaultValue uint16) uint16 {
	v, err := p.Uint16()
	if err != nil {
		return defaultValue
	}
	return v
}
func (p *Atype) Uint24() (Uint24, error) {
	return Uint24b(p.raw)
}

func (p *Atype) DefaultUint24(defaultValue Uint24) Uint24 {
	v, err := p.Uint24()
	if err != nil {
		return defaultValue
	}
	return v
}
func (p *Atype) Uint32() (uint32, error) {
	return Uint32(p.raw)
}

func (p *Atype) DefaultUint32(defaultValue uint32) uint32 {
	v, err := p.Uint32()
	if err != nil {
		return defaultValue
	}
	return v
}
func (p *Atype) Uint() (uint, error) {
	return Uint(p.raw)
}

func (p *Atype) DefaultUint(defaultValue uint) uint {
	v, err := p.Uint()
	if err != nil {
		return defaultValue
	}
	return v
}

func (p *Atype) Uint64() (uint64, error) {
	return Uint64(p.raw)
}

func (p *Atype) DefaultUint64(defaultValue uint64) uint64 {
	v, err := p.Uint64()
	if err != nil {
		return defaultValue
	}
	return v
}

func (p *Atype) Float32() (float32, error) {
	return Float32(p.raw)
}

func (p *Atype) DefaultFloat32(defaultValue float32) float32 {
	v, err := p.Float32()
	if err != nil {
		return defaultValue
	}
	return v
}

func (p *Atype) Float64() (float64, error) {
	return Float64(p.raw)
}

func (p *Atype) DefaultFloat64(defaultValue float64) float64 {
	v, err := p.Float64()
	if err != nil {
		return defaultValue
	}
	return v
}
