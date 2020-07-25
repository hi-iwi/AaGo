package com

import (
	"context"
	"github.com/kataras/iris"
	"github.com/luexu/AaGo/aa"
)

func ReqResp(ictx iris.Context, respType ...string) (*Req, *RespStruct, context.Context) {
	r := NewReq(ictx)
	resp := Resp(ictx, r)
	ctx := aa.Context(ictx)
	if len(respType) > 0 {
		resp.SetHeader("Content-Type", respType[0])
	}
	return r, resp, ctx
}
