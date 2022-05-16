package com

import (
	"errors"

	"github.com/hi-iwi/AaGo/ae"
)

func (resp *RespStruct) CatchErrors(es ...*ae.Error) error {
	for i := 0; i < len(es); i++ {
		if es[i] != nil {
			resp.WriteSafeE(*es[i])
			return errors.New(es[i].Text())
		}
	}
	return nil
}
