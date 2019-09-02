package com

import (
	"github.com/kataras/iris"
)

func ReqResp(ictx iris.Context, respType ...string) (*Req, *RespStruct) {
	r := NewReq(ictx)
	resp := Resp(ictx, r)
	if len(respType) > 0 {
		resp.SetHeader("Content-Type", respType[0])
	}
	return r, resp
}
