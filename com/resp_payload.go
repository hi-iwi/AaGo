package com

import (
	"github.com/hi-iwi/AaGo/ae"
)

func (resp *RespStruct) handlePayload(payload any, tagname string) (any, *ae.Error) {
	pf, e := resp.filterPayload(payload, tagname)
	if e != nil {
		return nil, e
	}

	p, e := resp.decoratePayload(pf, tagname)
	if e != nil {
		return nil, e
	}

	return p, nil
}
