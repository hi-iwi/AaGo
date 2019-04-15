package com

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/kataras/iris"
	"github.com/luexu/AaGo/aa"
	"github.com/luexu/AaGo/dict"
)

type RespStruct struct {
	beforeFlush []func(*RespStruct)
	writer      http.ResponseWriter
	ic          iris.Context
	req         *Req
	aa.Error
	Payload interface{} `json:"data"`

	code    int
	headers map[string]string
	content []byte

	headlck sync.RWMutex
}

var (
	beforeFlush       = make([]func(*RespStruct), 3)
	LogHandler        func(code int, msg string) error
	RespDebugableFunc func(req *Req) bool
)

func Resp(p interface{}, as ...interface{}) (resp RespStruct) {
	resp.code = 200
	resp.headers = make(map[string]string, 5)
	for i := 0; i < len(as); i++ {
		if r, ok := as[i].(*Req); ok {
			resp.req = r
		} else if mws, ok := as[i].([]func(*RespStruct)); ok {
			resp.beforeFlush = mws
		}
	}

	if w, ok := p.(http.ResponseWriter); ok {
		resp.writer = w
	} else if c, ok := p.(iris.Context); ok {
		resp.ic = c
		resp.writer = c.ResponseWriter()
		if resp.req == nil {
			resp.req = NewReq(c)
		}
	}

	return
}

func AddGlbRespMidwares(mws ...func(*RespStruct)) {
	beforeFlush = append(beforeFlush, mws...)
}
func (resp RespStruct) AddMidwares(mws ...func(*RespStruct)) {
	resp.beforeFlush = append(resp.beforeFlush, mws...)
}

func (resp RespStruct) WriteHeader(code interface{}) {

	if c, ok := code.(int); ok {
		resp.code = c
	} else if e, ok := code.(*aa.Error); ok {
		if e == nil {
			resp.code = 200
		} else {
			resp.code = e.Code
		}
	}
	resp.WriteRaw()
}

func (resp RespStruct) writeNotModified() {
	w := resp.writer

	if resp.ic != nil {
		resp.ic.StatusCode(403)
	} else {
		resp.DelHeader("Content-Type")
		resp.DelHeader("Content-Length")
		if w.Header().Get("Etag") != "" {
			resp.DelHeader("Last-Modified")
		}
		w.WriteHeader(resp.code)
	}
}

func (resp RespStruct) WriteRaw(ps ...interface{}) (int, error) {
	w := resp.writer
	for i := 0; i < len(ps); i++ {
		if bytes, ok := ps[i].([]byte); ok {
			resp.content = bytes
		} else if str, ok := ps[i].(string); ok {
			resp.content = []byte(str)
		}
	}

	for i := 0; i < len(beforeFlush); i++ {
		if w := beforeFlush[i]; w != nil {
			w(&resp)
		}
	}

	for i := 0; i < len(resp.beforeFlush); i++ {
		if w := resp.beforeFlush[i]; w != nil {
			w(&resp)
		}
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

	w.WriteHeader(resp.code)

	if resp.req.Method != "HEAD" {
		if len(resp.content) > 0 {
			return w.Write(resp.content)
		}

	}

	return 0, nil
}

/*
Write(404)
Write(404, "Not Found")
Write(aa.Error{})
Write(aa.Error{}, data)
Write(aa.Error{}, data)
Write(data)
*/
func (resp RespStruct) Write(a interface{}, d ...interface{}) error {

	if e, ok := a.(*aa.Error); ok {
		resp.Code = e.Code
		resp.Msg = e.Msg
		if len(d) > 0 {
			resp.Payload = d[0]
		}
	} else if code, ok := a.(int); ok {
		resp.Code = code
		if len(d) == 0 {
			resp.Msg = dict.Code2Msg(code)
		} else {
			resp.Msg = aa.NewDtype(d[0]).String()
		}
	} else {
		resp.Code = 200
		resp.Msg = "OK"
		resp.Payload = a
	}

	if resp.Code >= 500 {
		if LogHandler != nil {
			if err := LogHandler(resp.Code, resp.Msg); err != nil {
				log.Println("[error] LogHandler error: ", err.Error())
				log.Printf("[error] %d %s\n", resp.Code, resp.Msg)
			}
		} else {
			log.Printf("[error] %d %s\n", resp.Code, resp.Msg)
		}
		resp.writeDebugInfo(resp.Msg)
		resp.Msg = dict.Code2Msg(resp.Code)
	}
	resp.writer.WriteHeader(resp.Code)
	b, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	resp.SetHeader("Content-Type", "application/json")
	resp.content = b

	resp.WriteRaw()
	return nil
}

func (resp RespStruct) writeDebugInfo(info string) {
	if RespDebugableFunc != nil {
		if RespDebugableFunc(resp.req) {
			resp.writer.Header().Set("Debug", info)
		}
	}
}
