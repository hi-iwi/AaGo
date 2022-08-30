package com

import (
	"context"
	"github.com/hi-iwi/AaGo/aa"
	"github.com/kataras/iris/v12"
)

// 读取json buffer 的时候，会清空掉 r.Body，所以这个使用一次；
func ReqResp(ictx iris.Context, as ...interface{}) (*Req, *RespStruct, context.Context) {
	r := NewReq(ictx)
	resp := Resp(ictx, r, as...)
	glbCt := ictx.Values().GetString("Content-Type")
	if glbCt != "" {
		// 一定要load or set，因为 as 可能重新设置 header，如jsonp
		resp.LoadOrSetHeader("Content-Type", glbCt)
	}
	ctx := aa.Context(ictx)
	return r, resp, ctx
}
