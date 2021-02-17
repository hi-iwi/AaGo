package com

import (
	"reflect"
	"strconv"

	"github.com/hi-iwi/AaGo/dtype"

	"github.com/hi-iwi/AaGo/ae"
)

// ?_stringify=1  weak language, turn all fields into string
func (resp *RespStruct) decoratePayload(payload interface{}, tagname string) (interface{}, *ae.Error) {
	stringify, _ := resp.req.Query("_stringify", `^[01]$`, false)
	s := stringify.DefaultBool(false)
	if !s {
		cookie, err := resp.req.Cookie("_stringify")
		if err == nil {
			s, _ = strconv.ParseBool(cookie.Value)
		}
	}
	if stringify.DefaultBool(false) {
		return stringifyPayloadFields(payload, tagname)
	}
	return payload, nil
}

func stringifyPayloadFields(payload interface{}, tagname string) (interface{}, *ae.Error) {
	var e *ae.Error
	t := reflect.TypeOf(payload)
	v := reflect.ValueOf(payload)
	k := t.Kind()

	if k == reflect.Slice || k == reflect.Array {
		p := make([]interface{}, v.Len())
		for i := 0; i < v.Len(); i++ {
			p[i], e = stringifyPayloadFields(v.Index(i).Interface(), tagname)
			if e != nil {
				return nil, e
			}
		}
		return p, nil
	} else if k == reflect.Struct {
		p := make(map[string]interface{}, 0)
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			ks := f.Tag.Get(tagname)
			// 忽略json/xml 里面的  -
			if (tagname == "json" || tagname == "xml") && ks == "-" {
				continue
			}
			p[ks], e = stringifyPayloadFields(v.FieldByName(f.Name).Interface(), tagname)
			if e != nil {
				return nil, e
			}
		}
		return p, nil
	} else if k == reflect.Map {
		p := make(map[string]interface{}, v.Len())
		for _, key := range v.MapKeys() {
			ks := key.String()
			// 忽略json/xml 里面的  -
			if (tagname == "json" || tagname == "xml") && ks == "-" {
				continue
			}
			p[ks], e = stringifyPayloadFields(v.MapIndex(key).Interface(), tagname)
			if e != nil {
				return nil, e
			}
		}
		return p, nil
	}
	return dtype.String(payload), nil
}
