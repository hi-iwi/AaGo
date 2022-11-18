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
	ctx := aa.Context(ictx)
	return r, resp, ctx
}

// 强制返回 json
func ReqRespJson(ictx iris.Context, as ...interface{}) (*Req, *RespStruct, context.Context) {
	if as == nil {
		as = make([]interface{}, 0)
	}
	as = append(as, "application/json")
	return ReqResp(ictx, as...)
}

// 强制返回 html
func ReqRespHtml(ictx iris.Context, as ...interface{}) (*Req, *RespStruct, context.Context) {
	if as == nil {
		as = make([]interface{}, 0)
	}
	as = append(as, "text/html")
	return ReqResp(ictx, as...)
}

// 强制返回 xml
func ReqRespXml(ictx iris.Context, as ...interface{}) (*Req, *RespStruct, context.Context) {
	if as == nil {
		as = make([]interface{}, 0)
	}
	as = append(as, "application/xml")
	return ReqResp(ictx, as...)
}

// 强制返回 javascript/jsonp
func ReqRespJs(ictx iris.Context, as ...interface{}) (*Req, *RespStruct, context.Context) {
	if as == nil {
		as = make([]interface{}, 0)
	}
	as = append(as, "text/javascript")
	return ReqResp(ictx, as...)
}

func ReqRespJpeg(ictx iris.Context, as ...interface{}) (*Req, *RespStruct, context.Context) {
	if as == nil {
		as = make([]interface{}, 0)
	}
	as = append(as, "image/jpeg")
	return ReqResp(ictx, as...)
}

func ReqRespWebp(ictx iris.Context, as ...interface{}) (*Req, *RespStruct, context.Context) {
	if as == nil {
		as = make([]interface{}, 0)
	}
	as = append(as, "image/webp")
	return ReqResp(ictx, as...)
}

func ReqRespPng(ictx iris.Context, as ...interface{}) (*Req, *RespStruct, context.Context) {
	if as == nil {
		as = make([]interface{}, 0)
	}
	as = append(as, "image/png")
	return ReqResp(ictx, as...)
}

func ReqRespGif(ictx iris.Context, as ...interface{}) (*Req, *RespStruct, context.Context) {
	if as == nil {
		as = make([]interface{}, 0)
	}
	as = append(as, "image/gif")
	return ReqResp(ictx, as...)
}
