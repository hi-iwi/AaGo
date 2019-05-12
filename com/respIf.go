package com

import (
	"reflect"
	"strings"
)

type RespContentDTO struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Payload interface{} `json:"data"`
}

// @TODO
// ?_field=time,service,connections:[name,scheme],server_id,test:{a,b,c}
func (resp RespStruct) handlePayload(a interface{}, tagname string) interface{} {
	xm, _ := resp.req.Query("_field", false)
	if xm.IsEmpty() {
		return a
	}
	m := xm.String()
	if m[0] == '[' && m[len(m)-1] == ']' {
		return resp.handlePayloadArray(a, tagname, strings.Split(m[1:len(m)-1], ",")...)
	}
	return resp.handlePayloadMap(a, tagname, strings.Split(m, ",")...)
}

func (resp RespStruct) handlePayloadMap(u interface{}, tagname string, tags ...string) map[string]interface{} {
	ret := make(map[string]interface{}, 0)
	t := reflect.TypeOf(u)
	var found bool
	for _, tag := range tags {
		found = false
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			al := f.Tag.Get(tagname)
			if al == tag {
				found = true
				ret[tag] = reflect.ValueOf(u).FieldByName(f.Name).Interface()
			}
		}
		if !found {
			ret[tag] = nil
		}
	}
	return ret
}

func (resp RespStruct) handlePayloadArray(w interface{}, tagname string, tags ...string) []map[string]interface{} {
	t := reflect.TypeOf(w).Kind()
	if t != reflect.Slice && t != reflect.Array {
		return nil
	}
	v := reflect.ValueOf(w)
	ret := make([]map[string]interface{}, v.Len())

	for i := 0; i < v.Len(); i++ {
		ret[i] = resp.handlePayloadMap(v.Index(i).Interface(), tagname, tags...)
	}

	return ret
}
