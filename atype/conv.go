package atype

import (
	"errors"
	"fmt"
	"strconv"
)

// String convert into string
// @warn byte is a built-in alias of uint8, String('A') returns "97"; String(Abyte('A')) returns "A"
func String(d interface{}) string {
	if d == nil {
		return ""
	}
	switch v := d.(type) {
	case Abyte: // Name(Abyte('A')) returns "A"
		return string([]byte{byte(v)})
	case bool:
		return strconv.FormatBool(v)
	case []byte:
		return string(v)
	case string:
		return v
	case Date:
		return string(v)
	case Datetime:
		return string(v)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case rune: // is a built-in alias of int32, @notice 'A' is a rune(65), is different with byte('A') (alias of uint8(65))
		return strconv.FormatInt(int64(v), 10)
	case int:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case byte: // is a built-in alias of uint8, Name('A') returns "97"
		return strconv.FormatUint(uint64(v), 10)
	case Booln:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case Uint24:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	}
	// 有些类型type vt uint  var a, b vt 这样就无法识别为 uint；所以尝试通过字符串方式转一下
	return fmt.Sprint(d)
}

func Bytes(d interface{}) []byte {
	return []byte(String(d))
}

func Bool(d interface{}) (bool, error) {
	if d == nil {
		return false, errors.New("nil to bool")
	}

	switch v := d.(type) {
	case bool:
		return v, nil
	case string:
		return strconv.ParseBool(v)
	case Booln:
		return v.Bool(), nil
	case int8:
		return v > 0, nil
	case uint8:
		return v > 0, nil
	case int:
		return v > 0, nil
	}
	// 有些类型type vt uint  var a, b vt 这样就无法识别为 uint；所以尝试通过字符串方式转一下
	return strconv.ParseBool(String(d))
}

func Slice(d interface{}) ([]interface{}, error) {
	if d != nil {
		if v, ok := d.([]interface{}); ok {
			return v, nil
		}
	}
	return nil, errors.New("cast type error")
}

func Int8(d interface{}) (int8, error) {
	v, err := BaseInt64(d, 8)
	return int8(v), err
}

func Int16(d interface{}) (int16, error) {
	v, err := BaseInt64(d, 16)
	return int16(v), err
}
func Int32(d interface{}) (int32, error) {
	v, err := BaseInt64(d, 32)
	return int32(v), err
}

func Int(d interface{}) (int, error) {
	v, err := BaseInt64(d, 32)
	return int(v), err
}
func Int64(d interface{}) (int64, error) {
	v, err := BaseInt64(d, 64)
	return v, err
}
func BaseInt64(d interface{}, bitSize int) (int64, error) {
	if d == nil {
		return 0, errors.New("nil to integer")
	}

	switch v := d.(type) {
	case Abyte:
		return int64(v), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		return strconv.ParseInt(v, 10, bitSize)
	case int8:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case rune:
		return int64(v), nil
	case int:
		return int64(v), nil
	case int64:
		return v, nil
	case byte:
		return int64(v), nil
	case Booln:
		return int64(v), nil
	case uint16:
		return int64(v), nil
	case Uint24:
		return int64(v), nil
	case uint32:
		return int64(v), nil
	case uint:
		return int64(v), nil
	case uint64:
		return int64(v), nil
	case float32:
		return int64(v), nil
	case float64:
		return int64(v), nil
	}
	// 有些类型type vt uint  var a, b vt 这样就无法识别为 uint；所以尝试通过字符串方式转一下
	return strconv.ParseInt(String(d), 10, bitSize)
}

func Uint8(d interface{}) (uint8, error) {
	v, err := BaseUint64(d, 8)
	return uint8(v), err
}

func Uint16(d interface{}) (uint16, error) {
	v, err := BaseUint64(d, 16)
	return uint16(v), err
}
func Uint24b(d interface{}) (Uint24, error) {
	v, err := BaseUint64(d, 24)
	return Uint24(v), err
}
func Uint32(d interface{}) (uint32, error) {
	r, err := BaseUint64(d, 32)
	return uint32(r), err
}

func Uint(d interface{}) (uint, error) {
	r, err := BaseUint64(d, 32)
	return uint(r), err
}
func Uint64(d interface{}) (uint64, error) {
	return BaseUint64(d, 64)
}
func BaseUint64(d interface{}, bitSize int) (uint64, error) {
	if d == nil {
		return 0, errors.New("nil to uint64")
	}
	switch v := d.(type) {
	case Abyte:
		return uint64(v), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		return strconv.ParseUint(v, 10, bitSize)
	case int8:
		return uint64(v), nil
	case int16:
		return uint64(v), nil
	case rune:
		return uint64(v), nil
	case int:
		return uint64(v), nil
	case int64:
		return uint64(v), nil
	case byte: // 等同uint8
		return uint64(v), nil
	case Booln:
		return uint64(v), nil
	case uint16:
		return uint64(v), nil
	case Uint24:
		return uint64(v), nil
	case uint32:
		return uint64(v), nil
	case uint:
		return uint64(v), nil
	case uint64:
		return v, nil
	case float32:
		return uint64(v), nil
	case float64:
		return uint64(v), nil
	}
	// 有些类型type vt uint  var a, b vt 这样就无法识别为 uint；所以尝试通过字符串方式转一下
	return strconv.ParseUint(String(d), 10, bitSize)
}

func Float64(d interface{}, bitSize int) (float64, error) {
	if d == nil {
		return 0.0, errors.New("nil to float64")
	}
	switch v := d.(type) {
	case Abyte:
		return float64(v), nil
	case bool:
		if v {
			return 1.0, nil
		}
		return 0.0, nil

	case string:
		return strconv.ParseFloat(v, bitSize)
	case int8:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case rune:
		return float64(v), nil
	case int:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case byte:
		return float64(v), nil
	case Booln:
		return float64(v), nil
	case uint16:
		return float64(v), nil
	case Uint24:
		return float64(v), nil
	case uint32:
		return float64(v), nil
	case uint:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	}
	// 有些类型type vt uint  var a, b vt 这样就无法识别为 uint；所以尝试通过字符串方式转一下
	return strconv.ParseFloat(String(d), bitSize)
}
func Float32(d interface{}) (float32, error) {
	f, err := Float64(d, 32)
	if err != nil {
		return 0.0, err
	}
	return float32(f), nil
}
func IsEmpty(d interface{}) bool {
	return !NotEmpty(d)
}

// NotEmpty check if a value is empty
// @warn NotEmpty(byte(0)) == false,  NotEmpty(byte('0')) == true
//       NotEmpty(0) == false, NotEmpty(-1) == true, NotEmpty(1) == true
func NotEmpty(d interface{}) bool {
	if d == nil {
		return false
	}
	switch v := d.(type) {
	case Abyte:
		return v != 0
	case bool:
		return v
	case string:
		return v != ""
	case int8:
		return v != 0 // 复数，不算 empty，只有0 才算empty
	case int16:
		return v != 0
	case rune:
		return v != 0
	case int:
		return v != 0
	case int64:
		return v != 0
	case byte:
		return v > 0
	case Booln:
		return v > 0
	case uint16:
		return v > 0
	case Uint24:
		return v > 0
	case uint32:
		return v > 0
	case uint:
		return v > 0
	case uint64:
		return v > 0
	case float32:
		return v != 0
	case float64:
		return v != 0
	}

	return String(d) != ""
}
