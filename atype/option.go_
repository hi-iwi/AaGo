package atype

import (
	"encoding/json"
	"errors"
	"reflect"
)

type Option struct {
	Key   uint        `json:"key"` // <option key value>text</option>
	Value interface{} `json:"value"`
	Pid   interface{} `json:"pid"`
	Text  string      `json:"text"`

	Inherit Booln  `json:"inherit"` // 如果为true，下级会增加一个选项
	Prefix  string `json:"prefix"`  // 选取显示的时候，在text前面加的字符    特殊情况：<f8 f803/>  select 组件会转义成  <i class="f8 f803"></i> 之类的
	Suffix  string `json:"suffix"`  // 选取显示时候，在text后面加的字符 特殊情况：<fi fi32/>  select 组件会转义成  <i class="fi fi32"></i> 之类的
	Comment string `json:"comment"`
	Virtual Booln  `json:"virtual"` // 是否虚拟，虚拟的就不显示
}

// 这个用到映射，尽量仅用于初始化阶段
// @param names map[interface{}]string   {value:text}
func ToOptions(names interface{}) ([]Option, error) {
	t := reflect.TypeOf(names)
	v := reflect.ValueOf(names)
	k := t.Kind()
	if k != reflect.Map {
		return nil, errors.New("bad atype.ToOptions args")
	}
	opts := make([]Option, 0, v.Len())
	for _, key := range v.MapKeys() {
		value := key.Interface()
		text := v.MapIndex(key).String()
		opts = append(opts, Option{
			Value: value,
			Text:  text,
		})
	}
	return opts, nil
}
func ToOptionsJson(names interface{}) ([]byte, error) {
	opts, err := ToOptions(names)
	if err != nil {
		return nil, err
	}
	return json.Marshal(opts)
}
