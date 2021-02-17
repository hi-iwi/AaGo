package com

import (
	"github.com/hi-iwi/AaGo/dtype"
	"strconv"
	"strings"
)

type ReqProp struct {
	dtype.Dtype
	param string
}

func NewReqProp(param string, data interface{}) *ReqProp {
	var p ReqProp
	p.param = param
	p.Value = data
	return &p
}

func (p *ReqProp) Default(v interface{}) {
	if p.IsEmpty() {
		p.Value = v
	}
}
func UintsRegExp(set ...interface{}) string {
	elems := make([]string, len(set))
	for i, v := range set {
		elems[i] = strconv.FormatUint(v.(uint64), 10)
	}
	s := strings.Join(elems, "|")
	return "^(" + s + ")$"
}
func IntsRegExp(set ...interface{}) string {
	elems := make([]string, len(set))
	for i, v := range set {
		elems[i] = strconv.FormatInt(v.(int64), 10)
	}
	s := strings.Join(elems, "|")
	return "^(" + s + ")$"
}
func StringsRegExp(set ...interface{}) string {
	// @TODO 使用正则转义

	elems := make([]string, len(set))
	for i, v := range set {
		elems[i] = v.(string)
	}
	s := strings.Join(elems, "|")
	return "^(" + s + ")$"
}
