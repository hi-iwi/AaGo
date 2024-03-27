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

	"github.com/hi-iwi/AaGo/ae"
	"github.com/kataras/iris/v12"
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
	query  map[string]any
	header map[string]any
	body   map[string]any
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

func NewReq(ictx iris.Context) *Req {
	req := &Req{}
	req.ID = ictx.Values().GetString("traceid")
	req.r = ictx.Request()
	req.Method = req.r.Method
	if len(ictx.Params().Store) > 0 {
		req.data.qlck.Lock()
		req.data.query = make(map[string]any, len(ictx.Params().Store))
		req.data.qlck.Unlock()
	}
	for _, v := range ictx.Params().Store {
		req.data.qlck.Lock()
		req.data.query[v.Key] = v.ValueRaw
		req.data.qlck.Unlock()
	}
	return req
}

func (r *Req) ContentType() string {
	if r.contentType == "" {
		ct := r.r.Header.Get(ContentType)
		if ct == "" {
			ct = CtOctetStream.String()
		}
		r.contentType, _, _ = mime.ParseMediaType(ct)
	}
	return r.contentType
}

func (r *Req) Uri() string {
	if r.data.uri != "" {
		return r.data.uri
	}
	if r.r != nil {
		return r.r.URL.Path
	}
	return ""
}

func (r *Req) Headers() map[string]any {
	if !r.parsed && r.r != nil {
		r.data.hlck.Lock()
		// @note 这里必须要判断 map[string]string 是否为空，并分配内存空间
		if r.data.header == nil {
			//req.data.header = make(map[string]string, len(req.r.Header))
		}
		r.data.hlck.Unlock()
	}
	r.data.hlck.RLock()
	h := r.data.header
	rh := r.r.Header
	r.data.hlck.RUnlock()

	headers := make(map[string]any, len(rh)+len(h))
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

func (r *Req) Header(param string, patterns ...any) (*ReqProp, *ae.Error) {
	r.data.hlck.RLock()
	h := r.data.header
	r.data.hlck.RUnlock()
	for k, v := range h {
		if k == param {
			r := NewReqProp(param, v)
			return r, r.Filter(patterns...)
		}
	}
	if r.r != nil {
		v := r.r.Header.Get(param)
		r := NewReqProp(param, v)
		return r, r.Filter(patterns...)
	}
	p := NewReqProp(param, "")
	return p, p.Filter(patterns...)
}

func (r *Req) Queries() map[string]any {
	r.data.qlck.RLock()
	rq := r.r.URL.Query()
	dq := r.data.query
	r.data.qlck.RUnlock()
	queries := make(map[string]any, len(rq)+len(dq))
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

func (r *Req) Query(param string, patterns ...any) (*ReqProp, *ae.Error) {
	r.data.qlck.RLock()
	q := r.data.query
	r.data.qlck.RUnlock()

	for k, v := range q {
		if k == param {
			r := NewReqProp(param, v)
			return r, r.Filter(patterns...)
		}
	}
	if r.r != nil {
		v := r.r.URL.Query().Get(param)
		r := NewReqProp(param, v)
		return r, r.Filter(patterns...)
	}
	p := NewReqProp(param, "")
	return p, p.Filter(patterns...)
}

func (r *Req) loadFormBody(d url.Values) {
	if len(d) == 0 {
		return
	}
	if len(r.data.body) == 0 {
		r.data.body = make(map[string]any, len(d))
	}
	for k, vs := range d {

		if len(vs) > 0 {
			r.data.blck.Lock()
			r.data.body[k] = vs[0]
			r.data.blck.Unlock()
		}
	}
}
func (r *Req) Body(param string, patterns ...any) (*ReqProp, *ae.Error) {
	ct := r.ContentType()
	if !r.parsed {
		// 参考 http.parsePostForm()  request.go  ParseForm()
		switch ct {
		case CtJson.String(), CtJson.Utf8(), CtOctetStream.String(), CtOctetStream.Utf8():
			var reader io.Reader = r.r.Body
			maxFormSize := int64(1024)
			if _, ok := reader.(*maxBytesReader); !ok {
				maxFormSize = int64(10 << 20) // 10 MB is a lot of json.
				reader = io.LimitReader(r.r.Body, maxFormSize+1)
			}
			b, err := ioutil.ReadAll(r.r.Body)
			if err != nil {
				return NewReqProp(param, ""), ae.NewErr(err.Error())
			}
			if int64(len(b)) > maxFormSize {
				return NewReqProp(param, ""), ae.NewError(413, "Json body is too large")
			}
			r.raw = string(b)
			r.data.blck.Lock()
			json.Unmarshal(b, &r.data.body)
			r.data.blck.Unlock()
			r.parsed = true
		case CtForm.String(), CtForm.Utf8():
			if r.r.PostForm == nil {
				r.r.ParseMultipartForm(1 << 20) // 1M
			}
			r.loadFormBody(r.r.PostForm)
		case CtFormData.String(), CtFormData.Utf8():
			if r.r.MultipartForm != nil {
				r.loadFormBody(r.r.MultipartForm.Value)
			} else {
				r.loadFormBody(r.r.Form)
			}
		}
	}
	r.data.blck.RLock()
	b := r.data.body
	r.data.blck.RUnlock()

	if v, ok := b[param]; ok {
		r := NewReqProp(param, v)
		return r, r.Filter(patterns...)
	}

	fv := r.r.PostFormValue(param)
	p := NewReqProp(param, fv)
	return p, p.Filter(patterns...)
}
