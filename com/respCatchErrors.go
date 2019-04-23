package com

import (
	"errors"

	"github.com/luexu/AaGo/ae"
)

func (resp RespStruct) CatchErrors(es ...*ae.Error) error {
	for i := 0; i < len(es); i++ {
		if es[i] != nil {
			resp.Write(es[i])
			return errors.New(es[i].Error())
		}
	}
	return nil
}
