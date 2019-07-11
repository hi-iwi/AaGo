package com

import (
	"reflect"
	"strings"

	"github.com/luexu/AaGo/ae"
)

// @TODO
// ?_field=time,service,connections:[name,scheme],server_id,test:{a,b,c}
func (resp RespStruct) filterPayload(a interface{}, tagname string) (interface{}, *ae.Error) {
	xm, _ := resp.req.Query("_field", false)
	if xm.IsEmpty() {
		return a, nil
	}
	m := xm.String()
	if m[0] == '[' && m[len(m)-1] == ']' {
		return filterPayloadArray(a, tagname, strings.Split(m[1:len(m)-1], ",")...)
	}
	return filterPayloadMap(a, tagname, strings.Split(m, ",")...)
}

func filterPayloadMap(u interface{}, tagname string, tags ...string) (map[string]interface{}, *ae.Error) {
	var found bool
	ret := make(map[string]interface{}, 0)
	t := reflect.TypeOf(u)
	if t.Kind() == reflect.Map {
		for _, tag := range tags {
			found = false
			iter := reflect.ValueOf(u).MapRange()
			for iter.Next() {
				if iter.Key().String() == tag {
					found = true
					ret[tag] = iter.Value().Interface()
				}
			}
			if !found {
				ret[tag] = nil
			}
		}
		return ret, nil
	}

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
	return ret, nil
}

func filterPayloadArray(w interface{}, tagname string, tags ...string) (ret []map[string]interface{}, e *ae.Error) {
	t := reflect.TypeOf(w).Kind()
	if t != reflect.Slice && t != reflect.Array {
		return nil, ae.NewError(401, "invalid `_field`, not an array")
	}
	v := reflect.ValueOf(w)
	ret = make([]map[string]interface{}, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i], e = filterPayloadMap(v.Index(i).Interface(), tagname, tags...)
		if e != nil {
			return nil, e
		}
	}

	return ret, nil
}
