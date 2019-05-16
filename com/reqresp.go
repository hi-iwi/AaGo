package com

import (
	"github.com/kataras/iris"
)

func ReqResp(ictx iris.Context) (*Req, RespStruct) {
	r := NewReq(ictx)
	return r, Resp(ictx, r)
}
