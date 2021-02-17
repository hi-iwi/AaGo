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
		case byte: // is a built-in alias of uint8, Name('A') returns "97"
			return strconv.Itoa(int(v))
		case []byte:
			return string(v)
		case string:
			return v
		case int:
			return strconv.Itoa(v)
		case int8:
			return strconv.Itoa(int(v))
		case int16:
			return strconv.Itoa(int(v))
		case rune: // is a built-in alias of int32, @notice 'A' is a rune(65), is different with byte('A') (alias of uint8(65))
			return strconv.Itoa(int(v))
		case int64:
			return strconv.FormatInt(v, 10)
		case uint:
			return strconv.FormatUint(uint64(v), 10)
		case uint16:
			return strconv.FormatUint(uint64(v), 10)
		case uint32:
			return strconv.FormatUint(uint64(v), 10)
		case uint64:
			return strconv.FormatUint(v, 10)
		case float32:
			return strconv.FormatFloat(float64(v), 'f', -1, 64)
		case float64:
			return strconv.FormatFloat(v, 'f', -1, 64)
		default:
			return fmt.Sprint(d)
		}
	}
	return ""
}

func Bytes(d interface{}, errs ...error) []byte {
	return []byte(String(d))
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
	if (len(errs) == 0 || errs[0] == nil) && d != nil {

		switch v := d.(type) {
		case bool:
			if v {
				return 1, nil
			}
			return 0, nil
		case string:
			x, err := strconv.Atoi(v)
			if err != nil {
				return 0, err
			}
			return int8(x), nil
		case int:
			return int8(v), nil
		case int8:
			return v, nil
		case int16:
			return int8(v), nil
		case int32:
			return int8(v), nil
		case int64:
			return int8(v), nil
		case float32:
			return int8(v), nil
		case float64:
			return int8(v), nil
		case byte:
			return int8(v), nil
		case Dbyte:
			return int8(v), nil
		default:
			x, err := strconv.Atoi(String(v))
			if err != nil {
				return 0, err
			}
			return int8(x), nil
		}
	}
	return 0, errors.New("cast type error " + reflect.TypeOf(d).Kind().String() + "--> int8")
}

func Int16(d interface{}, errs ...error) (int16, error) {
	if (len(errs) == 0 || errs[0] == nil) && d != nil {

		switch v := d.(type) {
		case bool:
			if v {
				return 1, nil
			}
			return 0, nil
		case string:
			x, err := strconv.Atoi(v)
			if err != nil {
				return 0, err
			}
			return int16(x), nil
		case int:
			return int16(v), nil
		case int8:
			return int16(v), nil
		case int16:
			return v, nil
		case int32:
			return int16(v), nil
		case int64:
			return int16(v), nil
		case float32:
			return int16(v), nil
		case float64:
			return int16(v), nil
		case byte:
			return int16(v), nil
		case Dbyte:
			return int16(v), nil
		default:
			x, err := strconv.Atoi(String(v))
			if err != nil {
				return 0, err
			}
			return int16(x), nil
		}
	}
	return 0, errors.New("cast type error " + reflect.TypeOf(d).Kind().String() + "--> int16")
}
func Int32(d interface{}, errs ...error) (int32, error) {
	v, err := Int(d, errs...)
	return int32(v), err
}

func Int(d interface{}, errs ...error) (int, error) {
	if (len(errs) == 0 || errs[0] == nil) && d != nil {

		switch v := d.(type) {
		case bool:
			if v {
				return 1, nil
			}
			return 0, nil
		case string:
			return strconv.Atoi(v)
		case int:
			return v, nil
		case int8:
			return int(v), nil
		case int16:
			return int(v), nil
		case int32:
			return int(v), nil
		case int64:
			return int(v), nil
		case float32:
			return int(v), nil
		case float64:
			return int(v), nil
		case byte:
			return int(v), nil
		case Dbyte:
			return int(v), nil
		default:
			return strconv.Atoi(String(v))
		}
	}
	return 0, errors.New("cast type error " + reflect.TypeOf(d).Kind().String() + "--> int")
}

func Int64(d interface{}, errs ...error) (int64, error) {
	if (len(errs) == 0 || errs[0] == nil) && d != nil {
		switch v := d.(type) {
		case bool:
			if v {
				return int64(1), nil
			}
			return int64(0), nil
		case byte:
			return int64(v), nil
		case string:
			return strconv.ParseInt(v, 10, 64)
		case int:
			return int64(v), nil
		case int8:
			return int64(v), nil
		case int16:
			return int64(v), nil
		case int32:
			return int64(v), nil
		case int64:
			return v, nil
		case float32:
			return int64(v), nil
		case float64:
			return int64(v), nil
		default:
			return strconv.ParseInt(String(v), 10, 64)
		}
	}
	return 0, errors.New("cast type error " + reflect.TypeOf(d).Kind().String() + "--> int64")
}

func Uint8(d interface{}, errs ...error) (uint8, error) {
	if (len(errs) == 0 || errs[0] == nil) && d != nil {
		switch v := d.(type) {
		case bool:
			if v {
				return uint8(1), nil
			}
			return uint8(0), nil
		case byte:
			return uint8(v), nil
		case Dbyte:
			return uint8(v), nil
		case string:
			val, err := strconv.ParseUint(v, 10, 64)
			return uint8(val), err
		case int:
			return uint8(v), nil
		case int8:
			return uint8(v), nil
		case int16:
			return uint8(v), nil
		case int32:
			return uint8(v), nil
		case int64:
			return uint8(v), nil
		case float32:
			return uint8(v), nil
		case float64:
			return uint8(v), nil
		default:
			val, err := strconv.ParseUint(String(v), 10, 64)
			return uint8(val), err
		}
	}
	return 0, errors.New("cast type error " + reflect.TypeOf(d).Kind().String() + "--> uint8")
}

func Uint16(d interface{}, errs ...error) (uint16, error) {
	if (len(errs) == 0 || errs[0] == nil) && d != nil {
		switch v := d.(type) {
		case bool:
			if v {
				return uint16(1), nil
			}
			return uint16(0), nil
		case byte:
			return uint16(v), nil
		case Dbyte:
			return uint16(v), nil
		case string:
			val, err := strconv.ParseUint(v, 10, 64)
			return uint16(val), err
		case int:
			return uint16(v), nil
		case int8:
			return uint16(v), nil
		case int16:
			return uint16(v), nil
		case int32:
			return uint16(v), nil
		case int64:
			return uint16(v), nil
		case float32:
			return uint16(v), nil
		case float64:
			return uint16(v), nil
		default:
			val, err := strconv.ParseUint(String(v), 10, 64)
			return uint16(val), err
		}
	}
	return 0, errors.New("cast type error " + reflect.TypeOf(d).Kind().String() + "--> uint16")
}
func Uint32(d interface{}, errs ...error) (uint32, error) {
	r, err := Uint(d, errs...)
	return uint32(r), err
}

func Uint(d interface{}, errs ...error) (uint, error) {
	if (len(errs) == 0 || errs[0] == nil) && d != nil {
		switch v := d.(type) {
		case bool:
			if v {
				return uint(1), nil
			}
			return uint(0), nil
		case byte:
			return uint(v), nil
		case Dbyte:
			return uint(v), nil
		case string:
			val, err := strconv.ParseUint(v, 10, 64)
			return uint(val), err
		case int:
			return uint(v), nil
		case int8:
			return uint(v), nil
		case int16:
			return uint(v), nil
		case int32:
			return uint(v), nil
		case int64:
			return uint(v), nil
		case float32:
			return uint(v), nil
		case float64:
			return uint(v), nil
		default:
			val, err := strconv.ParseUint(String(v), 10, 64)
			return uint(val), err
		}
	}
	return 0, errors.New("cast type error " + reflect.TypeOf(d).Kind().String() + "--> uint")
}

func Uint64(d interface{}, errs ...error) (uint64, error) {
	if (len(errs) == 0 || errs[0] == nil) && d != nil {
		switch v := d.(type) {
		case bool:
			if v {
				return uint64(1), nil
			}
			return uint64(0), nil
		case byte:
			return uint64(v), nil
		case Dbyte:
			return uint64(v), nil
		case string:
			return strconv.ParseUint(v, 10, 64)
		case int:
			return uint64(v), nil
		case int8:
			return uint64(v), nil
		case int16:
			return uint64(v), nil
		case int32:
			return uint64(v), nil
		case int64:
			return uint64(v), nil
		case float32:
			return uint64(v), nil
		case float64:
			return uint64(v), nil
		default:
			return strconv.ParseUint(String(v), 10, 64)
		}
	}
	return 0, errors.New("cast type error " + reflect.TypeOf(d).Kind().String() + "--> uint64")
}

func Float32(d interface{}, errs ...error) (float32, error) {
	v, err := Float64(d, errs...)
	return float32(v), err
}

func Float64(d interface{}, errs ...error) (float64, error) {
	if (len(errs) == 0 || errs[0] == nil) && d != nil {
		switch v := d.(type) {
		case bool:
			if v {
				return float64(1), nil
			}
			return float64(0), nil
		case byte:
			return float64(v), nil
		case Dbyte:
			return float64(v), nil
		case string:
			return strconv.ParseFloat(v, 64)
		case int:
			return float64(v), nil
		case int8:
			return float64(v), nil
		case int16:
			return float64(v), nil
		case int32:
			return float64(v), nil
		case int64:
			return float64(v), nil
		case float32:
			return float64(v), nil
		case float64:
			return v, nil
		default:
			return strconv.ParseFloat(String(v), 64)
		}
	}
	return float64(0), errors.New("cast type error " + reflect.TypeOf(d).Kind().String() + "--> Float64")
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
		case bool:
			return v
		case byte:
			return v != byte(0)
		case Dbyte:
			return v != Dbyte(0)
		case string:
			return v != ""
		case int:
			return v != 0
		case int8:
			return v != int8(0)
		case int16:
			return v != int16(0)
		case rune:
			return v != rune(0)
		case int64:
			return v != int64(0)
		case float32:
			return v != float32(0)
		case float64:
			return v != float64(0)
		default:
			return String(d) != ""
		}
	}
	return false
}
