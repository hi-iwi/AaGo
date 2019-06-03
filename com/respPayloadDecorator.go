package com

import (
	"reflect"

	"github.com/luexu/dtype"

	"github.com/luexu/AaGo/ae"
)

// ?_stringify=1  weak language, turn all fields into string
func (resp RespStruct) decoratePayload(payload interface{}, tagname string) (interface{}, *ae.Error) {
	xweak, _ := resp.req.Query("_stringify", `^[01]$`, false)
	if xweak.IsEmpty() || payload == nil {
		return payload, nil
	}
	weak, _ := xweak.Bool()
	if weak {
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
			p[f.Tag.Get(tagname)], e = stringifyPayloadFields(v.FieldByName(f.Name).Interface(), tagname)
			if e != nil {
				return nil, e
			}
		}
		return p, nil
	} else if k == reflect.Map {
		p := make(map[string]interface{}, v.Len())
		for _, key := range v.MapKeys() {
			p[key.String()], e = stringifyPayloadFields(v.MapIndex(key).Interface(), tagname)
			if e != nil {
				return nil, e
			}
		}
		return p, nil
	}
	return dtype.String(payload), nil
}
