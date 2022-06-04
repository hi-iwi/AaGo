package atype

import (
	"fmt"
	"reflect"
	"strings"
)

// structs.Map(rsp) 可以将struct 转为 map[string]interface{}
type Map struct {
	Value interface{}
}

func NewMap(v interface{}) Map {
	return Map{
		Value: v,
	}
}
func splitInterfaces(key interface{}) []interface{} {
	switch k := key.(type) {
	case string:
		arr := make([]string, 0)
		arr = append(arr, strings.Split(k, ".")...)
		n := make([]interface{}, len(arr))
		for i, a := range arr {
			n[i] = a
		}
		return n
	}
	n := []interface{}{key}
	return n
}

// Get get key from a map[string]interface{}
// p.Get("users.1.name") is short for p.Get("user", "1", "name")
// @warn p.Get("user", "1", "name") is diffirent with p.Get("user", 1, "name")

func (m Map) Get(key interface{}, keys ...interface{}) (interface{}, error) {
	value := m.Value
	var val map[interface{}]interface{}
	var ok bool

	var rvalue reflect.Value
	keys = append([]interface{}{key}, keys...)
	if len(keys) == 1 {
		keys = splitInterfaces(keys[0])
	}

	for i, k := range keys {
		if val, ok = value.(map[interface{}]interface{}); !ok {
			val = make(map[interface{}]interface{})

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
