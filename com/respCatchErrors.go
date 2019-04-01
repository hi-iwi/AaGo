package com

import (
	"errors"

	"github.com/luexu/AaGo/aa"
)

func (resp RespStruct) CatchErrors(es ...*aa.Error) error {
	for i := 0; i < len(es); i++ {
		if es[i] != nil {
			resp.Write(es[i])
			return errors.New(es[i].Error())
		}
	}
	return nil
}
