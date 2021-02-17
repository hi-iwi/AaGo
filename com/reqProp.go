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
		w, err := dtype.Uint64(v)
		if err != nil {
			continue
		}
		elems[i] = strconv.FormatUint(w, 10)
	}
	s := strings.Join(elems, "|")
	return "^(" + s + ")$"
}
func IntsRegExp(set ...interface{}) string {
	elems := make([]string, len(set))
	for i, v := range set {
		w, err := dtype.Int64(v)
		if err != nil {
			continue
		}
		elems[i] = strconv.FormatInt(w, 10)
	}
	s := strings.Join(elems, "|")
	return "^(" + s + ")$"
}
func StringsRegExp(set ...interface{}) string {
	// @TODO 使用正则转义

	elems := make([]string, len(set))
	for i, v := range set {
		elems[i] = dtype.String(v)
	}
	s := strings.Join(elems, "|")
	return "^(" + s + ")$"
}
