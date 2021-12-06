package com

import (
	"github.com/hi-iwi/AaGo/aa"
	"github.com/kataras/iris/v12"
	"net/http"
	"runtime"
	"strconv"
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
	HideServerErr   = defaultHideServerErr
	beforeSerialize []func(*RespContentDTO)
	afterSerialize  []func([]byte) []byte
	beforeFlush     []func(*RespStruct)
)

func defaultHideServerErr(ictx iris.Context, cs *RespContentDTO, r *Req) {
	if cs.Code >= 500 {
		msg := cs.Msg
		ctx := aa.Context(ictx)
		cs.Msg = dict.Code2Msg(cs.Code)

		_, file, line, _ := runtime.Caller(2)
		a := strings.Split(file, "/")
		Log.Error(ctx, "file: %s, code: %d, msg: %s", a[len(a)-1]+":"+strconv.Itoa(line)+" ", cs.Code, msg)
	}
}

func Resp(ictx iris.Context, req *Req, as ...interface{}) *RespStruct {
	resp := &RespStruct{
		req:    req,
		code:   200,
		ictx:   ictx,
		writer: ictx.ResponseWriter(),
	}
	for _, a := range as {
		if mw, _ := a.(string); mw != "" {
			resp.SetHeader("Content-Type", mw)
		} else if mw, ok := a.(func(*RespStruct)); ok {
			resp.beforeFlush = append(resp.beforeFlush, mw)
		} else if mw, ok := a.(func(*RespContentDTO)); ok {
			resp.beforeSerialize = append(resp.beforeSerialize, mw)
		} else if mw, ok := a.(func([]byte) []byte); ok {
			resp.afterSerialize = append(resp.afterSerialize, mw)
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
