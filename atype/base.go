package atype

import (
	"reflect"
)

const MaxInt24 = 1<<23 - 1
const MinInt24 = -1 << 23
const MaxUint24 = 1<<24 - 1

const MaxInt8Len = 4    // -128 ~ 127
const MaxUint8Len = 3   // 0 ~ 256
const MaxInt16Len = 6   // -32768 ~ 32767
const MaxUint16Len = 5  // 0 ~ 65535
const MaxInt24Len = 8   // -8388608 ~ 8388607
const MaxUint24Len = 8  // 0 ~ 16777215
const MaxIntLen = 11    // -2147483648 ~ 2147483647
const MaxUintLen = 10   // 0 ~ 4294967295
const MaxInt64Len = 20  // -9223372036854775808 ~ 9223372036854775807
const MaxUint64Len = 20 // 0 ~ 18446744073709551615

const MaxDistriLen = MaxUint24Len

//Invalid Kind = iota
//Bool
//Int
//Int8
//Int16
//Int32
//Int64
//Uint
//Uint8
//Uint16
//Uint32
//Uint64
//Uintptr
//Float32
//Float64
//Complex64
//Complex128
//Array
//Chan
//Func
//Interface
//Map
//Ptr
//Slice
//String
//Struct
//UnsafePointer
// 获取原始类型  i 用指针
// @param i 必须为指针
// @return 除了 reflect.Ptr 外其他类型；包括 interface
func PrimitiveType(i interface{}) reflect.Kind {
	if i == nil {
		return reflect.Invalid // nil
	}
	k := reflect.TypeOf(i).Elem().Kind()
	if k == reflect.Invalid {
		return reflect.Invalid // nil
	}
	if k == reflect.Ptr {
		v := reflect.ValueOf(i).Elem()
		if !v.CanInterface() {
			return reflect.Invalid
		}
		return PrimitiveType(v.Interface())
	}
	if k == reflect.Interface {
		k = reflect.ValueOf(i).Kind()
		if k == reflect.Ptr {
			v := reflect.ValueOf(i).Elem()
			if !v.CanInterface() {
				return reflect.Invalid
			}
			return PrimitiveType(v.Interface())
		}
		return k
	}
	return k
}

// 可能为指针，或者其他
func PType(i interface{}) reflect.Kind {
	if i == nil {
		return reflect.UnsafePointer // nil
	}
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	k := t.Kind()
	if k == reflect.Invalid {
		return reflect.Invalid // nil
	}
	// 指针
	if k == reflect.Ptr {
		return PrimitiveType(i)
	}
	if k == reflect.Interface {
		k = reflect.ValueOf(i).Kind()
		if k == reflect.Ptr {
			v = reflect.ValueOf(i).Elem()
			if !v.CanInterface() {
				return reflect.Invalid
			}
			return PrimitiveType(v.Interface())
		}
		return k
	}
	return k
}
