package com

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

func (resp RespStruct) Header(head string) string {
	vs, ok := resp.writer.Header()[head]
	if ok && len(vs) > 0 {
		return vs[0]
	}
	resp.headlck.RLock()
	h := resp.headers
	resp.headlck.RUnlock()
	if v, ok := h[head]; ok {
		return v
	}
	return ""
}

func (resp RespStruct) DelHeader(head string) {
	if _, ok := resp.writer.Header()[head]; ok {
		delete(resp.writer.Header(), head)
	}
	resp.headlck.Lock()
	if _, ok := resp.headers[head]; ok {
		resp.headers[head] = ""
	}
	resp.headlck.Unlock()
}

func (resp RespStruct) SetHeader(head interface{}, values ...string) RespStruct {
	var key, value string
	if k, ok := head.(string); ok {
		for i := 0; i < len(values); i++ {
			key = k
			value = values[i]
		}
	} else if hs, ok := head.(map[string]string); ok {
		for k, v := range hs {
			key = k
			value = v
		}
	}
	if key == "Last-Modified" {
		value = fmtLastModified(value)
	}

	if value != "" {
		resp.headlck.Lock()
		resp.headers[key] = value
		resp.headlck.Unlock()
	}
	return resp
}
