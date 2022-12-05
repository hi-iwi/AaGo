package com

import (
	"errors"
	"github.com/hi-iwi/AaGo/ae"
	"github.com/hi-iwi/AaGo/dict"
	"github.com/hi-iwi/AaGo/util"
	"github.com/kataras/iris/v12/context"
	"net/http"
	"strings"
)

func (resp *RespStruct) write(cs RespContentDTO) error {

	for _, mw := range beforeSerialize {
		mw(&cs)
	}
	for _, mw := range resp.beforeSerialize {
		mw(&cs)
	}

	HideServerErr(resp.ictx, &cs, resp.req)

	ct, _ := resp.headers.Load(ContentType)
	var (
		b   []byte
		err error
	)
	switch ct.(type) {
	case string:
		if IsHtml(ct.(string)) {
			// 返回状态码，交给route层处理
			if context.StatusCodeNotSuccessful(cs.Code) {
				resp.ictx.Values().Set(ErrCodeKey, cs.Code)
				resp.ictx.Values().Set(ErrMsgKey, cs.Msg)
				resp.ictx.StatusCode(cs.Code)
				return nil
			}
		}
	}

	// json Marshal 不转译 HTML 字符
	b, err = util.JsonString(cs)
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
	_, err = resp.WriteRaw()
	return err
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
		resp.DelHeader(ContentType)
		resp.DelHeader(ContentLength)
		if w.Header().Get(Etag) != "" {
			resp.DelHeader(LastModified)
		}
		w.WriteHeader(resp.code)
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

func (resp *RespStruct) trySetContentType() {
	if _, ok := resp.headers.Load(ContentType); ok {
		return
	}

	// ④
	accept := resp.req.FastXheader("Accept").String()
	if accept != "" {
		accepts := strings.Split(accept, ",")
		for _, ac := range accepts {
			if _, ok := respContentTypes[ac]; ok {
				resp.headers.Store(ContentType, ac)
				return
			}
		}
	}
	// ⑤
	cliType := resp.req.ContentType()
	if _, ok := respContentTypes[cliType]; ok {
		resp.headers.Store(ContentType, cliType)
		return
	}
	// ⑥
	resp.headers.Store(ContentType, http.DetectContentType(resp.content)) // 这里需要解析 content，所以不要用 LoadOrStore()

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
	// @TODO 这里设置Content-Length之后，iris Gzip 就会异常
	resp.DelHeader(ContentLength) // 因为内容变更了，必须要把content-length设为空，不然客户端会读取错误
	resp.trySetContentType()

	resp.headers.Range(func(k, v interface{}) bool {
		s := v.(string)
		if s != "" {
			w.Header().Set(k.(string), s)
		}
		return true
	})

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

// 返回插入数据的ID，ID 可能是联合主键，或者字段不为id，那么就会以对象形式返回
// 如： {"id":12314}   {"id":"ADREDD"}   {"id":{"k":"i_am_prinary_key"}}  {"id": {"k":"", "uid":""}}
func (resp *RespStruct) WriteId(id string, ee ...*ae.Error) error {
	if len(ee) > 0 && ee[0] != nil {
		return resp.WriteE(ee[0])
	}
	return resp.Write(map[string]string{"id": id})
}
func (resp *RespStruct) WriteUint64Id(id uint64, ee ...*ae.Error) error {
	if len(ee) > 0 && ee[0] != nil {
		return resp.WriteE(ee[0])
	}
	return resp.Write(map[string]uint64{"id": id})
}
func (resp *RespStruct) WriteUintId(id uint, ee ...*ae.Error) error {
	if len(ee) > 0 && ee[0] != nil {
		return resp.WriteE(ee[0])
	}
	return resp.Write(map[string]uint{"id": id})
}
func (resp *RespStruct) WriteAliasId(alias string, id string, ee ...*ae.Error) error {
	if len(ee) > 0 && ee[0] != nil {
		return resp.WriteE(ee[0])
	}
	return resp.Write(map[string]string{alias: id})
}
func (resp *RespStruct) WriteUint64AliasId(alias string, id uint64, ee ...*ae.Error) error {
	if len(ee) > 0 && ee[0] != nil {
		return resp.WriteE(ee[0])
	}
	return resp.Write(map[string]uint64{alias: id})
}
func (resp *RespStruct) WriteUintAliasId(alias string, id uint, ee ...*ae.Error) error {
	if len(ee) > 0 && ee[0] != nil {
		return resp.WriteE(ee[0])
	}
	return resp.Write(map[string]uint{alias: id})
}

// k1,v1, k2, v2, k3,v3
func (resp *RespStruct) WriteJointId(args ...interface{}) error {
	l := len(args)
	if l < 2 || l%2 == 1 {
		return errors.New("no enough write joint id args  ")
	}
	id := make(map[string]interface{}, l/2)
	for i := 0; i < l; i += 2 {
		id[args[i].(string)] = args[i+1]
	}
	return resp.Write(id)
}

func (resp *RespStruct) WriteE(e *ae.Error) error {
	if e != nil {
		return resp.WriteSafeE(*e)
	}
	return resp.WriteCode(200)
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

func (resp *RespStruct) Write(a interface{}, ee ...*ae.Error) error {
	if len(ee) > 0 && ee[0] != nil {
		return resp.WriteE(ee[0])
	}

	//
	//v := reflect.ValueOf(a)
	//if a == nil || (v.Kind() == reflect.Ptr && v.IsNil()) {
	//	cs.Code = 200
	//	cs.Msg = dict.Code2Msg(200)
	//} else if e, ok := a.(*ae.Error); ok {
	//	cs.Code = e.Code
	//	cs.Msg = e.Msg
	//	if len(d) > 0 {
	//		cs.Payload = d[0]
	//	}
	//} else if code, ok := a.(int); ok {
	//	cs.Code = code
	//	if len(d) == 0 {
	//		cs.Msg = dict.Code2Msg(code)
	//	} else {
	//		cs.Msg = atype.New(d[0]).String()
	//	}
	//} else if (v.Kind() == reflect.Array || v.Kind() == reflect.Slice) && v.Len() == 0 {
	//	cs.Code = 204
	//	cs.Msg = dict.Code2Msg(cs.Code)
	//	cs.Payload = a
	//} else {

	//}
	payload, e := resp.handlePayload(a, "json")
	if e != nil {
		return resp.write(RespContentDTO{
			Code:    e.Code,
			Msg:     e.Msg,
			Payload: nil,
		})
	}
	return resp.write(RespContentDTO{
		Code:    200,
		Msg:     "OK",
		Payload: payload,
	})
}
func (resp *RespStruct) WriteJSONP(varname string, d map[string]interface{}) error {
	cs := RespContentDTO{
		Code: 200,
		Msg:  dict.Code2Msg(200),
	}
	if payload, e := resp.handlePayload(d, "json"); e != nil {
		cs.Code = e.Code
		cs.Msg = e.Msg
	} else {
		cs.Code = 200
		cs.Msg = dict.Code2Msg(200)
		cs.Payload = payload
	}
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
	c := []byte("<script>var " + varname + "=")
	c = append(c, b...)
	c = append(c, ";</script>"...)
	resp.SetHeader(ContentType, CtJsonp.String())
	resp.content = c
	_, err = resp.WriteRaw()
	return err
}
