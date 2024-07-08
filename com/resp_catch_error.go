package com

import (
	"github.com/hi-iwi/AaGo/ae"
)

func (resp *RespStruct) CatchErrors(es ...*ae.Error) *ae.Error {
	e := ae.Check(es...)
	if e != nil {
		resp.WriteE(e)
		return e
	}
	return nil
}
