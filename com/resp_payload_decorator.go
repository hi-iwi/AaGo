package com

import (
	"github.com/hi-iwi/AaGo/ae"
	"github.com/hi-iwi/AaGo/dtype"
	"reflect"
)

// ?_stringify=1  weak language, turn all fields into string
func (resp *RespStruct) decoratePayload(payload interface{}, tagname string) (interface{}, *ae.Error) {
	stringify, _ := resp.req.Query("_stringify", `^[01]$`, false)
	if stringify.DefaultBool(false) {
		return StringifyPayloadFields(payload, tagname)
	}
	return payload, nil
}

func StringifyPayloadFields(payload interface{}, tagname string) (interface{}, *ae.Error) {
	var e *ae.Error
	t := reflect.TypeOf(payload)
	v := reflect.ValueOf(payload)
	k := t.Kind()

	// 指针
	if k == reflect.Ptr || k == reflect.UnsafePointer {
		t = t.Elem()
		k = t.Kind()
		v = v.Elem()
	}

	if k == reflect.Slice || k == reflect.Array {
		// v 有可能是一个nil指针
		if v.Kind() == reflect.Invalid || v.Len() == 0 {
			return nil, nil
		}
		p := make([]interface{}, v.Len())
		for i := 0; i < v.Len(); i++ {
			p[i], e = StringifyPayloadFields(v.Index(i).Interface(), tagname)
			if e != nil {
				return nil, e
			}
		}
		return p, nil
	} else if k == reflect.Struct {
		// v 有可能是一个nil指针
		if v.Kind() == reflect.Invalid || t.NumField() == 0 {
			return nil, nil
		}

		p := make(map[string]interface{}, 0)
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			ks := f.Tag.Get(tagname)
			// 忽略json/xml 里面的  -
			if ks == "-" {
				continue
			}
			w, e := StringifyPayloadFields(v.FieldByName(f.Name).Interface(), tagname)
			if e != nil {
				return nil, e
			}

			//   struct in struct
			if ks == "" {
				m, ok := w.(map[string]interface{})
				if !ok {
					return nil, ae.NewErr("unsolved json stringify struct `%v`, maybe tag `json:` not defined", w)
				}
				for y, z := range m {
					p[y] = z
				}
			} else {
				p[ks] = w
			}

		}
		return p, nil
	} else if k == reflect.Map {
		// v 有可能是一个nil指针
		if v.Kind() == reflect.Invalid || len(v.MapKeys()) == 0 {
			return nil, nil
		}
		p := make(map[string]interface{}, v.Len())
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
				m, ok := w.(map[string]interface{})
				if !ok {
					return nil, ae.NewErr("unsolved json stringify map `%v`, maybe tag `json:` not defined", w)
				}
				for y, z := range m {
					p[y] = z
				}
			} else {
				p[ks] = w
			}
		}
		return p, nil
	}
	return dtype.String(payload), nil
}
