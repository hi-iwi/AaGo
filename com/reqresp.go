package com

import (
	"github.com/kataras/iris"
)

func ReqResp(ctx iris.Context) (*Req, RespStruct) {
	r := NewReq(ctx)
	return r, Resp(ctx, r)
}
