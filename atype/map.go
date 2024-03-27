package atype

import (
	"fmt"
	"reflect"
	"strings"
)

// structs.Map(rsp) 可以将struct 转为 map[string]any
type Map struct {
	Value any
}

func NewMap(v any) Map {
	return Map{
		Value: v,
	}
}
func splitInterfaces(key any) []any {
	switch k := key.(type) {
	case string:
		arr := make([]string, 0)
		arr = append(arr, strings.Split(k, ".")...)
		n := make([]any, len(arr))
		for i, a := range arr {
			n[i] = a
		}
		return n
	}
	n := []any{key}
	return n
}

// Get get key from a map[string]any
// p.Get("users.1.name") is short for p.Get("user", "1", "name")
// @warn p.Get("user", "1", "name") is diffirent with p.Get("user", 1, "name")

func (m Map) Get(key any, keys ...any) (any, error) {
	value := m.Value
	var val map[any]any
	var ok bool

	var rvalue reflect.Value
	keys = append([]any{key}, keys...)
	if len(keys) == 1 {
		keys = splitInterfaces(keys[0])
	}

	for i, k := range keys {
		if val, ok = value.(map[any]any); !ok {
			val = make(map[any]any)

			rvalue = reflect.ValueOf(value)
			switch rvalue.Kind() {
			case reflect.Map:
				for _, v := range rvalue.MapKeys() {
					val[v.Interface()] = rvalue.MapIndex(v).Interface()
				}
				ok = true
			}
		}

		if ok {
			if value, ok = val[k]; ok {
				if len(keys)-1 == i {
					return value, nil
				}
			}
		}
	}
	s := ""
	for _, key := range keys {
		s += "." + String(key)
	}
	return nil, fmt.Errorf("map atype field `%s` not found", s[1:])
}
