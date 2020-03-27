package com

import (
	"errors"

	"github.com/luexu/AaGo/ae"
)

func (resp *RespStruct) CatchErrors(es ...*ae.Error) error {
	for i := 0; i < len(es); i++ {
		if es[i] != nil {
			resp.WriteE(*es[i])
			return errors.New(es[i].Error())
		}
	}
	return nil
}
