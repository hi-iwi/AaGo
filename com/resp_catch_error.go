package com

import (
	"github.com/hi-iwi/AaGo/ae"
)

func (resp *RespStruct) CatchErrors(es ...*ae.Error) *ae.Error {
	for i := 0; i < len(es); i++ {
		if es[i] != nil {
			resp.WriteSafeE(*es[i])
			return es[i]
		}
	}
	return nil
}
