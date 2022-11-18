package com

type ContentTypes string

const (
	ErrCodeKey     = "com.error.code"
	ErrMsgKey      = "com.error.msg"
	ParamPage      = "page"
	ParamLimit     = "limit"
	ParamOffset    = "offset"
	ParamStringify = "_stringify"
	ParamField     = "_field"

	ContentRange  = "Content-Range"
	ContentType   = "Content-Type"
	ContentLength = "Content-Length"
	Etag          = "Etag"
	LastModified  = "Last-Modified"

	CtJson        ContentTypes = "application/json"
	CtHtml        ContentTypes = "text/html"
	CtJsonp       ContentTypes = "text/html"
	CtOctetStream ContentTypes = "application/octet-stream"
	CtForm        ContentTypes = "application/x-www-form-urlencoded"
	CtFormData    ContentTypes = "multipart/form-data"
)

func (t ContentTypes) String() string {
	return string(t)
}
func (t ContentTypes) Utf8() string {
	return t.String() + "; charset=utf-8"
}
func IsHtml(contentType string) bool {
	return contentType == CtHtml.String() || contentType == CtHtml.Utf8()
}
