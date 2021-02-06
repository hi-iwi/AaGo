package com

import (
	"github.com/hi-iwi/AaGo/aa"
	"github.com/hi-iwi/AaGo/util"
	"net/http"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/hi-iwi/AaGo/ae"
	"github.com/hi-iwi/AaGo/dict"
	"github.com/hi-iwi/dtype"
	"github.com/kataras/iris/v12"
)

type RespStruct struct {
	beforeFlush     []func(*RespStruct)
	beforeSerialize []func(*RespContentDTO)
	afterSerialize  []func([]byte) []byte
	writer          http.ResponseWriter
	ictx            iris.Context
	req             *Req

	code          int
	headers       map[string]string
	content       []byte
	contentStruct RespContentDTO

	headlck sync.RWMutex
}

var (
	HideServerErr   = defaultHideServerErr
	beforeSerialize []func(*RespContentDTO)
	afterSerialize  []func([]byte) []byte
	beforeFlush     []func(*RespStruct)
)

func defaultHideServerErr(ictx iris.Context, cs *RespContentDTO, r *Req) {
	if cs.Code >= 500 {
		ctx := aa.Context(ictx)
		cs.Msg = dict.Code2Msg(cs.Code)

		_, file, line, _ := runtime.Caller(1)
		a := strings.Split(file, "/")
		Log.Error(ctx, "file: %s, code: %d, msg: %s", a[len(a)-1]+":"+strconv.Itoa(line)+" ", cs.Code, cs.Msg)
	}
}

func Resp(p interface{}, as ...interface{}) *RespStruct {
	resp := &RespStruct{
		code: 200,
		headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	for _, a := range as {
		if r, ok := a.(*Req); ok {
			resp.req = r
		} else if mw, ok := a.(func(*RespStruct)); ok {
			resp.beforeFlush = append(resp.beforeFlush, mw)
		} else if mw, ok := a.(func(*RespContentDTO)); ok {
			resp.beforeSerialize = append(resp.beforeSerialize, mw)
		} else if mw, ok := a.(func([]byte) []byte); ok {
			resp.afterSerialize = append(resp.afterSerialize, mw)
		}
	}

	if w, ok := p.(http.ResponseWriter); ok {
		resp.writer = w
	} else if c, ok := p.(iris.Context); ok {
		resp.ictx = c
		resp.writer = c.ResponseWriter()
		if resp.req == nil {
			resp.req = NewReq(c)
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

func (resp *RespStruct) WriteHeader(code interface{}) {

	if c, ok := code.(int); ok {
		resp.code = c
	} else if e, ok := code.(*ae.Error); ok {
		if e == nil {
			resp.code = 200
		} else {
			resp.code = e.Code
		}
	}
	resp.WriteRaw()
}

func (resp *RespStruct) writeNotModified() {
	w := resp.writer

	if resp.ictx != nil {
		resp.ictx.StatusCode(403)
	} else {
		resp.DelHeader("Content-Type")
		resp.DelHeader("Content-Length")
		if w.Header().Get("Etag") != "" {
			resp.DelHeader("Last-Modified")
		}
		w.WriteHeader(resp.code)
	}
}

func (resp *RespStruct) WriteRaw(ps ...interface{}) (int, error) {
	w := resp.writer
	for i := 0; i < len(ps); i++ {
		if bytes, ok := ps[i].([]byte); ok {
			resp.content = bytes
		} else if str, ok := ps[i].(string); ok {
			resp.content = []byte(str)
		}
	}

	for _, mw := range beforeFlush {
		mw(resp)
	}

	for _, mw := range resp.beforeFlush {
		mw(resp)
	}

	if resp.code == 403 {
		resp.writeNotModified()
		return 0, nil
	}

	resp.headlck.Lock()
	resp.headers["Content-Length"] = ""
	if resp.headers["Content-Type"] == "" {
		resp.headers["Content-Type"] = http.DetectContentType(resp.content)
	}
	resp.headlck.Unlock()

	resp.headlck.RLock()
	hds := resp.headers
	resp.headlck.RUnlock()

	for k, v := range hds {
		if v != "" {
			w.Header().Set(k, v)
		}
	}
	resp.DelHeader("Content-Length")
	// @TODO 这里设置Content-Length之后，iris Gzip 就会异常
	//w.Header().Set("Content-Length", strconv.Itoa(len(resp.content)))

	//w.WriteHeader(resp.code)

	if resp.req.Method != "HEAD" {
		if len(resp.content) > 0 {
			return w.Write(resp.content)
		}

	}

	return 0, nil
}
func (resp *RespStruct) WriteOK() error {
	cs := RespContentDTO{
		Code: 200,
		Msg:  "OK",
	}
	return resp.write(cs)
}
func (resp *RespStruct) WriteE(e *ae.Error) error {
	if e != nil {
		return resp.WriteSafeE(*e)
	}
	return resp.Write(200)
}

func (resp *RespStruct) WriteSafeE(e ae.Error) error {
	cs := RespContentDTO{
		Code: e.Code,
		Msg:  e.Msg,
	}
	return resp.write(cs)
}
func (resp *RespStruct) WriteError(err error) error {
	cs := RespContentDTO{
		Code: 500,
		Msg:  err.Error(),
	}
	return resp.write(cs)
}

func (resp *RespStruct) WriteErr(code int, msg string) error {
	cs := RespContentDTO{
		Code: code,
		Msg:  msg,
	}
	return resp.write(cs)
}
func (resp *RespStruct) WriteCode(code int) error {
	cs := RespContentDTO{
		Code: code,
		Msg:  dict.Code2Msg(code),
	}
	return resp.write(cs)
}

func (resp *RespStruct) WriteErrMsg(msg string) error {
	cs := RespContentDTO{
		Code: 500,
		Msg:  msg,
	}
	return resp.write(cs)
}

/*
Write(404)
Write(404, "Not Found")
Write(ae.Error{})
Write(ae.Error{}, data)
Write(ae.Error{}, data)
Write(data)
*/
func (resp *RespStruct) Write(a interface{}, d ...interface{}) error {
	cs := RespContentDTO{}
	v := reflect.ValueOf(a)
	if a == nil || (v.Kind() == reflect.Ptr && v.IsNil()) {
		cs.Code = 200
		cs.Msg = "OK"
	} else if e, ok := a.(*ae.Error); ok {
		cs.Code = e.Code
		cs.Msg = e.Msg
		if len(d) > 0 {
			cs.Payload = d[0]
		}
	} else if code, ok := a.(int); ok {
		cs.Code = code
		if len(d) == 0 {
			cs.Msg = dict.Code2Msg(code)
		} else {
			cs.Msg = dtype.New(d[0]).String()
		}
	} else if (v.Kind() == reflect.Array || v.Kind() == reflect.Slice) && v.Len() == 0 {
		cs.Code = 204
		cs.Msg = dict.Code2Msg(cs.Code)
		cs.Payload = a
	} else {
		if payload, e := resp.handlePayload(a, "json"); e != nil {
			cs.Code = e.Code
			cs.Msg = e.Msg
		} else {
			cs.Code = 200
			cs.Msg = "OK"
			cs.Payload = payload
		}
	}

	return resp.write(cs)
}

func (resp *RespStruct) write(cs RespContentDTO) error {

	for _, mw := range beforeSerialize {
		mw(&cs)
	}
	for _, mw := range resp.beforeSerialize {
		mw(&cs)
	}

	HideServerErr(resp.ictx, &cs, resp.req)

	// json Marshal 不转译 HTML 字符
	b, err := util.JsonString(cs)
	if err != nil {
		return err
	}

	for _, mw := range afterSerialize {
		b = mw(b)
	}
	for _, mw := range resp.afterSerialize {
		b = mw(b)
	}

	resp.content = b
	resp.WriteRaw()
	return nil
}
