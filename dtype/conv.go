package dtype

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// String convert into string
// @warn byte is a built-in alias of uint8, String('A') returns "97"; String(Dbyte('A')) returns "A"
func String(d interface{}, errs ...error) string {
	if (len(errs) == 0 || errs[0] == nil) && d != nil {
		switch v := d.(type) {
		case Dbyte: // Name(Dbyte('A')) returns "A"
			return string([]byte{byte(v)})
		case bool:
			return strconv.FormatBool(v)
		case []byte:
			return string(v)
		case string:
			return v
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
		case uint16:
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
		default:
			return fmt.Sprint(d)
		}
	}
	return ""
}

func Bytes(d interface{}, errs ...error) []byte {
	return []byte(String(d, errs...))
}

func Bool(d interface{}, errs ...error) (bool, error) {
	if (len(errs) == 0 || errs[0] == nil) && d != nil {
		switch v := d.(type) {
		case bool:
			return v, nil
		default:
			return strconv.ParseBool(String(v))
		}
	}
	return false, errors.New("cast type error " + reflect.TypeOf(d).Kind().String() + "--> bool")
}

func Slice(d interface{}, errs ...error) ([]interface{}, error) {
	if (len(errs) == 0 || errs[0] == nil) && d != nil {
		if v, ok := d.([]interface{}); ok {
			return v, nil
		}
	}
	return nil, errors.New("cast type error")
}

func Int8(d interface{}, errs ...error) (int8, error) {
	v, err := BaseInt64(d, 8, errs...)
	return int8(v), err
}

func Int16(d interface{}, errs ...error) (int16, error) {
	v, err := BaseInt64(d, 16, errs...)
	return int16(v), err
}
func Int32(d interface{}, errs ...error) (int32, error) {
	v, err := BaseInt64(d, 32, errs...)
	return int32(v), err
}

func Int(d interface{}, errs ...error) (int, error) {
	v, err := BaseInt64(d, 32, errs...)
	return int(v), err
}
func Int64(d interface{}, errs ...error) (int64, error) {
	v, err := BaseInt64(d, 64, errs...)
	return v, err
}
func BaseInt64(d interface{}, bitSize int, errs ...error) (int64, error) {
	if (len(errs) == 0 || errs[0] == nil) && d != nil {
		switch v := d.(type) {
		case Dbyte:
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
		case uint16:
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
		default:
			return strconv.ParseInt(String(v), 10, bitSize)
		}
	}
	return 0, errors.New("cast type error " + reflect.TypeOf(d).Kind().String() + "--> int64")
}

func Uint8(d interface{}, errs ...error) (uint8, error) {
	v, err := BaseUint64(d, 8, errs...)
	return uint8(v), err
}

func Uint16(d interface{}, errs ...error) (uint16, error) {
	v, err := BaseUint64(d, 16, errs...)
	return uint16(v), err
}
func Uint32(d interface{}, errs ...error) (uint32, error) {
	r, err := BaseUint64(d, 32, errs...)
	return uint32(r), err
}

func Uint(d interface{}, errs ...error) (uint, error) {
	r, err := BaseUint64(d, 32, errs...)
	return uint(r), err
}
func Uint64(d interface{}, errs ...error) (uint64, error) {
	r, err := BaseUint64(d, 64, errs...)
	return r, err
}
func BaseUint64(d interface{}, bitSize int, errs ...error) (uint64, error) {
	if (len(errs) == 0 || errs[0] == nil) && d != nil {
		switch v := d.(type) {
		case Dbyte:
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
		case uint16:
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
		default:
			return strconv.ParseUint(String(v), 10, bitSize)
		}
	}
	return 0, errors.New("cast type error " + reflect.TypeOf(d).Kind().String() + "--> uint64")
}

// 由于精度不一样，就不要用 float32(float64) 了
func Float32(d interface{}, errs ...error) (float32, error) {
	if (len(errs) == 0 || errs[0] == nil) && d != nil {
		switch v := d.(type) {
		case Dbyte:
			return float32(v), nil
		case bool:
			if v {
				return 1.0, nil
			}
			return 0.0, nil
		case string:
			val, err := strconv.ParseFloat(v, 32)
			return float32(val), err
		case int8:
			return float32(v), nil
		case int16:
			return float32(v), nil
		case rune:
			return float32(v), nil
		case int:
			return float32(v), nil
		case int64:
			return float32(v), nil
		case byte:
			return float32(v), nil
		case uint16:
			return float32(v), nil
		case uint32:
			return float32(v), nil
		case uint:
			return float32(v), nil
		case uint64:
			return float32(v), nil
		case float32:
			return v, nil
		case float64:
			return float32(v), nil
		default:
			val, err := strconv.ParseFloat(String(v), 32)
			return float32(val), err
		}
	}
	return 0.0, errors.New("cast type error " + reflect.TypeOf(d).Kind().String() + "--> Float32")
}

func Float64(d interface{}, errs ...error) (float64, error) {
	if (len(errs) == 0 || errs[0] == nil) && d != nil {
		switch v := d.(type) {
		case Dbyte:
			return float64(v), nil
		case bool:
			if v {
				return 1.0, nil
			}
			return 0.0, nil

		case string:
			return strconv.ParseFloat(v, 64)
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
		case uint16:
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
		default:
			return strconv.ParseFloat(String(v), 64)
		}
	}
	return 0.0, errors.New("cast type error " + reflect.TypeOf(d).Kind().String() + "--> Float64")
}

func IsEmpty(d interface{}, errs ...error) bool {
	return !NotEmpty(d, errs...)
}

// NotEmpty check if a value is empty
// @warn NotEmpty(byte(0)) == false,  NotEmpty(byte('0')) == true
//       NotEmpty(0) == false, NotEmpty(-1) == true, NotEmpty(1) == true
func NotEmpty(d interface{}, errs ...error) bool {
	if (len(errs) == 0 || errs[0] == nil) && d != nil {
		switch v := d.(type) {
		case Dbyte:
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
		case uint16:
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
		default:
			return String(d) != ""
		}
	}
	return false
}
