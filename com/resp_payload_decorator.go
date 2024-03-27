package com

import (
	"github.com/hi-iwi/AaGo/ae"
	"reflect"
	"strconv"
)

// ?_stringify=1  weak language, turn int64/uint64 fields into string
func (resp *RespStruct) decoratePayload(payload any, tagname string) (any, *ae.Error) {
	stringify, _ := resp.req.QueryBool(ParamStringify)
	if stringify {
		return StringifyPayloadFields(payload, tagname)
	}
	return payload, nil
}
func stringifySlice(v reflect.Value, tagname string) (any, *ae.Error) {
	if v.Len() == 0 {
		return nil, nil
	}
	p := make([]any, v.Len())
	var e *ae.Error
	for i := 0; i < v.Len(); i++ {
		if !v.Index(i).CanInterface() {
			return nil, nil
		}
		p[i], e = StringifyPayloadFields(v.Index(i).Interface(), tagname)
		if e != nil {
			return nil, e
		}
	}
	return p, nil
}
func stringifyStruct(t reflect.Type, v reflect.Value, tagname string) (any, *ae.Error) {
	// v 有可能是一个nil指针
	if t.NumField() == 0 {
		return nil, nil
	}

	p := make(map[string]any, 0)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		ks := f.Tag.Get(tagname)
		// 忽略json/xml 里面的  -
		if ks == "-" {
			continue
		}
		v1 := v.FieldByName(f.Name)
		if !v1.CanInterface() {
			continue
		}
		w, e := StringifyPayloadFields(v1.Interface(), tagname)
		if e != nil {
			return nil, e
		}

		//   struct in struct
		if ks == "" {
			m, ok := w.(map[string]any)
			if !ok {
				//p[t.Name()] = w  // 忽略 json/xml tag不存在或为空
				continue
			} else {
				for y, z := range m {
					p[y], _ = StringifyPayloadFields(z, tagname)
				}
			}

		} else {
			p[ks] = w
		}

	}
	return p, nil
}
func stringifyMap(t reflect.Type, v reflect.Value, tagname string) (any, *ae.Error) {
	// v 有可能是一个nil指针
	if len(v.MapKeys()) == 0 {
		return nil, nil
	}
	p := make(map[string]any, v.Len())
	for _, key := range v.MapKeys() {
		ks := key.String()
		// 忽略json/xml 里面的  -
		if ks == "-" {
			continue
		}
		w, e := StringifyPayloadFields(v.MapIndex(key).Interface(), tagname)
		if e != nil {
			return nil, e
		}

		//   struct in struct
		if ks == "" {
			m, ok := w.(map[string]any)
			if !ok {
				p[t.Name()] = w
			} else {
				for y, z := range m {
					p[y], _ = StringifyPayloadFields(z, tagname)
				}
			}
		} else {
			p[ks] = w
		}
	}
	return p, nil
}

// 2.0 版本，仅针对初级原始类型为 int64/uint64 字段转为 string
func StringifyPayloadFields(payload any, tagname string) (any, *ae.Error) {
	if payload == nil {
		return nil, nil
	}

	t := reflect.TypeOf(payload)
	v := reflect.ValueOf(payload)
	k := v.Kind()
	// 指针
	if k == reflect.Ptr {
		v = v.Elem() // 必须放t=v.Type前面
		t = t.Elem() // 必须用 t.Elem()，不能用 v.Type()
		k = v.Kind()
	}
	if k == reflect.Interface {
		k = v.Kind()
		if k == reflect.Ptr {
			v = v.Elem() // 必须放t=v.Type前面
			t = t.Elem() // 必须用 t.Elem()，不能用 v.Type()
			k = v.Kind()
		}
	}
	if k == reflect.Invalid {
		return nil, nil
	}
	switch k {
	case reflect.Invalid:
		return nil, nil // v 有可能是一个nil指针
	case reflect.Slice, reflect.Array:
		return stringifySlice(v, tagname)
	case reflect.Struct:
		return stringifyStruct(t, v, tagname)
	case reflect.Map:
		return stringifyMap(t, v, tagname)
	case reflect.Int64: // 也有可能是 int64 变体，所以不能用 x.(int64) 转换，应该用强制类型转换
		return strconv.FormatInt(v.Int(), 10), nil
	case reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10), nil
	}
	return payload, nil
}

// 1.0 版本，针对所有字段全部替换
//func StringifyPayloadFields(payload any, tagname string) (any, *ae.Error) {
//	if payload == nil {
//		return nil, nil
//	}
//	var e *ae.Error
//	t := reflect.TypeOf(payload)
//	v := reflect.ValueOf(payload)
//	k := t.Kind()
//
//	// 指针
//	if k == reflect.Ptr || k == reflect.UnsafePointer {
//		v = v.Elem()
//		t = t.Elem()
//		k = t.Kind()
//	}
//	if k == reflect.Invalid || v.Kind() == reflect.Invalid {
//		return nil, nil
//	}
//	if k == reflect.Slice || k == reflect.Array {
//		// v 有可能是一个nil指针
//		if v.Len() == 0 {
//			return nil, nil
//		}
//		p := make([]any, v.Len())
//		for i := 0; i < v.Len(); i++ {
//			p[i], e = StringifyPayloadFields(v.Index(i).Interface(), tagname)
//			if e != nil {
//				return nil, e
//			}
//		}
//		return p, nil
//	} else if k == reflect.Struct {
//		// v 有可能是一个nil指针
//		if t.NumField() == 0 {
//			return nil, nil
//		}
//
//		p := make(map[string]any, 0)
//		for i := 0; i < t.NumField(); i++ {
//			f := t.Field(i)
//			ks := f.Tag.Get(tagname)
//			// 忽略json/xml 里面的  -
//			if ks == "-" {
//				continue
//			}
//			w, e := StringifyPayloadFields(v.FieldByName(f.Name).Interface(), tagname)
//			if e != nil {
//				return nil, e
//			}
//
//			//   struct in struct
//			if ks == "" {
//				m, ok := w.(map[string]any)
//				if !ok {
//					return nil, ae.NewErr("unsolved json stringify struct `%v`, maybe tag `json:` not defined", w)
//				}
//				for y, z := range m {
//					p[y] = z
//				}
//			} else {
//				p[ks] = w
//			}
//
//		}
//		return p, nil
//	} else if k == reflect.Map {
//		// v 有可能是一个nil指针
//		if v.Kind() == reflect.Invalid || len(v.MapKeys()) == 0 {
//			return nil, nil
//		}
//		p := make(map[string]any, v.Len())
//		for _, key := range v.MapKeys() {
//			ks := key.String()
//			// 忽略json/xml 里面的  -
//			if ks == "-" {
//				continue
//			}
//			w, e := StringifyPayloadFields(v.MapIndex(key).Interface(), tagname)
//			if e != nil {
//				return nil, e
//			}
//
//			//   struct in struct
//			if ks == "" {
//				m, ok := w.(map[string]any)
//				if !ok {
//					return nil, ae.NewErr("unsolved json stringify map `%v`, maybe tag `json:` not defined", w)
//				}
//				for y, z := range m {
//					p[y] = z
//				}
//			} else {
//				p[ks] = w
//			}
//		}
//		return p, nil
//	}
//	return atype.String(payload), nil
//}
