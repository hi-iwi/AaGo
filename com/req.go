package com

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"net/url"
	"sync"

	"github.com/kataras/iris"
	"github.com/luexu/AaGo/ae"
)

type Req struct {
	ID          string
	Method      string
	r           *http.Request
	contentType string
	data        reqData
	raw         string
	parsed      bool
}
type reqData struct {
	qlck   sync.RWMutex
	hlck   sync.RWMutex
	blck   sync.RWMutex
	uri    string
	query  map[string]interface{}
	header map[string]interface{}
	body   map[string]interface{}
}

type maxBytesReader struct {
	w   http.ResponseWriter
	r   io.ReadCloser // underlying reader
	n   int64         // max bytes remaining
	err error         // sticky error
}

func (l *maxBytesReader) Read(p []byte) (n int, err error) {
	if l.err != nil {
		return 0, l.err
	}
	if len(p) == 0 {
		return 0, nil
	}
	// If they asked for a 32KB byte read but only 5 bytes are
	// remaining, no need to read 32KB. 6 bytes will answer the
	// question of the whether we hit the limit or go past it.
	if int64(len(p)) > l.n+1 {
		p = p[:l.n+1]
	}
	n, err = l.r.Read(p)

	if int64(n) <= l.n {
		l.n -= int64(n)
		l.err = err
		return n, err
	}

	n = int(l.n)
	l.n = 0

	// The server code and client code both use
	// maxBytesReader. This "requestTooLarge" check is
	// only used by the server code. To prevent binaries
	// which only using the HTTP Client code (such as
	// cmd/go) from also linking in the HTTP server, don't
	// use a static type assertion to the server
	// "*response" type. Check this interface instead:
	type requestTooLarger interface {
		requestTooLarge()
	}
	if res, ok := l.w.(requestTooLarger); ok {
		res.requestTooLarge()
	}
	l.err = errors.New("http: request body too large")
	return n, l.err
}

func (l *maxBytesReader) Close() error {
	return l.r.Close()
}

func NewReq(p interface{}) *Req {
	req := &Req{}
	if r, ok := p.(*http.Request); ok {
		req.r = r
		req.Method = r.Method
	} else if c, ok := p.(iris.Context); ok {
		req.ID = c.Values().GetString("traceid")
		req.r = c.Request()
		req.Method = req.r.Method
		if len(c.Params().Store) > 0 {
			req.data.qlck.Lock()
			req.data.query = make(map[string]interface{}, len(c.Params().Store))
			req.data.qlck.Unlock()
		}
		for _, v := range c.Params().Store {
			req.data.qlck.Lock()
			req.data.query[v.Key] = v.ValueRaw
			req.data.qlck.Unlock()
		}
	}

	return req
}

func (req *Req) ContentType() string {
	if req.contentType == "" {
		ct := req.r.Header.Get("Content-Type")
		if ct == "" {
			ct = "application/octet-stream"
		}
		req.contentType, _, _ = mime.ParseMediaType(ct)
	}
	return req.contentType
}

func (req *Req) Uri() string {
	if req.data.uri != "" {
		return req.data.uri
	}
	if req.r != nil {
		return req.r.URL.Path
	}
	return ""
}

func (req *Req) Headers() map[string]interface{} {

	if !req.parsed && req.r != nil {
		req.data.hlck.Lock()
		// @note 这里必须要判断 map[string]string 是否为空，并分配内存空间
		if req.data.header == nil {
			//req.data.header = make(map[string]string, len(req.r.Header))
		}
		req.data.hlck.Unlock()
	}
	req.data.hlck.RLock()
	h := req.data.header
	rh := req.r.Header
	req.data.hlck.RUnlock()

	headers := make(map[string]interface{}, len(rh)+len(h))
	for k := range rh {
		vs := rh[k]
		for i := 0; i < len(vs); i++ {
			if vs[i] != "" {
				headers[k] = vs[i]
				break
			}
		}
	}
	for k, v := range h {
		headers[k] = v
	}
	return headers
}

func (req *Req) Header(param string, patterns ...interface{}) (*ReqProp, *ae.Error) {
	req.data.hlck.RLock()
	h := req.data.header
	req.data.hlck.RUnlock()
	for k, v := range h {
		if k == param {
			r := NewReqProp(param, v)
			return r, r.Filter(patterns...)
		}
	}
	if req.r != nil {
		v := req.r.Header.Get(param)
		r := NewReqProp(param, v)
		return r, r.Filter(patterns...)
	}
	r := NewReqProp(param, "")
	return r, r.Filter(patterns...)
}

func (req *Req) Queries() map[string]interface{} {
	req.data.qlck.RLock()
	rq := req.r.URL.Query()
	dq := req.data.query
	req.data.qlck.RUnlock()
	queries := make(map[string]interface{}, len(rq)+len(dq))
	for k := range rq {
		vs := rq[k]
		for i := 0; i < len(vs); i++ {
			if vs[i] != "" {
				queries[k] = vs[i]
				break
			}
		}
	}
	for k, v := range dq {
		queries[k] = v
	}
	return queries
}

func (req *Req) Query(param string, patterns ...interface{}) (*ReqProp, *ae.Error) {
	req.data.qlck.RLock()
	q := req.data.query
	req.data.qlck.RUnlock()

	for k, v := range q {
		if k == param {
			r := NewReqProp(param, v)
			return r, r.Filter(patterns...)
		}
	}
	if req.r != nil {
		v := req.r.URL.Query().Get(param)
		r := NewReqProp(param, v)
		return r, r.Filter(patterns...)
	}
	r := NewReqProp(param, "")
	return r, r.Filter(patterns...)
}

func (req *Req) loadFormBody(d url.Values) {
	if len(d) == 0 {
		return
	}
	if len(req.data.body) == 0 {
		req.data.body = make(map[string]interface{}, len(d))
	}
	for k, vs := range d {

		if len(vs) > 0 {
			req.data.blck.Lock()
			req.data.body[k] = vs[0]
			req.data.blck.Unlock()
		}
	}
}
func (req *Req) Body(param string, patterns ...interface{}) (*ReqProp, *ae.Error) {
	ct := req.ContentType()
	if !req.parsed {
		// 参考 http.parsePostForm()  request.go  ParseForm()
		switch ct {
		case "application/json", "application/octet-stream":
			var reader io.Reader = req.r.Body
			maxFormSize := int64(1024)
			if _, ok := reader.(*maxBytesReader); !ok {
				maxFormSize = int64(10 << 20) // 10 MB is a lot of json.
				reader = io.LimitReader(req.r.Body, maxFormSize+1)
			}
			b, e := ioutil.ReadAll(req.r.Body)
			if e != nil {
				return NewReqProp(param, ""), ae.NewError(500, e)
			}
			if int64(len(b)) > maxFormSize {
				return NewReqProp(param, ""), ae.NewError(413, "Json body is too large")
			}
			req.raw = string(b)
			req.data.blck.Lock()
			json.Unmarshal(b, &req.data.body)
			req.data.blck.Unlock()
			req.parsed = true
		case "application/x-www-form-urlencoded":
			if req.r.PostForm == nil {
				req.r.ParseMultipartForm(1 << 20) // 1M
			}
			req.loadFormBody(req.r.PostForm)
		case "multipart/form-data":
			if req.r.MultipartForm != nil {
				req.loadFormBody(req.r.MultipartForm.Value)
			} else {
				req.loadFormBody(req.r.Form)
			}
		}
	}
	req.data.blck.RLock()
	b := req.data.body
	req.data.blck.RUnlock()

	if v, ok := b[param]; ok {
		r := NewReqProp(param, v)
		return r, r.Filter(patterns...)
	}

	fv := req.r.PostFormValue(param)
	r := NewReqProp(param, fv)
	return r, r.Filter(patterns...)
}
