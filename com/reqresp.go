package com

import (
	"context"
	"github.com/hi-iwi/AaGo/aa"
	"github.com/kataras/iris/v12"
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
