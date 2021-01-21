package com

import "github.com/hi-iwi/dtype"

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
