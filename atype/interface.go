package atype

import "reflect"

/*
在Go语言中，一个any类型的变量包含了2个指针，一个指针指向值的在编译时确定的类型，另外一个指针指向实际的值。
*/
func IsNil(x any) bool {
	// @warn 断言和反射性能不是特别好，如果不得已再使用，控制使用有助于提升程序性能。
	return x == nil || (reflect.ValueOf(x).Kind() == reflect.Ptr && reflect.ValueOf(x).IsNil())
}

func XUint64s(v []uint64) []any {
	args := make([]any, len(v))
	for i, x := range v {
		args[i] = x
	}
	return args
}
func XUints(v []int64) []any {
	args := make([]any, len(v))
	for i, x := range v {
		args[i] = x
	}
	return args
}
func XInt64s(v []int64) []any {
	args := make([]any, len(v))
	for i, x := range v {
		args[i] = x
	}
	return args
}
func XInts(v []int64) []any {
	args := make([]any, len(v))
	for i, x := range v {
		args[i] = x
	}
	return args
}
