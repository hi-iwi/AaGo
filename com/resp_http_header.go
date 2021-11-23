package com

import (
	"github.com/hi-iwi/AaGo/aenum"
)

/*
http fs.go
func setLastModified(w ResponseWriter, modtime time.Time) {
	if !isZeroTime(modtime) {
		w.Header().Set("Last-Modified", modtime.UTC().Format(TimeFormat))
	}
}
*/
func fmtLastModified(value string) string {
	if value == "Thu, 01 Jan 1970 00:00:00 GMT" {
		return ""
	}
	return value
}

func (resp *RespStruct) Header(head string) string {
	vs, ok := resp.writer.Header()[head]
	if ok && len(vs) > 0 {
		return vs[0]
	}
	v, _ := resp.headers.Load(head)
	s := v.(string)
	return s
}

func (resp *RespStruct) DelHeader(head string) {
	resp.headers.Delete(head)
}

func (resp *RespStruct) storeHeader(k, v string, loadOrStore bool) {
	if v == "" {
		resp.headers.Delete(k)
	} else {
		if k == aenum.LastModified {
			v = fmtLastModified(v)
		}
		if loadOrStore {
			resp.headers.LoadOrStore(k, v)
		} else {
			resp.headers.Store(k, v)
		}
	}
}
func (resp *RespStruct) SetHeader(k, v string) {
	resp.storeHeader(k, v, false)
}
func (resp *RespStruct) LoadOrSetHeader(k, v string) {
	resp.storeHeader(k, v, true)
}

func (resp *RespStruct) SetHeaders(heads map[string]string) {
	for k, v := range heads {
		resp.SetHeader(k, v)
	}
}
