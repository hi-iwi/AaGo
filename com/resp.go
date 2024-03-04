package com

import (
	"github.com/hi-iwi/AaGo/aa"
	"github.com/hi-iwi/AaGo/ae"
	"github.com/kataras/iris/v12"
	"net/http"
	"strings"
	"sync"

	"github.com/hi-iwi/AaGo/dict"
)

type RespStruct struct {
	beforeFlush     []func(*RespStruct)
	beforeSerialize []func(*RespContentDTO)
	afterSerialize  []func([]byte) []byte
	writer          http.ResponseWriter
	ictx            iris.Context
	req             *Req

	code          int
	headers       sync.Map
	content       []byte
	contentStruct RespContentDTO
}

var (
	HideServerErr    = defaultHideServerErr
	beforeSerialize  []func(*RespContentDTO)
	afterSerialize   []func([]byte) []byte
	beforeFlush      []func(*RespStruct)
	respContentTypes = make(map[string]struct{}) // struct{} takes 0 bytes

)

func defaultHideServerErr(ictx iris.Context, cs *RespContentDTO, r *Req) {
	if cs.Code >= 500 {
		msg := cs.Msg
		ctx := aa.Context(ictx)
		m := ae.Caller(1) + " " + msg
		Log.Error(ctx, m)

		// hide errmsg
		cs.Msg = dict.Code2Msg(cs.Code)
	}
}

// 注册通用resp content types；

func RegisterRespContentTypes(ctypes ...string) {
	for _, ctype := range ctypes {
		// 会同时把 ; charset=utf-8 一起注册进去
		respContentTypes[ctype] = struct{}{}
		if strings.IndexByte(ctype, ';') == 0 {
			respContentTypes[ctype+"; charset=utf-8"] = struct{}{}
		}
	}
}

/*
  resp Content-Type 优先级：
    --> 在 resp.Write 之前
		① new Resp() 时，通过 as 参数指定header；
		② ictx.Values() 设定的；
		③ controller 里面 resp.SetHeader() 或 resp.LoadOrSetHeader() 设置
    --> 在 resp.Write 阶段
		④ 客户端 Accept 指定  -> 必须要通过 RegisterRespContentTypes()注册过的才可以
		⑤ 根据客户 Content Type 相同  -> 必须要通过 RegisterRespContentTypes()注册过的才可以
		⑥ 根据content内容自动判定
*/
// @param as:  string 表示 Content-Type

func Resp(ictx iris.Context, req *Req, as ...interface{}) *RespStruct {
	resp := &RespStruct{
		req:    req,
		code:   200,
		ictx:   ictx,
		writer: ictx.ResponseWriter(),
	}
	var accept string
	for _, a := range as {
		if accept, _ = a.(string); accept != "" {
			// ①
			resp.SetHeader(ContentType, accept)
		} else if mw, ok := a.(func(*RespStruct)); ok {
			resp.beforeFlush = append(resp.beforeFlush, mw)
		} else if mw, ok := a.(func(*RespContentDTO)); ok {
			resp.beforeSerialize = append(resp.beforeSerialize, mw)
		} else if mw, ok := a.(func([]byte) []byte); ok {
			resp.afterSerialize = append(resp.afterSerialize, mw)
		}
	}
	if accept == "" {
		// ②
		if accept = ictx.Values().GetString(ContentType); accept != "" {
			resp.SetHeader(ContentType, accept)
		}
	}

	return resp
}

func AddGlbRespMidwares(mws ...interface{}) {
	for _, a := range mws {
		if mw, ok := a.(func(*RespStruct)); ok {
			beforeFlush = append(beforeFlush, mw)
		} else if mw, ok := a.(func(*RespContentDTO)); ok {
			beforeSerialize = append(beforeSerialize, mw)
		} else if mw, ok := a.(func([]byte) []byte); ok {
			afterSerialize = append(afterSerialize, mw)
		} else {
			panic("undefined global response middleware")
		}
	}
}

func (resp *RespStruct) AddMidwares(mws ...interface{}) {
	for _, a := range mws {
		if mw, ok := a.(func(*RespStruct)); ok {
			resp.beforeFlush = append(resp.beforeFlush, mw)
		} else if mw, ok := a.(func(*RespContentDTO)); ok {
			resp.beforeSerialize = append(resp.beforeSerialize, mw)
		} else if mw, ok := a.(func([]byte) []byte); ok {
			resp.afterSerialize = append(resp.afterSerialize, mw)
		} else {
			panic("undefined global response middleware")
		}
	}
}
